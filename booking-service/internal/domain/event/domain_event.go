package event

type DomainEvent interface {
	EventName() string
}
