package services

import (
	"gorm.io/gorm"
	"nano/src/db/repositories"
	"nano/src/models"
)

// Services 服务管理
var (
	UserService    *userService
	DeviceService  *deviceService
	OTAService     *otaService
)

// InitServices 初始化服务
func InitServices(db *gorm.DB) {
	// 初始化仓库
	userRepo := repositories.NewUserRepository(db)
	deviceRepo := repositories.NewDeviceRepository(db)

	// 初始化服务
	UserService = NewUserService(userRepo)
	DeviceService = NewDeviceService(deviceRepo)
	OTAService = NewOTAService()
}

// userService 用户服务
type userService struct {
	repo repositories.UserRepository
}

// NewUserService 创建用户服务
func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo: repo}
}

// Register 注册用户
func (s *userService) Register(user *models.User) error {
	return s.repo.Create(user)
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

// GetUserByPhone 根据手机号获取用户
func (s *userService) GetUserByPhone(phone string) (*models.User, error) {
	return s.repo.FindByPhone(phone)
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

// ListUsers 获取用户列表
func (s *userService) ListUsers() ([]models.User, error) {
	return s.repo.List()
}

// deviceService 设备服务
type deviceService struct {
	repo repositories.DeviceRepository
}

// NewDeviceService 创建设备服务
func NewDeviceService(repo repositories.DeviceRepository) *deviceService {
	return &deviceService{repo: repo}
}

// CreateDevice 创建设备
func (s *deviceService) CreateDevice(device *models.Device) error {
	return s.repo.Create(device)
}

// GetDeviceByID 根据ID获取设备
func (s *deviceService) GetDeviceByID(id string) (*models.Device, error) {
	return s.repo.FindByID(id)
}

// GetDevicesByUserID 根据用户ID获取设备列表
func (s *deviceService) GetDevicesByUserID(userID string) ([]models.Device, error) {
	return s.repo.FindByUserID(userID)
}

// UpdateDevice 更新设备信息
func (s *deviceService) UpdateDevice(device *models.Device) error {
	return s.repo.Update(device)
}

// DeleteDevice 删除设备
func (s *deviceService) DeleteDevice(id string) error {
	return s.repo.Delete(id)
}

// ListDevices 获取设备列表
func (s *deviceService) ListDevices() ([]models.Device, error) {
	return s.repo.List()
}

// otaService OTA服务
type otaService struct{}

// NewOTAService 创建OTA服务
func NewOTAService() *otaService {
	return &otaService{}
}

// GetLatestVersion 获取最新OTA版本
func (s *otaService) GetLatestVersion() (string, error) {
	// 实现获取最新OTA版本逻辑
	return "1.1.0", nil
}

// GetDeviceVersion 获取设备当前OTA版本
func (s *otaService) GetDeviceVersion(deviceID string) (string, error) {
	// 实现获取设备当前OTA版本逻辑
	return "1.0.0", nil
}

// UpgradeDevice 升级设备OTA
func (s *otaService) UpgradeDevice(deviceID, version string) error {
	// 实现升级设备OTA逻辑
	return nil
}
