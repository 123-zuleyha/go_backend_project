package db

import (
	"fmt"
	"log"
	"os"

	"github.com/123-zuleyha/go_backend_project/db/seeders"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || password == "" || database == "" || port == "" {
		log.Fatalf("One or more database environment variables are not set. Please check DB_HOST, DB_USER, DB_PASSWORD, DB_DATABASE, and DB_PORT.")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, database, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Modeller buraya eklenecek :)
	if err := DB.AutoMigrate(); err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Connection to database established successfully")
}

func Seed() error {
	userTypeSeeder := seeders.UserTypeSeeder{}
	userSeeder := seeders.UserSeeder{}

	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(userTypeSeeder.Run()).Error
	if err != nil {
		return err
	}

	err = DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(userSeeder.Run()).Error
	return err
}
