package ordercontroller

import (
	"fmt"
	"net/http"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/orders"

	"github.com/gin-gonic/gin"
)

func CreateTransactionOrder(c *gin.Context) {
	var transaction orders.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	tx := database.DB.Begin()

	// Simpan transaksi utama
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusInternalServerError, "Gagal menyimpan transaksi")
		return
	}

	fmt.Printf("Transaction ID: %d", transaction.ID)

	// Simpan item-item transaksi
	for _, itemInput := range transaction.TransactionItems {
		item := &orders.TransactionItem{
			TransactionID: transaction.ID,
			ProductID:     itemInput.ProductID,
			Quantity:      itemInput.Quantity,
			UnitPrice:     itemInput.UnitPrice,
			Subtotal:      itemInput.Subtotal,
		}
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			helpers.JSONError(c, http.StatusInternalServerError, "Gagal menyimpan item transaksi")
			return
		}
	}

	tx.Commit()
	helpers.JSONResponse(c, http.StatusCreated, "Transaksi berhasil dibuat", transaction)
}

func GetTransactionOrders(c *gin.Context) {
	var transactions []orders.Transaction
	result := database.DB.Preload("TransactionItems").Find(&transactions)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Transaction is empty", transactions)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Transaction", transactions)
}
