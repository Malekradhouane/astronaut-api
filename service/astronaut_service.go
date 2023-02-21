package service

import (
	"context"
	"github/malekradhouane/test-cdi/store"
	"strconv"
)

//AstronautService
type AstronautService struct {
	AstronautStore store.AstronautStore
}

//NewAstronautService constructs a new AstronautService
func NewAstronautService(as store.AstronautStore) *AstronautService {
	return &AstronautService{
		AstronautStore: as,
	}
}

//Create creates a new Astronaut
func (as *AstronautService) Create(ctx context.Context, req *store.Astronaut) (*store.Astronaut, error) {
	astronaut, err := as.AstronautStore.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return astronaut, nil
}

//Get retrieves an astronaut
func (as *AstronautService) Get(ctx context.Context, id string) (*store.Astronaut, error) {
	requestId, _ := strconv.Atoi(id)
	astronaut, err := as.AstronautStore.Get(ctx, requestId)
	if err != nil {
		return nil, err
	}
	return astronaut, nil
}

//List lists all astronauts
func (as *AstronautService) List(ctx context.Context) ([]*store.Astronaut, error) {
	astronauts, err := as.AstronautStore.List(ctx)
	if err != nil {
		return nil, err
	}
	return astronauts, nil
}

//Update updates an Astronaut
func (as *AstronautService) Update(ctx context.Context, req *store.Astronaut, id string) error {
	requestId, _ := strconv.Atoi(id)
	err := as.AstronautStore.Update(ctx, req, requestId)
	if err != nil {
		return err
	}
	return nil
}

//Delete _
func (as *AstronautService) Delete(ctx context.Context, id string) error {
	requestId, _ := strconv.Atoi(id)
	err := as.AstronautStore.Delete(ctx, requestId)
	if err != nil {
		return err
	}
	return nil
}
