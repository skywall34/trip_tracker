package handlers

import (
	"net/http"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
)

type DeleteTripHandler struct {
	tripStore *db.TripStore
}

type DeleteTripHandlerParams struct {
	TripStore *db.TripStore
}

func NewDeleteTripHandler(params DeleteTripHandlerParams) (*DeleteTripHandler) {
	return &DeleteTripHandler{
		tripStore: params.TripStore,
	}
}


func (t *DeleteTripHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {

	tripID := r.URL.Query().Get("id")

	ctx := r.Context()
	_, ok := ctx.Value(m.UserKey).(int)
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
	err = t.tripStore.DeleteTrip(numTripId)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// 200 tells htmx to remove the element from the html
	w.WriteHeader(http.StatusOK)
}
