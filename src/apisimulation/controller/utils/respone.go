package utils

import (
	"fmt"
	"net/http"
	"strconv"

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

// FormatRupiahSimple memformat angka dengan titik dan koma
func FormatRupiah(amount float64) string {
    // Ubah ke string dengan dua desimal
    amountStr := strconv.FormatFloat(amount, 'f', 2, 64)

    // Pisahkan bagian integer dan desimal
    whole := amountStr[:len(amountStr)-3] // Bagian integer
    decimal := amountStr[len(amountStr)-3:] // Bagian desimal

    // Tambahkan titik sebagai pemisah ribuan
    formattedWhole := ""
    for i, digit := range whole {
        if i > 0 && (len(whole)-i)%3 == 0 {
            formattedWhole += "."
        }
        formattedWhole += string(digit)
    }

    // Gabungkan bagian integer dan desimal
    return fmt.Sprintf("%s,%s", formattedWhole, decimal[1:]) // Menghapus "0."
}