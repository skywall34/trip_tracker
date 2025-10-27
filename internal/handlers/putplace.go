package handlers

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
)

type PutPlaceHandler struct {
	placeStore *db.PlaceStore
}

type PutPlaceHandlerParams struct {
	PlaceStore *db.PlaceStore
}

func NewPutPlaceHandler(params PutPlaceHandlerParams) *PutPlaceHandler {
	return &PutPlaceHandler{
		placeStore: params.PlaceStore,
	}
}

func (h *PutPlaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get place ID
	placeIDStr := r.FormValue("id")
	placeID, err := strconv.Atoi(placeIDStr)
	if err != nil {
		http.Error(w, "Invalid place ID", http.StatusBadRequest)
		return
	}

	// Get existing place to preserve certain fields
	existingPlace, err := h.placeStore.GetPlaceByID(placeID, userID)
	if err != nil {
		http.Error(w, "Place not found", http.StatusNotFound)
		return
	}

	// Update fields
	name := r.FormValue("name")
	address := r.FormValue("address")
	visitDateStr := r.FormValue("visit_date")
	category := r.FormValue("category")
	notes := r.FormValue("notes")
	markerColor := r.FormValue("marker_color")

	visitTime, err := time.Parse("2006-01-02", visitDateStr)
	if err != nil {
		http.Error(w, "Invalid visit date", http.StatusBadRequest)
		return
	}

	// Update place object
	updatedPlace := models.Place{
		ID:          placeID,
		UserID:      userID,
		PlaceID:     existingPlace.PlaceID, // Preserve Google Place ID
		Name:        name,
		Address:     &address,
		Latitude:    existingPlace.Latitude,  // Preserve coordinates
		Longitude:   existingPlace.Longitude,
		VisitDate:   uint32(visitTime.Unix()),
		Category:    &category,
		Notes:       &notes,
		MarkerColor: markerColor,
	}

	err = h.placeStore.UpdatePlace(updatedPlace)
	if err != nil {
		http.Error(w, "Error updating place", http.StatusInternalServerError)
		return
	}

	// Render the updated place card
	err = templates.PlaceCard(updatedPlace).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
