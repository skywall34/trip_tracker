package handlers

import (
	"fmt"
	"net/http"

	"github.com/skywall34/trip-tracker/templates"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

	_, err := CreateUser(email, password, firstname, lastname)

	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// c := templates.Register()
		// c.Render(r.Context(), w)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// TODO: Generate a secure session ID via Session Store
	sessionID := "abc123" // This would typically be generated securely

	// Set the session cookie (max-age 1 day, httpOnly for security)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true, // Prevents JavaScript access (security best practice)
		MaxAge:   86400, // 1 day in seconds
	})

	fmt.Println("User Registered, session cookie set.")

	// Redirect user after setting cookie
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HtmxRegisterHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.Register()
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}