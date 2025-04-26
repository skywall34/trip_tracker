// Generate the ResetPassword form passing in the token
package handlers

import (
	"net/http"

	"github.com/skywall34/trip-tracker/templates"
)

type GetResetPasswordHandler struct {}

type GetResetPasswordHandlerParams struct {}

func NewGetResetPasswordHandlerParams() *GetResetPasswordHandler {
    return &GetResetPasswordHandler{}
}

func (h *GetResetPasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if token == "" {
		http.Error(w, "Token Not Found!", http.StatusBadRequest)
		return
	}

	c := templates.ResetPasswordPage(token) // TODO: Implement
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}