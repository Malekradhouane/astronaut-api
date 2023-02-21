package postgres

import (
	"context"
	"github.com/jinzhu/gorm"
	"github/malekradhouane/test-cdi/errs"
	. "github/malekradhouane/test-cdi/store"
)

//Create creates an astronaut
func (c *Client) Create(ctx context.Context, astronaut *Astronaut) (*Astronaut, error) {
	err := c.db.Save(astronaut).Error
	if err != nil {
		return nil, err
	}
	return astronaut, nil
}

//Get retrieve a user
func (c *Client) Get(ctx context.Context, id int) (*Astronaut, error) {
	astronaut := new(Astronaut)
	err := c.db.Debug().Model(Astronaut{}).Where("id = ?", id).Take(astronaut).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errs.ErrNoSuchEntity
		}
		return nil, err
	}
	return astronaut, err
}

//List retrieve all users
func (c *Client) List(ctx context.Context) ([]*Astronaut, error) {
	var astronauts []*Astronaut
	result := c.db.Find(&astronauts)
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return nil, errs.ErrNoSuchEntity
		}
		return nil, result.Error
	}
	return astronauts, result.Error
}

//Update User
func (c *Client) Update(ctx context.Context, astronaut *Astronaut, id int) error {
	if err := c.db.Debug().Model(Astronaut{}).Where("id= ?", id).Update(astronaut).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errs.ErrNoSuchEntity
		}
		return err
	}
	return nil
}

//Delete remove user
func (c *Client) Delete(ctx context.Context, id int) error {
	return c.db.Where("id = ?", id).Delete(Astronaut{}).Error
}
