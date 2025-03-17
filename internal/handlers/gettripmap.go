package handlers

import (
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetTripMapHandler struct {
    userStore *db.UserStore
}

type GetTripMapHandlerParams struct {
    UserStore *db.UserStore
}

func NewGetTripMapHandler(params GetTripMapHandlerParams) *GetTripMapHandler {
    return &GetTripMapHandler{
        userStore: params.UserStore,
    }
}


func (u *GetTripMapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	c := templates.TripMap(user.FirstName)
	err = templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}