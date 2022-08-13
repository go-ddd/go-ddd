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
	eventQueue chan event.IEvent
}

func NewHandler(option HandlerOptions) Handler {
	if option.EventQueueSize <= 0 {
		option.EventQueueSize = DefaultEventQueueSize
	}
	return Handler{
		Eventstore: option.Eventstore,
		eventQueue: make(chan event.IEvent, option.EventQueueSize),
	}
}

func (h *Handler) GetEventQueue() <-chan event.IEvent {
	return h.eventQueue
}

func (h *Handler) Subscribe(aggregateTypes ...vo.AggregateType) {
	h.Sub = event.SubscribeAggregates(h.eventQueue, aggregateTypes...)
}

func (h *Handler) SubscribeEventTypes(aggregateType vo.AggregateType, eventTypes ...vo.EventType) {
	h.Sub = event.SubscribeEventTypes(h.eventQueue, aggregateType, eventTypes...)
}

func (h *Handler) Unsubscribe() {
	if h.Sub == nil {
		return
	}
	h.Sub.Unsubscribe()
}
