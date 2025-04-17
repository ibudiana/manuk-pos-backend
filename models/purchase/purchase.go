package purchase

import (
	"time"

	"manuk-pos-backend/models/inventory"
	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
	"manuk-pos-backend/models/vendor"
)

type PurchaseOrder struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	PONumber   string `gorm:"unique;not null;type:varchar(100)" json:"po_number" binding:"required"`
	SupplierID int    `gorm:"not null" json:"supplier_id" binding:"required"`
	BranchID   int    `gorm:"not null" json:"branch_id" binding:"required"`
	UserID     int    `gorm:"not null" json:"user_id" binding:"required"`

	PODate         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"po_date"`
	ExpectedDate   *string   `gorm:"type:varchar(100)" json:"expected_date,omitempty"`
	Subtotal       float64   `gorm:"type:decimal(10,2);not null" json:"subtotal" binding:"required"`
	TaxAmount      float64   `gorm:"type:decimal(10,2);default:0" json:"tax_amount"`
	DiscountAmount float64   `gorm:"type:decimal(10,2);default:0" json:"discount_amount"`
	ShippingCost   float64   `gorm:"type:decimal(10,2);default:0" json:"shipping_cost"`
	OtherCosts     float64   `gorm:"type:decimal(10,2);default:0" json:"other_costs"`
	GrandTotal     float64   `gorm:"type:decimal(10,2);not null" json:"grand_total" binding:"required"`
	Status         string    `gorm:"default:'draft';type:varchar(100)" json:"status"`
	Notes          *string   `gorm:"type:varchar(100)" json:"notes,omitempty"`
	PaymentTerms   *int      `gorm:"type:int" json:"payment_terms,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Supplier *vendor.Supplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	Branch   *store.Branch    `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	User     *user.User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type PurchaseOrderItem struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	POID      int `gorm:"not null" json:"po_id" binding:"required"`
	ProductID int `gorm:"not null" json:"product_id" binding:"required"`

	Quantity         int     `gorm:"not null" json:"quantity" binding:"required"`
	ReceivedQuantity int     `gorm:"default:0" json:"received_quantity"`
	UnitPrice        float64 `gorm:"type:decimal(10,2);not null" json:"unit_price" binding:"required"`
	DiscountPercent  float64 `gorm:"type:decimal(10,2);default:0" json:"discount_percent"`
	TaxPercent       float64 `gorm:"type:decimal(10,2);default:0" json:"tax_percent"`
	Subtotal         float64 `gorm:"type:decimal(10,2)" json:"subtotal"`
	Notes            *string `gorm:"type:varchar(100)" json:"notes,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	PO      *PurchaseOrder     `gorm:"foreignKey:POID" json:"po,omitempty"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
