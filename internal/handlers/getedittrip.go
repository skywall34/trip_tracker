package handlers

import (
	"net/http"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetEditTripHandler struct {
	tripStore *db.TripStore
}

type GetEditTripHandlerParams struct {
	TripStore *db.TripStore
}

func NewGetEditTripHandlerParmas(params GetEditTripHandlerParams) *GetEditTripHandler {
	return &GetEditTripHandler{
		tripStore: params.TripStore,
	}
}

func (t *GetEditTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    tripID := r.URL.Query().Get("id")

	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		// redirect to home
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// Insert
	numTripId, err := strconv.Atoi(tripID)
	if err != nil {
		http.Error(w, "Error processing tripID for deletion", http.StatusInternalServerError)
		return
	}
	trip, err := t.tripStore.GetTripGivenId(numTripId, userID)
	if err != nil {
		http.Error(w, "Error getting trip", http.StatusInternalServerError)
		return
	}

	err = templates.EditTripForm(trip).Render(r.Context(), w)
	if err != nil {
		// Handle rendering error
		http.Error(w, "Error rendering GetEditTrip template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}