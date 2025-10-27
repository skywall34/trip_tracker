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

type PostPlaceHandler struct {
	placeStore *db.PlaceStore
}

type PostPlaceHandlerParams struct {
	PlaceStore *db.PlaceStore
}

func NewPostPlaceHandler(params PostPlaceHandlerParams) *PostPlaceHandler {
	return &PostPlaceHandler{
		placeStore: params.PlaceStore,
	}
}

func (h *PostPlaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	// Get form values
	placeID := r.FormValue("place_id")
	name := r.FormValue("name")
	address := r.FormValue("address")
	latStr := r.FormValue("latitude")
	lonStr := r.FormValue("longitude")
	visitDateStr := r.FormValue("visit_date")
	category := r.FormValue("category")
	notes := r.FormValue("notes")
	markerColor := r.FormValue("marker_color")

	// Convert and validate
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	// Parse visit date (assume format: 2006-01-02)
	visitTime, err := time.Parse("2006-01-02", visitDateStr)
	if err != nil {
		http.Error(w, "Invalid visit date", http.StatusBadRequest)
		return
	}

	// Set default marker color if not provided
	if markerColor == "" {
		markerColor = "#26e0b0"
	}

	// Create place object
	place := models.Place{
		UserID:      userID,
		PlaceID:     placeID,
		Name:        name,
		Address:     &address,
		Latitude:    lat,
		Longitude:   lon,
		VisitDate:   uint32(visitTime.Unix()),
		Category:    &category,
		Notes:       &notes,
		MarkerColor: markerColor,
	}

	// Save to database
	_, err = h.placeStore.CreatePlace(place)
	if err != nil {
		http.Error(w, "Error saving place", http.StatusInternalServerError)
		return
	}

	// Get updated list and re-render timeline
	places, err := h.placeStore.GetPlacesForUser(userID)
	if err != nil {
		http.Error(w, "Error fetching places", http.StatusInternalServerError)
		return
	}

	// Render the timeline component
	err = templates.TimelineFeed(places, nil).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
