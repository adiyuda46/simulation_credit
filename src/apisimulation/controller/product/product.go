package product

import (
	"log"
	"net/http"
	repository "simulation/src/apisimulation/controller/repo"
	"simulation/src/apisimulation/controller/utils"

	"github.com/gin-gonic/gin"
)

func Lob(c *gin.Context) {
	// get all LOB
	result, err := repository.GetAllLob()
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, gin.H{"message": "Data di Temukan", "result": result})
}

func LobById(c *gin.Context) {
	
}
