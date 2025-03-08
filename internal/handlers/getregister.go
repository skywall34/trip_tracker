package handlers

import (
	"net/http"

	"github.com/skywall34/trip-tracker/templates"
)


type GetRegisterHandler struct {}

type GetRegisterHandlerParams struct {}

func NewGetRegisterHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

func (h *GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Register()
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}