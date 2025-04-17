package inventorycontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/inventory"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []inventory.Product
	result := database.DB.Find(&products)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Product is empty", products)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List product", products)
}

func CreateProduct(c *gin.Context) {
	var product inventory.Product

	// Bind JSON ke struct Product
	if err := c.ShouldBindJSON(&product); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek SKU
	if err := tx.Where("sku = ?", product.SKU).First(&product).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "Product SKU already exists")
		return
	}

	// Cek Barcode (jika barcode tidak kosong/null)
	if product.Barcode != "" {
		if err := tx.Where("barcode = ?", product.Barcode).First(&product).Error; err == nil {
			tx.Rollback()
			helpers.JSONError(c, http.StatusConflict, "Product barcode already exists")
			return
		}
	}

	// Simpan ke database
	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create product successfuly", product)
}

// With Product ID
func GetProductByID(c *gin.Context) {
	var product inventory.Product

	// ambil id dari parameter
	id := c.Param("id")

	// cari product berdasarkan id
	if err := database.DB.First(&product, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Product not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Product info", product)
}

func UpdateProductByID(c *gin.Context) {
	var product inventory.Product

	// ambil id dari parameter
	id := c.Param("id")

	// cari product berdasarkan id
	if err := database.DB.First(&product, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Product not found")
		return
	}

	// Binding input JSON ke product
	if err := c.ShouldBindJSON(&product); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	product.ID, _ = strconv.Atoi(id)

	// Update product ke dalam database
	if err := database.DB.Omit("ID").Save(&product).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "Product update successfuly", product)
}

func DeleteProductByID(c *gin.Context) {
	var product inventory.Product

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari product berdasarkan id
	if err := database.DB.First(&product, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Product not found")
		return
	}

	// Delete product
	if err := database.DB.Delete(&product).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Product delete successfuly", nil)
}
