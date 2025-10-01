package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

func GenerateAutoMigration(db *gorm.DB) {
	println("Entering table creation")

	db.AutoMigrate(&Claim{})
	println("Table Created")
}
func MySQLConnectionHelper() *gorm.DB {

	_ = godotenv.Load(".env")
	username := os.Getenv("username")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	host := os.Getenv("host")
	port := os.Getenv("port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetTableInstance(db *gorm.DB) {
	//once.Do(func() { GenerateTable(db) })
	once.Do(func() { GenerateAutoMigration(db) })
}

func SaveClaimInfo(message string) {
	//store the claim info to db
	db := MySQLConnectionHelper()
	//defer db.Close()

	//parse the request body
	var claim Claim

	err := json.Unmarshal([]byte(message), &claim)
	if err != nil {
		log.Printf("Failed to parse claim info: %v", err)
	}

	//save the site info to db
	result := db.Create(&claim)
	if result.Error != nil {
		log.Printf("Failed to save claim info: %v", result.Error)
		return
	}
	log.Printf("Claim info saved successfully")

}
