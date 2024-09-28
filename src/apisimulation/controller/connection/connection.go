package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func ConnectToDatabase() (*sql.DB, error) {
	var (
		server   = viper.GetString("simulation.server")
		port     = viper.GetString("simulation.portdb")
		user     = viper.GetString("simulation.user")
		password = viper.GetString("simulation.password")
		scheme   = viper.GetString("simulation.scheme")
	)
	// Konfigurasi koneksi database
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		server, port, user, password, scheme)

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
