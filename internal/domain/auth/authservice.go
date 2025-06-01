package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(user *AuthUser, plainPassword string) error
	CreateUserWithRole(user *AuthUser, plainPassword string, roleID int) error
	Authenticate(identifier, password string) (*AuthUser, []string, error)
	AssignRole(userID, roleID int) error

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

func (s *service) CreateUser(user *AuthUser, plainPassword string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Secret = string(hashed)
	return s.repo.CreateUser(user)
}

func (s *service) CreateUserWithRole(user *AuthUser, plainPassword string, roleID int) error {
	if err := s.CreateUser(user, plainPassword); err != nil {
		return err
	}
	return s.AssignRole(user.ID, roleID)
}

func (s *service) Authenticate(identifier, password string) (*AuthUser, []string, error) {
	user, err := s.repo.GetUserByUsernameOrEmail(identifier)
	if err != nil {
		return nil, nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Secret), []byte(password))
	if err != nil {
		return nil, nil, errors.New("invalid password")
	}

	roles := []string{}
	for _, ur := range user.UserRoles {
		roles = append(roles, ur.Role.Name)
	}

	return user, roles, nil
}

func (s *service) AssignRole(userID, roleID int) error {
	return s.repo.AssignRoleToUser(userID, roleID)
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
