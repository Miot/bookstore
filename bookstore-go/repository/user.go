package repository

import (
	"bookstore/global"
	"bookstore/model"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		db: global.GetDB(),
	}
}

func (u *UserDAO) CreateUser(user *model.User) error {
	return u.db.Debug().Create(user).Error
}

func (u *UserDAO) CheckUserExisits(username, phone, email string) (bool, error) {
	var count int64

	err := u.db.Model(&model.User{}).Where("username = ? OR phone = ? OR email = ?", username, phone, email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
