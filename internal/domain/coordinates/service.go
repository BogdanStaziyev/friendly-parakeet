package coordinate

import "fmt"

type Service interface {
	FindAll() ([]Coordinate, error)
	FindOne(id int64) (*Coordinate, error)
	AddCoordinate(coordinate *Coordinate) (*Coordinate, error)
	UpdateCoordinate(coordinate *Coordinate) error
	InverseTask(coordinate1, coordinate2 *Coordinate) error
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Coordinate, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Coordinate, error) {
	return (*s.repo).FindOne(id)
}
func (s *service) AddCoordinate(coordinate *Coordinate) (*Coordinate, error) {
	coordinate, err := (*s.repo).AddCoordinate(coordinate)
	if err != nil {
		return nil, fmt.Errorf("servis AddCoordinate: %w", err)
	}
	return coordinate, err
}
func (s *service) UpdateCoordinate(coordinate *Coordinate) error {
	return (*s.repo).UpdateCoordinate(coordinate)
}
func (s *service) InverseTask(coordinate1, coordinate2 *Coordinate) error {
	return (*s.repo).InverseTask(coordinate1, coordinate2)
}
