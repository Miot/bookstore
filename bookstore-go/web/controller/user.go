package controller

import (
	"bookstore/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func UserRegister(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "参数绑定错误",
			"error": err,
		})
		return
	}

	svc := service.NewUserService()
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

	err := svc.UserRegister(req.Username, req.Password, req.Email, req.Phone)
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

func UserLogin(ctx *gin.Context) {
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
	svc := service.NewUserService()
	loginResponse, err := svc.UserLogin(req.Username, req.Password)
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
