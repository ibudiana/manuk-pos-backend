package controllers

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/vendor"

	"github.com/gin-gonic/gin"
)

func GetSuppliers(c *gin.Context) {
	var Suppliers []vendor.Supplier
	result := database.DB.Find(&Suppliers)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Supplier is empty", Suppliers)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Supplier", Suppliers)
}

func CreateSupplier(c *gin.Context) {
	var Supplier vendor.Supplier

	// Bind JSON ke struct Supplier
	if err := c.ShouldBindJSON(&Supplier); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek SKU
	if err := tx.Where("code = ?", Supplier.Code).First(&Supplier).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "Supplier Code already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&Supplier).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Supplier")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create Supplier successfuly", Supplier)
}

// With Supplier ID
func GetSupplierByID(c *gin.Context) {
	var Supplier vendor.Supplier

	// ambil id dari parameter
	id := c.Param("id")

	// cari Supplier berdasarkan id
	if err := database.DB.First(&Supplier, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Supplier not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Supplier info", Supplier)
}

func UpdateSupplierByID(c *gin.Context) {
	var Supplier vendor.Supplier

	// ambil id dari parameter
	id := c.Param("id")

	// cari Supplier berdasarkan id
	if err := database.DB.First(&Supplier, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Supplier not found")
		return
	}

	// Binding input JSON ke Supplier
	if err := c.ShouldBindJSON(&Supplier); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	Supplier.ID, _ = strconv.Atoi(id)

	// Update Supplier ke dalam database
	if err := database.DB.Omit("ID").Save(&Supplier).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "Supplier update successfuly", Supplier)
}

func DeleteSupplierByID(c *gin.Context) {
	var Supplier vendor.Supplier

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Supplier berdasarkan id
	if err := database.DB.First(&Supplier, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Supplier not found")
		return
	}

	// Delete Supplier
	if err := database.DB.Delete(&Supplier).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Supplier")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Supplier delete successfuly", nil)
}
