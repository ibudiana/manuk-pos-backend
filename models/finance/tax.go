package finance

import "time"

type Tax struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`

	Name      string  `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Rate      float64 `gorm:"type:decimal(10,2);not null" json:"rate" binding:"required"`
	IsDefault int     `gorm:"type:int;default:0" json:"is_default"`
	IsActive  bool    `gorm:"type:tinyint(1);default:1" json:"is_active"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
