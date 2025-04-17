package helpers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret Key
var jwtSecret = []byte(os.Getenv("JWT_SECRETE"))

// Claims adalah struktur payload JWT
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// ValidateToken memvalidasi JWT dan mengembalikan payload (Claims)
func ValidateToken(tokenString string) (*Claims, error) {
	// Parsing token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Validasi token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GenerateToken membuat JWT baru
func GenerateToken(userID uint, role string) (string, error) {
	validTimeStr := os.Getenv("JWT_EXPIRES_IN_HOUR")
	validTime, err := strconv.Atoi(validTimeStr)
	if err != nil {
		return "", fmt.Errorf("invalid JWT_EXPIRES_IN_HOUR: %v", err)
	}
	// Isi payload token (Claims)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(validTime))), // Token valid 24 jam
		},
	}

	// Buat token menggunakan algoritma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangani token dengan secret key
	return token.SignedString(jwtSecret)
}
