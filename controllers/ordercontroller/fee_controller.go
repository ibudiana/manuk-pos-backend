package ordercontroller

import (
	"net/http"
	"strconv"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/orders"

	"github.com/gin-gonic/gin"
)

func GetFees(c *gin.Context) {
	var fees []orders.Fee
	result := database.DB.Find(&fees)

	if result.RowsAffected == 0 {
		helpers.JSONResponse(c, http.StatusOK, "Fee is empty", fees)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "List Fee", fees)
}

func CreateFee(c *gin.Context) {
	var fee orders.Fee

	if err := c.ShouldBindJSON(&fee); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	tx := database.DB.Begin()

	if err := tx.Create(&fee).Error; err != nil {
		tx.Rollback()
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create Fee")
		return
	}

	tx.Commit()

	helpers.JSONResponse(c, http.StatusCreated, "Create Fee successfully", fee)
}

func GetFeeByID(c *gin.Context) {
	var fee orders.Fee
	id := c.Param("id")

	if err := database.DB.First(&fee, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Fee not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Fee info", fee)
}

func UpdateFeeByID(c *gin.Context) {
	var fee orders.Fee
	id := c.Param("id")

	if err := database.DB.First(&fee, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Fee not found")
		return
	}

	if err := c.ShouldBindJSON(&fee); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	fee.ID, _ = strconv.Atoi(id)

	if err := database.DB.Omit("ID").Save(&fee).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to update Fee")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Fee updated successfully", fee)
}

func DeleteFeeByID(c *gin.Context) {
	var fee orders.Fee
	id := c.Param("id")

	if err := database.DB.First(&fee, id).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Fee not found")
		return
	}

	if err := database.DB.Delete(&fee).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to delete Fee")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Fee deleted successfully", nil)
}
