package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Connect to database
	dsn := "host=localhost port=5432 dbname=test-database user=postgres password=" // replace fields with your own database credentials
	conn, err := sql.Open("pgx", dsn)                                              // Opening connection to database with the driver 'pgx'
	if err != nil {                                                                // Error handling
		log.Fatalf("Unable to open database: %v\n", err)
	}
	defer conn.Close() // This db connection will stop when our program finishes not before that.

	// Test Connection
	err = conn.Ping() // Ping verifies a connection to the database is still alive, establishing a connection if necessary.
	if err != nil {   // Error handling
		log.Fatalln("Unable to ping database: ", err)
	}
	log.Println("Pinged Database!")

	// View Rows
	fmt.Printf("Initialising with these two entries.\n")
	err = viewRows(conn) // Only function call to view all data in terminal.
	if err != nil {      // Error handling
		log.Println("Unable to view database: ", err)
	}

	// Insert Row
	query := `INSERT INTO users (id, first_name, last_name, email) values ($1, $2, $3, $4)`
	// query is the SQL statement to perform in database using GO
	_, err = conn.Exec(query, 3, "Amit", "Nandi", "amit@ac.org") // Executing a query with 4 placeholders.
	if err != nil {                                              // Error handling
		log.Println("Unable to insert into database: ", err)
	}

	// View Rows
	err = viewRows(conn) // View all data again.
	if err != nil {      // Error handling
		log.Println("Unable to view database: ", err)
	}

	// Update Row
	query = `UPDATE users SET email = $1 WHERE id = $2`
	// query to update something in our database, where id matches with our given id.
	_, err = conn.Exec(query, "deba@slg.org", 2) // Executing query with 2 placeholders, email and id.
	if err != nil {                              // Error handling
		log.Println("Unable to update database: ", err)
	}

	// View Rows
	err = viewRows(conn) // Viewing all rows again.
	if err != nil {      // Error handling
		log.Println("Unable to view database: ", err)
	}

	// Delete Row
	query = `DELETE FROM users WHERE id = $1`
	// query to delete a specific row fromd database where id matches.
	_, err = conn.Exec(query, 1) // Executing the specific query to delete thw row whose id is 1.
	if err != nil {              // Error handling
		log.Println("Unable to execute delete query.")
	}

	// View Rows
	viewRows(conn)

	// View Row by ID
	var id int
	var first_name, last_name, email string
	// declaring some variables to hold data and will come from database
	query = `SELECT id, first_name, last_name, email FROM users WHERE id = $1`
	// SELECT will show us data like id, first, last and email FROM a table called users WHERE id is placeholder.
	err = conn.QueryRow(query, 3).Scan(&id, &first_name, &last_name, &email)
	// QueryRow executes a query that is expected to return at most one row.
	// We are scaning our returned row to extract the values that came.
	fmt.Println("QUERYROW \t", first_name, last_name, email) // Printing everything in terminal.

	// Close Connection
}

func viewRows(conn *sql.DB) error {

	// Remember to use `` instead of '' or "" for multiline strings
	// Query to select all rows from table, SQL language is used here.
	query := `SELECT id, first_name, last_name, email FROM users ORDER BY id;`

	rows, err := conn.Query(query) // Query function will execute the query and return rows and error if any.
	if err != nil {                // Error handling
		log.Println("Unable to query result.")
		return err
	}

	// We will use defer to close the rows after the function finishes.
	defer rows.Close()

	var id int8
	var first_name, last_name, email string

	// Rows is like a list of row and looping through it to extract values from all indivisual row.
	for rows.Next() { // Looping through all rows
		// scanning the rows and storing in variables
		err := rows.Scan(&id, &first_name, &last_name, &email)
		if err != nil {
			log.Println("Something went wrong while scanning rows.")
			return err
		}
		fmt.Printf("ID: %v, FIRST: %v, LAST: %v, EMAIL: %v\n", id, first_name, last_name, email)
	}
	fmt.Println("----------------------------------")
	if err = rows.Err(); err != nil { // Rechecking for error in rows.
		log.Println("Something went wrong while iterating over rows.")
		return err
	}

	return nil
}
