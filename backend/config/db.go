package config

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDataBase() error {
	if _, err := os.Stat("../../app.db"); os.IsNotExist(err) {
		fmt.Println("Creating new database file...")
		db, err := sql.Open("sqlite3", "../../app.db")
		if err != nil {
			return fmt.Errorf("failed to open database: %v", err)
		}
		defer db.Close()

		sqlFile, err := os.ReadFile("./db.sql")
		if err != nil {
			return fmt.Errorf("failed to read SQL file: %v", err)
		}
		_, err = db.Exec("PRAGMA foreign_keys = ON;")
		if err != nil {
			return fmt.Errorf("failed to enable foreign keys: %v", err)
		}

		_, err = db.Exec(string(sqlFile))
		if err != nil {
			return fmt.Errorf("failed to execute SQL: %v", err)
		}

	}
	return nil
}
