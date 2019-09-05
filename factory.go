package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/neuronlabs/neuron-core/mapping"
	"github.com/neuronlabs/neuron-core/repository"
)

/**

Factory

*/

// Factory is the repository.Factory mock implementation
type Factory struct {
	mock.Mock
}

// New creates new repository
// Implements repository.Factory method
func (f *Factory) New(s repository.Controller, model *mapping.ModelStruct) (repository.Repository, error) {
	return &Repository{}, nil
}

// DriverName returns the factory repository name
// Implements repository.Repository
func (f *Factory) DriverName() string {
	return DriverName
}

// Close closes the factory
func (f *Factory) Close(ctx context.Context, done chan<- interface{}) {
	done <- struct{}{}
}
