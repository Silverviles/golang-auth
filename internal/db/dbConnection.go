package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func GetDatabaseConnection() *sql.DB {
	var configuration *DatabaseConfig
	plan, err := os.ReadFile("internal/config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(plan, &configuration)
	if err != nil {
		return nil
	}
	connectionInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.Host,
		configuration.Port,
		configuration.User,
		configuration.Password,
		configuration.DbName,
	)
	dbConnection, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	return dbConnection
}
