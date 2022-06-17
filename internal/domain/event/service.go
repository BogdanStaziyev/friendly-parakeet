package coordinate

type Service interface {
	FindAll() ([]Coordinate, error)
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
