package handlers

import (
	"context"
	"net/http"

	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

    userId := m.GetUserUsingContext(r.Context())

	if userId > -1 {
		ctx := context.WithValue(r.Context(), m.UserKey, userId)
		c := templates.Home()
		err := templates.Layout(c, "Mia's Trips").Render(ctx, w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	} else {
		c := templates.Home()
		err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}