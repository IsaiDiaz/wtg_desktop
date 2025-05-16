package projectemployee

type Service interface {
	CreateProjectEmployee(input *ProjectEmployee) error
	GetProjectEmployeeByID(id int) (*ProjectEmployee, error)
	GetProjectEmployeesByProjectID(projectID int) ([]*ProjectEmployee, error)
	GetProjectEmployeesByEmployeeID(employeeID int) ([]*ProjectEmployee, error)
	GetAllProjectEmployees() ([]*ProjectEmployee, error)
	UpdateProjectEmployee(input *ProjectEmployee) error
	DeleteProjectEmployee(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateProjectEmployee(input *ProjectEmployee) error {
	return s.repo.Create(input)
}

func (s *service) GetProjectEmployeeByID(id int) (*ProjectEmployee, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetProjectEmployeesByProjectID(projectID int) ([]*ProjectEmployee, error) {
	return s.repo.GetByProjectID(projectID)
}

func (s *service) GetProjectEmployeesByEmployeeID(employeeID int) ([]*ProjectEmployee, error) {
	return s.repo.GetByEmployeeID(employeeID)
}

func (s *service) GetAllProjectEmployees() ([]*ProjectEmployee, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateProjectEmployee(input *ProjectEmployee) error {
	return s.repo.Update(input)
}

func (s *service) DeleteProjectEmployee(id int) error {
	return s.repo.Delete(id)
}
