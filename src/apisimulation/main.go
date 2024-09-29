package main

import (
	"fmt"
	"log"
	Connect "simulation/src/apisimulation/controller/connection"
	router "simulation/src/apisimulation/controller/router"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("src/config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
}

func main() {
	// init ke database
	db, err := Connect.ConnectToDatabase()
	if err != nil {
		// handle error jika gagal terhubung ke database
		fmt.Println("Terjadi kesalahan:", err)
		return
	}
	defer db.Close()
	fmt.Println("Terhubung ke server PostgreSQL!")
	addr := ":8080" // port local host
	host := "http://localhost"

	// Membuat router
	r := router.Router()
	// init router
	router.InitRouter(r)

	// Menjalankan server HTTP
	fmt.Println("Server berjalan di", host)
	if err := r.Run(addr); err != nil {

		fmt.Println("Server berjalan di port :", host)
		fmt.Println("Gagal menjalankan server:", err)
		log.Fatal("Gagal menjalankan server:", err)
	}
}
