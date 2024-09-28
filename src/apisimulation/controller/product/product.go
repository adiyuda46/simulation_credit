package product

import (
	"log"
	"net/http"
	repository "simulation/src/apisimulation/controller/repo"
	"simulation/src/apisimulation/controller/utils"
modelApp "simulation/src/apisimulation/model"
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
	var input modelApp.GetLob
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	// get lob by id
	lob ,err := repository.GetLobByid(input.Id)
	if err != nil {
		log.Println("data not found", err) // Log error
		utils.ResponseError(c, http.StatusNotFound, "data not found: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, gin.H{"message": "Data di Temukan", "LOB product": lob})
}