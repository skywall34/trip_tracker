package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

// Create the tables if they don't exist
func CreateTable(db *sql.DB) error {
	userTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        first_name TEXT,
        last_name TEXT,
        email TEXT
    );`

    tripTable := `CREATE TABLE IF NOT EXISTS trips (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        departure TEXT,
        arrival TEXT,
        departure_time INTEGER,
        arrival_time INTEGER,
        airline TEXT,
        flight_number TEXT,
        reservation TEXT,
        terminal TEXT,
        gate TEXT,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`

    _, err := db.Exec(userTable)
    if err != nil {
        return err
    }

    _, err = db.Exec(tripTable)
    if err != nil {
        return err
    }

    return nil
}