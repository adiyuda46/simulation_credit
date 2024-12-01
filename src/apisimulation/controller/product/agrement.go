package product

import (
	"net/http"
	repository "simulation/src/apisimulation/controller/repo"
	"simulation/src/apisimulation/controller/utils"

	"github.com/gin-gonic/gin"
)

func GetAgrement(c *gin.Context) {
	// get id by token
	
	userID := c.MustGet("userID").(int)

	result, err := repository.GetListaAgrement(userID)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, gin.H{"List Agrment": result})

}
