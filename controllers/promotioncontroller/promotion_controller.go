package promotioncontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/promotion"

	"github.com/gin-gonic/gin"
)

func GetPromotions(c *gin.Context) {
	var Promotions []promotion.Promotion
	result := database.DB.Find(&Promotions)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Promotion is empty", Promotions)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Promotion", Promotions)
}

func CreatePromotion(c *gin.Context) {
	var Promotion promotion.Promotion

	// Bind JSON ke struct Promotion
	if err := c.ShouldBindJSON(&Promotion); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()

	// Simpan ke database
	if err := tx.Create(&Promotion).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Promotion")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create Promotion successfuly", Promotion)
}

// With Promotion ID
func GetPromotionByID(c *gin.Context) {
	var Promotion promotion.Promotion

	// ambil id dari parameter
	id := c.Param("id")

	// cari Promotion berdasarkan id
	if err := database.DB.First(&Promotion, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Promotion not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Promotion info", Promotion)
}

func UpdatePromotionByID(c *gin.Context) {
	var Promotion promotion.Promotion

	// ambil id dari parameter
	id := c.Param("id")

	// cari Promotion berdasarkan id
	if err := database.DB.First(&Promotion, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Promotion not found")
		return
	}

	// Binding input JSON ke Promotion
	if err := c.ShouldBindJSON(&Promotion); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	Promotion.ID, _ = strconv.Atoi(id)

	// Update Promotion ke dalam database
	if err := database.DB.Omit("ID").Save(&Promotion).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "Promotion update successfuly", Promotion)
}

func DeletePromotionByID(c *gin.Context) {
	var Promotion promotion.Promotion

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Promotion berdasarkan id
	if err := database.DB.First(&Promotion, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Promotion not found")
		return
	}

	// Delete Promotion
	if err := database.DB.Delete(&Promotion).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Promotion")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Promotion delete successfuly", nil)
}
