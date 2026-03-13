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
	utils.Info(c, "健康检查请求")
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "service is running",
	})
	utils.Info(c, "健康检查响应成功")
}

// Register 账号注册
func Register(c *gin.Context) {

	utils.Info(c, "开始处理注册请求")

	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "解析注册请求失败", map[string]interface{}{"error": err.Error()})
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

	if err := services.UserService.Register(c, user); err != nil {
		utils.Error(c, "注册用户失败", map[string]interface{}{"phone": req.Phone, "error": err.Error()})
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "注册失败",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.Error(c, "生成token失败", map[string]interface{}{"user_id": user.ID, "error": err.Error()})
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.RegisterResponse{
		Token:  token,
		UserID: user.ID,
	})
	utils.Info(c, "注册请求处理成功", map[string]interface{}{"user_id": user.ID, "phone": req.Phone})
}

// Login 账号登录
func Login(c *gin.Context) {

	utils.Info(c, "开始处理登录请求")

	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "解析登录请求失败", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := services.UserService.GetUserByPhone(c, req.Phone)
	if err != nil {
		utils.Error(c, "用户登录失败", map[string]interface{}{"phone": req.Phone, "error": err.Error()})
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "用户不存在或密码错误",
		})
		return
	}

	if user.Password != req.Password {
		utils.Error(c, "密码错误", map[string]interface{}{"phone": req.Phone, "user_id": user.ID})
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "用户不存在或密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.Error(c, "生成token失败", map[string]interface{}{"user_id": user.ID, "error": err.Error()})
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token:  token,
		UserID: user.ID,
	})
	utils.Info(c, "登录请求处理成功", map[string]interface{}{"user_id": user.ID, "phone": req.Phone})
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {

	utils.Info(c, "开始处理获取用户信息请求")

	// 从context中获取用户ID（由AuthMiddleware设置）
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, "用户未认证")
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "用户未认证",
		})
		return
	}

	// 获取用户信息
	user, err := services.UserService.GetUserByID(c, userID.(string))
	if err != nil {
		utils.Error(c, "获取用户信息失败", map[string]interface{}{"user_id": userID.(string), "error": err.Error()})
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "获取用户信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, user)
	utils.Info(c, "获取用户信息成功", map[string]interface{}{"user_id": userID.(string)})
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(c *gin.Context) {

	utils.Info(c, "开始处理更新用户信息请求")
	// 实现更新用户信息逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "更新用户信息请求处理成功")
}

// BindDevice 绑定设备
func BindDevice(c *gin.Context) {

	utils.Info(c, "开始处理绑定设备请求")
	// 实现绑定设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "绑定设备请求处理成功")
}

// UnbindDevice 解绑设备
func UnbindDevice(c *gin.Context) {

	utils.Info(c, "开始处理解绑设备请求")
	// 实现解绑设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "解绑设备请求处理成功")
}

// ListDevices 获取设备列表
func ListDevices(c *gin.Context) {

	utils.Info(c, "开始处理获取设备列表请求")
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
	utils.Info(c, "获取设备列表请求处理成功")
}

// GetDeviceInfo 获取设备信息
func GetDeviceInfo(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理获取设备信息请求", map[string]interface{}{"device_id": deviceID})
	// 实现获取设备信息逻辑
	c.JSON(http.StatusOK, gin.H{
		"id":      "device1",
		"version": "1.0.0",
		"status":  "active",
		"user_id": "user1",
	})
	utils.Info(c, "获取设备信息请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// ControlDevice 控制设备
func ControlDevice(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理控制设备请求", map[string]interface{}{"device_id": deviceID})
	// 实现控制设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "控制设备请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// UpdateDeviceInfo 更新设备信息
func UpdateDeviceInfo(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理更新设备信息请求", map[string]interface{}{"device_id": deviceID})
	// 实现更新设备信息逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "更新设备信息请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// GetLatestOTA 获取最新OTA版本
func GetLatestOTA(c *gin.Context) {

	utils.Info(c, "开始处理获取最新OTA版本请求")
	// 实现获取最新OTA版本逻辑
	c.JSON(http.StatusOK, gin.H{
		"version": "1.1.0",
		"url":     "/api/v1/ota/download/1.1.0",
	})
	utils.Info(c, "获取最新OTA版本请求处理成功")
}

// GetDeviceOTA 获取设备OTA版本
func GetDeviceOTA(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理获取设备OTA版本请求", map[string]interface{}{"device_id": deviceID})
	// 实现获取设备OTA版本逻辑
	c.JSON(http.StatusOK, gin.H{
		"current_version": "1.0.0",
		"latest_version":  "1.1.0",
	})
	utils.Info(c, "获取设备OTA版本请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// UpgradeDeviceOTA 升级设备OTA
func UpgradeDeviceOTA(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理升级设备OTA请求", map[string]interface{}{"device_id": deviceID})
	// 实现升级设备OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "upgrade request sent",
	})
	utils.Info(c, "升级设备OTA请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// DeviceRegister 设备注册
func DeviceRegister(c *gin.Context) {

	utils.Info(c, "开始处理设备注册请求")
	// 实现设备注册逻辑
	c.JSON(http.StatusOK, gin.H{
		"device_id": "device1",
		"secret":    "device_secret_123",
	})
	utils.Info(c, "设备注册请求处理成功", map[string]interface{}{"device_id": "device1"})
}

// ReportDeviceStatus 上报设备状态
func ReportDeviceStatus(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理上报设备状态请求", map[string]interface{}{"device_id": deviceID})
	// 实现上报设备状态逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "上报设备状态请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// CheckOTA 检查OTA版本
func CheckOTA(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理检查OTA版本请求", map[string]interface{}{"device_id": deviceID})
	// 实现检查OTA版本逻辑
	c.JSON(http.StatusOK, gin.H{
		"version": "1.1.0",
		"url":     "/api/v1/ota/download/1.1.0",
	})
	utils.Info(c, "检查OTA版本请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// DownloadOTA 下载OTA文件
func DownloadOTA(c *gin.Context) {

	version := c.Param("version")
	utils.Info(c, "开始处理下载OTA文件请求", map[string]interface{}{"version": version})
	// 实现下载OTA文件逻辑
	c.String(http.StatusOK, "OTA file content")
	utils.Info(c, "下载OTA文件请求处理成功", map[string]interface{}{"version": version})
}

// ReportOTAStatus 上报OTA状态
func ReportOTAStatus(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理上报OTA状态请求", map[string]interface{}{"device_id": deviceID})
	// 实现上报OTA状态逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "上报OTA状态请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// ListUsers 列出用户
func ListUsers(c *gin.Context) {

	utils.Info(c, "开始处理列出用户请求")
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
	utils.Info(c, "列出用户请求处理成功")
}

// ManageUser 管理用户
func ManageUser(c *gin.Context) {

	userID := c.Param("id")
	utils.Info(c, "开始处理管理用户请求", map[string]interface{}{"user_id": userID})
	// 实现管理用户逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "管理用户请求处理成功", map[string]interface{}{"user_id": userID})
}

// ListAllDevices 列出所有设备
func ListAllDevices(c *gin.Context) {

	utils.Info(c, "开始处理列出所有设备请求")
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
	utils.Info(c, "列出所有设备请求处理成功")
}

// ManageDevice 管理设备
func ManageDevice(c *gin.Context) {

	deviceID := c.Param("id")
	utils.Info(c, "开始处理管理设备请求", map[string]interface{}{"device_id": deviceID})
	// 实现管理设备逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "管理设备请求处理成功", map[string]interface{}{"device_id": deviceID})
}

// UploadOTA 上传OTA
func UploadOTA(c *gin.Context) {

	utils.Info(c, "开始处理上传OTA请求")
	// 实现上传OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "上传OTA请求处理成功")
}

// ManageOTA 管理OTA
func ManageOTA(c *gin.Context) {

	version := c.Param("version")
	utils.Info(c, "开始处理管理OTA请求", map[string]interface{}{"version": version})
	// 实现管理OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	utils.Info(c, "管理OTA请求处理成功", map[string]interface{}{"version": version})
}

// ListOTA 列出OTA
func ListOTA(c *gin.Context) {

	utils.Info(c, "开始处理列出OTA请求")
	// 实现列出OTA逻辑
	c.JSON(http.StatusOK, gin.H{
		"ota": []gin.H{
			{
				"version": "1.1.0",
				"url":     "/api/v1/ota/download/1.1.0",
			},
		},
	})
	utils.Info(c, "列出OTA请求处理成功")
}
