package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func openDB() error {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func closeDB() error {
	return DB.Close()
}

func setupDB() error {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id integer not null primary key,
        title text,
        completed boolean default false,
        position integer
    );
    `

	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
