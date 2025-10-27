package handlers

import (
	"net/http"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
)

type DeletePlaceHandler struct {
	placeStore *db.PlaceStore
}

type DeletePlaceHandlerParams struct {
	PlaceStore *db.PlaceStore
}

func NewDeletePlaceHandler(params DeletePlaceHandlerParams) *DeletePlaceHandler {
	return &DeletePlaceHandler{
		placeStore: params.PlaceStore,
	}
}

func (h *DeletePlaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	placeIDStr := r.URL.Query().Get("id")
	placeID, err := strconv.Atoi(placeIDStr)
	if err != nil {
		http.Error(w, "Invalid place ID", http.StatusBadRequest)
		return
	}

	err = h.placeStore.DeletePlace(placeID, userID)
	if err != nil {
		http.Error(w, "Error deleting place", http.StatusInternalServerError)
		return
	}

	// Return empty response for HTMX to swap out the element
	w.WriteHeader(http.StatusOK)
}
