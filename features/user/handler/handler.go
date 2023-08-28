package handler

import (
	"net/http"
	"strconv"

	"strings"
	"yoga/clean/app/middlewares"
	user "yoga/clean/features/user"
	helpers "yoga/clean/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	response := map[string]any{
		"token":   token,
		"user_id": dataLogin.ID,
		"name":    dataLogin.Name,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success login", response))
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	result, err := handler.userService.GetAll()
	if err != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]any{
		// 	"code": 500,
		// 	"message": "hello world",
		// })
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
	var usersResponse []UserResponse
	for _, value := range result {
		usersResponse = append(usersResponse, UserResponse{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			Address:   value.Address,
			CreatedAt: value.CreatedAt,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", usersResponse))
	// return c.JSON(http.StatusOK, map[string]any{
	// 	"code":    200,
	// 	"message": "success read data",
	// 	"data":    usersResponse,
	// })

}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	userCore := RequestToCore(*userInput)
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *UserHandler) FindUserByID(c echo.Context) error {
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "ID harus berupa angka", nil))
	}

	user, err := handler.userService.GetByID(uint(idParam))
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "User tidak ditemukan", nil))
		}

		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Berhasil", userResponse))
}

func (handler *UserHandler) DeleteUserByID(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	err := handler.userService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User berhasil terhapus", nil))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	updatedUser := user.Core{
		Name:        userInput.Name,
		Email:       userInput.Email,
		Password:    userInput.Password,
		Address:     userInput.Address,
		PhoneNumber: userInput.PhoneNumber,
	}

	err := handler.userService.Update(uint(id), updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating user", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User updated successfully", nil))
}
