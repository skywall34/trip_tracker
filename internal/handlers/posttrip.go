package handlers

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
)

type PostTripHandler struct {
	tripStore *db.TripStore
}

type PostTripHandlerParams struct {
	TripStore *db.TripStore
}

func NewPostTripHandler(params PostTripHandlerParams) (*PostTripHandler) {
	return &PostTripHandler{
		tripStore: params.TripStore,
	}
}

const layout = "2006-01-02T15:04"

func parseLocalToUTC(input, timezone string)(time.Time, error) {
	// Parse the user-provided time (local time format)
	localTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}

	// Load the provided timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("Error loading timezone, defaulting to UTC:", err)
		loc = time.UTC
	}

	// Convert the local time to the correct timezone
	localTime = time.Date(
		localTime.Year(), localTime.Month(), localTime.Day(),
		localTime.Hour(), localTime.Minute(), 0, 0, loc,
	)

	// Convert to UTC
	return localTime.UTC(), nil
}

// TODO: Some of this input is going to have to come from the API
func (t *PostTripHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {

	departure := r.FormValue("departure")
	arrival := r.FormValue("arrival")
	departureTimeString := r.FormValue("departuretime") // Will receive as datetime (2024-05-06T14:30:25)
	arrivalTimeString := r.FormValue("arrivaltime") // Will receive as datetime (2024-05-06T14:30:25)
	airline := r.FormValue("airline")
	flightNumber := r.FormValue("flightnumber")
	reservation := r.FormValue("reservation")
	terminal := r.FormValue("terminal")
	gate := r.FormValue("gate")
	timezone := r.FormValue("timezone") // Hjidden field to get timezone of user

	ctx := r.Context()
	userId, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		// redirect to home
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	parsedDepartureTime, err := parseLocalToUTC(departureTimeString, timezone)
	if err != nil {
		fmt.Println("Error parsing departure time string:", err)
		return
	}
	parsedArrivalTime, err := parseLocalToUTC(arrivalTimeString, timezone)
	if err != nil {
		fmt.Println("Error parsing arrival time string:", err)
		return
	}

	newTrip := models.Trip{
		UserId: userId,
		Departure: departure,
		Arrival: arrival,
		DepartureTime: uint32(parsedDepartureTime.Unix()), // Save the data as UTC for uniform datetime, Frontend takes care of timezones
		ArrivalTime: uint32(parsedArrivalTime.Unix()),
		Airline: airline,
		FlightNumber: flightNumber,
		Reservation: reservation,
		Terminal: terminal,
		Gate: gate,
	}

	// Insert
	_, err = t.tripStore.CreateTrip(newTrip)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// HTMX Redirect Response
	w.WriteHeader(http.StatusOK)
}
