package handlers

import (
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetPlacesHandler struct {
	placeStore *db.PlaceStore
	tripStore  *db.TripStore
}

type GetPlacesHandlerParams struct {
	PlaceStore *db.PlaceStore
	TripStore  *db.TripStore
}

func NewGetPlacesHandler(params GetPlacesHandlerParams) *GetPlacesHandler {
	return &GetPlacesHandler{
		placeStore: params.PlaceStore,
		tripStore:  params.TripStore,
	}
}

func (h *GetPlacesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get all places
	places, err := h.placeStore.GetPlacesForUser(userID)
	if err != nil {
		http.Error(w, "Error fetching places", http.StatusInternalServerError)
		return
	}

	// Get all trips
	trips, err := h.tripStore.GetTripsGivenUser(userID)
	if err != nil {
		http.Error(w, "Error fetching trips", http.StatusInternalServerError)
		return
	}

	// Get stats
	stats, err := h.placeStore.GetPlaceStats(userID)
	if err != nil {
		http.Error(w, "Error fetching stats", http.StatusInternalServerError)
		return
	}

	// Render the page
	c := templates.PlacesPage(places, trips, stats)
	err = templates.Layout(c, "Places").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
