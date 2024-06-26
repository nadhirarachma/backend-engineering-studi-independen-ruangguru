package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	
	u.db.Create(&user)
	
	return nil 
}

func (u *userRepository) CheckAvail(user model.User) error {

	var result model.User
	response := u.db.Model(&model.User{}).Where("username = ?", user.Username).Where("password = ?", user.Password).Take(&result)

	return response.Error
}
