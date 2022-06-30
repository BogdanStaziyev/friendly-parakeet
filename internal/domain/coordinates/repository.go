package coordinate

import (
	"fmt"
	"github.com/upper/db/v4"
	"log"
	"math"
)

type coordinate struct {
	Id      int64   `db:"id,omitempty"`
	MT      int64   `db:"mt"`
	Axis    string  `db:"axis"`
	Horizon string  `db:"horizon"`
	X       float64 `db:"x"`
	Y       float64 `db:"y"`
}

type Repository interface {
	AddCoordinate(coordinate *Coordinate) (*Coordinate, error)
	UpdateCoordinate(coordinate *Coordinate) error
	DeleteCoordinate(id int64) error
	FindAll() ([]Coordinate, error)
	FindOne(id int64) (*Coordinate, error)
	InverseTask(firstId, secondId int64) (string, error)
}

type repository struct {
	coll db.Collection
}

func NewRepository(dbSession *db.Session) Repository {
	return &repository{
		coll: (*dbSession).Collection("coordinate"),
	}
}

func (r *repository) AddCoordinate(coordinate *Coordinate) (*Coordinate, error) {
	coordinates := mapCoordinateDbModel(coordinate)

	err := r.coll.InsertReturning(coordinates)
	if err != nil {
		return nil, fmt.Errorf("Coordinaterepository Create: %w", err)
	}

	return mapCoordinateDbModelToDomain(coordinates), nil
}

func (r *repository) UpdateCoordinate(coordinate *Coordinate) error {
	coordinates := mapCoordinateDbModel(coordinate)

	err := r.coll.Find(coordinates.Id).Update(coordinates)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (r *repository) DeleteCoordinate(id int64) error {

	err := r.coll.Find("id", id).Delete()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
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

func (r *repository) InverseTask(firstId, secondId int64) (string, error) {
	var coordinateOne, coordinateTwo coordinate

	firstErr := r.coll.Find("id", firstId).One(&coordinateOne)
	if firstErr != nil {
		log.Fatal("repository Invert first: ", firstErr)
	}
	secondErr := r.coll.Find("id", secondId).One(&coordinateTwo)
	if secondErr != nil {
		log.Fatal("repository Invert second: ", secondErr)
	}
	n, u, m := atanNumber(coordinateOne.X, coordinateOne.Y, coordinateTwo.X, coordinateTwo.Y)
	return fmt.Sprint("result: ", n, "° ", u, "′ ", m, "″ "), nil
}

func atanNumber(x1, y1, x2, y2 float64) (int, int, int) {
	x := x2 - x1
	y := y2 - y1
	num := y / x
	res := math.Atan(num)
	res *= 180 / math.Pi
	deg := int(res)
	min1 := (res - float64(deg)) * 60
	min := int(min1)
	sec1 := int((min1 - float64(min)) * 60)
	sec := int(sec1)
	if x < 0 && y > 0 {
		deg = 179 + deg
		min = 59 + min
		sec = 60 + sec
	} else if x < 0 && y < 0 {
		deg = 180 + deg
	} else if x > 0 && y < 0 {
		deg = 359 + deg
		min = 59 + min
		sec = 60 + sec
	}
	fmt.Println(deg, min, sec)
	return deg, min, sec
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
