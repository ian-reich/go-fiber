package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Memuat variabel lingkungan dari file .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}

	// Mengambil nilai variabel lingkungan untuk konfigurasi koneksi
	mode := os.Getenv("DB_MODE")
	var host, port, user, password, dbName string
	if mode == "DEV" {
		decodedHost, _ := decodeParams(os.Getenv("DB_HOST_DEV"))
		decodedPort, _ := decodeParams(os.Getenv("DB_PORT_DEV"))
		decodedUser, _ := decodeParams(os.Getenv("DB_USER_DEV"))
		decodedPassword, _ := decodeParams(os.Getenv("DB_PASSWORD_DEV"))
		decodedDbName, _ := decodeParams(os.Getenv("DB_NAME_DEV"))
		host = decodedHost
		port = decodedPort
		user = decodedUser
		password = decodedPassword
		dbName = decodedDbName
	} else if mode == "TEST" {
		decodedHost, _ := decodeParams(os.Getenv("DB_HOST_TEST"))
		decodedPort, _ := decodeParams(os.Getenv("DB_PORT_TEST"))
		decodedUser, _ := decodeParams(os.Getenv("DB_USER_TEST"))
		decodedPassword, _ := decodeParams(os.Getenv("DB_PASSWORD_TEST"))
		decodedDbName, _ := decodeParams(os.Getenv("DB_NAME_TEST"))
		host = decodedHost
		port = decodedPort
		user = decodedUser
		password = decodedPassword
		dbName = decodedDbName
	} else {
		decodedHost, _ := decodeParams(os.Getenv("DB_HOST_PRO"))
		decodedPort, _ := decodeParams(os.Getenv("DB_PORT_PRO"))
		decodedUser, _ := decodeParams(os.Getenv("DB_USER_PRO"))
		decodedPassword, _ := decodeParams(os.Getenv("DB_PASSWORD_PRO"))
		decodedDbName, _ := decodeParams(os.Getenv("DB_NAME_PRO"))
		host = decodedHost
		port = decodedPort
		user = decodedUser
		password = decodedPassword
		dbName = decodedDbName
	}

	// Membentuk string koneksi DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbName)

	// Membuat koneksi database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	DB = db
	return DB, nil
}

// Decode Param
func decodeParams(param string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
