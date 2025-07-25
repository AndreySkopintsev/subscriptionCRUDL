package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func Init() {
	dbConnection, err := Connect()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	PostgresDB = dbConnection
}

func Connect() (*sql.DB, error) {
	log.Println("about to connect to db")
	host := getEnv("POSTGRES_HOST", "postgres")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "123")
	password := getEnv("POSTGRES_PASSWORD", "123")
	dbname := getEnv("POSTGRES_DB", "subscriptions")

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("pinged db successfully")

	return db, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
