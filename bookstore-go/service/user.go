package service

import (
	"bookstore/jwt"
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

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	UserInfo     *UserInfo `json:"user_info"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (u *UserService) UserLogin(username, password string) (*LoginResponse, error) {
	user, err := u.UserDB.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if user.Password != u.encodePassword(password) {
		return nil, errors.New("密码错误")
	}

	token, err := jwt.GenerateTokenPair(uint(user.ID), user.Username)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	response := &LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    int64(jwt.AccessTokenExpire.Seconds()),
		UserInfo: &UserInfo{
			ID:       int(user.ID),
			Username: user.Username,
			Phone:    user.Phone,
			Email:    user.Email,
		},
	}
	return response, nil
}

func (u *UserService) GetUserByID(id int) (*model.User, error) {
	user, err := u.UserDB.GetUserByID(id)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return user, nil
}

func (u *UserService) UpdateUserInfo(user *model.User) error {
	oldUser, err := u.UserDB.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	oldUser.Username = user.Username
	oldUser.Email = user.Email
	oldUser.Phone = user.Phone
	oldUser.Avatar = user.Avatar

	// 更新数据库
	if err := u.UserDB.UpdateUser(oldUser); err != nil {
		return err
	}
	return nil
}
