package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
)


type EditTripHandler struct {
	tripStore *db.TripStore
}

type EditTripHandlerParams struct {
	TripStore *db.TripStore
}

func NewEditTripHandler(params EditTripHandlerParams) (*EditTripHandler) {
	return &EditTripHandler{
		tripStore: params.TripStore,
	}
}


func (t *EditTripHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {

	tripID := r.URL.Query().Get("id")
	numTripID, err := strconv.Atoi(tripID)
	if err != nil {
		http.Error(w, "Error parsing tripID ", http.StatusInternalServerError)
		return
	}

	departure := r.FormValue("departure")
	arrival := r.FormValue("arrival")
	departureTimeString := r.FormValue("departuretime") // Will receive as datetime (2024-05-06T14:30:25)
	arrivalTimeString := r.FormValue("arrivaltime") // Will receive as datetime (2024-05-06T14:30:25)
	airline := r.FormValue("airline")
	flightNumber := r.FormValue("flightnumber")
	reservation := r.FormValue("reservation")
	terminal := r.FormValue("terminal")
	gate := r.FormValue("gate")
	timezone := r.FormValue("timezone") // Hidden field to get timezone of user

	// We would get the form values

	// Get the trip, we need to take in the params of what the user wants to change
	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
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
		ID: numTripID,
		UserId: userID,
		Departure: departure,
		Arrival: arrival,
		DepartureTime: uint32(parsedDepartureTime.Unix()), // Save the data as UTC for uniform datetime, Frontend takes care of timezones
		ArrivalTime: uint32(parsedArrivalTime.Unix()),
		Airline: airline,
		FlightNumber: flightNumber,
		Reservation: &reservation,
		Terminal: &terminal,
		Gate: &gate,
	}

	err = t.tripStore.EditTrip(newTrip)
	if err != nil {
		http.Error(w, "Error editing trip", http.StatusInternalServerError)
		return
	}

	// HTMX Redirect Response
	// We just need the refreshed page with the udpated trip
	w.Header().Set("HX-Redirect", "/trips")
	w.WriteHeader(http.StatusOK)


}