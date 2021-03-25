package util

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234567"
	dbname   = "taskpad"
)

// GetDB method returns a DB instance
func GetDB() (*sql.DB, error) {
	//connectionString := "user=madhanganesh dbname=taskpad sslmode=disable"
	//connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	if connectionString == "" {
		return nil, errors.New("'POSTGRES_CONNECTION_STRING' environment variable not set")
	}
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	return conn, nil
}
