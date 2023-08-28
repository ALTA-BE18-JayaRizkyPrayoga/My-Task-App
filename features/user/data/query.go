package data

import (
	"errors"
	"fmt"

	"yoga/clean/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func (repo *userQuery) Login(email string, password string) (dataLogin user.Core, err error) {
	var data User
	tx := repo.db.Where("email = ? and password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}
	dataLogin = ModelToCore(data)
	return dataLogin, nil
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// mapping dari struct core to struct gorm model
	userGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&userGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements user.UserDataInterface.
func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersData []User
	tx := repo.db.Find(&usersData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println("users:", usersData)
	//mapping dari struct gorm model ke struct core (entity)
	var usersCore []user.Core
	for _, value := range usersData {
		var user = user.Core{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Password:    value.Password,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		usersCore = append(usersCore, user)
	}
	return usersCore, nil
}

func (repo *userQuery) GetByID(id uint) (user.Core, error) {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {

			return user.Core{}, fmt.Errorf("User with ID %d not found", id)
		}
		return user.Core{}, tx.Error
	}

	userCore := user.Core{
		ID:          userGorm.ID,
		Name:        userGorm.Name,
		Email:       userGorm.Email,
		Password:    userGorm.Password,
		Address:     userGorm.Address,
		PhoneNumber: userGorm.PhoneNumber,
		CreatedAt:   userGorm.CreatedAt,
		UpdatedAt:   userGorm.UpdatedAt,
	}

	return userCore, nil
}

func (repo *userQuery) Delete(id uint) error {
	var userGorm User
	tx := repo.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *userQuery) Update(id uint, updatedUser user.Core) error {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("User with ID %d not found", id)
		}
		return tx.Error
	}

	userGorm.Name = updatedUser.Name
	userGorm.Email = updatedUser.Email
	userGorm.Password = updatedUser.Password
	userGorm.Address = updatedUser.Address
	userGorm.PhoneNumber = updatedUser.PhoneNumber

	tx = repo.db.Save(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
