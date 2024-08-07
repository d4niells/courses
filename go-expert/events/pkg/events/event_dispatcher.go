package events

import (
	"errors"
	"sync"
)

var ErrEventHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{handlers: make(map[string][]EventHandlerInterface)}
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, exists := ed.handlers[event.GetName()]; exists {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, exists := ed.handlers[eventName]; exists {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrEventHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, exists := ed.handlers[eventName]; exists {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, exists := ed.handlers[eventName]; exists {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
