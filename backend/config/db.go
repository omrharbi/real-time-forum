package config

import (
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	Connection *sql.DB
}

func InitDataBase() (*Database, error) {
	if _, err := os.Stat("../../app.db"); os.IsNotExist(err) {
		fmt.Println("Creating new database file...")
	}
	db, err := sql.Open("sqlite3", "../../app.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %v", err)
	}

	sqlFile, err := os.ReadFile("./db.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL file: %v", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL: %v", err)
	}

	return &Database{Connection: db}, nil
}

func (db *Database) Close() error {
	if db.Connection != nil {
		return db.Connection.Close()
	}
	return nil
}
