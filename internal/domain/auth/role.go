package auth

type AuthRole struct {
	ID   int    `gorm:"column:auth_role_id;primaryKey;autoIncrement"`
	Name string `gorm:"column:auth_role_name;size:30;not null"`

	RolePermissions []AuthRolePermission `gorm:"foreignKey:RoleID"`
	RoleUsers       []AuthUserRole       `gorm:"foreignKey:RoleID"`
}

func (AuthRole) TableName() string {
	return "auth_role"
}
