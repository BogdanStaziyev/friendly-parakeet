package coordinate

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
	FindAll() ([]Coordinate, error)
}

type repository struct {
	//Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Coordinate, error) {
	coordinates := make([]Coordinate, 1)
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	res := db.Collection("Coordinate")
	err = res.Find().All(&coordinates)
	return coordinates, nil
}
