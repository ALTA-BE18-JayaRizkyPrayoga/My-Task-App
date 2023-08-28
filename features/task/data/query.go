package data

import (
	"errors"
	"yoga/clean/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

var ErrNotFound = errors.New("task not found")

func New(db *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: db,
	}

}

func (repo *taskQuery) SelectAll() ([]task.Core, error) {
	var tasksData []Task
	tx := repo.db.Find(&tasksData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var tasksCore []task.Core
	for _, taskGorm := range tasksData {
		taskCore := task.Core{
			ID:        taskGorm.ID,
			Status:    taskGorm.Status,
			ProjectID: taskGorm.ProjectID,
		}
		tasksCore = append(tasksCore, taskCore)
	}

	return tasksCore, nil
}

func (repo *taskQuery) Insert(input task.Core) error {

	taskGorm := Task{
		Status:    input.Status,
		ProjectID: input.ProjectID,
	}

	tx := repo.db.Create(&taskGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *taskQuery) GetByID(id uint) (task.Core, error) {
	var taskGorm Task
	tx := repo.db.First(&taskGorm, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return task.Core{}, ErrNotFound
		}
		return task.Core{}, tx.Error
	}

	taskCore := task.Core{
		ID:        taskGorm.ID,
		Status:    taskGorm.Status,
		ProjectID: taskGorm.ProjectID,
	}

	return taskCore, nil
}

func (repo *taskQuery) Update(id uint, input task.Core) error {

	existingTask, err := repo.GetByID(id)
	if err != nil {
		return err
	}

	existingTask.Status = input.Status
	existingTask.ProjectID = input.ProjectID

	tx := repo.db.Save(&existingTask)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *taskQuery) Delete(id uint) error {

	existingTask, err := repo.GetByID(id)
	if err != nil {
		return err
	}

	tx := repo.db.Delete(&existingTask)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
