package handlers

import (
	"net/http"

	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

type GetCreateTripHandler struct {}

type GetCreateTripHandlerParams struct {}

func NewGetCreateTripHandler() *GetCreateTripHandler {
	return &GetCreateTripHandler{}
}

func (h *GetCreateTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    ctx := r.Context()
    _, ok := ctx.Value(m.UserKey).(int)

	if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

	c := templates.CreateTripPage()
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w) 
	if err != nil {
		// Handle rendering error
		http.Error(w, "Error rendering GetFlightHandler template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}