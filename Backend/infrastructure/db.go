package infrastructure

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("POSTGRES_URI")
	if dsn == "" {
		log.Fatal("POSTGRES_URI environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Test the connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get database instance: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	log.Println("✅ Connected to PostgreSQL!")
}
