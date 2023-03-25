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
		os.Mkdir("sqlite", os.ModePerm)
		fmt.Println(err.Error())
	}
	file.Close()
}

func createTable(db *sql.DB) {
	fmt.Println("Create table...")
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT NOT NULL,
		"last_name" TEXT,
		"email" TEXT);`

	statement, err := db.Prepare(createUserTableSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec()

	createUserTableSQLSession := `CREATE TABLE IF NOT EXISTS sessions (
		"id" TEXT NOT NULL,		
		"user" TEXT NOT NULL,
		"expire" DATETIME DEFAULT current_timestamp);`

	statement, err = db.Prepare(createUserTableSQLSession)
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec()
	fmt.Println("table created")
}
