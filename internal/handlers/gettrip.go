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
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // if htmx is requesting, return only the fragment
    if headerVal == "true" {

        userTrips, err := t.tripStore.GetTripsGivenUser(userId)
        if err != nil {
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

            // Filter only the past trips before now
            for _, trip := range userTrips {
                if trip.DepartureTime < uint32(currentUnixTime) {
                    pastTrips = append(pastTrips, trip)
                }
            }
            renderErr := templates.RenderPastTrips(pastTrips).Render(r.Context(), w)
            if renderErr != nil {
                http.Error(w, "Error rendering template", http.StatusInternalServerError)
                return
            }
        } else {
            // We want upcoming, which are trips that are coming in the future (up to 1 week)
            var filteredTrips []models.Trip
            for _, trip := range userTrips {
                // get the unix time for 1 year from now
                oneWeekFromNow := currentUnixTime + (365 * 24 * 60 * 60) 
                if trip.DepartureTime > uint32(currentUnixTime) && trip.DepartureTime < uint32(oneWeekFromNow) {
                    filteredTrips = append(filteredTrips, trip)
                }
            }
            renderErr := templates.RenderTrips(filteredTrips).Render(r.Context(), w)
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

