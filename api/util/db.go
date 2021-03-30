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
	port     = 64076
	user     = "postgres"
	password = "1234567"
	dbname   = "TaskPad"
)

// GetDB method returns a DB instance
func GetDB() (*sql.DB, error) {
	//connectionString := "user=madhanganesh dbname=taskpad sslmode=disable"
	//connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname)

	if connectionString == "" {
		return nil, errors.New("'POSTGRES_CONNECTION_STRING' environment variable not set")
	}
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	return conn, nil
}
