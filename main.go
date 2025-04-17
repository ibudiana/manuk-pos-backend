package main

import (
	"log"
	"os"

	"manuk-pos-backend/database"
	"manuk-pos-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Membuka koneksi ke database
	database.ConnectDB()
	// Menutup koneksi database
	defer database.CloseDB()

	// Buat tabel baru jika belum ada dan seed data
	database.MigrateTables()
	database.Seed()

	// Setup router
	r := routes.SetupRouter()

	serverAddr := os.Getenv("SERVER_PORT")

	log.Printf("Server berjalan di http://localhost%s", serverAddr)
	r.Run(serverAddr)
}
