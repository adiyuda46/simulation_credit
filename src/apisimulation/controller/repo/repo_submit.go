package repo

import (
	"log"
	conn "simulation/src/apisimulation/controller/connection"
	"simulation/src/apisimulation/model"
	"strconv"
	"strings"
	"time"
)

func SubmitPengajuan(param model.MsubmitPengajuan, userId int, agrement string) (string, error) {
	// Bersihkan string input
	cleanAmountInstalment := strings.ReplaceAll(param.AmountIntalment, ".", "")
	cleanAmountInstalment = strings.ReplaceAll(cleanAmountInstalment, ",", ".") // Ubah koma menjadi titik untuk desimal

	cleanTotalAmount := strings.ReplaceAll(param.TotalAmopunt, ".", "")
	cleanTotalAmount = strings.ReplaceAll(cleanTotalAmount, ",", ".") // Ubah koma menjadi titik untuk desimal

	// Konversi string yang dibersihkan ke float64 terlebih dahulu, lalu ke integer jika perlu
	newAmountIntalment, err := strconv.ParseFloat(cleanAmountInstalment, 64)
	if err != nil {
		log.Printf("Kesalahan saat mengonversi AmountInstalment: %v", err)
		return "", err
	}

	newTotalAmopunt, err := strconv.ParseFloat(cleanTotalAmount, 64)
	if err != nil {
		log.Printf("Kesalahan saat mengonversi TotalAmount: %v", err)
		return "", err
	}

	// Konversi ke integer jika perlu (misalnya, pembulatan)
	newAmountIntalmentInt := int(newAmountIntalment)
	newTotalAmopuntInt := int(newTotalAmopunt)

	// Koneksi ke DB
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Koneksi database gagal: %v", errConn)
		return "", errConn
	}
	defer db.Close()

	// Tetapkan DueDate
	dueDate := time.Now().Format("2006-01-02")

	query := `INSERT INTO public."AGREMENT" (
		"USER_ID", "AGREMENT", "AMOUNT_INSTALMENT", "PRODUCT", "INSTALMENT", "DUE_DATE", "DTM_CRT", "DTM_UPD", "TOTAL_AMOUNT"
	) VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW(), $7);`

	_, err = db.Exec(query, userId, agrement, newAmountIntalmentInt, param.TypeProduct, param.Instalment, dueDate, newTotalAmopuntInt)
	if err != nil {
		log.Printf("Kesalahan saat memasukkan data: %v", err)
		return "Kesalahan saat memasukkan data", err
	}

	return "Data berhasil dimasukkan", nil
}