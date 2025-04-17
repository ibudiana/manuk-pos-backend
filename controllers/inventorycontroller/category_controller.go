package inventorycontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/inventory"

	"github.com/gin-gonic/gin"
)

func GetProductCategories(c *gin.Context) {
	var ProductCategorys []inventory.Category
	result := database.DB.Find(&ProductCategorys)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "ProductCategory is empty", ProductCategorys)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List ProductCategory", ProductCategorys)
}

func CreateProductCategory(c *gin.Context) {
	var ProductCategory inventory.Category

	// Bind JSON ke struct ProductCategory
	if err := c.ShouldBindJSON(&ProductCategory); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek SKU
	if err := tx.Where("code = ?", ProductCategory.Code).First(&ProductCategory).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "ProductCategory Code already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&ProductCategory).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create ProductCategory")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create ProductCategory successfuly", ProductCategory)
}

// With ProductCategory ID
func GetProductCategoryByID(c *gin.Context) {
	var ProductCategory inventory.Category

	// ambil id dari parameter
	id := c.Param("id")

	// cari ProductCategory berdasarkan id
	if err := database.DB.First(&ProductCategory, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "ProductCategory not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "ProductCategory info", ProductCategory)
}

func UpdateProductCategoryByID(c *gin.Context) {
	var ProductCategory inventory.Category

	// ambil id dari parameter
	id := c.Param("id")

	// cari ProductCategory berdasarkan id
	if err := database.DB.First(&ProductCategory, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "ProductCategory not found")
		return
	}

	// Binding input JSON ke ProductCategory
	if err := c.ShouldBindJSON(&ProductCategory); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	ProductCategory.ID, _ = strconv.Atoi(id)

	// Update ProductCategory ke dalam database
	if err := database.DB.Omit("ID").Save(&ProductCategory).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "ProductCategory update successfuly", ProductCategory)
}

func DeleteProductCategoryByID(c *gin.Context) {
	var ProductCategory inventory.Category

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari ProductCategory berdasarkan id
	if err := database.DB.First(&ProductCategory, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "ProductCategory not found")
		return
	}

	// Delete ProductCategory
	if err := database.DB.Delete(&ProductCategory).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete ProductCategory")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "ProductCategory delete successfuly", nil)
}
