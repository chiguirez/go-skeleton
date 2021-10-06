package domain_name

//go:generate moq -out event_dispatcher_mock.go . EventDispatcher
type EventDispatcher interface {
	Dispatch(domainEvent []DomainEvent)
}

type DomainEvent interface {
	GetAggregateID() string
}

type AggregateRoot interface {
	Record(DomainEvent)
	GetEvents() []DomainEvent
	ClearEvents()
}

type BaseAggregate struct {
	events []DomainEvent
}

func (a *BaseAggregate) GetEvents() []DomainEvent {
	return a.events
}

func (a *BaseAggregate) ClearEvents() {
	a.events = []DomainEvent{}
}

func (a *BaseAggregate) Record(event DomainEvent) {
	if a.events == nil {
		a.events = make([]DomainEvent, 0)
	}

	a.events = append(a.events, event)
}
