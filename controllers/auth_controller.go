package controllers

import (
	"log"
	"net/http"

	"manuk-pos-backend/database"
	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/user"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var auth_user user.UserInput
	var existingUser user.User

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&auth_user); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err)
		return
	}

	// Cek apakah username sudah terdaftar
	if err := database.DB.Where("username = ?", auth_user.Username).First(&existingUser).Error; err == nil {
		helpers.JSONError(c, http.StatusConflict, "Username already exists")
		return
	} else {
		log.Println("Username available")
	}

	// Cek apakah email sudah terdaftar
	if err := database.DB.Where("email = ?", auth_user.Email).First(&existingUser).Error; err == nil {
		helpers.JSONError(c, http.StatusConflict, "Email already exists")
		return
	} else {
		log.Println("Email available")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(auth_user.Password)
	if err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Konversi ke model user.User
	newUser := user.User{
		Username: auth_user.Username,
		Password: hashedPassword,
		Email:    auth_user.Email,
	}

	// Simpan ke database
	if err := database.DB.Create(&newUser).Error; err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	newUser.Password = ""
	helpers.JSONResponse(c, http.StatusCreated, "User created", newUser)
}

func LoginUser(c *gin.Context) {
	var auth_user user.User
	var login_user user.UserLogin

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&login_user); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err)
		return
	}

	// Cari user berdasarkan email atau username
	if err := database.DB.Preload("Role").
		Where("email = ? OR username = ?", login_user.UsernameOrEmail, login_user.UsernameOrEmail).
		First(&auth_user).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "User not found")
		return
	}

	// Cek password
	if !helpers.CheckPassword(login_user.Password, auth_user.Password) {
		helpers.JSONError(c, http.StatusNotFound, "Invalid password")
		return
	}

	// Generate token
	token, err := helpers.GenerateToken(uint(auth_user.ID), auth_user.Role.Name)
	if err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "token", token)
}
