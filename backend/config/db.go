package config

import (
	"database/sql"
	"log"
	"os"
)

type Database struct {
	Connection *sql.DB
}

func InitDataBase() error {
	db, err := sql.Open("sqlite3", "../../app.db")
	if err != nil {
		log.Fatal("error opening database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	query, err := os.ReadFile("../config/db.sql")
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	return nil
}

func Config() *Database {
	db, err := sql.Open("sqlite3", "../../app.db")
	if err != nil {
		log.Fatal("error opening database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	return &Database{Connection: db}
}

func (db *Database) Close() error {
	if db.Connection != nil {
		return db.Connection.Close()
	}
	return nil
}
