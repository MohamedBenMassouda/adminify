package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(fileName ...string) {
	// Load the .env file
	err := godotenv.Load(fileName...)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func ConnectDB(driver string) (*sql.DB, error) {
	// Get the database URL from environment variable
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	db, err := sql.Open(driver, dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
