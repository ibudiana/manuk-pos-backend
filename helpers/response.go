package helpers

import "github.com/gin-gonic/gin"

type JSONResponseStruct struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSONResponse mengirimkan respons JSON umum
func JSONResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := JSONResponseStruct{
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}

// JSONError mengirimkan respons JSON untuk error
func JSONError(c *gin.Context, statusCode int, err interface{}) {
	var errorMessage string

	// Deteksi tipe data dan log
	switch e := err.(type) {
	case error:
		errorMessage = e.Error()
	case string:
		errorMessage = e
	default:
		errorMessage = "Unknown error"
	}

	c.JSON(statusCode, gin.H{
		"error": errorMessage,
	})
}
