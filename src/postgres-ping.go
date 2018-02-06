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
	if err != nil {
		panic(err)
	}

	dbname := args[3]
	user := args[4]
	password := args[5]

	message := fmt.Sprintf("Checking connection to %s:%d/%s with user %s ",
		host, port, dbname, user)
	fmt.Println(message)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select 1")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)

	}
	err = rows.Err() // get any error encountered during iteration
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	db.Close()
}
