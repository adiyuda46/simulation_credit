package product

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetAgrement(c *gin.Context) {
	// get id by token
	userID := c.MustGet("userID").(int)
	log.Printf("User ID from token: %d", userID)

	
}