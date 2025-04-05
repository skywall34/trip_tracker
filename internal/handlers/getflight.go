package handlers

import (
	"net/http"

	"github.com/skywall34/trip-tracker/internal/api"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetFlightHandler struct {}

type GetFlightHandlerParams struct {}

func NewGetFlightHandler() *GetFlightHandler {
	return &GetFlightHandler{}
}

func (h *GetFlightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
    _, ok := ctx.Value(m.UserKey).(int)

	if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

	flightIATA := r.URL.Query().Get("flight_iata")

	// Render
	if flightIATA == "" {
		// If no flight IATA provided, render the home page without flight data
		http.Error(w, "Flight IATA is required", http.StatusBadRequest)
		return
	}

	// Get the flight data from the api
	flightData, err := api.GetFlight(flightIATA)
	if err != nil {
		// Handle the error and return an internal server error
		http.Error(w, "Failed to retrieve flight data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template with flight data
	if flightData == nil {
		// Handle the case where no flight data was returned
		http.Error(w, "No flight data found for the provided flight IATA", http.StatusNotFound)
		return
	}

	err = templates.TripForm(*flightData).Render(r.Context(), w)

	if err != nil {
		// Handle rendering error
		http.Error(w, "Error rendering GetFlightHandler template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}