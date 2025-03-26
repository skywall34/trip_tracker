package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetStatisticsHandler struct {
    tripStore *db.TripStore
}

type GetStatisticsHandlerParams struct {
    TripStore *db.TripStore
}

func NewGetStatisticsHandlerParams(params GetStatisticsHandlerParams) *GetStatisticsHandler {
    return &GetStatisticsHandler{
        tripStore: params.TripStore,
    }
}

func (t *GetStatisticsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	agg := r.URL.Query().Get("agg")
	yearStr := r.URL.Query().Get("year")

	if yearStr == "" {
		yearStr = strconv.Itoa(time.Now().Year())
	}

    ctx := r.Context()
    userId, ok := ctx.Value(m.UserKey).(int)
    fmt.Printf("User ID: %d, Agg: %s, OK: %t \n", userId, agg, ok)

	if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

	// Aggregation will be m for month or y for year
	if agg != "m" && agg != "y" {
		http.Error(w, "Invalid aggregation", http.StatusBadRequest)
		return
	}

	flights, airline, country, err := t.tripStore.GetTripsPerAggregation(userId, yearStr, agg)
	if err != nil {
		http.Error(w, "Error getting trips", http.StatusInternalServerError)
		return
	}

	c := templates.AggregationComponent(flights, airline, country)
    templates.Layout(c, "Statistics").Render(r.Context(), w)


}