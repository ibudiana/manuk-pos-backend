package finance

import (
	"time"

	"manuk-pos-backend/models/customer"
)

// Loan represents a customer's loan record
type Loan struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID int `gorm:"not null" json:"customer_id" binding:"required"`

	LoanAmount        float64 `gorm:"not null" json:"loan_amount" binding:"required"`
	InterestRate      float64 `gorm:"default:0" json:"interest_rate"`
	LoanTerm          int     `gorm:"not null" json:"loan_term" binding:"required"`
	InstallmentAmount float64 `gorm:"not null" json:"installment_amount" binding:"required"`
	RemainingAmount   float64 `gorm:"not null" json:"remaining_amount" binding:"required"`
	StartDate         string  `gorm:"type:varchar(100);not null" json:"start_date" binding:"required"`
	DueDate           string  `gorm:"type:varchar(100);not null" json:"due_date" binding:"required"`
	Status            string  `gorm:"type:varchar(100);default:'active'" json:"status"`
	Notes             string  `gorm:"type:varchar(100)" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi dengan tabel lain
	Customer *customer.Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}

// LoanPayment represents a payment made against a loan
type LoanPayment struct {
	ID     int `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanID int `gorm:"not null" json:"loan_id" binding:"required"`

	PaymentAmount float64   `gorm:"not null" json:"payment_amount" binding:"required"`
	PaymentDate   time.Time `gorm:"autoCreateTime" json:"payment_date"`
	Notes         string    `gorm:"type:varchar(100)" json:"notes"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relasi dengan tabel lain
	Loan Loan `gorm:"foreignKey:LoanID" json:"loan"`
}
