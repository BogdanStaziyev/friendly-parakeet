package coordinate

import (
	"fmt"
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
	FindOne(id int64) (*Coordinate, error)
	AddCoordinate(coordinate Coordinate) error
	UpdateCoordinate(coordinate *Coordinate) error
	InverseTask(coordinate1, coordinate2 *Coordinate) error
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

func (r *repository) FindOne(id int64) (*Coordinate, error) {
	var coordinate Coordinate
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	res := db.Collection("Coordinate").Find("id", id).One(&coordinate)
	fmt.Println(res)
	return &coordinate, nil
}

func (r *repository) AddCoordinate(coordinate Coordinate) error {
	return nil
}

func (r *repository) UpdateCoordinate(coordinate *Coordinate) error {
	return nil
}

func (r *repository) InverseTask(coordinate1, coordinate2 *Coordinate) error {
	fmt.Println(coordinate1.X + coordinate2.X)
	return nil
}
