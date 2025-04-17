package orders

import "time"

type Fee struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`

	Name      string  `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	FeeType   string  `gorm:"type:varchar(100);not null" json:"fee_type" binding:"required"`
	FeeValue  float64 `gorm:"type:decimal(10,2);not null" json:"fee_value" binding:"required"`
	IsDefault int     `gorm:"type:int;default:0" json:"is_default"`
	IsActive  bool    `gorm:"type:tinyint(1);default:1" json:"is_active"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
