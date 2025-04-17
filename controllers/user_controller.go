package controllers

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/user"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var Users []user.User
	result := database.DB.Find(&Users)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "User is empty", Users)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List User", Users)
}

func CreateUser(c *gin.Context) {
	var User user.User

	// Bind JSON ke struct User
	if err := c.ShouldBindJSON(&User); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek Username
	if err := tx.Where("username = ?", User.Username).First(&User).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "Username already exists")
		return
	}

	// Cek Email
	if err := tx.Where("email = ?", User.Email).First(&User).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "User Email already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&User).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create User")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create User successfuly", User)
}

// With User ID
func GetUserByID(c *gin.Context) {
	var User user.User

	// ambil id dari parameter
	id := c.Param("id")

	// cari User berdasarkan id
	if err := database.DB.First(&User, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "User not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "User info", User)
}

func UpdateUserByID(c *gin.Context) {
	var User user.User

	// ambil id dari parameter
	id := c.Param("id")

	// cari User berdasarkan id
	if err := database.DB.First(&User, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "User not found")
		return
	}

	// Binding input JSON ke User
	if err := c.ShouldBindJSON(&User); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	User.ID, _ = strconv.Atoi(id)

	// Update User ke dalam database
	if err := database.DB.Omit("ID").Save(&User).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "User update successfuly", User)
}

func DeleteUserByID(c *gin.Context) {
	var User user.User

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari User berdasarkan id
	if err := database.DB.First(&User, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "User not found")
		return
	}

	// Delete User
	if err := database.DB.Delete(&User).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete User")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "User delete successfuly", nil)
}
