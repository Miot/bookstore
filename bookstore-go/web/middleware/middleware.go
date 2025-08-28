package middleware

import (
	"bookstore/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 验证jwt token中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "未提供认证令牌",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "认证令牌格式错误",
			})
			c.Abort()
			return
		}

		// 提取token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析JWT token
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "认证令牌无效",
			})
			c.Abort()
			return
		}

		// 检查token类型
		if claims.TokenType != "access" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "token类型错误，请使用access token",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", int(claims.UserID))
		c.Set("username", claims.Username)
		c.Next()
	}
}
