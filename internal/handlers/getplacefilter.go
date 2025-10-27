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

type GetPlaceFilterHandler struct {
	placeStore *db.PlaceStore
	tripStore  *db.TripStore
}

type GetPlaceFilterHandlerParams struct {
	PlaceStore *db.PlaceStore
	TripStore  *db.TripStore
}

func NewGetPlaceFilterHandler(params GetPlaceFilterHandlerParams) *GetPlaceFilterHandler {
	return &GetPlaceFilterHandler{
		placeStore: params.PlaceStore,
		tripStore:  params.TripStore,
	}
}

func (h *GetPlaceFilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(m.UserKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Parse form values
	r.ParseForm()

	showTrips := r.FormValue("show_trips") == "true"
	showPlaces := r.FormValue("show_places") == "true"
	years := r.Form["year"]
	categories := r.Form["category"]

	// Get all places and trips
	allPlaces := []models.Place{}
	allTrips := []models.Trip{}

	if showPlaces {
		places, err := h.placeStore.GetPlacesForUser(userID)
		if err == nil {
			allPlaces = places
		}
	}

	if showTrips {
		trips, err := h.tripStore.GetTripsGivenUser(userID)
		if err == nil {
			allTrips = trips
		}
	}

	// Filter places by year and category
	filteredPlaces := []models.Place{}
	for _, place := range allPlaces {
		// Check year filter
		yearMatch := len(years) == 0 // If no years selected, show all
		if len(years) > 0 {
			placeYear := time.Unix(int64(place.VisitDate), 0).Year()
			for _, year := range years {
				yearInt, _ := strconv.Atoi(year)
				if placeYear == yearInt {
					yearMatch = true
					break
				}
			}
		}

		// Check category filter
		categoryMatch := len(categories) == 0 // If no categories selected, show all
		if len(categories) > 0 && place.Category != nil {
			for _, cat := range categories {
				if *place.Category == cat {
					categoryMatch = true
					break
				}
			}
		}

		// Add to filtered list if matches both filters
		if yearMatch && categoryMatch {
			filteredPlaces = append(filteredPlaces, place)
		}
	}

	// Filter trips by year
	filteredTrips := []models.Trip{}
	for _, trip := range allTrips {
		yearMatch := len(years) == 0 // If no years selected, show all
		if len(years) > 0 {
			tripYear := time.Unix(int64(trip.DepartureTime), 0).Year()
			for _, year := range years {
				yearInt, _ := strconv.Atoi(year)
				if tripYear == yearInt {
					yearMatch = true
					break
				}
			}
		}

		if yearMatch {
			filteredTrips = append(filteredTrips, trip)
		}
	}

	// Render the filtered timeline
	err := templates.TimelineFeed(filteredPlaces, filteredTrips).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
