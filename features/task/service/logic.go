package service

import (
	"errors"
	"yoga/clean/features/task"
)

type taskService struct {
	taskData task.TaskDataInterface
}

func New(repo task.TaskDataInterface) task.TaskServiceInterface {
	return &taskService{
		taskData: repo,
	}
}

func (service *taskService) Create(input task.Core) error {
	if input.Status == "" {
		return errors.New("validation error. name/description required")
	}
	err := service.taskData.Insert(input)
	return err
}

func (service *taskService) GetAll() ([]task.Core, error) {
	result, err := service.taskData.SelectAll()
	return result, err
}

func (service *taskService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validation error: invalid ID")
	}
	return nil
}

func (service *taskService) GetByID(id uint) (task.Core, error) {
	if id == 0 {
		return task.Core{}, errors.New("validation error: invalid ID")
	}

	result, err := service.taskData.GetByID(id)
	return result, err
}

func (service *taskService) Update(id uint, input task.Core) error {
	if id == 0 {
		return errors.New("validation error: invalid ID")
	}

	if input.Status == "" {
		return errors.New("validation error")
	}

	return nil
}
