package data

import (
	_projectData "yoga/clean/features/project/data"

	"gorm.io/gorm"
)

// struct item gorm model
type Task struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Status    string
	ProjectID uint
	Project   _projectData.Project
}
