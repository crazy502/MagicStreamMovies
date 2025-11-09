package middleware

import (
	"net/http"

	"github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/utils"
	"github.com/gin-gonic/gin"
)

// 函数将被使用验证访问令牌并授予或禁止访问受保护的端点
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetAccessToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"eror": err.Error()})
			c.Abort()
			return
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Set("role", claims.Role)

		c.Next()

	}
}
