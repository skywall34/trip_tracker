package handlers

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
)

type GetTripHandler struct {
    tripStore *db.TripStore
}

type GetTripHandlerParams struct {
    TripStore *db.TripStore
}

func NewGetTripHandler(params GetTripHandlerParams) *GetTripHandler {
    return &GetTripHandler{
        tripStore: params.TripStore,
    }
}

func (t *GetTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    filterPast := r.URL.Query().Get("past")
    headerVal := r.Header.Get("HX-Request")

    ctx := r.Context()
    userId, ok := ctx.Value(m.UserKey).(int)
    fmt.Printf("User ID: %d, Filter Past: %s HeaderVal: %s, OK: %t \n", userId, filterPast, headerVal, ok)
    if !ok {
        // redirect to home
        w.WriteHeader(http.StatusNoContent)
        return
    }

    // if htmx is requesting, return only the fragment
    if headerVal == "true" {

        userTrips, userConnectingTrips, err := t.tripStore.GetConnectingTripsGivenUser(userId)
        if err != nil {
            fmt.Printf("Error getting trips for user %d: %v\n", userId, err)
            http.Error(w, "Error getting trips", http.StatusInternalServerError)
            return
        }
        if len(userTrips) == 0 {
            http.Error(w, "No trips found for this user", http.StatusNotFound)
            return
        }

        currentUnixTime := time.Now().Unix()

        if filterPast == "true" {
            var pastTrips []models.Trip
            var pastConnectingTrips []models.ConnectingTrip

            // Filter only the past trips before now
            for _, trip := range userTrips {
                if trip.DepartureTime < uint32(currentUnixTime) {
                    pastTrips = append(pastTrips, trip)
                }
            }

            // Filter the connecting trips to only include those that are past
            for _, connectingTrip := range userConnectingTrips {
                if connectingTrip.FromTrip.DepartureTime < uint32(currentUnixTime) {
                    pastConnectingTrips = append(pastConnectingTrips, connectingTrip)
                }
            }

            renderErr := templates.RenderPastTrips(pastTrips, pastConnectingTrips).Render(r.Context(), w)
            if renderErr != nil {
                http.Error(w, "Error rendering template", http.StatusInternalServerError)
                return
            }
        } else {
            // We want upcoming, which are trips that are coming in the future (up to 1 year)
            var filteredTrips []models.Trip
            var filteredConnectingTrips []models.ConnectingTrip
            oneYearFromNow := currentUnixTime + (365 * 24 * 60 * 60) 
            for _, trip := range userTrips {
                // get the unix time for 1 year from now
                if trip.DepartureTime > uint32(currentUnixTime) && trip.DepartureTime < uint32(oneYearFromNow) {
                    filteredTrips = append(filteredTrips, trip)
                }
            }
            for _, connectingTrip := range userConnectingTrips {
                if connectingTrip.FromTrip.DepartureTime > uint32(currentUnixTime) && connectingTrip.FromTrip.DepartureTime < uint32(oneYearFromNow) {
                    filteredConnectingTrips = append(filteredConnectingTrips, connectingTrip)
                }
            }
            renderErr := templates.RenderTrips(filteredTrips, filteredConnectingTrips).Render(r.Context(), w)
            if renderErr != nil {
                http.Error(w, "Error rendering template", http.StatusInternalServerError)
                return
            }
        }
    } else {
        // Otherwise, return the full page
        c := templates.TripsPage()
        templates.Layout(c, "Trips").Render(r.Context(), w)
    }
}

