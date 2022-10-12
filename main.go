package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

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
	err = conn.Ping()
	if err != nil {
		log.Fatalln("Unable to ping database: ", err)
	}
	log.Println("Pinged Database!")

	// View Rows
	err = viewRows(conn)
	if err != nil {
		log.Println("Unable to view database: ", err)
	}

	// Insert Row
	query := `INSERT INTO users (id, first_name, last_name, email) values ($1, $2, $3, $4)`
	if 1 == 0 {
		fmt.Print(query)
	}
	_, err = conn.Exec(query, rand.Int(), "Amit", "Nandi", "amit@ac.org")
	if err != nil {
		log.Println("Unable to insert into database: ", err)
	}

	// View Rows
	err = viewRows(conn)
	if err != nil {
		log.Println("Unable to view database: ", err)
	}

	// Update Row
	query = `UPDATE users SET email = $1 WHERE id = $2`
	_, err = conn.Exec(query, "deba@slg.org", 2)
	if err != nil {
		log.Println("Unable to update database: ", err)
	}

	// View Rows
	err = viewRows(conn)
	if err != nil {
		log.Println("Unable to view database: ", err)
	}

	// Delete Row
	query = `DELETE FROM users WHERE id = $1`
	_, err = conn.Exec(query, 1)
	if err != nil {
		log.Println("Unable to execute delete query.")
	}

	// View Rows
	viewRows(conn)

	// View Row by ID
	var id int
	var first_name, last_name, email string
	query = `SELECT id, first_name, last_name, email FROM users WHERE id = $1`
	err = conn.QueryRow(query, 3).Scan(&id, &first_name, &last_name, &email)
	fmt.Println("QUERYROW \t", first_name, last_name, email)

	// Close Connection
}

func viewRows(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, first_name, last_name, email FROM users ORDER BY id;")
	if err != nil {
		log.Println("Unable to query result.")
		return err
	}
	defer rows.Close()

	var id int8
	var first_name, last_name, email string
	for rows.Next() {
		err := rows.Scan(&id, &first_name, &last_name, &email)
		if err != nil {
			log.Println("Seomthing went wrong while scaning rows.")
			return err
		}
		fmt.Printf("ID: %v, FIRST: %v, LAST: %v, EMAIL: %v\n", id, first_name, last_name, email)
	}
	fmt.Println("----------------------------------")
	if err = rows.Err(); err != nil {
		log.Println("Something went wrong while iterating over rows.")
		return err
	}

	return nil
}
