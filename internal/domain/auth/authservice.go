package auth

type Service interface {
	CreateUser(user *AuthUser) error
	GetUserByID(id int) (*AuthUser, error)
	GetAllUsers() ([]*AuthUser, error)
	UpdateUser(user *AuthUser) error
	DeleteUser(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateUser(user *AuthUser) error {
	return s.repo.CreateUser(user)
}

func (s *service) GetUserByID(id int) (*AuthUser, error) {
	return s.repo.GetUserByID(id)
}

func (s *service) GetAllUsers() ([]*AuthUser, error) {
	return s.repo.GetAllUsers()
}

func (s *service) UpdateUser(user *AuthUser) error {
	return s.repo.UpdateUser(user)
}

func (s *service) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
