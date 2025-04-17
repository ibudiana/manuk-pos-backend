package financecontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/finance"

	"github.com/gin-gonic/gin"
)

func GetLoans(c *gin.Context) {
	var Loans []finance.Loan
	result := database.DB.Find(&Loans)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Loan is empty", Loans)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Loan", Loans)
}

func CreateLoan(c *gin.Context) {
	var Loan finance.Loan

	// Bind JSON ke struct Loan
	if err := c.ShouldBindJSON(&Loan); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()

	// Simpan ke database
	if err := tx.Create(&Loan).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Loan")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create Loan successfuly", Loan)
}

// With Loan ID
func GetLoanByID(c *gin.Context) {
	var Loan finance.Loan

	// ambil id dari parameter
	id := c.Param("id")

	// cari Loan berdasarkan id
	if err := database.DB.First(&Loan, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Loan not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Loan info", Loan)
}

func UpdateLoanByID(c *gin.Context) {
	var Loan finance.Loan

	// ambil id dari parameter
	id := c.Param("id")

	// cari Loan berdasarkan id
	if err := database.DB.First(&Loan, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Loan not found")
		return
	}

	// Binding input JSON ke Loan
	if err := c.ShouldBindJSON(&Loan); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	Loan.ID, _ = strconv.Atoi(id)

	// Update Loan ke dalam database
	if err := database.DB.Omit("ID").Save(&Loan).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "Loan update successfuly", Loan)
}

func DeleteLoanByID(c *gin.Context) {
	var Loan finance.Loan

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Loan berdasarkan id
	if err := database.DB.First(&Loan, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Loan not found")
		return
	}

	// Delete Loan
	if err := database.DB.Delete(&Loan).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Loan")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Loan delete successfuly", nil)
}
