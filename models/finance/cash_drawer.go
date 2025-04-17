package finance

import (
	"time"

	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
)

// CashDrawer represents the cash drawer entity
type CashDrawer struct {
	ID       int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   int `gorm:"not null" json:"user_id" binding:"required"`
	BranchID int `gorm:"not null" json:"branch_id" binding:"required"`

	OpeningTime    string  `gorm:"type:varchar(100)" json:"opening_time"`
	ClosingTime    string  `gorm:"type:varchar(100)" json:"closing_time"`
	OpeningAmount  float64 `gorm:"type:decimal(10,2);not null" json:"opening_amount" binding:"required"`
	ClosingAmount  float64 `gorm:"type:decimal(10,2)" json:"closing_amount"`
	ExpectedAmount float64 `gorm:"type:decimal(10,2)" json:"expected_amount"`
	PhysicalAmount float64 `gorm:"type:decimal(10,2)" json:"physical_amount"`
	Difference     float64 `gorm:"type:decimal(10,2)" json:"difference"`
	Status         string  `gorm:"type:varchar(100);default:'open'" json:"status"`
	Notes          string  `gorm:"type:varchar(100)" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	SyncStatus string `gorm:"type:varchar(100);default:'pending'" json:"sync_status"`

	// Relasi dengan tabel lain
	User   *user.User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Branch *store.Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
}

// CashDrawerTransaction represents a transaction for a cash drawer
type CashDrawerTransaction struct {
	ID           int `gorm:"primaryKey;autoIncrement" json:"id"`
	CashDrawerID int `gorm:"not null" json:"cash_drawer_id" binding:"required"`
	UserID       int `gorm:"not null" json:"user_id" binding:"required"`

	TransactionType string  `gorm:"type:varchar(100);not null" json:"transaction_type" binding:"required"`
	Amount          float64 `gorm:"type:decimal(10,2);not null" json:"amount" binding:"required"`
	ReferenceID     int     `gorm:"" json:"reference_id"`
	Notes           string  `gorm:"type:varchar(100)" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	SyncStatus string `gorm:"type:varchar(100);default:'pending'" json:"sync_status"`

	// Relasi dengan tabel lain
	CashDrawer *CashDrawer `gorm:"foreignKey:CashDrawerID" json:"cash_drawer,omitempty"`
	User       *user.User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
