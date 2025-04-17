package database

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var DB *gorm.DB

// ConnectDB menghubungkan ke database MySQL
func ConnectDB() *gorm.DB {
	if DB == nil {
		once.Do(func() {
			var err error
			dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
			DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatal("Failed connect to database:", err)
			} else {
				log.Println("Success connect to database")
			}
		})
	} else {
		log.Println("Database already connected")
	}
	return DB
}

// CloseDB - Menutup koneksi database
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database:", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal("Failed to close the database connection:", err)
	} else {
		log.Println("Database connection closed successfully")
	}
}
