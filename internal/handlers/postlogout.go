package handlers

import (
	"net/http"
	"time"
)

type PostLogoutHandler struct {
	sessionCookieName string
}

type PostLogoutHandlerParams struct {
	SessionCookieName string
}

func NewPostLogoutHandler(params PostLogoutHandlerParams) *PostLogoutHandler {
	return &PostLogoutHandler{
		sessionCookieName: params.SessionCookieName,
	}
}

func (h *PostLogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    h.sessionCookieName,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
	})

	// HTMX Redirect Response
    w.Header().Set("HX-Redirect", "/") // This makes HTMX handle the redirect
    w.WriteHeader(http.StatusSeeOther) // HTTP 303 See Other (optional but recommended)
}