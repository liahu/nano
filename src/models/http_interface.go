package models

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Wechat   string `json:"wechat"`
}

// RegisterResponse 注册响应
type RegisterResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

// UpdateUserInfoRequest 更新用户信息请求
type UpdateUserInfoRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Wechat   string `json:"wechat"`
}

// BindDeviceRequest 绑定设备请求
type BindDeviceRequest struct {
	DeviceID string `json:"device_id" binding:"required"`
}

// UnbindDeviceRequest 解绑设备请求
type UnbindDeviceRequest struct {
	DeviceID string `json:"device_id" binding:"required"`
}

// ControlDeviceRequest 控制设备请求
type ControlDeviceRequest struct {
	Command string                 `json:"command" binding:"required"`
	Params  map[string]interface{} `json:"params"`
}

// UpdateDeviceInfoRequest 更新设备信息请求
type UpdateDeviceInfoRequest struct {
	Name string `json:"name"`
}

// UpgradeDeviceOTARequest 升级设备OTA请求
type UpgradeDeviceOTARequest struct {
	Version string `json:"version" binding:"required"`
}

// DeviceRegisterRequest 设备注册请求
type DeviceRegisterRequest struct {
	HardwareID string `json:"hardware_id" binding:"required"`
	Version    string `json:"version" binding:"required"`
}

// DeviceRegisterResponse 设备注册响应
type DeviceRegisterResponse struct {
	DeviceID string `json:"device_id"`
	Secret   string `json:"secret"`
}

// ReportDeviceStatusRequest 上报设备状态请求
type ReportDeviceStatusRequest struct {
	Status  string  `json:"status" binding:"required"`
	Battery float64 `json:"battery"`
	Version string  `json:"version"`
}

// ReportOTAStatusRequest 上报OTA状态请求
type ReportOTAStatusRequest struct {
	Status   string  `json:"status" binding:"required"`
	Progress float64 `json:"progress"`
	Error    string  `json:"error"`
}

// ManageUserRequest 管理用户请求
type ManageUserRequest struct {
	Action string `json:"action" binding:"required"`
	User   User   `json:"user"`
}

// ManageDeviceRequest 管理设备请求
type ManageDeviceRequest struct {
	Action string `json:"action" binding:"required"`
	Device Device `json:"device"`
}

// ManageOTARequest 管理OTA请求
type ManageOTARequest struct {
	Action  string `json:"action" binding:"required"`
	Version string `json:"version" binding:"required"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse 成功响应
type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// HealthCheckResponse 健康检查响应
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// DeviceListResponse 设备列表响应
type DeviceListResponse struct {
	Devices []Device `json:"devices"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Users []User `json:"users"`
}

// OTAListResponse OTA列表响应
type OTAListResponse struct {
	OTA []OTAVersion `json:"ota"`
}

// OTAVersion OTA版本
type OTAVersion struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

// LatestOTAResponse 最新OTA响应
type LatestOTAResponse struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

// DeviceOTAResponse 设备OTA响应
type DeviceOTAResponse struct {
	CurrentVersion string `json:"current_version"`
	LatestVersion  string `json:"latest_version"`
}
