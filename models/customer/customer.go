package customer

import (
	"time"
)

type Customer struct {
	ID    int `gorm:"primaryKey;autoIncrement" json:"id"`
	TaxID int `json:"tax_id"` // tanpa relasi ke tax

	Code           string     `gorm:"type:varchar(50);unique" json:"code"`
	Name           string     `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Phone          string     `gorm:"type:varchar(25)" json:"phone"`
	Email          string     `gorm:"type:varchar(100)" json:"email"`
	Address        string     `gorm:"type:varchar(255)" json:"address"`
	City           string     `gorm:"type:varchar(100)" json:"city"`
	PostalCode     string     `gorm:"type:varchar(10)" json:"postal_code"`
	Birthdate      *time.Time `json:"birthdate"` // Nullable DATE
	JoinDate       time.Time  `gorm:"autoCreateTime" json:"join_date"`
	CustomerType   string     `gorm:"type:varchar(25);default:regular" json:"customer_type"`
	CreditLimit    float64    `gorm:"type:decimal(10,2);default:0" json:"credit_limit"`
	CurrentBalance float64    `gorm:"type:decimal(10,2);default:0" json:"current_balance"`
	IsActive       bool       `gorm:"default:true" json:"is_active"`
	Notes          string     `gorm:"type:varchar(100)" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
