package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
	"github.com/skywall34/trip-tracker/templates"
	"golang.org/x/crypto/bcrypt"
)

type PostLoginHandler struct {
	userStore    *db.UserStore
	sessionStore *db.SessionStore
}

type PostLoginHandlerParams struct {
	UserStore    *db.UserStore
	SessionStore *db.SessionStore
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		userStore:    params.UserStore,
		sessionStore: params.SessionStore,
	}
}

func comparePasswords(password, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Creating a duration to prevent timing attacks
	// https://en.wikipedia.org/wiki/Timing_attack#:~:text=In%20cryptography%2C%20a%20timing%20attack,of%20the%20timing%20measurements%2C%20etc.
	const duration = 1 * time.Second
	startTime := time.Now()

	email := r.FormValue("email")
	password := r.FormValue("password")

	// TODO: Proper authentication logic here
	user, err := h.userStore.GetUserGivenEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	passwordIsValid, err := comparePasswords(password, user.Password)

	if err != nil || !passwordIsValid {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	sessionID, err := h.sessionStore.CreateSession(strconv.Itoa(user.ID))

	if err != nil {
		log.Fatal("Error Creating session: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	// Set the session cookie (max-age 1 day, httpOnly for security)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,  // Prevents JavaScript access (security best practice)
		MaxAge:   86400, // 1 day in seconds
	})

	fmt.Println("User logged in, session cookie set.")

	// Measure login time, if less than duration invoke time.Sleep to ensure handler
	// responds exactly after that duration
	if time.Since(startTime) < duration {
		time.Sleep(duration - time.Since(startTime))
	}

	// Redirect user after setting cookie
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}
