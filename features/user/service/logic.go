package service

import (
	"errors"
	"yoga/clean/app/middlewares"
	"yoga/clean/features/user"
)

type userService struct {
	userData user.UserDataInterface
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

func (service *userService) Login(email string, password string) (dataLogin user.Core, token string, err error) {
	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return user.Core{}, "", err
	}
	token, err = middlewares.CreateToken(int(dataLogin.ID))
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("validation error. name/email/password required")
	}
	err := service.userData.Insert(input)
	return err
}

// GetAll implements user.UserServiceInterface.
func (service *userService) GetAll() ([]user.Core, error) {
	result, err := service.userData.SelectAll()
	return result, err
}

func (service *userService) GetByID(id uint) (user.Core, error) {
	if id == 0 {
		return user.Core{}, errors.New("ID tidak valid")
	}
	result, err := service.userData.GetByID(id)
	return result, err
}

func (service *userService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validation error. invalid ID")
	}
	err := service.userData.Delete(id)
	return err
}

func (service *userService) Update(id uint, updatedUser user.Core) error {
	if id == 0 {
		return errors.New("validation error. invalid ID")
	}
	err := service.userData.Update(id, updatedUser)
	return err
}
