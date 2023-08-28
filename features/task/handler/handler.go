package handler

import (
	"net/http"
	"strconv"
	"strings"

	task "yoga/clean/features/task"
	helpers "yoga/clean/helper"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.TaskServiceInterface
}

func New(service task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}

func (handler *TaskHandler) CreateTask(c echo.Context) error {
	taskInput := new(task.Core)
	errBind := c.Bind(&taskInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	err := handler.taskService.Create(*taskInput)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success insert data", nil))
}

func (handler *TaskHandler) GetAllTask(c echo.Context) error {
	result, err := handler.taskService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	var tasksResponse []task.Core
	for _, value := range result {
		tasksResponse = append(tasksResponse, task.Core{
			ID:        value.ID,
			Status:    value.Status,
			ProjectID: value.ProjectID,
		})
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", tasksResponse))
}

func (handler *TaskHandler) GetTaskByID(c echo.Context) error {
	id := c.Param("task_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "ID harus berupa angka", nil))
	}

	task, err := handler.taskService.GetByID(uint(idParam))
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "task tidak ditemukan", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	taskResponse := TaskResponse{
		ID:        task.ID,
		Status:    task.Status,
		ProjectID: task.ProjectID,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Berhasil", taskResponse))
}

func (handler *TaskHandler) UpdateTaskByID(c echo.Context) error {
	id := c.Param("task_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "ID harus berupa angka", nil))
	}

	taskInput := new(task.Core)
	errBind := c.Bind(&taskInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	err := handler.taskService.Update(uint(idParam), *taskInput)
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "task tidak ditemukan", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "task berhasil diperbarui", nil))
}

func (handler *TaskHandler) DeleteTaskByID(c echo.Context) error {
	id := c.Param("task_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "ID harus berupa angka", nil))
	}

	err := handler.taskService.Delete(uint(idParam))
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "task tidak ditemukan", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "task berhasil terhapus", nil))
}
