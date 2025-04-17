package orders

import (
	"time"

	"manuk-pos-backend/models/customer"
)

// Points represents the points balance and history for a customer
type Points struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID int `gorm:"not null" json:"customer_id" binding:"required"`

	PointsBalance    int    `gorm:"default:0" json:"points_balance"`
	PointsEarned     int    `gorm:"default:0" json:"points_earned"`
	PointsUsed       int    `gorm:"default:0" json:"points_used"`
	PointsExpired    int    `gorm:"default:0" json:"points_expired"`
	LastActivityDate string `gorm:"type:varchar(100)" json:"last_activity_date"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	SyncStatus string `gorm:"type:varchar(100);default:'pending'" json:"sync_status"`

	// Relasi dengan tabel lain
	Customer customer.Customer `gorm:"foreignKey:CustomerID" json:"customer"`
}

// PointsHistory represents the change in points for a customer over time
type PointsHistory struct {
	ID            int `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID    int `gorm:"not null" json:"customer_id" binding:"required"`
	TransactionID int `gorm:"" json:"transaction_id"`

	PointsChange       int    `gorm:"not null" json:"points_change" binding:"required"`
	PointsBalanceAfter int    `gorm:"not null" json:"points_balance_after" binding:"required"`
	Description        string `gorm:"type:varchar(100)" json:"description"`
	ExpiryDate         string `gorm:"type:varchar(100)" json:"expiry_date"`
	IsExpired          int    `gorm:"default:0" json:"is_expired"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	SyncStatus string `gorm:"type:varchar(100);default:'pending'" json:"sync_status"`

	// Relasi dengan tabel lain
	Customer    customer.Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	Transaction Transaction       `gorm:"foreignKey:TransactionID" json:"transaction"`
}
