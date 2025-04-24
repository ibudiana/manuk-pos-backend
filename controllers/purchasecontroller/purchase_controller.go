package purchasecontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/purchase"

	"github.com/gin-gonic/gin"
)

func GetPurchaseOrders(c *gin.Context) {
	var PurchaseOrders []purchase.PurchaseOrder
	result := database.DB.Find(&PurchaseOrders)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "PurchaseOrder is empty", PurchaseOrders)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List PurchaseOrder", PurchaseOrders)
}

func CreatePurchaseOrder(c *gin.Context) {
	var PurchaseOrder purchase.PurchaseOrder

	// Bind JSON ke struct PurchaseOrder
	if err := c.ShouldBindJSON(&PurchaseOrder); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek SKU
	if err := tx.Where("ponumber = ?", PurchaseOrder.PONumber).First(&PurchaseOrder).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "PurchaseOrder PON umber already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&PurchaseOrder).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create PurchaseOrder")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create PurchaseOrder successfuly", PurchaseOrder)
}

// With PurchaseOrder ID
func GetPurchaseOrderByID(c *gin.Context) {
	var PurchaseOrder purchase.PurchaseOrder

	// ambil id dari parameter
	id := c.Param("id")

	// cari PurchaseOrder berdasarkan id
	if err := database.DB.First(&PurchaseOrder, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "PurchaseOrder not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "PurchaseOrder info", PurchaseOrder)
}

func UpdatePurchaseOrderByID(c *gin.Context) {
	var PurchaseOrder purchase.PurchaseOrder

	// ambil id dari parameter
	id := c.Param("id")

	// cari PurchaseOrder berdasarkan id
	if err := database.DB.First(&PurchaseOrder, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "PurchaseOrder not found")
		return
	}

	// Binding input JSON ke PurchaseOrder
	if err := c.ShouldBindJSON(&PurchaseOrder); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	PurchaseOrder.ID, _ = strconv.Atoi(id)

	// Update PurchaseOrder ke dalam database
	if err := database.DB.Omit("ID").Save(&PurchaseOrder).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "PurchaseOrder update successfuly", PurchaseOrder)
}

func DeletePurchaseOrderByID(c *gin.Context) {
	var PurchaseOrder purchase.PurchaseOrder

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari PurchaseOrder berdasarkan id
	if err := database.DB.First(&PurchaseOrder, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "PurchaseOrder not found")
		return
	}

	// Delete PurchaseOrder
	if err := database.DB.Delete(&PurchaseOrder).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete PurchaseOrder")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "PurchaseOrder delete successfuly", nil)
}
