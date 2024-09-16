package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess mengembalikan respons sukses
func ResponseSuccess(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "data":    data,
    })
}

// ResponseError mengembalikan respons error
func ResponseError(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, gin.H{
        "status":  "error",
        "message": message,
    })
}