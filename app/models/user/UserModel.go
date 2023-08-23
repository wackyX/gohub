// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"size:255;index:idx_name" json:"name,omitempty"`
	Email    string `gorm:"size:255;" json:"-"`
	Phone    string `gorm:"size:255;" json:"-"`
	Password string `gorm:"size:255;" json:"-"`

	models.CommonTimestampsField
}
