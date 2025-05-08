package project

type Service interface {
	CreateProject(input *Project) error
	GetProjectByID(id uint) (*Project, error)
	GetAllProjects() ([]*Project, error)
	UpdateProject(input *Project) error
	DeleteProject(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}
func (s *service) CreateProject(input *Project) error {
	return s.repo.Create(input)
}

func (s *service) GetProjectByID(id uint) (*Project, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetAllProjects() ([]*Project, error) {
	return s.repo.FindAll()
}

func (s *service) UpdateProject(input *Project) error {
	return s.repo.Update(input)
}

func (s *service) DeleteProject(id uint) error {
	return s.repo.Delete(id)
}
