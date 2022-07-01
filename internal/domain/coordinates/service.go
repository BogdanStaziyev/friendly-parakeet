package coordinate

import (
	"fmt"
	"log"
	"math"
)

type Service interface {
	AddCoordinate(coordinate *Coordinate) (*Coordinate, error)
	UpdateCoordinate(coordinate *Coordinate) error
	DeleteCoordinate(id int64) error
	FindAll() ([]Coordinate, error)
	FindOne(id int64) (*Coordinate, error)
	InverseTask(firstId, secondId int64) (string, error, *Coordinate, *Coordinate)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) AddCoordinate(coordinate *Coordinate) (*Coordinate, error) {
	coordinates, err := (*s.repo).AddCoordinate(coordinate)
	if err != nil {
		return nil, fmt.Errorf("servis AddCoordinate: %w", err)
	}
	return coordinates, err
}

func (s *service) UpdateCoordinate(coordinate *Coordinate) error {
	err := (*s.repo).UpdateCoordinate(coordinate)
	if err != nil {
		return fmt.Errorf("servis UppdateCoordinate: %w", err)
	}
	return err
}

func (s *service) DeleteCoordinate(id int64) error {
	err := (*s.repo).DeleteCoordinate(id)
	if err != nil {
		return fmt.Errorf("servis Deletecoordinate: %w", err)
	}
	return err
}

func (s *service) FindAll() ([]Coordinate, error) {
	coordinates, err := (*s.repo).FindAll()
	if err != nil {
		return nil, fmt.Errorf("serves FindeAll: %w", err)
	}
	return coordinates, nil
}

func (s *service) FindOne(id int64) (*Coordinate, error) {
	coordinates, err := (*s.repo).FindOne(id)
	if err != nil {
		return nil, fmt.Errorf("servis FindeOne: %w", err)
	}
	return coordinates, nil
}

func (s *service) InverseTask(firstId, secondId int64) (string, error, *Coordinate, *Coordinate) {
	res, err, coordinateOne, coordinateTwo := (*s.repo).InverseTask(firstId, secondId)
	if err != nil {
		return res, fmt.Errorf("servis Invertask: %w", err), coordinateOne, coordinateTwo
	}
	n, u, m := atanNumber(coordinateOne.X, coordinateOne.Y, coordinateTwo.X, coordinateTwo.Y)
	return fmt.Sprint(res, n, "° ", u, "′ ", m, "″ "), nil, coordinateOne, coordinateTwo
}

func atanNumber(x1, y1, x2, y2 float64) (int, int, int) {
	const radius float64 = 180
	const degree, minutes, seconds int = 180, 60, 60
	x := x2 - x1
	y := y2 - y1
	num := y / x
	res := math.Atan(num)
	res *= radius / math.Pi
	deg := int(res)
	min1 := (res - float64(deg)) * 60
	min := int(min1)
	sec := int((min1 - float64(min)) * 60)
	if x < 0 && y > 0 {
		deg = (degree - 1) + deg
		min = (minutes - 1) + min
		sec = seconds + sec
	} else if x < 0 && y < 0 {
		deg = degree + deg
	} else if x > 0 && y < 0 {
		deg = ((degree * 2) - 1) + deg
		min = (minutes - 1) + min
		sec = seconds + sec
	}
	if deg > 360 || min > 60 || sec > 60 {
		log.Println("error compilation atan")
	}
	return deg, min, sec
}
