package main

// Interfaces to allow switching of implementation.

type Repository interface{}

type Logger interface{}

type MessageBroker interface{}

type Services struct {
	repository Repository
	logger     Logger
	broker     MessageBroker
}

// Method Injection
func (s *Services) SetRepository(repo Repository) {
	s.repository = repo
}

func (s *Services) SetLogger(logger Logger) {
	s.logger = logger
}

func (s *Services) SetBroker(broker MessageBroker) {
	s.broker = broker
}
