package handlers

import (
	"log"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetStatisticsPageHandler struct {
    userStore *db.UserStore
	tripStore *db.TripStore
}

type GetStatisticsPageHandlerParams struct {
    UserStore *db.UserStore
	TripStore *db.TripStore
}

func NewGetStatisticsPageHandler(params GetStatisticsPageHandlerParams) *GetStatisticsPageHandler {
    return &GetStatisticsPageHandler{
        userStore: params.UserStore,
		tripStore: params.TripStore,
    }
}


func (u *GetStatisticsPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
    userID, ok := ctx.Value(m.UserKey).(int)
    if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

	user, err := u.userStore.GetUserGivenID(userID)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	// Get the total trips for the user
	// TODO: Get the mileage, and hours for each flight
	tsAggregation, err := u.tripStore.GetTotalMileageAndTime(userID)
	if err != nil {
		log.Printf("Error getting totals: %v", err)
		http.Error(w, "Error getting Totals", http.StatusInternalServerError)
		return
	}

	c := templates.Statistics(user.FirstName, tsAggregation)
	err = templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}