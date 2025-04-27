package handlers

import (
	"net/http"

	"github.com/skywall34/trip-tracker/templates"
)

type GetForgotPasswordHandler struct {}

type GetForgotPasswordHandlerParams struct {}

func NewGetForgotPasswordHandler() *GetForgotPasswordHandler {
	return &GetForgotPasswordHandler{}
}

func (h *GetForgotPasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Just return the page
	c := templates.ForgotPasswordPage()
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w) 
	if err != nil {
		// Handle rendering error
		http.Error(w, "Error rendering GetForgotPasswordHandler template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}