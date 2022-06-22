package coordinate

type Service interface {
	FindAll() ([]Coordinate, error)
	FindOne(id int64) (*Coordinate, error)
	AddCoordinate(coordinate Coordinate) error
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
func (s *service) AddCoordinate(coordinate Coordinate) error {
	return (*s.repo).AddCoordinate(coordinate)
}
func (s *service) UpdateCoordinate(coordinate *Coordinate) error {
	return (*s.repo).UpdateCoordinate(coordinate)
}
func (s *service) InverseTask(coordinate1, coordinate2 *Coordinate) error {
	return (*s.repo).InverseTask(coordinate1, coordinate2)
}
