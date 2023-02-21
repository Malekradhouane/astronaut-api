package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//Config db config
type Config struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
	Env      string
}

func (cfg *Config) URI() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s timeZone=Europe/Paris  sslmode=disable",
		cfg.Host,
		cfg.DB,
		cfg.User,
		cfg.Password,
	)
}

//Client postgres client
type Client struct {
	db     *gorm.DB
	models []interface{}
}

//NewClient db client constructor
func NewClient(c *Config, models []interface{}) (*Client, error) {
	db, err := gorm.Open("postgres", c.URI())
	if err != nil {
		return nil, errors.Wrap(err, "gorm Open")
	}

	db = db.Debug()
	err = db.AutoMigrate(models...).Error
	if err != nil {
		return nil, errors.Wrap(err, "gorm AutoMigrate")
	}
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(10)

	return &Client{db: db, models: models}, nil
}

// Teardown teardown all db tables
func (c *Client) Teardown() error {
	return errors.Wrap(c.db.DropTableIfExists(c.models...).Error, "DropTableIfExists")
}

// Shutdown close client connection to db
func (c *Client) Shutdown() error {
	return errors.Wrap(c.db.Close(), "Close")
}
