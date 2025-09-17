package store

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionHelper() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	userName := os.Getenv("username")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbName := os.Getenv("dbname")

	dsn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	log.Println("Database connected successfully")
	//auto migrate the schema
	db.AutoMigrate(&Member{})
	log.Println("Database migrated successfully")
	return db

}


//save member data
