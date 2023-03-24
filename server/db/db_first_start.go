package db

import (
	"database/sql"
	"fmt"
	"os"
)

// USED ONLY TO CREATE DATABASE
func createDB(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	file.Close()
}

func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE users (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT NOT NULL,
		"last_name" TEXT,
		"email" TEXT
	  );`

	fmt.Println("Create table...")
	statement, err := db.Prepare(createUserTableSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec()
	fmt.Println("table created")
}
