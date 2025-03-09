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

	ctx := r.Context()
    userId, ok := ctx.Value(m.UserKey).(int)
	if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

	parsedDepartureTime, err := time.Parse(layout, departureTimeString)
	if err != nil {
		fmt.Println("Error parsing departure time string:", err)
		return
	}
	parsedArrivalTime, err := time.Parse(layout, arrivalTimeString) 
	if err != nil {
		fmt.Println("Error parsing arrival time string:", err)
		return
	}

	newTrip := models.Trip{
		UserId: int64(userId),
		Departure: departure,
		Arrival: arrival,
		DepartureTime: uint32(parsedDepartureTime.UTC().Unix()),
		ArrivalTime: uint32(parsedArrivalTime.UTC().Unix()),
		Airline: airline,
		FlightNumber: flightNumber,
		Reservation: reservation,
		Terminal: terminal,
		Gate: gate,
	}

	// Insert
	_, err = t.tripStore.CreateTrip(newTrip)
	if err != nil {
		// c := templates.CreateTripFailure()
		// err = c.Render(r.Context(), w)

		// if err != nil {
		// 	http.Error(w, "Error Rendering Template", http.StatusInternalServerError)
		// 	return
		// }
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// HTMX Redirect Response
    w.Header().Set("HX-Redirect", "/trips") // This makes HTMX handle the redirect
    w.WriteHeader(http.StatusSeeOther) // HTTP 303 See Other (optional but recommended)

}