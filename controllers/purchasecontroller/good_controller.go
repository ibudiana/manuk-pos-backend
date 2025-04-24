package purchasecontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/purchase"

	"github.com/gin-gonic/gin"
)

func GetGoodsReceivings(c *gin.Context) {
	var GoodsReceivings []purchase.GoodsReceiving
	result := database.DB.Find(&GoodsReceivings)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "GoodsReceiving is empty", GoodsReceivings)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List GoodsReceiving", GoodsReceivings)
}

func CreateGoodsReceiving(c *gin.Context) {
	var GoodsReceiving purchase.GoodsReceiving

	// Bind JSON ke struct GoodsReceiving
	if err := c.ShouldBindJSON(&GoodsReceiving); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek SKU
	if err := tx.Where("referencenumber = ?", GoodsReceiving.ReferenceNumber).First(&GoodsReceiving).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "GoodsReceiving ReferenceNumber already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&GoodsReceiving).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create GoodsReceiving")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create GoodsReceiving successfuly", GoodsReceiving)
}

// With GoodsReceiving ID
func GetGoodsReceivingByID(c *gin.Context) {
	var GoodsReceiving purchase.GoodsReceiving

	// ambil id dari parameter
	id := c.Param("id")

	// cari GoodsReceiving berdasarkan id
	if err := database.DB.First(&GoodsReceiving, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "GoodsReceiving not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "GoodsReceiving info", GoodsReceiving)
}

func UpdateGoodsReceivingByID(c *gin.Context) {
	var GoodsReceiving purchase.GoodsReceiving

	// ambil id dari parameter
	id := c.Param("id")

	// cari GoodsReceiving berdasarkan id
	if err := database.DB.First(&GoodsReceiving, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "GoodsReceiving not found")
		return
	}

	// Binding input JSON ke GoodsReceiving
	if err := c.ShouldBindJSON(&GoodsReceiving); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	GoodsReceiving.ID, _ = strconv.Atoi(id)

	// Update GoodsReceiving ke dalam database
	if err := database.DB.Omit("ID").Save(&GoodsReceiving).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "GoodsReceiving update successfuly", GoodsReceiving)
}

func DeleteGoodsReceivingByID(c *gin.Context) {
	var GoodsReceiving purchase.GoodsReceiving

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari GoodsReceiving berdasarkan id
	if err := database.DB.First(&GoodsReceiving, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "GoodsReceiving not found")
		return
	}

	// Delete GoodsReceiving
	if err := database.DB.Delete(&GoodsReceiving).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete GoodsReceiving")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "GoodsReceiving delete successfuly", nil)
}
