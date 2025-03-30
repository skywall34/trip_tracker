package handlers

import (
	"encoding/json"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
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

	// Fetch trips from the store
	trips, err := t.tripStore.GetTripsGivenUser(userId)
	if err != nil {
		http.Error(w, "Error Fetchin Tripos", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
}