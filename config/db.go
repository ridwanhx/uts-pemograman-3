package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()

	dsn := os.Getenv("SUPABASE_DSN")

	if (dsn == "") {
		log.Fatalf("SUPABASE_DSN tidak ditemukan. Pastikan SUPABASE_DSN ada di file .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if (err != nil) {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	DB = db

	fmt.Println("Koneksi ke PostgreSQL (Supabase) berhasil.")
}


func GetDB() *gorm.DB {
	if (DB == nil) {
		log.Fatal("DB belum diinisialisasi. Panggil config.InitDB() terlebih dahulu.")
	}

	return DB;
}