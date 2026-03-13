package db

import (
	"nano/src/db/repositories"
	"nano/src/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestUserOperations 测试用户相关操作
func TestUserOperations(t *testing.T) {
	// 初始化数据库
	db, err := NewDB(":memory:") // 使用内存数据库进行测试
	assert.NoError(t, err)

	// 自动迁移表结构
	err = AutoMigrate(db, &models.User{}, &models.Device{})
	assert.NoError(t, err)

	// 创建用户仓库
	userRepo := repositories.NewUserRepository(db)

	// 1. 插入用户
	t.Log("测试插入用户...")
	user := &models.User{
		ID:        "user1",
		Username:  "testuser",
		Phone:     "13800138000",
		Wechat:    "testwechat",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = userRepo.Create(user)
	assert.NoError(t, err)

	// 2. 查询用户
	t.Log("测试查询用户...")
	foundUser, err := userRepo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, foundUser.Username)
	assert.Equal(t, user.Phone, foundUser.Phone)

	// 3. 更新用户
	t.Log("测试更新用户...")
	foundUser.Username = "updateduser"
	foundUser.Wechat = "updatedwechat"
	err = userRepo.Update(foundUser)
	assert.NoError(t, err)

	updatedUser, err := userRepo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "updateduser", updatedUser.Username)
	assert.Equal(t, "updatedwechat", updatedUser.Wechat)

	// 4. 删除用户
	t.Log("测试删除用户...")
	err = userRepo.Delete(user.ID)
	assert.NoError(t, err)

	// 验证用户已删除
	_, err = userRepo.FindByID(user.ID)
	assert.Error(t, err)

	t.Log("用户操作测试完成")
}

// TestDeviceOperations 测试设备相关操作
func TestDeviceOperations(t *testing.T) {
	// 初始化数据库
	db, err := NewDB(":memory:") // 使用内存数据库进行测试
	assert.NoError(t, err)

	// 自动迁移表结构
	err = AutoMigrate(db, &models.User{}, &models.Device{})
	assert.NoError(t, err)

	// 创建仓库
	userRepo := repositories.NewUserRepository(db)
	deviceRepo := repositories.NewDeviceRepository(db)

	// 先创建一个用户
	user := &models.User{
		ID:        "user1",
		Username:  "testuser",
		Phone:     "13800138000",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = userRepo.Create(user)
	assert.NoError(t, err)

	// 1. 插入设备
	t.Log("测试插入设备...")
	device := &models.Device{
		ID:        "device1",
		Version:   "1.0.0",
		Status:    "active",
		UserID:    user.ID,
		Secret:    "device_secret_123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = deviceRepo.Create(device)
	assert.NoError(t, err)

	// 2. 查询设备
	t.Log("测试查询设备...")
	foundDevice, err := deviceRepo.FindByID(device.ID)
	assert.NoError(t, err)
	assert.Equal(t, device.Version, foundDevice.Version)
	assert.Equal(t, device.Status, foundDevice.Status)
	assert.Equal(t, device.UserID, foundDevice.UserID)

	// 3. 更新设备
	t.Log("测试更新设备...")
	foundDevice.Version = "1.1.0"
	foundDevice.Status = "inactive"
	err = deviceRepo.Update(foundDevice)
	assert.NoError(t, err)

	updatedDevice, err := deviceRepo.FindByID(device.ID)
	assert.NoError(t, err)
	assert.Equal(t, "1.1.0", updatedDevice.Version)
	assert.Equal(t, "inactive", updatedDevice.Status)

	// 4. 测试绑定设备到用户（通过设置UserID）
	t.Log("测试绑定设备到用户...")
	// 创建第二个用户
	user2 := &models.User{
		ID:        "user2",
		Username:  "testuser2",
		Phone:     "13900139000",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = userRepo.Create(user2)
	assert.NoError(t, err)

	// 绑定设备到新用户
	updatedDevice.UserID = user2.ID
	err = deviceRepo.Update(updatedDevice)
	assert.NoError(t, err)

	// 验证设备已绑定到新用户
	boundDevice, err := deviceRepo.FindByID(device.ID)
	assert.NoError(t, err)
	assert.Equal(t, user2.ID, boundDevice.UserID)

	// 5. 删除设备
	t.Log("测试删除设备...")
	err = deviceRepo.Delete(device.ID)
	assert.NoError(t, err)

	// 验证设备已删除
	_, err = deviceRepo.FindByID(device.ID)
	assert.Error(t, err)

	t.Log("设备操作测试完成")
}
