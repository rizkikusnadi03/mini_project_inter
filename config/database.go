package config

import (
	"fmt"
	"log"
	"os"

	"backend_golang/internal/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected to Database")

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Store{},
		&domain.Address{},
		&domain.Category{},
		&domain.Product{},
		&domain.Transaction{},
		&domain.ProductLog{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database. \n", err)
	}

	return db
}
