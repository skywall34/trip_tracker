package handlers

import (
	"net/http"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
)

type EditTripHandler struct {
	tripStore *db.TripStore
}

type EditTripHandlerParams struct {
	TripStore *db.TripStore
}

func NewEditTripHandler(params EditTripHandlerParams) *EditTripHandler {
	return &EditTripHandler{
		tripStore: params.TripStore,
	}
}

// Get the single trip, edit it in code, then reupload to DB as edited trip
// TODO: Bug -> Editing the trip with the same timestamps changes the time of the trip (timezone issue)
func (t *EditTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tripID := r.URL.Query().Get("id")
	numTripID, err := strconv.Atoi(tripID)
	if err != nil {
		http.Error(w, "Error parsing tripID", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// First, get the existing trip from the database
	existingTrip, err := t.tripStore.GetTripGivenId(numTripID, userID)
	if err != nil {
		http.Error(w, "Trip not found", http.StatusNotFound)
		return
	}

	// Parse form values and only update if they're provided
	if departure := r.FormValue("departure"); departure != "" {
		existingTrip.Departure = departure
	}

	if arrival := r.FormValue("arrival"); arrival != "" {
		existingTrip.Arrival = arrival
	}

	if departureTimeString := r.FormValue("departuretime"); departureTimeString != "" {
		timezone := r.FormValue("timezone")
		parsedDepartureTime, err := ParseLocalToUTC(departureTimeString, existingTrip.Departure, timezone)
		if err != nil {
			http.Error(w, "Error parsing departure time", http.StatusBadRequest)
			return
		}
		existingTrip.DepartureTime = uint32(parsedDepartureTime.Unix())
	}

	if arrivalTimeString := r.FormValue("arrivaltime"); arrivalTimeString != "" {
		timezone := r.FormValue("timezone")
		parsedArrivalTime, err := ParseLocalToUTC(arrivalTimeString, existingTrip.Arrival, timezone)
		if err != nil {
			http.Error(w, "Error parsing arrival time", http.StatusBadRequest)
			return
		}
		existingTrip.ArrivalTime = uint32(parsedArrivalTime.Unix())
	}

	if airline := r.FormValue("airline"); airline != "" {
		existingTrip.Airline = airline
	}

	if flightNumber := r.FormValue("flightnumber"); flightNumber != "" {
		existingTrip.FlightNumber = flightNumber
	}

	// Handle optional fields (pointers) - update even if empty to allow clearing
	if r.PostForm.Has("reservation") {
		reservation := r.FormValue("reservation")
		if reservation != "" {
			existingTrip.Reservation = &reservation
		} else {
			existingTrip.Reservation = nil
		}
	}

	if r.PostForm.Has("terminal") {
		terminal := r.FormValue("terminal")
		if terminal != "" {
			existingTrip.Terminal = &terminal
		} else {
			existingTrip.Terminal = nil
		}
	}

	if r.PostForm.Has("gate") {
		gate := r.FormValue("gate")
		if gate != "" {
			existingTrip.Gate = &gate
		} else {
			existingTrip.Gate = nil
		}
	}

	// Update the trip in the database
	err = t.tripStore.EditTrip(existingTrip)
	if err != nil {
		http.Error(w, "Error editing trip", http.StatusInternalServerError)
		return
	}

	// HTMX Redirect Response
	w.Header().Set("HX-Redirect", "/trips")
	w.WriteHeader(http.StatusOK)
}
