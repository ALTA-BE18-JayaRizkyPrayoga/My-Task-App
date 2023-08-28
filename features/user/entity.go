package user

import "time"

type Core struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserDataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
	Delete(id uint) error
	GetByID(id uint) (Core, error)
	Update(id uint, updatedUser Core) error
	Login(email string, password string) (dataLogin Core, err error)
}

type UserServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core) error
	GetByID(id uint) (Core, error)
	Delete(id uint) error
	Update(id uint, updatedUser Core) error
	Login(email string, password string) (dataLogin Core, token string, err error)
}
