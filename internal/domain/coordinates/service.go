package coordinate

import "fmt"

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
	return res, nil, coordinateOne, coordinateTwo
}
