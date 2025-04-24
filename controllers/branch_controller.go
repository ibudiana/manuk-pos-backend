package controllers

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/store"

	"github.com/gin-gonic/gin"
)

func GetBranches(c *gin.Context) {
	var Branch []store.Branch
	result := database.DB.Find(&Branch)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Branch is empty", Branch)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Branch", Branch)
}

func CreateBranch(c *gin.Context) {
	var Branch store.Branch

	// Bind JSON ke struct Branch
	if err := c.ShouldBindJSON(&Branch); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek SKU
	if err := tx.Where("code = ?", Branch.Code).First(&Branch).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "Branch Code already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&Branch).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Branch")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create Branch successfuly", Branch)
}

// With Branch ID
func GetBranchByID(c *gin.Context) {
	var Branch store.Branch

	// ambil id dari parameter
	id := c.Param("id")

	// cari Branch berdasarkan id
	if err := database.DB.First(&Branch, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Branch not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Branch info", Branch)
}

func UpdateBranchByID(c *gin.Context) {
	var Branch store.Branch

	// ambil id dari parameter
	id := c.Param("id")

	// cari Branch berdasarkan id
	if err := database.DB.First(&Branch, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Branch not found")
		return
	}

	// Binding input JSON ke Branch
	if err := c.ShouldBindJSON(&Branch); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	Branch.ID, _ = strconv.Atoi(id)

	// Update Branch ke dalam database
	if err := database.DB.Omit("ID").Save(&Branch).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Branch update successfuly", Branch)
}

func DeleteBranchByID(c *gin.Context) {
	var Branch store.Branch

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Branch berdasarkan id
	if err := database.DB.First(&Branch, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Branch not found")
		return
	}

	// Delete Branch
	if err := database.DB.Delete(&Branch).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Branch")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Branch delete successfuly", nil)
}
