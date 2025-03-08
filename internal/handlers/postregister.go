package handlers

import (
	"fmt"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
	"golang.org/x/crypto/bcrypt"
)


type PostRegisterHandler struct {
	userStore *db.UserStore
}

type PostRegisterHandlerParams struct {
	UserStore *db.UserStore
}

func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{
		userStore: params.UserStore,
	}
}

func hashPassword(password string) (string, error) {
    hashedBytes, err  := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedBytes), nil
}

func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

	// Check if the user already exists
    _, err := h.userStore.GetUser(email)
    if err == nil {
        http.Error(w, "Error creating user: User Already Exists!", http.StatusInternalServerError)
		return
    }

    // Hash the password
    hashedPassword, err := hashPassword(password)
    if err != nil {
        http.Error(w, "Error creating user: Error Hashing Password", http.StatusInternalServerError)
		return
    }

    newUser := models.User{
        Username: email,
        Password: hashedPassword,
        FirstName: firstname,
        LastName: lastname,
        Email: email,
    }

	_, err = h.userStore.CreateUser(newUser)

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