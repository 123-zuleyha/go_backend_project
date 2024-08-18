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

	var err error

	dsn := fmt.Sprintf("host:%s user=%s password=%s database=%s port=%s", host, user, password, database, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panic("Failed to connect database:%v", err)
	}

	//Modeller buraya eklenecek :)
	if err := DB.AutoMigrate(); err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Connection to databse")

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
