package data

import (
	"errors"
	"yoga/clean/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

var ErrNotFound = errors.New("project not found")

func New(db *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: db,
	}

}

func (repo *projectQuery) SelectAll() ([]project.Core, error) {
	var projectsData []Project
	tx := repo.db.Find(&projectsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var projectsCore []project.Core
	for _, projectGorm := range projectsData {
		projectCore := project.Core{
			ID:          projectGorm.ID,
			Name:        projectGorm.Name,
			UserID:      projectGorm.UserID,
			Description: projectGorm.Description,
		}
		projectsCore = append(projectsCore, projectCore)
	}

	return projectsCore, nil
}

func (repo *projectQuery) Insert(input project.Core) error {

	projectGorm := Project{
		Name:        input.Name,
		UserID:      input.UserID,
		Description: input.Description,
	}

	tx := repo.db.Create(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *projectQuery) GetByID(id uint) (project.Core, error) {
	var projectGorm Project
	tx := repo.db.First(&projectGorm, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return project.Core{}, ErrNotFound
		}
		return project.Core{}, tx.Error
	}

	projectCore := project.Core{
		ID:          projectGorm.ID,
		Name:        projectGorm.Name,
		UserID:      projectGorm.UserID,
		Description: projectGorm.Description,
	}

	return projectCore, nil
}

func (repo *projectQuery) Update(id uint, input project.Core) error {
	// Find the item by ID.
	existingProject, err := repo.GetByID(id)
	if err != nil {
		return err
	}

	existingProject.Name = input.Name
	existingProject.UserID = input.UserID
	existingProject.Description = input.Description

	tx := repo.db.Save(&existingProject)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *projectQuery) Delete(id uint) error {

	existingProject, err := repo.GetByID(id)
	if err != nil {
		return err
	}

	tx := repo.db.Delete(&existingProject)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
