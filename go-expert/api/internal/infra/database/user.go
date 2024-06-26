package database

import (
	"github.com/d4niells/api/internal/entity"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type User struct {
	DB *gorm.DB
}

func NewUser(DB *gorm.DB) *User {
	return &User{DB}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
