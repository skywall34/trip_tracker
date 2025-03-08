package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	m "github.com/skywall34/trip-tracker/internal/models"
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

func Createuser(db *sql.DB, user m.User) (int64, error) {
    result, err := db.Exec("INSERT INTO users (username, password, first_name, last_name, email) VALUES (?, ?, ?, ?, ?)",
        user.Username, user.Password, user.FirstName, user.LastName, user.Email)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func GetUser(db *sql.DB, id int) (m.User, error) {
    var user m.User
    row := db.QueryRow("SELECT id, username, password, first_name, last_name, email FROM users WHERE id = ?", id)
    err := row.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email)
    if err != nil {
        return user, err
    }
    return user, nil
}

func UpdateUser(db *sql.DB, user m.User) (int64, error) {
    result, err := db.Exec("UPDATE users SET username = ?, password = ?, first_name = ?, last_name = ?, email = ? WHERE id = ?",
        user.Username, user.Password, user.FirstName, user.LastName, user.Email, user.ID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

func DeleteUser(db *sql.DB, id int) (int64, error) {
    result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

func CreateTrip(db *sql.DB, trip m.Trip) (int64, error) {
    result, err := db.Exec("INSERT INTO trips (user_id, departure, arrival, departure_time, arrival_time, airline, flight_number, reservation, terminal, gate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
        trip.UserId, trip.Departure, trip.Arrival, trip.DepartureTime, trip.ArrivalTime, trip.Airline, trip.FlightNumber, trip.Reservation, trip.Terminal, trip.Gate)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func GetTrip(db *sql.DB, id int) (m.Trip, error) {
    var trip m.Trip
    row := db.QueryRow("SELECT id, user_id, departure, arrival, departure_time, arrival_time, airline, flight_number, reservation, terminal, gate FROM trips WHERE id = ?", id)
    err := row.Scan(&trip.ID, &trip.UserId, &trip.Departure, &trip.Arrival, &trip.DepartureTime, &trip.ArrivalTime, &trip.Airline, &trip.FlightNumber, &trip.Reservation, &trip.Terminal, &trip.Gate)
    if err != nil {
        return trip, err
    }
    return trip, nil
}

func updateTrip(db *sql.DB, trip m.Trip) error {
    query := `
        UPDATE trips
        SET user_id = ?, departure = ?, arrival = ?, departure_time = ?, arrival_time = ?, airline = ?, flight_number = ?, reservation = ?, terminal = ?, gate = ?
        WHERE id = ?;
    `
    _, err := db.Exec(query, trip.UserId, trip.Departure, trip.Arrival, trip.DepartureTime, trip.ArrivalTime, trip.Airline, trip.FlightNumber, trip.Reservation, trip.Terminal, trip.Gate, trip.ID)
    if err != nil {
        return fmt.Errorf("could not update trip: %v", err)
    }
    return nil
}

func deleteTrip(db *sql.DB, tripID int) error {
    query := "DELETE FROM trips WHERE id = ?;"
    _, err := db.Exec(query, tripID)
    if err != nil {
        return fmt.Errorf("could not delete trip: %v", err)
    }
    return nil
}