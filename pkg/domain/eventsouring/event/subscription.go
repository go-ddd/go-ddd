package event

import (
	"sync"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

var (
	subscriptions sync.Map // map[vo.AggregateType][]*Subscription
)

type Subscription struct {
	Events chan IEvent
	types  map[vo.AggregateType][]vo.EventType
}

func (s *Subscription) Unsubscribe() {
	for aggregate := range s.types {
		list, ok := subscriptions.Load(aggregate)
		if !ok {
			continue
		}
		subs := list.([]*Subscription)
		ok = false
		for i := len(subs) - 1; i >= 0; i-- {
			if subs[i] == s {
				subs = append(subs[:i], subs[i+1:]...)
				ok = true
			}
		}
		if !ok {
			continue
		}
		if len(subs) > 0 {
			subscriptions.Store(aggregate, subs)
		} else {
			subscriptions.Delete(aggregate)
		}
	}
	_, ok := <-s.Events
	if ok {
		close(s.Events)
	}
}

// SubscribeAggregates subscribes for all events on the given aggregates
func SubscribeAggregates(eventQueue chan IEvent, aggregateTypes ...vo.AggregateType) *Subscription {
	types := make(map[vo.AggregateType][]vo.EventType, len(aggregateTypes))
	for _, aggregate := range aggregateTypes {
		types[aggregate] = nil
	}
	sub := &Subscription{
		Events: eventQueue,
		types:  types,
	}

	for _, aggregate := range aggregateTypes {
		actual, loaded := subscriptions.LoadOrStore(aggregate, []*Subscription{sub})
		if loaded {
			subscriptions.Store(aggregate, append(actual.([]*Subscription), sub))
		}
	}

	return sub
}

// SubscribeEventTypes subscribes for the given event types
// if no event types are provided the subscription is for all events of the aggregate
func SubscribeEventTypes(eventQueue chan IEvent, aggregateType vo.AggregateType, eventTypes ...vo.EventType) *Subscription {
	sub := &Subscription{
		Events: eventQueue,
		types: map[vo.AggregateType][]vo.EventType{
			aggregateType: eventTypes,
		},
	}

	actual, loaded := subscriptions.LoadOrStore(aggregateType, []*Subscription{sub})
	if loaded {
		subscriptions.Store(aggregateType, append(actual.([]*Subscription), sub))
	}
	return sub
}

func Notify(events []IEvent) {
	for _, event := range events {
		aggregateType := event.GetAggregate().Type
		subs, ok := subscriptions.Load(aggregateType)
		if !ok {
			continue
		}
		for _, sub := range subs.([]*Subscription) {
			eventTypes := sub.types[aggregateType]
			// subscription for all events
			if len(eventTypes) == 0 {
				sub.Events <- event
				continue
			}
			// subscription for certain events
			for _, eventType := range eventTypes {
				if event.GetType() == eventType {
					sub.Events <- event
					break
				}
			}
		}
	}
}
