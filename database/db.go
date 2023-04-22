package database

import (
	"fmt"
	"h8-assignment-final-project/helpers"
	"h8-assignment-final-project/models"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {

	log.Println("[START] Connecting to database...")

	helpers.LoadEnv()

	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	helpers.LogIfError(err, "Error converting port to integer")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[ERROR] Failed to connect to database!")
	}

	log.Println("[SUCCESS] Connected to database...")

	if os.Getenv("DB_MIGRATE") == "true" {
		log.Println("[START] Migrating database...")
		err := db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
		helpers.LogIfError(err, "Error migrating database")
		log.Println("[SUCCESS] Migrated database...")
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	helpers.LogIfError(err, "Error closing database connection")

	dbSQL.Close()
}
