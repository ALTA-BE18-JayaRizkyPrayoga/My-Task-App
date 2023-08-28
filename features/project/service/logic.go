package service

import (
	"errors"
	"yoga/clean/features/project"
)

type projectService struct {
	projectData project.ProjectDataInterface
}

func New(repo project.ProjectDataInterface) project.ProjectServiceInterface {
	return &projectService{
		projectData: repo,
	}
}

func (service *projectService) Create(input project.Core) error {
	if input.Name == "" || input.Description == "" {
		return errors.New("validation error. name/description required")
	}
	err := service.projectData.Insert(input)
	return err
}

func (service *projectService) GetAll() ([]project.Core, error) {
	result, err := service.projectData.SelectAll()
	return result, err
}

func (service *projectService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validation error: invalid ID")
	}
	return nil
}

func (service *projectService) GetByID(id uint) (project.Core, error) {
	if id == 0 {
		return project.Core{}, errors.New("validation error: invalid ID")
	}

	result, err := service.projectData.GetByID(id)
	return result, err
}

func (service *projectService) Update(id uint, input project.Core) error {
	if id == 0 {
		return errors.New("validation error: invalid ID")
	}

	if input.Name == "" || input.Description == "" {
		return errors.New("validation error: Name and Description are required")
	}

	return nil
}
