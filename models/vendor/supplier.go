package vendor

import "time"

type Supplier struct {
	ID    int `gorm:"primaryKey;autoIncrement" json:"id"`
	TaxID int `json:"tax_id"`

	Code          string `gorm:"type:varchar(50);unique" json:"code"`
	Name          string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	ContactPerson string `gorm:"type:varchar(25)" json:"contact_person"`
	Phone         string `gorm:"type:varchar(25)" json:"phone"`
	Email         string `gorm:"type:varchar(100)" json:"email"`
	Address       string `gorm:"type:varchar(255)" json:"address"`
	PaymentTerms  int    `json:"payment_terms"`
	IsActive      bool   `gorm:"default:true" json:"is_active"`
	Notes         string `gorm:"type:varchar(100)" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
