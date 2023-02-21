package store

import (
	"context"
)

//AstronautStore represents the interface to manage User storage
type AstronautStore interface {
	Create(context.Context, *Astronaut) (*Astronaut, error)
	Get(context.Context, int) (*Astronaut, error)
	List(context.Context) ([]*Astronaut, error)
	Update(context.Context, *Astronaut, int) error
	Delete(context.Context, int) error
}
