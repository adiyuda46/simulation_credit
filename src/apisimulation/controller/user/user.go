package user

import (
	"log"
	"net/http"
	repository "simulation/src/apisimulation/controller/repo"
	"simulation/src/apisimulation/controller/utils"
	modelApp "simulation/src/apisimulation/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input modelApp.Register
	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	// cek email dan phone
	checkRegis, err := repository.RegisterCheck(input.Email, input.Phone)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	if checkRegis != 0 {
		log.Printf("email or phone number are already to use") // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Email atau Nomor sudah di gunakan di akun lain, silakan masukan email dan nomor yang belum di gunakan!")
		return
	} else {
		regisErr := repository.RegisterRepository(input.Name, input.Password, input.Email, input.Phone)
		if regisErr != nil {
			log.Printf("Registration failed: %v", regisErr) // Log error
			utils.ResponseError(c, http.StatusInternalServerError, "Registration failed: "+regisErr.Error())
			return
		}

		utils.ResponseSuccess(c, gin.H{"message": "Registration successful"})
	}

}
