package main

import "context"

// Interfaces to allow switching of implementation.

type Repository interface{}

type Logger interface{}

type MessageBroker interface{}

type Services struct {
	repository Repository
	logger     Logger
	broker     MessageBroker
}

// Constructor Injection
func NewServices(repository Repository, logger Logger, broker MessageBroker) *Services {
	return &Services{
		repository: repository,
		logger:     logger,
		broker:     broker,
	}
}

func (s *Services) Do(ctx context.Context) error {
	// do action associated with injected services.
	return nil
}
