package promotion

import (
	"time"

	"manuk-pos-backend/models/inventory"
)

type Promotion struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`

	Name        string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	PromoType   string `gorm:"type:varchar(100);not null" json:"promo_type" binding:"required"`
	StartDate   string `gorm:"type:varchar(100);not null" json:"start_date" binding:"required"`
	EndDate     string `gorm:"type:varchar(100);not null" json:"end_date" binding:"required"`
	IsActive    bool   `gorm:"type:tinyint(1);default:1" json:"is_active"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type PromotionRule struct {
	ID          int `gorm:"primaryKey;autoIncrement" json:"id"`
	PromotionID int `gorm:"not null" json:"promotion_id" binding:"required"`

	MinQuantity   int     `gorm:"type:int" json:"min_quantity"`
	MinAmount     float64 `gorm:"type:decimal(10,2)" json:"min_amount"`
	DiscountType  string  `gorm:"type:varchar(100)" json:"discount_type"`
	DiscountValue float64 `gorm:"type:decimal(10,2)" json:"discount_value"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Promotion *Promotion `gorm:"foreignKey:PromotionID;references:ID" json:"promotion,omitempty"`
}

type PromotionProduct struct {
	PromotionID int `gorm:"primaryKey" json:"promotion_id"`
	ProductID   int `gorm:"primaryKey" json:"product_id"`

	IsTrigger bool `gorm:"type:tinyint(1);default:0" json:"is_trigger"`
	IsTarget  bool `gorm:"type:tinyint(1);default:0" json:"is_target"`
	Quantity  int  `gorm:"type:int;default:1" json:"quantity"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relasi dengan tabel lain
	Product   *inventory.Product `gorm:"foreignKey:ProductID;references:ID" json:"product,omitempty"`
	Promotion *Promotion         `gorm:"foreignKey:PromotionID;references:ID" json:"promotion,omitempty"`
}
