package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	"github.com/skywall34/trip-tracker/templates"
	"golang.org/x/crypto/bcrypt"
)

type PostLoginHandler struct {
	userStore *db.UserStore
	sessionStore *db.SessionStore
}

type PostLoginHandlerParams struct {
	UserStore *db.UserStore
	SessionStore *db.SessionStore
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		userStore: params.UserStore,
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

	email := r.FormValue("email")
	password := r.FormValue("password")

	// TODO: Proper authentication logic here
	user, err := h.userStore.GetUser(email)
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

	sessionID, err := h.sessionStore.CreateSession(strconv.FormatInt(user.ID, 10))

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
		HttpOnly: true, // Prevents JavaScript access (security best practice)
		MaxAge:   86400, // 1 day in seconds
	})

	fmt.Println("User logged in, session cookie set.")

	// Redirect user after setting cookie
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}