package coordinate

import (
	"fmt"
	"github.com/upper/db/v4"
	"log"
)

type coordinate struct {
	Id      int64   `db:"id, omitempty"`
	MT      int64   `db:"mt"`
	Axis    string  `db:"axis"`
	Horizon string  `db:"horizon"`
	X       float64 `db:"x"`
	Y       float64 `db:"y"`
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
		coll: (*dbSession).Collection("coordinate"),
	}
}

func (r *repository) FindAll() ([]Coordinate, error) {
	var coordinates []coordinate

	err := r.coll.Find().All(&coordinates)
	if err != nil {
		log.Fatal("coordinate.Find: ", err)
	}
	return mapCoordinateDbModelToDomainCollection(coordinates), nil
}

func (r *repository) FindOne(id int64) (*Coordinate, error) {
	var coordinates coordinate

	err := r.coll.Find("id", id).One(&coordinates)
	if err != nil {
		log.Fatal("coordinateCol.Find: ", err)
	}
	return mapCoordinateDbModelToDomain(&coordinates), nil
}

func (r *repository) AddCoordinate(coordinate *Coordinate) (*Coordinate, error) {
	coordinetes := mapCoordinateDbModel(coordinate)
	err := r.coll.InsertReturning(coordinetes)
	if err != nil {
		return nil, fmt.Errorf("Coordinaterepository Create: %w", err)
	}

	return mapCoordinateDbModelToDomain(coordinetes), nil
}

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
