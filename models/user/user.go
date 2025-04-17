package user

import (
	"time"

	"manuk-pos-backend/models/store"
)

type User struct {
	ID       int `gorm:"primaryKey" json:"id"`
	RoleID   int `gorm:"default:1;column:role_id" json:"role_id"`     //Admin
	BranchID int `gorm:"default:1;column:branch_id" json:"branch_id"` //Default Branch

	Username   string     `gorm:"type:varchar(50);unique;not null" json:"username"`
	Password   string     `gorm:"type:varchar(255);not null" json:"-"`
	Name       string     `gorm:"type:varchar(100);not null" json:"name"`
	Email      string     `gorm:"type:varchar(100);unique" json:"email"`
	Phone      string     `gorm:"type:varchar(25)" json:"phone"`
	IsActive   bool       `gorm:"type:tinyint(1);default:1" json:"is_active"`
	LastLogin  *time.Time `json:"last_login,omitempty"`
	LoginCount int        `gorm:"default:0" json:"login_count"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Role   *Role         `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	Branch *store.Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserLogin struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}
