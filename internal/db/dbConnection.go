package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
		"%s:%s@tcp(%s:%s)/%s",
		configuration.User,
		configuration.Password,
		configuration.Host,
		configuration.Port,
		configuration.DbName,
	)
	dbConnection, err := sql.Open("mysql", connectionInfo)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	return dbConnection
}
