package coordinate

import (
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type coordinate struct {
	Id      int     `json:"id" db:"id"`
	MT      int     `json:"mt" db:"mt"`
	Axis    string  `json:"axis" db:"axis"`
	Horizon string  `json:"horizon" db:"horizon"`
	X       float64 `json:"x" db:"x"`
	Y       float64 `json:"y" db:"y"`
}

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `localhost:54322`,
	User:     `postgres`,
	Password: `password`,
}

type Repository interface {
	FindAll() ([]Coordinate, error)
	FindOne(id int64) (*Coordinate, error)
	AddCoordinate(coordinate *Coordinate) (*Coordinate, error)
	UpdateCoordinate(coordinate *Coordinate) error
	InverseTask(coordinate1, coordinate2 *Coordinate) error
}

type repository struct {
	coll db.Collection
}

func NewRepository(dbSession *db.Session) Repository {
	return &repository{
		coll: (*dbSession).Collection("Coordinate"),
	}
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

func (r *repository) AddCoordinate(coordinate *Coordinate) (*Coordinate, error) {
	coord := mapCoordinateDbModel(coordinate)
	err := r.coll.InsertReturning(coord)
	if err != nil {
		return nil, fmt.Errorf("Coordinaterepository Create: %w", err)
	}

	return mapCoordinateDbModelToDomain(coord), nil
}

//db, err := postgresql.Open(settings)
//if err != nil {
//	log.Fatal("Open: ", err)
//}
//defer db.Close()
//res, err := db.Collection("Coordinate").Insert(coordinate)
//if err != nil {
//	fmt.Printf("Insert filed: %s", err)
//}
//fmt.Println(res)
//return nil

func (r *repository) UpdateCoordinate(coordinate *Coordinate) error {
	return nil
}

func (r *repository) InverseTask(coordinate1, coordinate2 *Coordinate) error {
	fmt.Println(coordinate1.X + coordinate2.X)
	return nil
}

func mapCoordinateDbModelToDomain(coordinate *coordinate) *Coordinate {
	return &Coordinate{
		Id:      coordinate.Id,
		MT:      coordinate.MT,
		Axis:    coordinate.Axis,
		Horizon: coordinate.Horizon,
		X:       coordinate.X,
		Y:       coordinate.Y,
	}
}

func mapCoordinateDbModelToDomainCollection(coordinate []coordinate) []Coordinate {
	var result []Coordinate
	for _, c := range coordinate {
		newCoordinate := mapCoordinateDbModelToDomain(&c)
		result = append(result, *newCoordinate)
	}
	return result
}

func mapCoordinateDbModel(coord *Coordinate) *coordinate {
	return &coordinate{
		Id:      coord.Id,
		MT:      coord.MT,
		Axis:    coord.Axis,
		Horizon: coord.Horizon,
		X:       coord.X,
		Y:       coord.Y,
	}
}
