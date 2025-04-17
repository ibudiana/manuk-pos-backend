package orders

import (
	"time"

	"manuk-pos-backend/models/customer"
	"manuk-pos-backend/models/finance"
	"manuk-pos-backend/models/inventory"
	"manuk-pos-backend/models/promotion"
	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
)

type Transaction struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID int `gorm:"not null" json:"customer_id" binding:"required"`
	UserID     int `gorm:"not null" json:"user_id" binding:"required"`
	BranchID   int `gorm:"not null" json:"branch_id" binding:"required"`
	DiscountID int `gorm:"not null" json:"discount_id" binding:"required"`
	TaxID      int `gorm:"not null" json:"tax_id" binding:"required"`
	FeeID      int `gorm:"not null" json:"fee_id" binding:"required"`

	InvoiceNumber    string    `gorm:"type:varchar(100);unique;not null" json:"invoice_number" binding:"required"`
	InvoiceDate      time.Time `gorm:"autoCreateTime" json:"invoice_date"`
	TransactionDate  time.Time `gorm:"autoCreateTime" json:"transaction_date"`
	DueDate          string    `gorm:"type:varchar(100)" json:"due_date"`
	Subtotal         float64   `gorm:"type:decimal(10,2);not null" json:"subtotal" binding:"required"`
	DiscountAmount   float64   `gorm:"type:decimal(10,2);default:0" json:"discount_amount"`
	TaxAmount        float64   `gorm:"type:decimal(10,2);default:0" json:"tax_amount"`
	FeeAmount        float64   `gorm:"type:decimal(10,2);default:0" json:"fee_amount"`
	ShippingCost     float64   `gorm:"type:decimal(10,2);default:0" json:"shipping_cost"`
	GrandTotal       float64   `gorm:"type:decimal(10,2);not null" json:"grand_total" binding:"required"`
	AmountPaid       float64   `gorm:"type:decimal(10,2);default:0" json:"amount_paid"`
	AmountReturned   float64   `gorm:"type:decimal(10,2);default:0" json:"amount_returned"`
	PaymentStatus    string    `gorm:"type:varchar(100);default:'unpaid'" json:"payment_status"`
	PointsEarned     int       `gorm:"type:int;default:0" json:"points_earned"`
	PointsUsed       int       `gorm:"type:int;default:0" json:"points_used"`
	Notes            string    `gorm:"type:varchar(100)" json:"notes"`
	Status           string    `gorm:"type:varchar(100);default:'completed'" json:"status"`
	ReferenceID      int       `gorm:"not null" json:"reference_id" binding:"required"`
	ShippingAddress  string    `gorm:"type:varchar(100)" json:"shipping_address"`
	ShippingTracking string    `gorm:"type:varchar(100)" json:"shipping_tracking"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	SyncStatus       string    `gorm:"type:varchar(100);default:'pending'" json:"sync_status"`

	// Relasi dengan tabel lain
	Customer *customer.Customer  `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	User     *user.User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Branch   *store.Branch       `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Discount *promotion.Discount `gorm:"foreignKey:DiscountID" json:"discount,omitempty"`
	Tax      *finance.Tax        `gorm:"foreignKey:TaxID" json:"tax,omitempty"`
	Fee      *Fee                `gorm:"foreignKey:FeeID" json:"fee,omitempty"`

	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID" json:"transaction_items" binding:"required,dive"`
}

type TransactionItem struct {
	ID            int `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID int `gorm:"not null" json:"transaction_id"`
	ProductID     int `gorm:"not null" json:"product_id" binding:"required"`

	Quantity        float64 `gorm:"type:decimal(10,2);not null" json:"quantity" binding:"required"`
	UnitPrice       float64 `gorm:"type:decimal(10,2);not null" json:"unit_price" binding:"required"`
	OriginalPrice   float64 `gorm:"type:decimal(10,2)" json:"original_price"`
	DiscountPercent float64 `gorm:"type:decimal(10,2);default:0" json:"discount_percent"`
	DiscountAmount  float64 `gorm:"type:decimal(10,2);default:0" json:"discount_amount"`
	TaxPercent      float64 `gorm:"type:decimal(10,2);default:0" json:"tax_percent"`
	TaxAmount       float64 `gorm:"type:decimal(10,2);default:0" json:"tax_amount"`
	Subtotal        float64 `gorm:"type:decimal(10,2);not null" json:"subtotal" binding:"required"`
	Notes           string  `gorm:"type:varchar(100)" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	SyncStatus string `gorm:"type:varchar(100);default:'pending'" json:"sync_status"`

	// Relasi dengan tabel lain
	Product     *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Transaction *Transaction       `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
}
