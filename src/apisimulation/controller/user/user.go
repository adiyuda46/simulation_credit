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
	err := c.BindJSON(&input)
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

		// Generate JWT token
		token, err := utils.GenerateToken(100)
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
			return
		}
		// insert token ke db

		utils.ResponseSuccess(c, gin.H{"message": "Registration successful","Token": token})
	}

}


func Login(c *gin.Context)  {
	var input modelApp.Login
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	// validasi nomor hp
	validatePhone ,err := repository.CheckPhoneNumber(input.Phone)
	if err != nil {
		log.Printf("Nomor atau Password salah: %v", validatePhone)
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Nomor atau Password salah"})
        return
	}
	// validasi password
	validatePassword := utils.VerifyPassword(validatePhone, input.Password)
    if validatePassword != nil {
        log.Printf("Nomor atau Password salah: %v", validatePassword)
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Nomor atau Password salah"})
        return
    }

	//get user id
	userId ,err := repository.GetUserbyPhone(input.Phone)
	if err != nil {
		log.Println("data not found")
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Nomor atau Password salah"})
        return
	}


	// Generate JWT token
    token, err := utils.GenerateToken(userId.Id)
    if err != nil {
        log.Printf("Failed to generate token: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
        return
    }

	utils.ResponseSuccess(c, gin.H{"message": "login berhasil","Token":token})
}

// Fungsi untuk mendapatkan detail pengguna
func GetUserDetails(c *gin.Context) {
    phone := c.MustGet("phone").(string) // Ambil phone dari konteks
    userDetails := modelApp.Login{
        Phone: phone,
        Password: "koko", // Ganti dengan data sebenarnya dari database
    }

    c.JSON(http.StatusOK, userDetails)
}


func Tes(c *gin.Context)  {
	// validasi nomor hp
	validatePhone ,err := repository.Supabase()
	if err != nil {
		log.Printf("Nomor atau Password salah: %v", validatePhone)
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Nomor atau Password salah"})
        return
	}
	utils.ResponseSuccess(c, gin.H{"message": "login berhasil","nama":validatePhone})
}