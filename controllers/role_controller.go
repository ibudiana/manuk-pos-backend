package controllers

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/user"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	var roles []user.Role
	result := database.DB.Find(&roles)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Roles is empty", roles)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List of Roles", roles)
}

func CreateRole(c *gin.Context) {
	var role user.Role

	// Bind JSON ke struct Role
	if err := c.ShouldBindJSON(&role); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek apakah role dengan nama yang sama sudah ada
	if err := tx.Where("name = ?", role.Name).First(&role).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "Role already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Role")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Role created successfully", role)
}

// With Role ID
func GetRoleByID(c *gin.Context) {
	var role user.Role

	// ambil id dari parameter
	id := c.Param("id")

	// cari Role berdasarkan id
	if err := database.DB.First(&role, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Role not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Role info", role)
}

func UpdateRoleByID(c *gin.Context) {
	var role user.Role

	// ambil id dari parameter
	id := c.Param("id")

	// cari Role berdasarkan id
	if err := database.DB.First(&role, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Role not found")
		return
	}

	// Binding input JSON ke Role
	if err := c.ShouldBindJSON(&role); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	role.ID, _ = strconv.Atoi(id)

	// Update Role ke dalam database
	if err := database.DB.Omit("ID").Save(&role).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to update Role")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Role updated successfully", role)
}

func DeleteRoleByID(c *gin.Context) {
	var role user.Role

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Role berdasarkan id
	if err := database.DB.First(&role, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Role not found")
		return
	}

	// Delete Role
	if err := database.DB.Delete(&role).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Role")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Role deleted successfully", nil)
}
