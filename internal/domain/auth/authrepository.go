package auth

import "gorm.io/gorm"

type Repository interface {
	GetUserByEmail(email string) (*AuthUser, error)
	GetUserRoles(userID int) ([]AuthRole, error)
	AssignRoleToUser(userID, roleID int) error
	GetUserByUsernameOrEmail(identifier string) (*AuthUser, error)

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

func (r *repository) GetUserByEmail(email string) (*AuthUser, error) {
	var user AuthUser
	err := r.db.Preload("UserRoles.Role").Where("auth_user_email = ?", email).First(&user).Error
	return &user, err
}

func (r *repository) GetUserRoles(userID int) ([]AuthRole, error) {
	var roles []AuthRole
	err := r.db.
		Joins("JOIN auth_user_role ON auth_user_role.auth_role_id = auth_role.auth_role_id").
		Where("auth_user_role.auth_user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

func (r *repository) AssignRoleToUser(userID, roleID int) error {
	userRole := AuthUserRole{
		UserID: userID,
		RoleID: roleID,
	}
	return r.db.Create(&userRole).Error
}

func (r *repository) GetUserByUsernameOrEmail(identifier string) (*AuthUser, error) {
	var user AuthUser
	err := r.db.
		Where("auth_user_email = ? OR auth_user_name = ?", identifier, identifier).
		Preload("UserRoles.Role").
		First(&user).Error
	return &user, err
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
