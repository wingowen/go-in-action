// Sample program to show how to show you how to briefly work
// with the sql package.
package main

import (
	"database/sql"

	_ "go-in-action/chapter3/dbdriver/postgres"
)

// main is the entry point for the application.
func main() {
	sql.Open("postgres", "mydb")
}
