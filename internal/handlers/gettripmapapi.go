package handlers

import (
	"encoding/json"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
)

type GetTripMapApiHandler struct {
    tripStore *db.TripStore
}

type GetTripMapApiHandlerParams struct {
    TripStore *db.TripStore
}

func NewGetTripMapApiHandler(params GetTripMapApiHandlerParams) *GetTripMapApiHandler {
    return &GetTripMapApiHandler{
        tripStore: params.TripStore,
    }
}


func (t *GetTripMapApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
    userId, ok := ctx.Value(m.UserKey).(int)
    if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

	// Fetch trips and connecting trips from the store
	standaloneTrips, connectingTrips, err := t.tripStore.GetConnectingTripsGivenUser(userId)
	if err != nil {
		http.Error(w, "Error fetching trips", http.StatusInternalServerError)
		return
	}

	// Structure the response to include both standalone and connecting flights
	response := struct {
		StandaloneTrips  []models.Trip           `json:"standalone_trips"`
		ConnectingTrips  []models.ConnectingTrip `json:"connecting_trips"`
	}{
		StandaloneTrips: standaloneTrips,
		ConnectingTrips: connectingTrips,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}