package data

import (
	"yoga/clean/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	Address     string
	PhoneNumber string
	// Items       []Item
}

// // struct item gorm model
//
//	type Item struct {
//		gorm.Model
//		// ID          uint `gorm:"primaryKey"`
//		// CreatedAt   time.Time
//		// UpdatedAt   time.Time
//		// DeletedAt   gorm.DeletedAt `gorm:"index"`
//		Name        string
//		UserID      uint
//		Brand       string
//		Description string
//		Price       int
//		Weight      int
//	}
func CoreToModel(dataCore user.Core) User {
	return User{
		Name:        dataCore.Name,
		Email:       dataCore.Email,
		Password:    dataCore.Password,
		Address:     dataCore.Address,
		PhoneNumber: dataCore.PhoneNumber,
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel User) user.Core {
	return user.Core{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		Address:     dataModel.Address,
		PhoneNumber: dataModel.PhoneNumber,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

// mapping []model to []core
func ListModelToCore(dataModel []User) []user.Core {
	var result []user.Core
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
