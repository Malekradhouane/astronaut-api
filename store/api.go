package store

import "github.com/jinzhu/gorm"

type Astronaut struct {
	gorm.Model
	FirstName        string    `json:"firstName" gorm:"column:firstname;"`
	LastName         string    `json:"lastName" gorm:"column:lastname;"`
	Email            string    `json:"email" gorm:"email:varchar(120);unique"`
}
