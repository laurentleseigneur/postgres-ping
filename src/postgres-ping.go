package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args
	argsCount := len(args)
	if argsCount != 6 {
		panic("wrong parameters count. Expected parameters are: host, port, dbname, user, password")
	}
	host := args[1]
	port, err := strconv.ParseInt(args[2], 10, 0)
	checkAndPanic(err)
	dbname := args[3]
	user := args[4]
	password := args[5]
	fmt.Printf("Checking connection to %s:%d/%s with user %s\n", host, port, dbname, user)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", connStr)
	checkAndPanic(err)
	defer closeDb(db)

	rows, err := db.Query("select 1")
	checkAndPanic(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
	}
	err = rows.Err() // get any error encountered during iteration
	checkAndPanic(err)
	fmt.Println("Successfully connected!")
}

func checkAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func closeDb(db *sql.DB) {
	err := db.Close()
	checkAndPanic(err)
	println("DB Connection closed.")
}
