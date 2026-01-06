package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Use environment variables in production
	dsn := "host=localhost user=admin password=password dbname=warroom_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("🚀 Connected to PostgreSQL database successfully")
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})
	if err != nil {
		log.Fatal("❌ Database Migration Failed: ", err)
	}
	fmt.Println("✅ Database Migration Completed: Users, Wallets, Transactions")
}
