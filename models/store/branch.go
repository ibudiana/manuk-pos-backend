package store

import (
	"time"
)

type Branch struct {
	ID int `gorm:"primaryKey" json:"id"`

	Code         string `gorm:"type:varchar(50);unique;not null" json:"code" binding:"required"`
	Name         string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Address      string `gorm:"type:varchar(255)" json:"address"`
	Phone        string `gorm:"type:varchar(25)" json:"phone"`
	Email        string `gorm:"type:varchar(100)" json:"email"`
	IsMainBranch bool   `gorm:"type:tinyint(1);default:0" json:"is_main_branch"`
	IsActive     bool   `gorm:"type:tinyint(1);default:1" json:"is_active"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
