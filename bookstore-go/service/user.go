package service

import (
	"bookstore/model"
	"bookstore/repository"
	"encoding/base64"
	"errors"
)

type UserService struct {
	UserDB *repository.UserDAO
}

// service --> repository --> db
func NewUserService() *UserService {
	return &UserService{
		UserDB: repository.NewUserDAO(),
	}
}

func (u *UserService) UserRegister(username, password, phone, email string) error {
	// 校验
	exists, err := u.UserDB.CheckUserExisits(username, phone, email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("用户名、邮箱或手机号已存在")
	}

	// 密码加密
	encodePassword := u.encodePassword(password)

	if err = u.CreateUser(username, encodePassword, phone, email); err != nil {
		return err
	}

	return nil
}

func (u *UserService) encodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func (u *UserService) CreateUser(username, password, phone, email string) error {
	user := &model.User{
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
	}
	return u.UserDB.CreateUser(user)
}
