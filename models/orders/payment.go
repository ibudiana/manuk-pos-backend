package orders

import (
	"time"

	"manuk-pos-backend/models/user"
)

type Payment struct {
	ID            int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        int `gorm:"not null" json:"user_id" binding:"required"`
	TransactionID int `gorm:"not null" json:"transaction_id" binding:"required"`

	PaymentMethod   string    `gorm:"type:varchar(100);not null" json:"payment_method" binding:"required"`
	Amount          float64   `gorm:"type:decimal(10,2);not null" json:"amount" binding:"required"`
	ReferenceNumber string    `gorm:"type:varchar(100)" json:"reference_number"`
	PaymentDate     time.Time `gorm:"autoCreateTime" json:"payment_date"`
	Status          string    `gorm:"type:varchar(100);default:'completed'" json:"status"`
	CardLast4       string    `gorm:"type:varchar(100)" json:"card_last4"`
	CardType        string    `gorm:"type:varchar(100)" json:"card_type"`
	EWalletProvider string    `gorm:"type:varchar(100)" json:"e_wallet_provider"`
	ChequeNumber    string    `gorm:"type:varchar(100)" json:"cheque_number"`
	ChequeDate      string    `gorm:"type:varchar(100)" json:"cheque_date"`
	AccountName     string    `gorm:"type:varchar(100)" json:"account_name"`
	Notes           string    `gorm:"type:varchar(100)" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	SyncStatus string `gorm:"type:varchar(100);default:'pending'" json:"sync_status"`

	// Relasi dengan tabel lain
	Transaction Transaction `gorm:"foreignKey:TransactionID" json:"transaction"`
	User        user.User   `gorm:"foreignKey:UserID" json:"user"`
}
