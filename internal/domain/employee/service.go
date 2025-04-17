package employee

type Service interface {
	CreateEmployee(input *Employee) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateEmployee(input *Employee) error {
	return s.repo.Create(input)
}
