package server

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

const dsn string = "file:db.sq3"
const readQuery string = "SELECT * FROM COUNTER;"
const resetDB string = "UPDATE COUNTER SET count = 0;"

var conn *sql.DB = nil

func ReadDB() (int, error) {
	// If we don't already have a connection to the DB try to initialize it
	if conn == nil {
		err := connect()
		if err != nil {
			return -1, errors.New(fmt.Sprintf("Error connecting to DB: %s", err))
		}
	}

	rows, err := conn.Query(readQuery, "")
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Error querying database: %s", err))
	}
	/* close is idempotent and doesn't affect the result of err according to the docs.
	We'll just handle this to keep the warnings clean and MAYBE get something out of it if it ever
	fails.  This is just a toy app*/
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	countRows := make([]int, 0)
	for rows.Next() {
		var count int
		err := rows.Scan(&count)
		if err != nil {
			return -1, errors.New(fmt.Sprintf("Error parsing data from DB: %s\n possible DB corruption", err))
		}
		countRows = append(countRows, count)
	}
	if err := rows.Err(); err != nil {
		return -1, errors.New(fmt.Sprintf("Error retrieving data from DB: %s", err))
	}
	if len(countRows) > 1 {
		countErr := errors.New("more than one row found in COUNTER table, possible database corruption")
		return -1, countErr
	} else {
		return countRows[0], nil
	}
}

func WriteDB(newVal int) error {
	if conn == nil {
		err := connect()
		if err != nil {
			return errors.New(fmt.Sprintf("Error connecting to DB: %s", err))
		}
	}
	statement := fmt.Sprintf("UPDATE COUNTER SET COUNT = %s", strconv.Itoa(newVal))
	_, err := conn.Exec(statement, "")
	if err != nil {
		return err
	}
	return nil
}

func ResetDB() error {
	if conn == nil {
		err := connect()
		if err != nil {
			return errors.New(fmt.Sprintf("Error connecting to DB: %s", err))
		}
	}
	_, err := conn.Exec(resetDB, "")
	if err != nil {
		return err
	}
	return nil
}

func connect() error {
	/* opening a nonexistent DB will create one without generating errors.  This DB won't have the table we need.
	We would want to guard against this in a real app */

	// make sure the connection is working, if not we'll fall through and do the connection work
	if conn != nil {
		err := conn.Ping()
		if err == nil {
			return nil
		}
	}
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	conn = db
	err = conn.Ping()
	if err != nil {
		// we had an error, so conn is probably bad, nil it out
		conn = nil
		return nil
	}
	return nil
}
