package simulation

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	repository "simulation/src/apisimulation/controller/repo"
	"simulation/src/apisimulation/controller/utils"
	modelApp "simulation/src/apisimulation/model"
	"time"

	"github.com/gin-gonic/gin"
)

func Simulation(c *gin.Context) {
	var input modelApp.MdownPaymemt
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}
	// Ubah int ke float
	var (
		newPrice         float64 = float64(input.Price) // Harga motor
		newDp            float64 = float64(input.DP) / 100.0 // Persentase DP
		tenor            int     = input.Tenor // Tenor dalam bulan
		sukuBungaTahunan float64 = 0.12 // Suku bunga tahunan
	)
	// Hitungan
	finalDp := newPrice * newDp
	jumlahPembiayaan := newPrice - finalDp
	sukuBungaBulanan := sukuBungaTahunan / 12
	// Menghitung cicilan bulanan
	cicilanBulanan := hitungCicilanBulanan(jumlahPembiayaan, sukuBungaBulanan, tenor)
	// Format rupiah
	formattedDp := utils.FormatRupiah(finalDp)
	formattedCicilan := utils.FormatRupiah(cicilanBulanan)
	// Hitung total pembayaran
	totalPembayaran := cicilanBulanan * float64(tenor)
	formattedTotalPembayaran := utils.FormatRupiah(totalPembayaran)
	// Hitung total bunga
	totalBunga := totalPembayaran - jumlahPembiayaan
	formattedTotalBunga := utils.FormatRupiah(totalBunga)

	// Pesan sukses
	utils.ResponseSuccess(c, gin.H{
		"message":               "Simulation calculated successfully.",
		"Total Down Payment":    formattedDp,
		"Cicilan Bulanan":       formattedCicilan,
		"Total Pembayaran":      formattedTotalPembayaran,
		"Total Bunga":           formattedTotalBunga,
	})
}

// Fungsi untuk menghitung cicilan bulanan
func hitungCicilanBulanan(jumlahPembiayaan float64, sukuBungaBulanan float64, tenor int) float64 {
	return (jumlahPembiayaan * sukuBungaBulanan * math.Pow(1+sukuBungaBulanan, float64(tenor))) / (math.Pow(1+sukuBungaBulanan, float64(tenor)) - 1)
}

func SubmitPengajuan(c *gin.Context) {
	var input modelApp.MsubmitPengajuan
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("Invalid input data: %v", err) // Log error
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}

	// Get user ID from token
	userID := c.MustGet("userID").(int)

	// Generate agreement number based on product type
	var agrement string
	switch input.TypeProduct {
	case "MOTOR BARU":
		agrement = generateAgreementNumber("NMB")
	case "MOTOR BEKAS":
		agrement = generateAgreementNumber("MB")
	case "MOBIL":
		agrement = generateAgreementNumber("CAR")
	case "MULTYPORDUCT":
		agrement = generateAgreementNumber("MP")
	default:
		utils.ResponseError(c, http.StatusBadRequest, "Invalid product type")
		return
	}

	// Insert into the database
	result, err := repository.SubmitPengajuan(input, userID, agrement)
	if err != nil {
		log.Printf("Error inserting record: %v", err)
		utils.ResponseError(c, http.StatusInternalServerError, "Error inserting record")
		return
	}

	utils.ResponseSuccess(c, gin.H{"Message": result, "AgreementNumber": agrement})
}

// generateAgreementNumber creates a unique agreement number
func generateAgreementNumber(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100000) // Generate a number between 0 and 99999
	return fmt.Sprintf("%s%05d", prefix, number) // Format to include prefix and 5 digits
}