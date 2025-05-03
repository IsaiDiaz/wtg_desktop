package employee

type Service interface {
	CreateEmployee(input *Employee) error
	GetEmployeeByID(id int) (*Employee, error)
	GetAllEmployees() ([]*Employee, error)
	UpdateEmployee(input *Employee) error
	DeleteEmployee(id int) error
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

func (s *service) GetEmployeeByID(id int) (*Employee, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAllEmployees() ([]*Employee, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateEmployee(input *Employee) error {
	return s.repo.Update(input)
}

func (s *service) DeleteEmployee(id int) error {
	return s.repo.Delete(id)
}
