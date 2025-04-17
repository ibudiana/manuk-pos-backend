package inventory

import (
	"time"

	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
)

type StockOpname struct {
	ID       int `gorm:"primaryKey;autoIncrement" json:"id"`
	BranchID int `gorm:"not null" json:"branch_id" binding:"required"`
	UserID   int `gorm:"not null" json:"user_id" binding:"required"`

	OpnameDate      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"opname_date"`
	ReferenceNumber *string   `gorm:"type:varchar(100)" json:"reference_number,omitempty"`
	Notes           *string   `gorm:"type:varchar(100)" json:"notes,omitempty"`
	Status          string    `gorm:"default:'draft';type:varchar(100)" json:"status"`
	CompletedAt     *string   `gorm:"type:varchar(100)" json:"completed_at,omitempty"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Branch *store.Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	User   *user.User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type StockOpnameItem struct {
	ID            int `gorm:"primaryKey;autoIncrement" json:"id"`
	StockOpnameID int `gorm:"not null" json:"stock_opname_id" binding:"required"`
	ProductID     int `gorm:"not null" json:"product_id" binding:"required"`

	SystemStock     float64   `gorm:"not null;type:decimal(10,2)" json:"system_stock" binding:"required"`
	PhysicalStock   float64   `gorm:"not null;type:decimal(10,2)" json:"physical_stock" binding:"required"`
	Difference      float64   `gorm:"type:decimal(10,2)" json:"difference,omitempty"`
	AdjustmentValue float64   `gorm:"type:decimal(10,2)" json:"adjustment_value,omitempty"`
	Notes           *string   `gorm:"type:varchar(100)" json:"notes,omitempty"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	StockOpname *StockOpname `gorm:"foreignKey:StockOpnameID" json:"stock_opname,omitempty"`
	Product     *Product     `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
