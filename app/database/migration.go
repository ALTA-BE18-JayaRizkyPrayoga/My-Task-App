package database

import (
	_projectData "yoga/clean/features/project/data"
	_taskData "yoga/clean/features/task/data"
	_userData "yoga/clean/features/user/data"

	"gorm.io/gorm"
)

// db migration
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_projectData.Project{})
	db.AutoMigrate(&_taskData.Task{})
}
