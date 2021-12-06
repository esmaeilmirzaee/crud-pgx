package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", "host=192.168.101.2 port=5234 dbname=pgx user=pgdmn password=secret")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect %v.", err))
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	log.Println("Connect to database")

	if err = conn.Ping(); err != nil {
		log.Println("Cannot ping the database.", err)
	}
	log.Println("Pinged the database")
}

// `delete from users where id = $1`

// findARow
// `select id, first_name, last_name from users where id = $1`
func findARow(conn *sql.DB, query string, values ...string) error {
	var id int
	row := conn.QueryRow(query, values)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

// updateARow updates a row
// `updates users set first_name = $1 where first_name = $2`
func updateARow(conn *sql.DB, statement string, values ...string) error {
	if _, err := conn.Exec(statement, statement, values); err != nil {
		return err
	}

	return nil
}

// addRow inserts a row into db
// insert into users (first_name, last_name) values($1, $2);
func addARow(conn *sql.DB, query string, values ...string) error {
	if _, err := conn.Exec(query, values); err != nil {
		return err
	}

	return nil
}

// getAllRows prints all the row in the database
func getAllRows(conn *sql.DB, query string) error {
	rows, err := conn.Query(query)

	defer func(rows *sql.Rows){
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}(rows)

	var id int
	var firstName, lastName string

	for rows.Next() {
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			log.Println(err)
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}