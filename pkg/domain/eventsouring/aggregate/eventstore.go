package aggregate

import (
	"context"
	"errors"
	"sync"

	"github.com/galaxyobe/go-ddd/pkg/contexts"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/entity"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"

	"github.com/Masterminds/squirrel"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/repository"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// Eventstore abstracts all functions needed to store valid events
// and filters the stored events
type Eventstore struct {
	repo              repository.IEventStore
	eventInterceptors sync.Map // map[vo.EventType]eventTypeInterceptor
}

type eventTypeInterceptor struct {
	eventMapper func(*entity.Event) (event.IEvent, error)
}

func NewEventstore(repo repository.IEventStore) *Eventstore {
	return &Eventstore{
		repo: repo,
	}
}

// Health checks if the eventstore can properly work
// It checks if the repository can serve load
func (es *Eventstore) Health(ctx context.Context) error {
	return es.repo.Health(ctx)
}

// Push pushes the events in a single transaction
// an event needs at least an aggregate
func (es *Eventstore) Push(ctx context.Context, cmds ...event.ICommand) ([]event.IEvent, error) {
	events, constraints, err := commandsToRepository(contexts.GetInstanceID(ctx), cmds)
	if err != nil {
		return nil, err
	}
	err = es.repo.Push(ctx, events, constraints...)
	if err != nil {
		return nil, err
	}

	eventReaders, err := es.mapEvents(events)
	if err != nil {
		return nil, err
	}

	go event.Notify(eventReaders)
	return eventReaders, nil
}

func (es *Eventstore) NewInstance(ctx context.Context, instanceID string) error {
	return es.repo.CreateInstance(ctx, instanceID)
}

func commandsToRepository(instanceID string, cmds []event.ICommand) (events []*entity.Event, constraints []*vo.UniqueConstraint, err error) {
	events = make([]*entity.Event, len(cmds))
	for i, cmd := range cmds {
		if err = cmd.Validate(); err != nil {
			return nil, nil, err
		}
		data, err := cmd.GetData().MarshalData()
		if err != nil {
			return nil, nil, err
		}
		aggregate := cmd.GetAggregate()
		events[i] = &entity.Event{
			Type:          cmd.GetType(),
			AggregateID:   aggregate.ID,
			AggregateType: aggregate.Type,
			OrgID:         aggregate.OrgID,
			InstanceID:    instanceID,
			Version:       aggregate.Version,
			Metadata:      cmd.GetMetadata(),
			Data:          data,
			Service:       cmd.GetService(),
			Creator:       cmd.GetCreator(),
		}
		uniqueConstraints := cmd.GetUniqueConstraints()
		if len(uniqueConstraints) > 0 {
			constraints = append(constraints, uniqueConstraints...)
		}
	}

	return events, constraints, nil
}

// Filter filters the stored events based on the searchQuery
// and maps the events to the defined event structs
func (es *Eventstore) Filter(ctx context.Context, query squirrel.SelectBuilder) ([]event.IEvent, error) {
	if instanceID := contexts.GetInstanceID(ctx); instanceID != "" {
		query = query.Where("instance_id = ?", instanceID)
	}
	events, err := es.repo.Filter(ctx, query)
	if err != nil {
		return nil, err
	}
	return es.mapEvents(events)
}

func (es *Eventstore) mapEvents(events []*entity.Event) (mappedEvents []event.IEvent, err error) {
	mappedEvents = make([]event.IEvent, len(events))

	for i, e := range events {
		interceptor := es.getEventTypeInterceptor(e.GetType())
		if interceptor.eventMapper == nil {
			return nil, errors.New("event mapper not defined")
		}
		mappedEvents[i], err = interceptor.eventMapper(e)
		if err != nil {
			return nil, err
		}
	}
	return mappedEvents, nil
}

func (es *Eventstore) getEventTypeInterceptor(eventType vo.EventType) eventTypeInterceptor {
	value, ok := es.eventInterceptors.Load(eventType)
	if !ok {
		return eventTypeInterceptor{}
	}
	interceptor, ok := value.(eventTypeInterceptor)
	if !ok {
		return eventTypeInterceptor{}
	}
	return interceptor
}

type Reducer interface {
	// Reduce handles the events of the internal events list
	// it only appends the newly added events
	Reduce() error
	// AppendEvents appends the passed events to an internal list of events
	AppendEvents(...event.IEvent)
}

// FilterToReducer filters the events based on the search query, appends all events to the reducer and calls it's reduce function
func (es *Eventstore) FilterToReducer(ctx context.Context, query squirrel.SelectBuilder, r Reducer) error {
	events, err := es.Filter(ctx, query)
	if err != nil {
		return err
	}
	r.AppendEvents(events...)
	return r.Reduce()
}

// LatestSequence filters the latest sequence for the given search query
func (es *Eventstore) LatestSequence(ctx context.Context, query squirrel.SelectBuilder) (uint64, error) {
	if instanceID := contexts.GetInstanceID(ctx); instanceID != "" {
		query = query.Where("instance_id = ?", instanceID)
	}
	return es.repo.LatestSequence(ctx, query)
}

// InstanceIDs returns the instance ids found by the search query
func (es *Eventstore) InstanceIDs(ctx context.Context, query squirrel.SelectBuilder) ([]string, error) {
	if instanceID := contexts.GetInstanceID(ctx); instanceID != "" {
		query = query.Where("instance_id = ?", instanceID)
	}
	return es.repo.InstanceIDs(ctx, query)
}

type QueryReducer interface {
	Reducer
	// Query returns the SearchQueryFactory for the events needed in reducer
	Query() squirrel.SelectBuilder
}

// FilterToQueryReducer filters the events based on the search query of the query function,
// appends all events to the reducer and calls it's reduce function
func (es *Eventstore) FilterToQueryReducer(ctx context.Context, r QueryReducer) error {
	events, err := es.Filter(ctx, r.Query())
	if err != nil {
		return err
	}
	r.AppendEvents(events...)

	return r.Reduce()
}

// RegisterFilterEventMapper registers a function for mapping an eventstore event to an event
func (es *Eventstore) RegisterFilterEventMapper(eventType vo.EventType, mapper func(*entity.Event) (event.IEvent, error)) *Eventstore {
	if mapper == nil || eventType == "" {
		return es
	}
	es.eventInterceptors.Store(eventType, eventTypeInterceptor{
		eventMapper: mapper,
	})
	return es
}
