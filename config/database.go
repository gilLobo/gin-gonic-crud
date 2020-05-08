package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DB ...
var DB *gorm.DB

// DBConfig ...
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// BuildDBConfig ...
func BuildDBConfig() *DBConfig {

	envFile := godotenv.Load()

	if envFile != nil {
		fmt.Print(envFile)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConfig := DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     username,
		DBName:   dbName,
		Password: password,
	}
	return &dbConfig
}

// DbURL ...
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}
