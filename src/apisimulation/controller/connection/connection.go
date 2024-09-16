package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	// Konfigurasi koneksi database
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "masuksaja", "simulasi")

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		fmt.Println("Gagal terhubung ke database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Gagal melakukan ping ke database:", err)
		return nil, err
	}

	fmt.Println("Berhasil terhubung ke database!")
	return db, nil
}
