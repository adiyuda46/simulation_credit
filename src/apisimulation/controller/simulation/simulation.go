package simulation

import (
	"log"
	"math"
	"net/http"
	"simulation/src/apisimulation/controller/utils"
	modelApp "simulation/src/apisimulation/model"

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
