package main

import (
	"log"
	"os"

	"manuk-pos-backend/database"
	"manuk-pos-backend/routes"

	"github.com/gin-contrib/cors"
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

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	serverAddr := os.Getenv("SERVER_PORT")

	log.Printf("Server berjalan di http://localhost%s", serverAddr)
	r.Run(serverAddr)
}
