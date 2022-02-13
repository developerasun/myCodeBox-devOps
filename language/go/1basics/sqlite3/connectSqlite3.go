package main

import (
	"database/sql"
	"fmt"

	// _ : Create package-level variables and
	// execute the init function of that package.
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	// Create a goCaffeine.db file and open sqlite3 database.
	database, _ := sql.Open("sqlite3", "./goCaffeine.db")

	// Create a prepared statement for later quries or executions : Create
	statement, _ := database.Prepare("create table if not exists daily_coffee (date text, name text, amount int)")

	// Execute a prepared statement.
	statement.Exec()

	// Create a prepared statement : Insert
	statement, _ = database.Prepare("insert into daily_coffee values(?,?,?)")
	statement.Exec("July 26th", "jake", 55)
	statement.Exec("July 27th", "jake", 66)

	// Execute a query, typically Select.
	rows, _ := database.Query("select * from daily_coffee")

	fmt.Println(rows)

}
