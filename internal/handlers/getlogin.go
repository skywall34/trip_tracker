package handlers

import (
	"net/http"

	"github.com/skywall34/trip-tracker/templates"
)

type GetLoginHandler struct {}

type GetLoginHandlerParams struct {}

func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

func (h *GetLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Login()
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}