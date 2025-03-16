package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	m "github.com/skywall34/trip-tracker/internal/models"
)


type TripStore struct {
	db *sql.DB
}

type NewTripStoreParams struct {
	DB *sql.DB
}

func NewTripStore(params NewTripStoreParams) *TripStore {
	return &TripStore{db: params.DB}
}

func (t *TripStore) CreateTrip(newTrip m.Trip) (int64, error) {
	stmt, err := t.db.Prepare("INSERT INTO trips (user_id, departure, arrival, departure_time, arrival_time, airline, flight_number, reservation, terminal, gate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		newTrip.UserId, 
		newTrip.Departure, 
		newTrip.Arrival, 
		newTrip.DepartureTime, 
		newTrip.ArrivalTime, 
		newTrip.Airline, 
		newTrip.FlightNumber, 
		newTrip.Reservation, 
		newTrip.Terminal, 
		newTrip.Gate,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TripStore) GetTrip(id int) (m.Trip, error) {
	var trip m.Trip
	err := t.db.QueryRow("SELECT id, user_id, departure, arrival, departure_time, arrival_time, airline, flight_number, reservation, terminal, gate FROM trips WHERE id = ?", id).Scan(&trip.ID, &trip.UserId, &trip.Departure, &trip.Arrival, &trip.DepartureTime, &trip.ArrivalTime, &trip.Airline, &trip.FlightNumber, &trip.Reservation, &trip.Terminal, &trip.Gate)
	if err != nil {
		return trip, err
	}
	return trip, nil
}

func (t *TripStore) GetTripsGivenUser(userID int) ([]m.Trip, error) {
    var trips []m.Trip
    rows, err := t.db.Query("SELECT id, user_id, departure, arrival, departure_time, arrival_time, airline, flight_number, reservation, terminal, gate FROM trips WHERE user_id = ?", userID)
    if err != nil {
        return trips, err
    }
    defer rows.Close()

    for rows.Next() {
        var trip m.Trip
        err := rows.Scan(&trip.ID, &trip.UserId, &trip.Departure, &trip.Arrival, &trip.DepartureTime, &trip.ArrivalTime, &trip.Airline, &trip.FlightNumber, &trip.Reservation, &trip.Terminal, &trip.Gate)
        if err != nil {
            return trips, err
        }
        trips = append(trips, trip)
    }

    return trips, nil
}

func (t *TripStore) DeleteTrip(id int) (error) {
	_, err := t.db.Exec("DELETE FROM trips WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}