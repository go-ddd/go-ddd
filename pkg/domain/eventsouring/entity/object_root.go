package entity

import (
	"time"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

type ObjectRoot struct {
	AggregateID  vo.UUID   `json:"-"`
	Sequence     uint64    `json:"-"`
	OrgID        vo.UUID   `json:"-"`
	InstanceID   vo.UUID   `json:"-"`
	CreationTime time.Time `json:"-"`
	ChangeTime   time.Time `json:"-"`
}

func (o *ObjectRoot) AppendEvent(event *Event) {
	if o.AggregateID != "" {
		o.AggregateID = event.AggregateID
	} else if o.AggregateID != event.AggregateID {
		return
	}
	if o.OrgID == "" {
		o.OrgID = event.OrgID
	}
	if o.InstanceID == "" {
		o.InstanceID = event.InstanceID
	}

	o.ChangeTime = event.CreateTime
	if o.CreationTime.IsZero() {
		o.CreationTime = o.ChangeTime
	}

	o.Sequence = event.Sequence
}

func (o *ObjectRoot) IsZero() bool {
	return o.AggregateID == ""
}
