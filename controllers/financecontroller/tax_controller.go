package financecontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/finance"

	"github.com/gin-gonic/gin"
)

func GetTaxes(c *gin.Context) {
	var Taxs []finance.Tax
	result := database.DB.Find(&Taxs)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Tax is empty", Taxs)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Tax", Taxs)
}

func CreateTax(c *gin.Context) {
	var Tax finance.Tax

	// Bind JSON ke struct Tax
	if err := c.ShouldBindJSON(&Tax); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()

	// Simpan ke database
	if err := tx.Create(&Tax).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Tax")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create Tax successfuly", Tax)
}

// With Tax ID
func GetTaxByID(c *gin.Context) {
	var Tax finance.Tax

	// ambil id dari parameter
	id := c.Param("id")

	// cari Tax berdasarkan id
	if err := database.DB.First(&Tax, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Tax not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Tax info", Tax)
}

func UpdateTaxByID(c *gin.Context) {
	var Tax finance.Tax

	// ambil id dari parameter
	id := c.Param("id")

	// cari Tax berdasarkan id
	if err := database.DB.First(&Tax, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Tax not found")
		return
	}

	// Binding input JSON ke Tax
	if err := c.ShouldBindJSON(&Tax); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	Tax.ID, _ = strconv.Atoi(id)

	// Update Tax ke dalam database
	if err := database.DB.Omit("ID").Save(&Tax).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "Tax update successfuly", Tax)
}

func DeleteTaxByID(c *gin.Context) {
	var Tax finance.Tax

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Tax berdasarkan id
	if err := database.DB.First(&Tax, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Tax not found")
		return
	}

	// Delete Tax
	if err := database.DB.Delete(&Tax).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Tax")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Tax delete successfuly", nil)
}
