package controller

import (
	"bookstore/model"
	"bookstore/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	CaptchaID       string `json:"captcha_id"`
	CaptchaValue    string `json:"captcha_value"`
}

type LoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaID    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}

func (u *UserController) UserRegister(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "参数绑定错误",
			"error": err,
		})
		return
	}

	captchaSvc := service.NewCaptchaService()
	if !captchaSvc.VerifyCaptcha(req.CaptchaID, req.CaptchaValue) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}

	// 校验密码
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "两次密码不一致",
		})
		return
	}

	err := u.UserService.UserRegister(req.Username, req.Password, req.Phone, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "注册成功",
	})
}

func (u *UserController) UserLogin(ctx *gin.Context) {
	// 验证图片验证码
	var req LoginRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "请求参数错误",
			"error": err.Error(),
		})
		return
	}
	captchaSvc := service.NewCaptchaService()
	if !captchaSvc.VerifyCaptcha(req.CaptchaID, req.CaptchaValue) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}
	// 校验用户信息（是否有这个用户，校验密码）,返回JWT信息给用户
	loginResponse, err := u.UserService.UserLogin(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": loginResponse,
	})
}

func (u *UserController) GetUserprofile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}
	user, err := u.UserService.GetUserByID(userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	response := gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"phone":      user.Phone,
		"avatar":     user.Avatar,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取用户信息成功",
		"data": response,
	})
}

func (u *UserController) UpdateUserprofile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	var updateData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "请求参数错误",
			"error": err.Error(),
		})
		return
	}

	user := &model.User{
		ID:       userID.(int),
		Username: updateData.Username,
		Email:    updateData.Email,
		Phone:    updateData.Phone,
		Avatar:   updateData.Avatar,
	}

	if err := u.UserService.UpdateUserInfo(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "更新用户信息失败",
			"error": err.Error(),
		})
		return
	}

	// 获取更新后的用户信息
	updatedUser, err := u.UserService.GetUserByID(userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取更新后的用户信息失败",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "更新用户信息成功",
		"data": updatedUser,
	})
}
