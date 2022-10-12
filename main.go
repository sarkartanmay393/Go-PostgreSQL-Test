package main

import (
	"database/sql"
	"log"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	// Connect to database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test-database user=postgres password=")
	if err != nil {
		log.Fatalf("Unable to open database: %v\n", err)
	}
	defer conn.Close() // Close connection when main() completes.

	// Test Connection

	// View Rows

	// Insert Row

	// View Rows

	// Update Row

	// View Rows

	// Delete Row

	// View Rows

	// Close Connection

}
