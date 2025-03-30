package database

import (
	"database/sql"
	"errors"

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

func (t *TripStore) getFlightsForYears(user_id int) ([]m.FlightAggregation, error) {
	var flights []m.FlightAggregation
	var total int

	rows, err := t.db.Query(`
		WITH RECURSIVE years(label) AS (
                SELECT strftime('%Y', datetime('now', '-10 years')) AS label
                UNION ALL
                SELECT CAST(CAST(label AS INTEGER) + 1 AS TEXT)
                FROM years
                WHERE CAST(label AS INTEGER) < CAST(strftime('%Y', 'now') AS INTEGER)
        ),
        trips_by_year AS (
                SELECT
					strftime('%Y', datetime(departure_time, 'unixepoch')) AS label,
					COUNT(*) AS count
                FROM trips
                WHERE user_id = ?
                AND departure_time >= strftime('%s', datetime('now', '-10 years'))
                GROUP BY label
        )
		SELECT 
			years.label AS label,
			COALESCE(trips_by_year.count, 0) AS count
		FROM years
		LEFT JOIN trips_by_year ON years.label = trips_by_year.label
		ORDER BY CAST(years.label as INTEGER);`, user_id)

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
		total += flight.Count
		flights = append(flights, flight)
	}

	// Add the total number of trips to each flight.Total
	// TODO: Probably a better data structure to store the total as a single element
	for i := range flights {
		flights[i].Total = total
	}

	return flights, nil
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
			WHERE user_id = ?
			AND strftime('%Y', datetime(departure_time, 'unixepoch')) = ?
			GROUP BY month
		),
		total_trips AS (
			SELECT COUNT(ID) AS total
			FROM trips
			WHERE user_id = ?
			AND strftime('%Y', datetime(departure_time, 'unixepoch')) = ?
		)
		SELECT 
			m.month AS label,
			COALESCE(tc.trip_count, 0) AS count,
			(SELECT total FROM total_trips) AS total
		FROM months m
		LEFT JOIN trip_counts tc ON m.month = tc.month
		ORDER BY m.month`, user_id, year, user_id, year)

	if err != nil {
		return flights, err
	}
	defer rows.Close()

	for rows.Next() {
		var flight m.FlightAggregation
		err := rows.Scan(
			&flight.Label, 
			&flight.Count,
			&flight.Total,
		)
		if err != nil {
			return flights, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}


// TODO: Maybe add an argument to filter by year for airlines and countries as well
func (t *TripStore) getAirlinesCount(user_id int) ([]m.AirlineAggregation, error) {
	var airlines []m.AirlineAggregation

	rows, err := t.db.Query(`
		SELECT
			airline,
			COUNT(*) AS flight_count
		FROM trips
		WHERE user_id = ?
		GROUP BY airline
		ORDER BY flight_count DESC`, user_id)
	
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

func (t *TripStore) getCountriesCount(user_id int) ([]m.CountryAggregation, error) {
	var countries []m.CountryAggregation

	rows, err := t.db.Query(`
	SELECT 
		strftime('%Y', datetime(t.departure_time, 'unixepoch')) AS label,
		d.country AS country,
		COUNT(DISTINCT t.arrival) AS country_count
	FROM trips t
	JOIN airports d ON t.arrival = d.iata_code
	WHERE user_id = ?
	GROUP BY label
	ORDER BY label`, user_id)

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

	if agg != "m" && agg != "y" {
		return nil, nil, nil, errors.New("proper aggregation not given")
	}
	var flights []m.FlightAggregation
	var airlines []m.AirlineAggregation
	var countries []m.CountryAggregation
	var err error

	if agg == "m" {
		flights, err = t.getFlightsPerMonthForYear(user_id, year)
		if err != nil {
			return nil, nil, nil, err
		}

		airlines, err = t.getAirlinesCountForYear(user_id, year)
		if err != nil {
			return nil, nil, nil, err
		}
		countries, err = t.getCountriesCountForYear(user_id, year)
		if err != nil {
			return nil, nil, nil, err
		}
	} else {
		flights, err = t.getFlightsForYears(user_id)
		if err != nil {
			return nil, nil, nil, err
		}
		airlines, err = t.getAirlinesCount(user_id)
		if err != nil {
			return nil, nil, nil, err
		}
		countries, err = t.getCountriesCount(user_id)
		if err != nil {
			return nil, nil, nil, err
		}
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


func (t *TripStore) GetTotalMileageAndTime(user_id int) (m.TimeSpaceAggregation, error) {
	var tsAggregation m.TimeSpaceAggregation
	row := t.db.QueryRow(`
		WITH trip_data AS (
			SELECT 
				t.departure_time, 
				t.arrival_time, 
				d.latitude AS departure_lat, 
				d.longitude AS departure_lon, 
				a.latitude AS arrival_lat, 
				a.longitude AS arrival_lon
			FROM trips t
			JOIN airports d ON t.departure = d.iata_code
			JOIN airports a ON t.arrival = a.iata_code
			WHERE t.user_id = 12
		)
		SELECT
		SUM((arrival_time - departure_time) / 3600.0) AS total_hours,
		CAST(
			SUM(
			6371 * 2 * ASIN(
				SQRT(
				POWER(SIN(((arrival_lat - departure_lat) * 3.141592653589793 / 180) / 2), 2) +
				COS(departure_lat * 3.141592653589793 / 180) * COS(arrival_lat * 3.141592653589793 / 180) *
				POWER(SIN(((arrival_lon - departure_lon) * 3.141592653589793 / 180) / 2), 2)
				)
			)
			) AS INTEGER
		) AS total_km
		FROM trip_data;`, user_id)

	err := row.Scan(
		&tsAggregation.TotalHours,
		&tsAggregation.TotalKm,
	)
	if err != nil {
		return tsAggregation, err
	}

	return tsAggregation, nil
}