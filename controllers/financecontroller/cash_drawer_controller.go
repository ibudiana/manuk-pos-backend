package financecontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/finance"

	"github.com/gin-gonic/gin"
)

func GetCashDrawers(c *gin.Context) {
	var CashDrawers []finance.CashDrawer
	result := database.DB.Find(&CashDrawers)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "CashDrawer is empty", CashDrawers)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List CashDrawer", CashDrawers)
}

func CreateCashDrawer(c *gin.Context) {
	var CashDrawer finance.CashDrawer

	// Bind JSON ke struct CashDrawer
	if err := c.ShouldBindJSON(&CashDrawer); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()

	// Simpan ke database
	if err := tx.Create(&CashDrawer).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create CashDrawer")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create CashDrawer successfuly", CashDrawer)
}

// With CashDrawer ID
func GetCashDrawerByID(c *gin.Context) {
	var CashDrawer finance.CashDrawer

	// ambil id dari parameter
	id := c.Param("id")

	// cari CashDrawer berdasarkan id
	if err := database.DB.First(&CashDrawer, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "CashDrawer not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "CashDrawer info", CashDrawer)
}

func UpdateCashDrawerByID(c *gin.Context) {
	var CashDrawer finance.CashDrawer

	// ambil id dari parameter
	id := c.Param("id")

	// cari CashDrawer berdasarkan id
	if err := database.DB.First(&CashDrawer, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "CashDrawer not found")
		return
	}

	// Binding input JSON ke CashDrawer
	if err := c.ShouldBindJSON(&CashDrawer); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	CashDrawer.ID, _ = strconv.Atoi(id)

	// Update CashDrawer ke dalam database
	if err := database.DB.Omit("ID").Save(&CashDrawer).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "CashDrawer update successfuly", CashDrawer)
}

func DeleteCashDrawerByID(c *gin.Context) {
	var CashDrawer finance.CashDrawer

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari CashDrawer berdasarkan id
	if err := database.DB.First(&CashDrawer, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "CashDrawer not found")
		return
	}

	// Delete CashDrawer
	if err := database.DB.Delete(&CashDrawer).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete CashDrawer")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "CashDrawer delete successfuly", nil)
}
