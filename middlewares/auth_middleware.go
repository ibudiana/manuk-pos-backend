package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"manuk-pos-backend/helpers"

	"slices"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware memvalidasi JWT di header Authorization
func AuthMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Ambil token (format: "Bearer <token>")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validasi token
		claims, err := helpers.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set UserID dan Role di context
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		// Check if the user's role is in the allowed roles array
		roleValid := slices.Contains(roles, claims.Role)

		// If the role is not valid, return an unauthorized error
		if !roleValid {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			fmt.Println((claims.Role))
			c.Abort()
			return
		}

		// If the role is "customer", ensure that the user is only allowed to update their own data
		if claims.Role == "Customer" {
			userIDParam := c.Param("id")
			userID, err := strconv.Atoi(userIDParam) // Convert string to integer
			if err != nil || userID != int(claims.UserID) {
				c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own data"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
