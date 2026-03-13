package handlers

import (
	"nano/src/models"
	"nano/src/services"
	"nano/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "service is running",
	})
}

// Register 账号注册
func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user := &models.User{
		ID:       utils.GenerateUUID(),
		Username: req.Username,
		Phone:    req.Phone,
		Password: req.Password,
		Wechat:   req.Wechat,
	}

	if err := services.UserService.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "注册失败",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.RegisterResponse{
		Token:  token,
		UserID: user.ID,
	})
}

// Login 账号登录
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := services.UserService.GetUserByPhone(req.Phone)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "用户不存在或密码错误",
		})
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "用户不存在或密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token:  token,
		UserID: user.ID,
	})
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	// 从context中获取用户ID（由AuthMiddleware设置）
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "用户未认证",
		})
		return
	}

	// 获取用户信息
	user, err := services.UserService.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "获取用户信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	// 实现更新用户信息逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// BindDevice 绑定设备
func BindDevice(c *gin.Context) {
	// 实现绑定设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// UnbindDevice 解绑设备
func UnbindDevice(c *gin.Context) {
	// 实现解绑设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// ListDevices 获取设备列表
func ListDevices(c *gin.Context) {
	// 实现获取设备列表逻辑
	c.JSON(http.StatusOK, gin.H{
		"devices": []gin.H{
			{
				"id":      "device1",
				"version": "1.0.0",
				"status":  "active",
			},
		},
	})
}

// GetDeviceInfo 获取设备信息
func GetDeviceInfo(c *gin.Context) {
	// 实现获取设备信息逻辑
	c.JSON(http.StatusOK, gin.H{
		"id":      "device1",
		"version": "1.0.0",
		"status":  "active",
		"user_id": "user1",
	})
}

// ControlDevice 控制设备
func ControlDevice(c *gin.Context) {
	// 实现控制设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// UpdateDeviceInfo 更新设备信息
func UpdateDeviceInfo(c *gin.Context) {
	// 实现更新设备信息逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// GetLatestOTA 获取最新OTA版本
func GetLatestOTA(c *gin.Context) {
	// 实现获取最新OTA版本逻辑
	c.JSON(http.StatusOK, gin.H{
		"version": "1.1.0",
		"url":     "/api/v1/ota/download/1.1.0",
	})
}

// GetDeviceOTA 获取设备OTA版本
func GetDeviceOTA(c *gin.Context) {
	// 实现获取设备OTA版本逻辑
	c.JSON(http.StatusOK, gin.H{
		"current_version": "1.0.0",
		"latest_version":  "1.1.0",
	})
}

// UpgradeDeviceOTA 升级设备OTA
func UpgradeDeviceOTA(c *gin.Context) {
	// 实现升级设备OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "upgrade request sent",
	})
}

// DeviceRegister 设备注册
func DeviceRegister(c *gin.Context) {
	// 实现设备注册逻辑
	c.JSON(http.StatusOK, gin.H{
		"device_id": "device1",
		"secret":    "device_secret_123",
	})
}

// ReportDeviceStatus 上报设备状态
func ReportDeviceStatus(c *gin.Context) {
	// 实现上报设备状态逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// CheckOTA 检查OTA版本
func CheckOTA(c *gin.Context) {
	// 实现检查OTA版本逻辑
	c.JSON(http.StatusOK, gin.H{
		"version": "1.1.0",
		"url":     "/api/v1/ota/download/1.1.0",
	})
}

// DownloadOTA 下载OTA文件
func DownloadOTA(c *gin.Context) {
	// 实现下载OTA文件逻辑
	c.String(http.StatusOK, "OTA file content")
}

// ReportOTAStatus 上报OTA状态
func ReportOTAStatus(c *gin.Context) {
	// 实现上报OTA状态逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// ListUsers 列出用户
func ListUsers(c *gin.Context) {
	// 实现列出用户逻辑
	c.JSON(http.StatusOK, gin.H{
		"users": []gin.H{
			{
				"id":       "user1",
				"username": "testuser",
				"phone":    "13800138000",
			},
		},
	})
}

// ManageUser 管理用户
func ManageUser(c *gin.Context) {
	// 实现管理用户逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// ListAllDevices 列出所有设备
func ListAllDevices(c *gin.Context) {
	// 实现列出所有设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"devices": []gin.H{
			{
				"id":      "device1",
				"version": "1.0.0",
				"status":  "active",
				"user_id": "user1",
			},
		},
	})
}

// ManageDevice 管理设备
func ManageDevice(c *gin.Context) {
	// 实现管理设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// UploadOTA 上传OTA
func UploadOTA(c *gin.Context) {
	// 实现上传OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// ManageOTA 管理OTA
func ManageOTA(c *gin.Context) {
	// 实现管理OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// ListOTA 列出OTA
func ListOTA(c *gin.Context) {
	// 实现列出OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"ota": []gin.H{
			{
				"version": "1.1.0",
				"url":     "/api/v1/ota/download/1.1.0",
			},
		},
	})
}
