package event

import (
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `localhost:54322`,
	User:     `postgres`,
	Password: `password`,
}

type Repository interface {
	FindAll() ([]Event, error)
}

const EventsCount int64 = 10

type repository struct {
	//Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	events := make([]Event, EventsCount)
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	res := db.Collection("Event")
	err = res.Find().All(&events)
	return events, nil
}
