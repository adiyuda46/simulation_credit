package product

import (
	"log"
	"net/http"
	repository "simulation/src/apisimulation/controller/repo"
	"simulation/src/apisimulation/controller/utils"
	modelApp "simulation/src/apisimulation/model"

	"github.com/gin-gonic/gin"
)

func MasterMultiproduct(c *gin.Context) {
	// get id by token
	userID := c.MustGet("userID").(int)
	log.Printf("User ID from token: %d", userID)

	// Ambil kategori dari database
	category, err := repository.GetCatMultiproduct()
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, gin.H{"message": "Data ditemukan", "category": category})
}

func ProductMultiproduct(c *gin.Context) {
	var input modelApp.ProductName
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	productName, err := repository.GetProductMultiproduct(input.Category)
	if err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Data not found: "+err.Error())
		return
	}
	
	if len(productName) == 0 {
		utils.ResponseError(c, http.StatusNotFound, "Data not found for category: "+input.Category)
		return
	}

	utils.ResponseSuccess(c, gin.H{"message": "Data ditemukan", "product name": productName})
}

func PriceMultiproduct(c *gin.Context) {
	var input modelApp.Price
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	price, err := repository.GetPriceMultiproduct(input)
	if err != nil {
		utils.ResponseError(c, http.StatusNotFound, "data not found: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, gin.H{"message": "Data ditemukan", "price product": price})
}
