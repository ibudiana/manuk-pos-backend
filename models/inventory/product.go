package inventory

import (
	"time"

	"manuk-pos-backend/models/vendor"
)

type Product struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID int `gorm:"not null" json:"product_category_id" binding:"required"`

	SKU             string   `gorm:"type:varchar(25);unique;not null" json:"sku" binding:"required"`
	Barcode         string   `gorm:"type:varchar(25);unique" json:"barcode"`
	Name            string   `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Description     string   `gorm:"type:text" json:"description"`
	BuyingPrice     float64  `gorm:"type:decimal(10,2);not null" json:"buying_price" binding:"required"`
	SellingPrice    float64  `gorm:"type:decimal(10,2);not null" json:"selling_price" binding:"required"`
	MinStock        int      `gorm:"default:1" json:"min_stock"`
	DiscountPrice   *float64 `gorm:"type:decimal(10,2)" json:"discount_price,omitempty"`
	Weight          *float64 `gorm:"type:decimal(10,2)" json:"weight,omitempty"`
	DimensionLength *float64 `gorm:"type:decimal(10,2)" json:"dimension_length,omitempty"`
	DimensionWidth  *float64 `gorm:"type:decimal(10,2)" json:"dimension_width,omitempty"`
	DimensionHeight *float64 `gorm:"type:decimal(10,2)" json:"dimension_height,omitempty"`
	IsService       bool     `gorm:"type:tinyint(1);default:0" json:"is_service"`
	IsActive        bool     `gorm:"type:tinyint(1);default:1" json:"is_active"`
	IsFeatured      bool     `gorm:"type:tinyint(1);default:0" json:"is_featured"`
	AllowFractions  int      `gorm:"default:0" json:"allow_fractions"`
	ImageURL        string   `gorm:"type:varchar(255)" json:"image_url"`
	Tags            string   `gorm:"type:varchar(100)" json:"tags"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

type ProductSupplier struct {
	ProductID  int `gorm:"primaryKey;not null" json:"product_id"`
	SupplierID int `gorm:"primaryKey;not null" json:"supplier_id" binding:"required"`

	BuyingPrice          *float64   `gorm:"type:decimal(10,2)" json:"buying_price,omitempty"`
	LeadTime             *int       `json:"lead_time,omitempty"` // in days
	MinimumOrderQuantity int        `gorm:"default:1" json:"minimum_order_quantity"`
	IsPrimary            bool       `gorm:"type:tinyint(1);default:0" json:"is_primary"`
	LastSupplyDate       *time.Time `json:"last_supply_date,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Product  *Product         `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Supplier *vendor.Supplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
}
