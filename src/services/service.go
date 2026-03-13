package services

import (
	"context"
	"nano/src/db/repositories"
	"nano/src/models"
	"nano/src/utils"

	"gorm.io/gorm"
)

// Services 服务管理
var (
	UserService   *userService
	DeviceService *deviceService
	OTAService    *otaService
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
func (s *userService) Register(ctx context.Context, user *models.User) error {
	utils.Info(ctx, "开始注册用户", map[string]interface{}{"user_id": user.ID, "phone": user.Phone})

	err := s.repo.Create(user)
	if err != nil {
		utils.Error(ctx, "注册用户失败", map[string]interface{}{"user_id": user.ID, "error": err.Error()})
		return err
	}

	utils.Info(ctx, "注册用户成功", map[string]interface{}{"user_id": user.ID, "phone": user.Phone})
	return nil
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	utils.Info(ctx, "开始根据ID获取用户", map[string]interface{}{"user_id": id})

	user, err := s.repo.FindByID(id)
	if err != nil {
		utils.Error(ctx, "根据ID获取用户失败", map[string]interface{}{"user_id": id, "error": err.Error()})
		return nil, err
	}

	utils.Info(ctx, "根据ID获取用户成功", map[string]interface{}{"user_id": id})
	return user, nil
}

// GetUserByPhone 根据手机号获取用户
func (s *userService) GetUserByPhone(ctx context.Context, phone string) (*models.User, error) {
	utils.Info(ctx, "开始根据手机号获取用户", map[string]interface{}{"phone": phone})

	user, err := s.repo.FindByPhone(phone)
	if err != nil {
		utils.Error(ctx, "根据手机号获取用户失败", map[string]interface{}{"phone": phone, "error": err.Error()})
		return nil, err
	}

	utils.Info(ctx, "根据手机号获取用户成功", map[string]interface{}{"phone": phone, "user_id": user.ID})
	return user, nil
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	utils.Info(ctx, "开始更新用户信息", map[string]interface{}{"user_id": user.ID})

	err := s.repo.Update(user)
	if err != nil {
		utils.Error(ctx, "更新用户信息失败", map[string]interface{}{"user_id": user.ID, "error": err.Error()})
		return err
	}

	utils.Info(ctx, "更新用户信息成功", map[string]interface{}{"user_id": user.ID})
	return nil
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	utils.Info(ctx, "开始删除用户", map[string]interface{}{"user_id": id})

	err := s.repo.Delete(id)
	if err != nil {
		utils.Error(ctx, "删除用户失败", map[string]interface{}{"user_id": id, "error": err.Error()})
		return err
	}

	utils.Info(ctx, "删除用户成功", map[string]interface{}{"user_id": id})
	return nil
}

// ListUsers 获取用户列表
func (s *userService) ListUsers(ctx context.Context) ([]models.User, error) {
	utils.Info(ctx, "开始获取用户列表")

	users, err := s.repo.List()
	if err != nil {
		utils.Error(ctx, "获取用户列表失败", map[string]interface{}{"error": err.Error()})
		return nil, err
	}

	utils.Info(ctx, "获取用户列表成功", map[string]interface{}{"count": len(users)})
	return users, nil
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
func (s *deviceService) CreateDevice(ctx context.Context, device *models.Device) error {
	utils.Info(ctx, "开始创建设备", map[string]interface{}{"device_id": device.ID, "user_id": device.UserID})

	err := s.repo.Create(device)
	if err != nil {
		utils.Error(ctx, "创建设备失败", map[string]interface{}{"device_id": device.ID, "error": err.Error()})
		return err
	}

	utils.Info(ctx, "创建设备成功", map[string]interface{}{"device_id": device.ID, "user_id": device.UserID})
	return nil
}

// GetDeviceByID 根据ID获取设备
func (s *deviceService) GetDeviceByID(ctx context.Context, id string) (*models.Device, error) {
	utils.Info(ctx, "开始根据ID获取设备", map[string]interface{}{"device_id": id})

	device, err := s.repo.FindByID(id)
	if err != nil {
		utils.Error(ctx, "根据ID获取设备失败", map[string]interface{}{"device_id": id, "error": err.Error()})
		return nil, err
	}

	utils.Info(ctx, "根据ID获取设备成功", map[string]interface{}{"device_id": id, "user_id": device.UserID})
	return device, nil
}

// GetDevicesByUserID 根据用户ID获取设备列表
func (s *deviceService) GetDevicesByUserID(ctx context.Context, userID string) ([]models.Device, error) {
	utils.Info(ctx, "开始根据用户ID获取设备列表", map[string]interface{}{"user_id": userID})

	devices, err := s.repo.FindByUserID(userID)
	if err != nil {
		utils.Error(ctx, "根据用户ID获取设备列表失败", map[string]interface{}{"user_id": userID, "error": err.Error()})
		return nil, err
	}

	utils.Info(ctx, "根据用户ID获取设备列表成功", map[string]interface{}{"user_id": userID, "count": len(devices)})
	return devices, nil
}

// UpdateDevice 更新设备信息
func (s *deviceService) UpdateDevice(ctx context.Context, device *models.Device) error {
	utils.Info(ctx, "开始更新设备信息", map[string]interface{}{"device_id": device.ID, "user_id": device.UserID})

	err := s.repo.Update(device)
	if err != nil {
		utils.Error(ctx, "更新设备信息失败", map[string]interface{}{"device_id": device.ID, "error": err.Error()})
		return err
	}

	utils.Info(ctx, "更新设备信息成功", map[string]interface{}{"device_id": device.ID, "user_id": device.UserID})
	return nil
}

// DeleteDevice 删除设备
func (s *deviceService) DeleteDevice(ctx context.Context, id string) error {
	utils.Info(ctx, "开始删除设备", map[string]interface{}{"device_id": id})

	err := s.repo.Delete(id)
	if err != nil {
		utils.Error(ctx, "删除设备失败", map[string]interface{}{"device_id": id, "error": err.Error()})
		return err
	}

	utils.Info(ctx, "删除设备成功", map[string]interface{}{"device_id": id})
	return nil
}

// ListDevices 获取设备列表
func (s *deviceService) ListDevices(ctx context.Context) ([]models.Device, error) {
	utils.Info(ctx, "开始获取设备列表")

	devices, err := s.repo.List()
	if err != nil {
		utils.Error(ctx, "获取设备列表失败", map[string]interface{}{"error": err.Error()})
		return nil, err
	}

	utils.Info(ctx, "获取设备列表成功", map[string]interface{}{"count": len(devices)})
	return devices, nil
}

// otaService OTA服务
type otaService struct{}

// NewOTAService 创建OTA服务
func NewOTAService() *otaService {
	return &otaService{}
}

// GetLatestVersion 获取最新OTA版本
func (s *otaService) GetLatestVersion(ctx context.Context) (string, error) {
	utils.Info(ctx, "开始获取最新OTA版本")

	// 实现获取最新OTA版本逻辑
	version := "1.1.0"
	utils.Info(ctx, "获取最新OTA版本成功", map[string]interface{}{"version": version})
	return version, nil
}

// GetDeviceVersion 获取设备当前OTA版本
func (s *otaService) GetDeviceVersion(ctx context.Context, deviceID string) (string, error) {
	utils.Info(ctx, "开始获取设备当前OTA版本", map[string]interface{}{"device_id": deviceID})

	// 实现获取设备当前OTA版本逻辑
	version := "1.0.0"
	utils.Info(ctx, "获取设备当前OTA版本成功", map[string]interface{}{"device_id": deviceID, "version": version})
	return version, nil
}

// UpgradeDevice 升级设备OTA
func (s *otaService) UpgradeDevice(ctx context.Context, deviceID, version string) error {
	utils.Info(ctx, "开始升级设备OTA", map[string]interface{}{"device_id": deviceID, "version": version})

	// 实现升级设备OTA逻辑
	utils.Info(ctx, "升级设备OTA成功", map[string]interface{}{"device_id": deviceID, "version": version})
	return nil
}
