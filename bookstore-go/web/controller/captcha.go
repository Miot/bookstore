package controller

import (
	"bookstore/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateCaptcha(ctx *gin.Context) {
	// 生成图片验证码
	svc := service.NewCaptchaService()
	res, err := svc.GenerateCaptcha()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "生成验证码失败",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "生成验证码成功",
		"data": res,
	})
}
