package permission

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"size:255;not null"`
	Resources   string `gorm:"size:100;not null"`
	Action      string `gorm:"size:50;not null"`
}
