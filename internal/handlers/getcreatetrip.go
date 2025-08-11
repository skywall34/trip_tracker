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
        // Check if this is an HTMX request
        if r.Header.Get("HX-Request") != "" {
            // Return just the login component for HTMX
            c := templates.Login(true)
            err := c.Render(r.Context(), w)
            if err != nil {
                http.Error(w, "Error rendering Login template: "+err.Error(), http.StatusInternalServerError)
                return
            }
            return
        }
        // redirect to login page for non-HTMX requests
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Check if this is an HTMX request
    if r.Header.Get("HX-Request") != "" {
        // Return just the form component for HTMX
        c := templates.CreateTripForm()
        err := c.Render(r.Context(), w)
        if err != nil {
            http.Error(w, "Error rendering CreateTripForm template: "+err.Error(), http.StatusInternalServerError)
            return
        }
        return
    }

    // Return full page for non-HTMX requests
	c := templates.CreateTripPage()
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w) 
	if err != nil {
		// Handle rendering error
		http.Error(w, "Error rendering CreateTripPage template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}