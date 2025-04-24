package inventorycontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/inventory"
	"manuk-pos-backend/models/vendor"

	"github.com/gin-gonic/gin"
)

func AddProductSuppliersByIdProduct(c *gin.Context) {
	var productSupplier inventory.ProductSupplier

	// ambil id dari parameter
	id := c.Param("id")

	// Bind JSON ke struct Product
	if err := c.ShouldBindJSON(&productSupplier); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Cek apakah ProductID dan SupplierID valid
	var product inventory.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		helpers.JSONError(c, http.StatusBadRequest, "Product not found")
		return
	}

	var supplier vendor.Supplier
	if err := database.DB.First(&supplier, productSupplier.SupplierID).Error; err != nil {
		helpers.JSONError(c, http.StatusBadRequest, "Supplier not found")
		return
	}

	// Cek apakah kombinasi sudah ada
	var existing inventory.ProductSupplier

	// Force override ID
	productSupplier.ProductID, _ = strconv.Atoi(id)

	if err := database.DB.
		Where("product_id = ? AND supplier_id = ?", productSupplier.ProductID, productSupplier.SupplierID).
		First(&existing).Error; err == nil {
		helpers.JSONError(c, http.StatusConflict, "This supplier is already linked to the product")
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()

	// Simpan ke database
	if err := tx.Create(&productSupplier).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create product successfuly", productSupplier)
}

func UpdateProductSupplierByIdProduct(c *gin.Context) {
	var productSupplier inventory.ProductSupplier

	// ambil id dari parameter
	id := c.Param("id")

	// cari product berdasarkan id
	if err := database.DB.First(&productSupplier, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Product not found")
		return
	}

	// Binding input JSON ke product
	if err := c.ShouldBindJSON(&productSupplier); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	// productSupplier.ID, _ = strconv.Atoi(id)

	// Update product ke dalam database
	if err := database.DB.Omit("ID").Save(&productSupplier).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Product update successfuly", productSupplier)
}

func DeleteProductSupplierByIdProduct(c *gin.Context) {
	var productSupplier inventory.ProductSupplier

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari product berdasarkan id
	if err := database.DB.First(&productSupplier, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Product not found")
		return
	}

	// Delete product
	if err := database.DB.Delete(&productSupplier).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Product delete successfuly", nil)
}
