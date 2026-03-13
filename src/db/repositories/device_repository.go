package repositories

import (
	"nano/src/models"

	"gorm.io/gorm"
)

// DeviceRepository 设备数据访问接口
type DeviceRepository interface {
	Create(device *models.Device) error
	FindByID(id string) (*models.Device, error)
	FindByUserID(userID string) ([]models.Device, error)
	Update(device *models.Device) error
	Delete(id string) error
	List() ([]models.Device, error)
}

// deviceRepository 设备数据访问实现
type deviceRepository struct {
	db *gorm.DB
}

// NewDeviceRepository 创建设备数据访问实例
func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{db: db}
}

// Create 创建设备
func (r *deviceRepository) Create(device *models.Device) error {
	return r.db.Create(device).Error
}

// FindByID 根据ID查找设备
func (r *deviceRepository) FindByID(id string) (*models.Device, error) {
	var device models.Device
	err := r.db.First(&device, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

// FindByUserID 根据用户ID查找设备列表
func (r *deviceRepository) FindByUserID(userID string) ([]models.Device, error) {
	var devices []models.Device
	err := r.db.Where("user_id = ?", userID).Find(&devices).Error
	return devices, err
}

// Update 更新设备
func (r *deviceRepository) Update(device *models.Device) error {
	return r.db.Save(device).Error
}

// Delete 删除设备
func (r *deviceRepository) Delete(id string) error {
	return r.db.Delete(&models.Device{}, "id = ?", id).Error
}

// List 获取设备列表
func (r *deviceRepository) List() ([]models.Device, error) {
	var devices []models.Device
	err := r.db.Find(&devices).Error
	return devices, err
}
