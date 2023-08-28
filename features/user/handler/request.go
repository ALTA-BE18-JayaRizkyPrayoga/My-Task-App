package handler

import "yoga/clean/features/user"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
	}
}
