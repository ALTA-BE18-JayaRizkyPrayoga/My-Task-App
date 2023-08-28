package handler

import (
	"net/http"
	"strconv"
	"strings"
	project "yoga/clean/features/project"
	helpers "yoga/clean/helper"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ProjectServiceInterface
}

func New(service project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

func (handler *ProjectHandler) CreateProject(c echo.Context) error {
	projectInput := new(project.Core)
	errBind := c.Bind(&projectInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	err := handler.projectService.Create(*projectInput)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success insert data", nil))
}

func (handler *ProjectHandler) GetAllProject(c echo.Context) error {
	result, err := handler.projectService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	var projectsResponse []project.Core
	for _, value := range result {
		projectsResponse = append(projectsResponse, project.Core{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
		})
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", projectsResponse))
}

func (handler *ProjectHandler) GetProjectByID(c echo.Context) error {
	id := c.Param("project_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "ID harus berupa angka", nil))
	}

	project, err := handler.projectService.GetByID(uint(idParam))
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "project tidak ditemukan", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	projectResponse := ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		UserID:      project.UserID,
		Description: project.Description,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Berhasil", projectResponse))
}

func (handler *ProjectHandler) UpdateProjectByID(c echo.Context) error {
	id := c.Param("project_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "ID harus berupa angka", nil))
	}

	projectInput := new(project.Core)
	errBind := c.Bind(&projectInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	err := handler.projectService.Update(uint(idParam), *projectInput)
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "project tidak ditemukan", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "project berhasil diperbarui", nil))
}

func (handler *ProjectHandler) DeleteProjectByID(c echo.Context) error {
	id := c.Param("project_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "ID harus berupa angka", nil))
	}

	err := handler.projectService.Delete(uint(idParam))
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "project tidak ditemukan", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "project berhasil terhapus", nil))
}
