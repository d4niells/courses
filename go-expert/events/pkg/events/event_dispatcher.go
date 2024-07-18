package events

import "errors"

var ErrEventHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{handlers: make(map[string][]EventHandlerInterface)}
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
