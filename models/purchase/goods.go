package purchase

import (
	"time"

	"manuk-pos-backend/models/inventory"
	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
	"manuk-pos-backend/models/vendor"
)

type GoodsReceiving struct {
	ID         int  `gorm:"primaryKey;autoIncrement" json:"id"`
	POID       *int `gorm:"type:int" json:"po_id,omitempty"`
	SupplierID int  `gorm:"not null" json:"supplier_id" binding:"required"`
	BranchID   int  `gorm:"not null" json:"branch_id" binding:"required"`
	UserID     int  `gorm:"not null" json:"user_id" binding:"required"`

	ReferenceNumber string    `gorm:"unique;not null;type:varchar(100)" json:"reference_number" binding:"required"`
	ReceivingDate   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"receiving_date"`
	Subtotal        float64   `gorm:"type:decimal(10,2);not null" json:"subtotal" binding:"required"`
	TaxAmount       float64   `gorm:"type:decimal(10,2);default:0" json:"tax_amount"`
	DiscountAmount  float64   `gorm:"type:decimal(10,2);default:0" json:"discount_amount"`
	ShippingCost    float64   `gorm:"type:decimal(10,2);default:0" json:"shipping_cost"`
	OtherCosts      float64   `gorm:"type:decimal(10,2);default:0" json:"other_costs"`
	GrandTotal      float64   `gorm:"type:decimal(10,2);not null" json:"grand_total" binding:"required"`
	Notes           *string   `gorm:"type:varchar(100)" json:"notes,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	PO       *PurchaseOrder   `gorm:"foreignKey:POID" json:"po,omitempty"`
	Supplier *vendor.Supplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	Branch   *store.Branch    `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	User     *user.User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type GoodsReceivingItem struct {
	ID        int  `gorm:"primaryKey;autoIncrement" json:"id"`
	GRID      int  `gorm:"not null" json:"gr_id" binding:"required"`
	POItemID  *int `gorm:"type:int" json:"po_item_id,omitempty"`
	ProductID int  `gorm:"not null" json:"product_id" binding:"required"`

	Quantity        int     `gorm:"not null" json:"quantity" binding:"required"`
	UnitPrice       float64 `gorm:"type:decimal(10,2);not null" json:"unit_price" binding:"required"`
	DiscountPercent float64 `gorm:"type:decimal(10,2);default:0" json:"discount_percent"`
	TaxPercent      float64 `gorm:"type:decimal(10,2);default:0" json:"tax_percent"`
	Subtotal        float64 `gorm:"type:decimal(10,2)" json:"subtotal"`
	ExpiryDate      *string `gorm:"type:varchar(100)" json:"expiry_date,omitempty"`
	BatchNumber     *string `gorm:"type:varchar(100)" json:"batch_number,omitempty"`
	Notes           *string `gorm:"type:varchar(100)" json:"notes,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	GoodsReceiving *GoodsReceiving    `gorm:"foreignKey:GRID" json:"goods_receiving,omitempty"`
	Product        *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	POItem         *PurchaseOrderItem `gorm:"foreignKey:POItemID" json:"po_item,omitempty"`
}
