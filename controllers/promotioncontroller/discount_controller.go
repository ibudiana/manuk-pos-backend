package promotioncontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/promotion"

	"github.com/gin-gonic/gin"
)

func GetDiscounts(c *gin.Context) {
	var Discounts []promotion.Discount
	result := database.DB.Find(&Discounts)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Discount is empty", Discounts)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Discount", Discounts)
}

func CreateDiscount(c *gin.Context) {
	var Discount promotion.Discount

	// Bind JSON ke struct Discount
	if err := c.ShouldBindJSON(&Discount); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek SKU
	if err := tx.Where("code = ?", Discount.Code).First(&Discount).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "Discount CODE already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&Discount).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Discount")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create Discount successfuly", Discount)
}

// With Discount ID
func GetDiscountByID(c *gin.Context) {
	var Discount promotion.Discount

	// ambil id dari parameter
	id := c.Param("id")

	// cari Discount berdasarkan id
	if err := database.DB.First(&Discount, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Discount not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Discount info", Discount)
}

func UpdateDiscountByID(c *gin.Context) {
	var Discount promotion.Discount

	// ambil id dari parameter
	id := c.Param("id")

	// cari Discount berdasarkan id
	if err := database.DB.First(&Discount, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Discount not found")
		return
	}

	// Binding input JSON ke Discount
	if err := c.ShouldBindJSON(&Discount); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	Discount.ID, _ = strconv.Atoi(id)

	// Update Discount ke dalam database
	if err := database.DB.Omit("ID").Save(&Discount).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Discount update successfuly", Discount)
}

func DeleteDiscountByID(c *gin.Context) {
	var Discount promotion.Discount

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Discount berdasarkan id
	if err := database.DB.First(&Discount, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Discount not found")
		return
	}

	// Delete Discount
	if err := database.DB.Delete(&Discount).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Discount")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Discount delete successfuly", nil)
}
