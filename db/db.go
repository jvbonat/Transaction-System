package db

import (
	"desafio-transacoes/models"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	loadingError := godotenv.Load()
    if loadingError != nil {
        log.Fatal("Error when loading .ENV file")
    }
	user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    hostname := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, hostname, port, dbname)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database")
		return err
	}

	return nil
}


func Migrate() {

	if err := DB.AutoMigrate(&models.Account{}); err != nil {
		log.Fatal("Error migrating Account:", err)
	} else {
		log.Println("Table Account created successfully")
	}

	if err := DB.AutoMigrate(&models.OperationsType{}); err != nil {
		log.Fatal("Error migrating OperationsType:", err)
	} else {
		log.Println("Table OperationsType created successfully")
	}

	if err := DB.AutoMigrate(&models.Transaction{}); err != nil {
		log.Fatal("Error migrating Transaction:", err)
	} else {
		log.Println("Table Transaction created successfully")
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	sqlDB.Close()
}
