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
	err := t.db.QueryRow(`
		SELECT 
			id, 
			user_id, 
			departure, 
			arrival, 
			departure_time, 
			arrival_time, 
			airline, 
			flight_number, 
			reservation, 
			terminal, 
			gate 
		FROM trips WHERE id = ?`, id).Scan(
			&trip.ID, 
			&trip.UserId, 
			&trip.Departure,
			&trip.Arrival, 
			&trip.DepartureTime, 
			&trip.ArrivalTime, 
			&trip.Airline, 
			&trip.FlightNumber, 
			&trip.Reservation,
			&trip.Terminal, 
			&trip.Gate)
	if err != nil {
		return trip, err
	}
	return trip, nil
}

func (t *TripStore) GetTripsGivenUser(userID int) ([]m.Trip, error) {
    var trips []m.Trip
    rows, err := t.db.Query(`
		SELECT 
			t.id, 
			t.user_id, 
			t.departure, 
			t.arrival, 
			t.departure_time, 
			t.arrival_time, 
			t.airline, 
			t.flight_number,
			t.reservation,
			t.terminal,
			t.gate,
			d.latitude, 
			d.longitude, 
			a.latitude, 
			a.longitude
        FROM trips t
        JOIN airports d ON t.departure = d.iata_code
        JOIN airports a ON t.arrival = a.iata_code
        WHERE t.user_id = ?`, userID)
    if err != nil {
        return trips, err
    }
    defer rows.Close()

    for rows.Next() {
        var trip m.Trip
        err := rows.Scan(
			&trip.ID, 
			&trip.UserId, 
			&trip.Departure, 
			&trip.Arrival, 
			&trip.DepartureTime, 
			&trip.ArrivalTime, 
			&trip.Airline, 
			&trip.FlightNumber, 
			&trip.Reservation, 
			&trip.Terminal, 
			&trip.Gate,
			&trip.DepartureLat,
			&trip.DepartureLon,
			&trip.ArrivalLat,
			&trip.ArrivalLon,
		)
        if err != nil {
            return trips, err
        }
        trips = append(trips, trip)
    }

    return trips, nil
}


func (t *TripStore) getFlightsPerMonthForYear(user_id int, year string) ([]m.FlightAggregation, error) {
	var flights []m.FlightAggregation

	rows, err := t.db.Query(`
		WITH months AS (
			SELECT '01' AS month UNION ALL
			SELECT '02' UNION ALL
			SELECT '03' UNION ALL
			SELECT '04' UNION ALL
			SELECT '05' UNION ALL
			SELECT '06' UNION ALL
			SELECT '07' UNION ALL
			SELECT '08' UNION ALL
			SELECT '09' UNION ALL
			SELECT '10' UNION ALL
			SELECT '11' UNION ALL
			SELECT '12'
		),
		trip_counts AS (
			SELECT 
				strftime('%m', datetime(departure_time, 'unixepoch')) AS month,
				COUNT(*) AS trip_count
			FROM trips
			WHERE user_id = 1
			AND strftime('%Y', datetime(departure_time, 'unixepoch')) = '2025'
			GROUP BY month
		)
		SELECT 
			m.month AS label,
			COALESCE(tc.trip_count, 0) AS count
		FROM months m
		LEFT JOIN trip_counts tc ON m.month = tc.month
		ORDER BY m.month
	`, user_id, year)

	if err != nil {
		return flights, err
	}
	defer rows.Close()

	for rows.Next() {
		var flight m.FlightAggregation
		err := rows.Scan(
			&flight.Label, 
			&flight.Count,
		)
		if err != nil {
			return flights, err
		}
		flights = append(flights, flight)
	}

	return flights, nil

}

func (t *TripStore) getAirlinesCountForYear(user_id int, year string) ([]m.AirlineAggregation, error) {
	var airlines []m.AirlineAggregation

	rows, err := t.db.Query(`
	SELECT 
		airline,                                                       
		COUNT(*) AS flight_count
	FROM trips
	WHERE user_id = ?
		AND strftime('%Y', datetime(departure_time, 'unixepoch')) = ?
	GROUP BY airline
	ORDER BY flight_count DESC`, user_id, year)

	if err != nil {
		return airlines, err
	}
	defer rows.Close()

	for rows.Next() {
		var airline m.AirlineAggregation
		err := rows.Scan(
			&airline.Label,
			&airline.Count,
		)
		if err != nil {
			return airlines, err
		}
		airlines = append(airlines, airline)
	}

	return airlines, nil
}

func (t *TripStore) getCountriesCountForYear(user_id int, year string) ([]m.CountryAggregation, error) {
	var countries []m.CountryAggregation

	rows, err := t.db.Query(`
	SELECT 
		strftime('%m', datetime(t.departure_time, 'unixepoch')) AS label,
		d.country AS country,
		COUNT(DISTINCT t.arrival) AS country_count
	FROM trips t
	JOIN airports d ON t.arrival = d.iata_code
	WHERE user_id = ?
		AND strftime('%Y', datetime(t.departure_time, 'unixepoch')) = ?
	GROUP BY label
	ORDER BY label`, user_id, year)

	if err != nil {
		return countries, err
	}
	defer rows.Close()

	for rows.Next() {
		var country m.CountryAggregation
		err := rows.Scan(
			&country.Label,
			&country.Country,
			&country.Count,
		)
		if err != nil {
			return countries, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}


func (t *TripStore) GetTripsPerAggregation(user_id int, year string, agg string) ([]m.FlightAggregation, []m.AirlineAggregation, []m.CountryAggregation, error) {
	// TODO: switch case for agg
	flights, err := t.getFlightsPerMonthForYear(user_id, year)
	if err != nil {
		return nil, nil, nil, err
	}
	
	airlines, err := t.getAirlinesCountForYear(user_id, year)
	if err != nil {
		return nil, nil, nil, err
	}

	countries, err := t.getCountriesCountForYear(user_id, year)
	if err != nil {
		return nil, nil, nil, err
	}

	return flights, airlines, countries, nil
}

func (t *TripStore) DeleteTrip(id int) (error) {
	_, err := t.db.Exec("DELETE FROM trips WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}