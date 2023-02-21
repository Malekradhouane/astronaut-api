package service

import (
	"context"
	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github/malekradhouane/test-cdi/store"
	"testing"
)

type MockAstronautStore struct {
	mock.Mock
}

func (m *MockAstronautStore) Create(ctx context.Context, astronaut *store.Astronaut) (*store.Astronaut, error) {
	args := m.Called(ctx, astronaut)
	return args.Get(0).(*store.Astronaut), args.Error(1)
}

func (m *MockAstronautStore) Get(ctx context.Context, id int) (*store.Astronaut, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*store.Astronaut), args.Error(1)
}

func (m *MockAstronautStore) List(ctx context.Context) ([]*store.Astronaut, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*store.Astronaut), args.Error(1)
}

func (m *MockAstronautStore) Update(ctx context.Context, astronaut *store.Astronaut, id int)  error {
	args := m.Called(ctx, astronaut, id)
	return  args.Error(0)
}

func (m *MockAstronautStore) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

var validAstronaut = &store.Astronaut{
	FirstName:        "astronaut1",
	LastName: "astronaut1",
	Email:   "astronaut@gmail.com",
}

var flagTestInsert = []struct {
	name          string
	astronautParam     *store.Astronaut
	expectedRes   *store.Astronaut
	returnedError error
	expectedError error
}{
	{
		"ok",
		validAstronaut,
		validAstronaut,
		nil,
		nil,
	},
}

func TestServiceCreate(t *testing.T) {
	for _, tt := range flagTestInsert {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := tassert.New(t)
			ctx := context.Background()

			mockStore := &MockAstronautStore{}
			mockStore.On("Create", ctx, tt.astronautParam).Return(tt.astronautParam, nil)

			service := NewAstronautService(mockStore)

			_, err := service.Create(ctx, tt.astronautParam)

			if tt.expectedError == nil {
				assert.NoError(err)
			} else {
				assert.ErrorIs(err, tt.expectedError)
			}
		})
	}
}

var astronauts []*store.Astronaut

var flagTestList = []struct {
	name          string
	astronautParam     *store.Astronaut
	expectedRes   []*store.Astronaut
	returnedError error
	expectedError error
}{
	{
		"ok",
		validAstronaut,
		append(astronauts, validAstronaut),
		nil,
		nil,
	},
}

func TestServiceList(t *testing.T) {
	for _, tt := range flagTestList {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := tassert.New(t)
			ctx := context.Background()
			mockStore := &MockAstronautStore{}
			mockStore.On("List", ctx).Return(tt.expectedRes, nil)

			service := NewAstronautService(mockStore)

			machines, err := service.List(ctx)
			assert.Equal(machines, tt.expectedRes)

			if tt.expectedError == nil {
				assert.NoError(err)
			} else {
				assert.ErrorIs(err, tt.expectedError)
			}
		})
	}
}

var flagTestDelete = []struct {
	name          string
	params        string
	returnedError error
	expectedError error
}{
	{
		"ok",
		"1",
		nil,
		nil,
	},
}

func TestServiceDelete(t *testing.T) {
	for _, tt := range flagTestDelete {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := tassert.New(t)
			ctx := context.Background()
			mockStore := &MockAstronautStore{}
			mockStore.On("Delete", ctx, 1).Return(nil)

			service := NewAstronautService(mockStore)

			err := service.Delete(ctx, tt.params)

			if tt.expectedError == nil {
				assert.NoError(err)
			} else {
				assert.ErrorIs(err, tt.expectedError)
			}
		})
	}
}