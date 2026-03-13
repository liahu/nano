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
			utils.Error(c, "缺少认证token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少认证token",
			})
			c.Abort()
			return
		}

		userID, err := utils.ValidateJWTToken(token)
		if err != nil {
			utils.Error(c, "token无效或已过期", map[string]interface{}{"error": err.Error()})
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token无效或已过期",
			})
			c.Abort()
			return
		}

		utils.Info(c, "用户认证成功", map[string]interface{}{"user_id": userID})
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
			utils.Error(c, "缺少设备认证信息")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少设备认证信息",
			})
			c.Abort()
			return
		}

		// 验证signature逻辑（待实现）
		utils.Info(c, "设备认证成功", map[string]interface{}{"device_id": deviceID})
		c.Set("device_id", deviceID)
		c.Next()
	}
}

// AdminAuthMiddleware 管理员认证中间件
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			utils.Error(c, "缺少认证token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少认证token",
			})
			c.Abort()
			return
		}

		userID, err := utils.ValidateJWTToken(token)
		if err != nil {
			utils.Error(c, "token无效或已过期", map[string]interface{}{"error": err.Error()})
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token无效或已过期",
			})
			c.Abort()
			return
		}

		// 这里可以添加管理员权限检查
		utils.Info(c, "管理员认证成功", map[string]interface{}{"user_id": userID})
		c.Set("user_id", userID)
		c.Next()
	}
}
