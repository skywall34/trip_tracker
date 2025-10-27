package handlers

import (
	"net/http"

	"github.com/skywall34/trip-tracker/templates"
)

type GetPlaceModalHandler struct{}

type GetPlaceModalHandlerParams struct{}

func NewGetPlaceModalHandler(params GetPlaceModalHandlerParams) *GetPlaceModalHandler {
	return &GetPlaceModalHandler{}
}

func (h *GetPlaceModalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Path

	// Check if this is a close action
	if action == "/api/places/modal/close" {
		err := templates.AddPlaceModal().Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// Otherwise, open an empty modal
	err := templates.AddPlaceModalForm("", "", "", 0, 0, "").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
