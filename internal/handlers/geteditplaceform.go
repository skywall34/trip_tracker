package handlers

import (
	"net/http"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetEditPlaceFormHandler struct {
	placeStore *db.PlaceStore
}

type GetEditPlaceFormHandlerParams struct {
	PlaceStore *db.PlaceStore
}

func NewGetEditPlaceFormHandler(params GetEditPlaceFormHandlerParams) *GetEditPlaceFormHandler {
	return &GetEditPlaceFormHandler{
		placeStore: params.PlaceStore,
	}
}

func (h *GetEditPlaceFormHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get place ID from query parameter
	placeIDStr := r.URL.Query().Get("id")
	placeID, err := strconv.Atoi(placeIDStr)
	if err != nil {
		http.Error(w, "Invalid place ID", http.StatusBadRequest)
		return
	}

	// Fetch place from database
	place, err := h.placeStore.GetPlaceByID(placeID, userID)
	if err != nil {
		http.Error(w, "Place not found", http.StatusNotFound)
		return
	}

	// Render the edit form
	err = templates.EditPlaceForm(place).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
