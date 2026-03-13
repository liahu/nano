package middlewares

import (
	"nano/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少认证token",
			})
			c.Abort()
			return
		}

		userID, err := utils.ValidateJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token无效或已过期",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

// DeviceAuthMiddleware 设备认证中间件
func DeviceAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceID := c.GetHeader("Device-ID")
		signature := c.GetHeader("Signature")
		if deviceID == "" || signature == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少设备认证信息",
			})
			c.Abort()
			return
		}

		// 验证signature逻辑（待实现）
		c.Set("device_id", deviceID)
		c.Next()
	}
}

// AdminAuthMiddleware 管理员认证中间件
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少认证token",
			})
			c.Abort()
			return
		}

		userID, err := utils.ValidateJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token无效或已过期",
			})
			c.Abort()
			return
		}

		// 这里可以添加管理员权限检查
		c.Set("user_id", userID)
		c.Next()
	}
}
