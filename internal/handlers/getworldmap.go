package handlers

import (
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
)


type GetWorldMapHandler struct {
	tripStore *db.TripStore
}

type GetWorldMapHandlerParams struct {
	TripStore *db.TripStore
}

func NewGetWorldMapHandler(parmas GetWorldMapHandlerParams) *GetWorldMapHandler {
	return &GetWorldMapHandler{
		tripStore: parmas.TripStore,
	}
}

func (t *GetWorldMapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
    userId, ok := ctx.Value(m.UserKey).(int)

	if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

	// Assume you've determined visited country ISO codes for the current user
    visited, err := t.tripStore.GetVisitedCountryMap(userId)
	if err != nil {
		// Handle error, e.g., log it and return an error response
		http.Error(w, "Error fetching visited countries", http.StatusInternalServerError)
		// Redirect to home
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

    // Deep copy of models.CountryMap so we don't overwrite global state
    userCountries := make([]models.Country, len(models.CountryMap))
    for i, c := range models.CountryMap {
        userCountries[i] = c
        userCountries[i].Visited = visited[c.ISOCode]
    }

	c := templates.WorldMap(userCountries)
    templates.Layout(c, "World Map").Render(r.Context(), w)
}