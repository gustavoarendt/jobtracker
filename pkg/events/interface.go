package events

import "sync"

type EventInterface interface {
	GetName() string
	GetData() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Add(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string) bool
	Clear() error
}
