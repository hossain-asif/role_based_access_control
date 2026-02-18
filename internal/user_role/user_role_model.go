package userrole

import (
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID uint
	RoleID uint
}
