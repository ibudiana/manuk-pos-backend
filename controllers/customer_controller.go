package controllers

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/customer"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	var Customers []customer.Customer
	result := database.DB.Find(&Customers)

	// Periksa apakah hasilnya kosong
	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Customer is empty", Customers)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Customer", Customers)
}

func CreateCustomer(c *gin.Context) {
	var Customer customer.Customer

	// Bind JSON ke struct Customer
	if err := c.ShouldBindJSON(&Customer); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Mulai Transaksi
	tx := database.DB.Begin()
	// Cek Customer Code
	if err := tx.Where("code = ?", Customer.Code).First(&Customer).Error; err == nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusConflict, "Customer Code already exists")
		return
	}

	// Simpan ke database
	if err := tx.Create(&Customer).Error; err != nil {
		tx.Rollback() // Batalkan jika gagal
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Customer")
		return
	}

	tx.Commit()

	// Respon sukses
	helpers.JSONResponse(c, http.StatusCreated, "Create Customer successfuly", Customer)
}

// With Customer ID
func GetCustomerByID(c *gin.Context) {
	var Customer customer.Customer

	// ambil id dari parameter
	id := c.Param("id")

	// cari Customer berdasarkan id
	if err := database.DB.First(&Customer, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Customer not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Customer info", Customer)
}

func UpdateCustomerByID(c *gin.Context) {
	var Customer customer.Customer

	// ambil id dari parameter
	id := c.Param("id")

	// cari Customer berdasarkan id
	if err := database.DB.First(&Customer, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Customer not found")
		return
	}

	// Binding input JSON ke Customer
	if err := c.ShouldBindJSON(&Customer); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Force override ID
	Customer.ID, _ = strconv.Atoi(id)

	// Update Customer ke dalam database
	if err := database.DB.Omit("ID").Save(&Customer).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to insert data into database")
		return
	}

	helpers.JSONResponse(c, http.StatusCreated, "Customer update successfuly", Customer)
}

func DeleteCustomerByID(c *gin.Context) {
	var Customer customer.Customer

	// Ambil id dari parameter
	id := c.Param("id")

	// Cari Customer berdasarkan id
	if err := database.DB.First(&Customer, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Customer not found")
		return
	}

	// Delete Customer
	if err := database.DB.Delete(&Customer).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Customer")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Customer delete successfuly", nil)
}
