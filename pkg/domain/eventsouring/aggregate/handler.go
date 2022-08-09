package aggregate

import (
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

const (
	DefaultEventQueueSize = 100
)

type HandlerOptions struct {
	Eventstore     *Eventstore
	EventQueueSize int
}

type Handler struct {
	Eventstore *Eventstore
	Sub        *event.Subscription
	EventQueue chan event.IEvent
}

func NewHandler(option HandlerOptions) Handler {
	if option.EventQueueSize <= 0 {
		option.EventQueueSize = DefaultEventQueueSize
	}
	return Handler{
		Eventstore: option.Eventstore,
		EventQueue: make(chan event.IEvent, option.EventQueueSize),
	}
}

func (h *Handler) Subscribe(aggregateTypes ...vo.AggregateType) {
	h.Sub = event.SubscribeAggregates(h.EventQueue, aggregateTypes...)
}

func (h *Handler) SubscribeEventTypes(aggregateType vo.AggregateType, eventTypes ...vo.EventType) {
	h.Sub = event.SubscribeEventTypes(h.EventQueue, aggregateType, eventTypes...)
}
