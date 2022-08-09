package aggregate

import (
	"time"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/entity"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
	"github.com/galaxyobe/go-ddd/pkg/infrastructure/tools"
)

// WriteModel is the minimum representation of a command side write model.
// It implements a basic reducer
// it's purpose is to reduce events to create new ones
type WriteModel struct {
	AggregateID       vo.GUID        `json:"-"`
	ProcessedSequence uint64         `json:"-"`
	Events            []event.IEvent `json:"-"`
	OrgID             string         `json:"-"`
	InstanceID        string         `json:"-"`
	ChangeTime        time.Time      `json:"-"`
}

// AppendEvents adds all the events to the read model.
// The function doesn't compute the new state of the read model
func (rm *WriteModel) AppendEvents(events ...event.IEvent) {
	rm.Events = append(rm.Events, events...)
}

// Reduce is the basic implementaion of reducer
// If this function is extended the extending function should be the last step
func (wm *WriteModel) Reduce() error {
	if len(wm.Events) == 0 {
		return nil
	}

	aggregate := wm.Events[0].GetAggregate()
	if tools.IsGUIDNull(wm.AggregateID) {
		wm.AggregateID = aggregate.ID
	}
	if wm.OrgID == "" {
		wm.OrgID = aggregate.OrgID
	}
	if wm.InstanceID == "" {
		wm.InstanceID = aggregate.InstanceID
	}

	wm.ProcessedSequence = wm.Events[len(wm.Events)-1].GetSequence()
	wm.ChangeTime = wm.Events[len(wm.Events)-1].GetCreateTime()

	// all events processed and not needed anymore
	wm.Events = nil
	wm.Events = []event.IEvent{}

	return nil
}

// NewAggregateFromWriteModel maps the given WriteModel to an Aggregate.
func NewAggregateFromWriteModel(
	wm *WriteModel,
	typ vo.AggregateType,
	version vo.Version,
) *entity.Aggregate {
	return &entity.Aggregate{
		ID:         wm.AggregateID,
		Type:       typ,
		OrgID:      wm.OrgID,
		InstanceID: wm.InstanceID,
		Version:    version,
	}
}
