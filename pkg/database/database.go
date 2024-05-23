package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes and returns a MySQL database connection
func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Successfully connected to the database.")
	return db
}
