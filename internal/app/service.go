package app

import (
	"fmt"
	"math"
	"startUp/internal/domain"
	"startUp/internal/infra/database"
)

type Service interface {
	AddCoordinate(coordinate *domain.Coordinate) (*domain.Coordinate, error)
	UpdateCoordinate(coordinate *domain.Coordinate) error
	DeleteCoordinate(id int64) error
	FindAll() ([]domain.Coordinate, error)
	FindOne(id int64) (*domain.Coordinate, error)
	InverseTask(firstId, secondId int64) (string, error, *domain.Coordinate, *domain.Coordinate)
}

type service struct {
	repo *database.Repository
}

func NewService(r *database.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) AddCoordinate(coordinate *domain.Coordinate) (*domain.Coordinate, error) {
	coordinates, err := (*s.repo).AddCoordinate(coordinate)
	if err != nil {
		return nil, fmt.Errorf("servis AddCoordinate: %w", err)
	}
	return coordinates, err
}

func (s *service) UpdateCoordinate(coordinate *domain.Coordinate) error {
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

func (s *service) FindAll() ([]domain.Coordinate, error) {
	coordinates, err := (*s.repo).FindAll()
	if err != nil {
		return nil, fmt.Errorf("serves FindeAll: %w", err)
	}
	return coordinates, nil
}

func (s *service) FindOne(id int64) (*domain.Coordinate, error) {
	coordinates, err := (*s.repo).FindOne(id)
	if err != nil {
		return nil, fmt.Errorf("service FindOne: %w", err)
	}
	return coordinates, nil
}

func (s *service) InverseTask(firstId, secondId int64) (string, error, *domain.Coordinate, *domain.Coordinate) {
	var n, u, m int
	res, err, coordinateOne, coordinateTwo := (*s.repo).InverseTask(firstId, secondId)
	if err != nil {
		return res, fmt.Errorf("servis Invertask: %w", err), coordinateOne, coordinateTwo
	}
	if coordinateOne.X == coordinateTwo.X || coordinateOne.Y == coordinateTwo.Y {
		return "Error Service same values used: ", nil, coordinateOne, coordinateTwo
	}
	n, u, m = atanNumber(coordinateOne.X, coordinateOne.Y, coordinateTwo.X, coordinateTwo.Y)
	return fmt.Sprint(res, n, "° ", u, "′ ", m, "″ "), nil, coordinateOne, coordinateTwo
}

func atanNumber(x1, y1, x2, y2 float64) (int, int, int) {
	const radius float64 = 180
	const degree, minutes, seconds int = 180, 60, 60
	var deg, min, sec int
	x := x2 - x1
	y := y2 - y1
	subtractionCoordinate := y / x
	atanResult := math.Atan(subtractionCoordinate)
	atanResult *= radius / math.Pi
	deg = int(atanResult)
	minute := (atanResult - float64(deg)) * 60
	min = int(minute)
	sec = int((minute - float64(min)) * 60)
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
	return deg, min, sec
}
