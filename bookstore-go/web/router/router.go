package router

import (
	"bookstore/web/controller"
	"bookstore/web/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 跨域
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	userController := controller.NewUserController()
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", userController.UserRegister)
			user.POST("/login", userController.UserLogin)
		}
		auth := user.Group("")
		{
			auth.Use(middleware.JWTAuthMiddleware())
			{
				auth.GET("/profile", userController.GetUserprofile)
				auth.PUT("/profile", userController.UpdateUserprofile)
				auth.PUT("/password", userController.ChangePassword)
				auth.DELETE("/logout", userController.Logout)
			}
		}
	}

	captController := controller.NewCaptChaController()
	captcha := v1.Group("/captcha")
	{
		captcha.GET("/generate", captController.GenerateCaptcha)
	}
	return r
}
