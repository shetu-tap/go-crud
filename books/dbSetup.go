package books

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetDBString(host string, name string, username string, password string, port string) string {
	return fmt.Sprintf("host=%v database=%v user=%v password=%v port=%v sslmode=disable TimeZone=Asia/Dhaka", host, name, username, password, port)
}

func ConnectToDB() *gorm.DB {
	// load the env
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load env file")
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USERNAME := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DSN_STRING := GetDBString(DB_HOST, DB_NAME, DB_USERNAME, DB_PASSWORD, DB_PORT)

	log.Println(fmt.Sprintf("DSN String: %v", DSN_STRING))

	db, err := gorm.Open(postgres.Open(DSN_STRING), &gorm.Config{})
	if err != nil {
		panic("Error in database connection establishment")
	}

	db.AutoMigrate(&Book{})
	return db
}