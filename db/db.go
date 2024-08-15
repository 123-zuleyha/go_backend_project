package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
        Logger:logger.Default.LogModel(logger.Info),
    

	})
	if err != nil {
		log.Panic("Failed to connect database:%v", err)
	}


	//Modeller buraya eklenecek :)
	if err :=DB.AutoMigrate(); err !=nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Connection to databse")

}
