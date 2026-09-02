package main

import (
	"database/sql"

	_ "github.com/meow-io/go-sqlcipher"
)

func main() {
	for _, driver := range sql.Drivers() {
		println(driver)
	}
}
