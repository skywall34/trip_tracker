package handlers

import (
	"fmt"
	"net/http"

	"github.com/skywall34/trip-tracker/templates"
)

func HtmxLoginHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.Login()
	err := templates.Layout(c, "Mia's Trips").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	// TODO: Proper authentication logic here
	// TODO: 
	user, err := GetUser(email)
	if err != nil {
		fmt.Println("Invalid User")
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	passwordIsValid, err := ComparePasswords(password, user.Password)

	if err != nil || !passwordIsValid {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	// Example: Assume user logs in successfully and gets a session ID
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

	fmt.Println("User logged in, session cookie set.")

	// Redirect user after setting cookie
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}