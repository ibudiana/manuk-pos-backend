package inventory

import (
	"time"

	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
)

type Inventory struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int `gorm:"not null" json:"product_id" binding:"required"`
	BranchID  int `gorm:"not null" json:"branch_id" binding:"required"`

	Quantity         int     `gorm:"default:0" json:"quantity"`
	ReservedQuantity int     `gorm:"default:0" json:"reserved_quantity"`
	MinStockLevel    int     `gorm:"default:0" json:"min_stock_level"`
	MaxStockLevel    *int    `json:"max_stock_level,omitempty"`
	ReorderPoint     *int    `json:"reorder_point,omitempty"`
	ReorderQuantity  *int    `json:"reorder_quantity,omitempty"`
	ShelfLocation    *string `gorm:"type:varchar(100)" json:"shelf_location,omitempty"`
	LastStockUpdate  *string `gorm:"type:varchar(100)" json:"last_stock_update,omitempty"`
	LastCountingDate *string `gorm:"type:varchar(100)" json:"last_counting_date,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Product *Product      `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Branch  *store.Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
}

type InventoryTransaction struct {
	ID        int  `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int  `gorm:"not null" json:"product_id" binding:"required"`
	BranchID  int  `gorm:"not null" json:"branch_id" binding:"required"`
	UserID    *int `json:"user_id,omitempty"`

	TransactionDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"transaction_date"`
	ReferenceID     *int      `json:"reference_id,omitempty"`
	ReferenceType   *string   `gorm:"type:varchar(100)" json:"reference_type,omitempty"`
	TransactionType string    `gorm:"not null" json:"transaction_type" binding:"required"`
	Quantity        int       `gorm:"not null" json:"quantity" binding:"required"`
	UnitPrice       *float64  `gorm:"type:decimal(10,2)" json:"unit_price,omitempty"`
	Notes           *string   `gorm:"type:varchar(100)" json:"notes,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relasi dengan tabel lain
	Product *Product      `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Branch  *store.Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	User    *user.User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type InventoryTransfer struct {
	ID           int `gorm:"primaryKey;autoIncrement" json:"id"`
	FromBranchID int `gorm:"not null" json:"from_branch_id" binding:"required"`
	ToBranchID   int `gorm:"not null" json:"to_branch_id" binding:"required"`

	ReferenceNumber string    `gorm:"not null;type:varchar(100)" json:"reference_number" binding:"required"`
	InitiatedBy     int       `gorm:"not null" json:"initiated_by" binding:"required"`
	ReceivedBy      *int      `json:"received_by,omitempty"`
	TransferDate    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"transfer_date"`
	ReceivedDate    *string   `json:"received_date,omitempty"`
	Status          string    `gorm:"default:'draft';type:varchar(100)" json:"status"`
	ShippingCost    float64   `gorm:"default:0;type:decimal(10,2)" json:"shipping_cost"`
	Notes           *string   `gorm:"type:varchar(100)" json:"notes,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Initiator  user.User    `gorm:"foreignKey:InitiatedBy" json:"initiator"`
	Receiver   *user.User   `gorm:"foreignKey:ReceivedBy" json:"receiver"`
	FromBranch store.Branch `gorm:"foreignKey:FromBranchID" json:"from_branch"`
	ToBranch   store.Branch `gorm:"foreignKey:ToBranchID" json:"to_branch"`
}

type InventoryTransferItem struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	TransferID int `gorm:"not null" json:"transfer_id" binding:"required"`
	ProductID  int `gorm:"not null" json:"product_id" binding:"required"`

	QuantitySent     int       `gorm:"not null" json:"quantity_sent" binding:"required"`
	QuantityReceived *int      `json:"quantity_received,omitempty"`
	UnitCost         float64   `gorm:"not null;type:decimal(10,2)" json:"unit_cost" binding:"required"`
	Notes            *string   `gorm:"type:varchar(100)" json:"notes,omitempty"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Transfer InventoryTransfer `gorm:"foreignKey:TransferID" json:"transfer"`
	Product  Product           `gorm:"foreignKey:ProductID" json:"product"`
}
