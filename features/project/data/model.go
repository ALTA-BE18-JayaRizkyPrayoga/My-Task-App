package data

import (
	_userData "yoga/clean/features/user/data"

	"gorm.io/gorm"
)

// struct item gorm model
type Project struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	UserID      uint
	Description string
	User        _userData.User
}
