package projection

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"

	"github.com/galaxyobe/go-ddd/pkg/contexts"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/aggregate"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
	zap_helper "github.com/galaxyobe/go-ddd/pkg/infrastructure/helper/zap"
)

var (
	ErrSomeStmtsFailed = errors.New("some statements failed")
)

type HandlerOptions struct {
	aggregate.HandlerOptions
	ProjectionName      string
	RequeueEvery        time.Duration
	RetryFailedAfter    time.Duration
	Retries             uint
	ConcurrentInstances uint

	Logger *zap.Logger
}

// Update updates the projection with the given statements
type Update func(context.Context, []*vo.Statement, Reduce) (index int, err error)

// Reduce reduces the given event to a vo.Statement
// which is used to update the projection
type Reduce func(event.IEvent) (*vo.Statement, error)

// SearchQuery generates the search query to lookup for events
type SearchQuery func(ctx context.Context, instanceIDs []string) (query squirrel.SelectBuilder, queryLimit uint64, err error)

// Lock is used for mutex handling if needed on the projection
type Lock func(context.Context, time.Duration, ...string) <-chan error

// Unlock releases the mutex of the projection
type Unlock func(...string) error

type Handler struct {
	aggregate.Handler
	ProjectionName      string
	reduce              Reduce
	update              Update
	searchQuery         SearchQuery
	triggerProjection   *time.Timer
	lock                Lock
	unlock              Unlock
	requeueAfter        time.Duration
	retryFailedAfter    time.Duration
	retries             int
	concurrentInstances int

	Logger zap_helper.ILogger
}

func NewProjectionHandler(
	ctx context.Context,
	options HandlerOptions,
	reduce Reduce,
	update Update,
	query SearchQuery,
	lock Lock,
	unlock Unlock,
) *Handler {
	concurrentInstances := int(options.ConcurrentInstances)
	if concurrentInstances < 1 {
		concurrentInstances = 1
	}
	h := &Handler{
		Handler:             aggregate.NewHandler(options.HandlerOptions),
		ProjectionName:      options.ProjectionName,
		reduce:              reduce,
		update:              update,
		searchQuery:         query,
		lock:                lock,
		unlock:              unlock,
		requeueAfter:        options.RequeueEvery,
		triggerProjection:   time.NewTimer(0), // first trigger is instant on startup
		retryFailedAfter:    options.RetryFailedAfter,
		retries:             int(options.Retries),
		concurrentInstances: concurrentInstances,
		Logger:              zap_helper.Wrap(options.Logger.With(zap.String("projection", options.ProjectionName))),
	}

	go h.subscribe(ctx)

	go h.schedule(ctx)

	return h
}

// Trigger handles all events for the provided instances (or current instance from context if non specified)
// by calling FetchEvents and Process until the amount of events is smaller than the BulkLimit
func (h *Handler) Trigger(ctx context.Context, instances ...string) error {
	ids := []string{contexts.GetInstanceID(ctx)}
	if len(instances) > 0 {
		ids = instances
	}
	for {
		events, hasLimitExceeded, err := h.FetchEvents(ctx, ids...)
		if err != nil {
			return err
		}
		if len(events) == 0 {
			return nil
		}
		_, err = h.Process(ctx, events...)
		if err != nil {
			return err
		}
		if !hasLimitExceeded {
			return nil
		}
	}
}

// Process handles multiple events by reducing them to statements and updating the projection
func (h *Handler) Process(ctx context.Context, events ...event.IEvent) (index int, err error) {
	if len(events) == 0 {
		return 0, nil
	}
	index = -1
	statements := make([]*vo.Statement, len(events))
	for i, event := range events {
		statements[i], err = h.reduce(event)
		if err != nil {
			return index, err
		}
	}
	for retry := 0; retry <= h.retries; retry++ {
		index, err = h.update(ctx, statements[index+1:], h.reduce)
		if err != nil && !errors.Is(err, ErrSomeStmtsFailed) {
			return index, err
		}
		if err == nil {
			return index, nil
		}
		time.Sleep(h.retryFailedAfter)
	}
	return index, err
}

// FetchEvents checks the current sequences and filters for newer events
func (h *Handler) FetchEvents(ctx context.Context, instances ...string) ([]event.IEvent, bool, error) {
	eventQuery, eventsLimit, err := h.searchQuery(ctx, instances)
	if err != nil {
		return nil, false, err
	}
	events, err := h.Eventstore.Filter(ctx, eventQuery)
	if err != nil {
		return nil, false, err
	}
	return events, int(eventsLimit) == len(events), err
}

func (h *Handler) subscribe(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		err := recover()
		if err != nil {
			h.Handler.Unsubscribe()
			h.Logger.Error("subscription panicked", zap.Any("recover", err))
		}
		cancel()
	}()
	for firstEvent := range h.GetEventQueue() {
		events := checkAdditionalEvents(h.GetEventQueue(), firstEvent)

		index, err := h.Process(ctx, events...)
		if err != nil || index < len(events)-1 {
			h.Logger.OnError(err).Warn("unable to process all events from subscription")
		}
	}
}

func (h *Handler) schedule(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		err := recover()
		if err != nil {
			h.Logger.Error("schedule panicked",
				zap.Any("recover", err),
				zap.String("stack", string(debug.Stack())),
			)
		}
		cancel()
	}()
	for range h.triggerProjection.C {
		ids, err := h.Eventstore.InstanceIDs(ctx, squirrel.SelectBuilder{})
		if err != nil {
			h.Logger.Error("instance ids", zap.Error(err))
			h.triggerProjection.Reset(h.requeueAfter)
			continue
		}
		for i := 0; i < len(ids); i = i + h.concurrentInstances {
			max := i + h.concurrentInstances
			if max > len(ids) {
				max = len(ids)
			}
			instances := ids[i:max]
			lockCtx, cancelLock := context.WithCancel(ctx)
			errs := h.lock(lockCtx, h.requeueAfter, instances...)
			// wait until projection is locked
			if err, ok := <-errs; err != nil || !ok {
				cancelLock()
				h.Logger.Warn("initial lock failed", zap.Error(err))
				continue
			}
			go h.cancelOnErr(lockCtx, errs, cancelLock)
			err = h.Trigger(lockCtx, instances...)
			if err != nil {
				h.Logger.Warn("trigger failed",
					zap.Strings("instanceIDs", instances),
					zap.Error(err),
				)
			}

			cancelLock()
			unlockErr := h.unlock(instances...)
			if unlockErr != nil {
				h.Logger.Warn("unable to unlock",
					zap.Strings("instanceIDs", instances),
					zap.Error(unlockErr),
				)
			}
		}
		h.triggerProjection.Reset(h.requeueAfter)
	}
}

func (h *Handler) cancelOnErr(ctx context.Context, errs <-chan error, cancel func()) {
	for {
		select {
		case err := <-errs:
			if err != nil {
				h.Logger.Warn("bulk canceled", zap.Error(err))
				cancel()
				return
			}
		case <-ctx.Done():
			cancel()
			return
		}

	}
}

func checkAdditionalEvents(eventQueue <-chan event.IEvent, e event.IEvent) []event.IEvent {
	events := make([]event.IEvent, 1)
	events[0] = e
	for {
		select {
		case e := <-eventQueue:
			events = append(events, e)
		default:
			return events
		}
	}
}
