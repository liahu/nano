package models

import (
	"time"
)

// User 用户模型
type User struct {
	ID        string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Username  string    `gorm:"type:varchar(50);not null" json:"username"`
	Phone     string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
	Wechat    string    `gorm:"type:varchar(50)" json:"wechat"`
	Password  string    `gorm:"type:varchar(100);not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
