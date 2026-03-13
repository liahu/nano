package models

import (
	"time"
)

// Device 设备模型
type Device struct {
	ID        string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Version   string    `gorm:"type:varchar(20);not null" json:"version"`
	Status    string    `gorm:"type:varchar(20);not null;default:'inactive'" json:"status"`
	UserID    string    `gorm:"type:varchar(36);index" json:"user_id"`
	Secret    string    `gorm:"type:varchar(100);not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Device) TableName() string {
	return "devices"
}
