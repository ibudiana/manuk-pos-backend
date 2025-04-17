package promotion

import (
	"time"

	"manuk-pos-backend/models/customer"
	"manuk-pos-backend/models/inventory"
)

type Discount struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID int `gorm:"type:int" json:"category_id"`
	ProductID  int `gorm:"type:int" json:"product_id"`
	CustomerID int `gorm:"type:int" json:"customer_id"`

	Name          string  `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Code          string  `gorm:"type:varchar(100);unique" json:"code"`
	Description   string  `gorm:"type:varchar(100)" json:"description"`
	DiscountType  string  `gorm:"type:varchar(100);not null" json:"discount_type" binding:"required"`
	DiscountValue float64 `gorm:"type:decimal(10,2);not null" json:"discount_value" binding:"required"`
	MinPurchase   float64 `gorm:"type:decimal(10,2)" json:"min_purchase"`
	MaxDiscount   float64 `gorm:"type:decimal(10,2)" json:"max_discount"`
	StartDate     string  `gorm:"type:varchar(100)" json:"start_date"`
	EndDate       string  `gorm:"type:varchar(100)" json:"end_date"`
	UsageLimit    int     `gorm:"type:int" json:"usage_limit"`
	UsageCount    int     `gorm:"type:int;default:0" json:"usage_count"`
	IsActive      bool    `gorm:"type:tinyint(1);default:1" json:"is_active"`
	AppliesTo     string  `gorm:"type:varchar(100)" json:"applies_to"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Category *inventory.Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Product  *inventory.Product  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Customer *customer.Customer  `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}
