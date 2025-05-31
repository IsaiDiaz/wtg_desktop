package auth

import "gorm.io/gorm"

type Repository interface {
	CreateUser(user *AuthUser) error
	GetUserByID(id int) (*AuthUser, error)
	GetAllUsers() ([]*AuthUser, error)
	UpdateUser(user *AuthUser) error
	DeleteUser(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(user *AuthUser) error {
	return r.db.Create(user).Error
}

func (r *repository) GetUserByID(id int) (*AuthUser, error) {
	var user AuthUser
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetAllUsers() ([]*AuthUser, error) {
	var users []*AuthUser
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) UpdateUser(user *AuthUser) error {
	return r.db.Save(user).Error
}

func (r *repository) DeleteUser(id int) error {
	var user AuthUser
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	return r.db.Delete(&user).Error
}
