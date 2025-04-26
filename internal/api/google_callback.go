package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	db "github.com/skywall34/trip-tracker/internal/database"
	"github.com/skywall34/trip-tracker/internal/models"
	"golang.org/x/oauth2"
)

type GoogleCallbackHandler struct {
	userStore *db.UserStore
	sessionStore *db.SessionStore
    googleOauthConfig *oauth2.Config
}

type GoogleCallbackHandlerParams struct {
	UserStore *db.UserStore
	SessionStore *db.SessionStore
    GoogleOauthConfig *oauth2.Config
}

func NewGoogleCallbackHandlerParams(params GoogleCallbackHandlerParams) *GoogleCallbackHandler {
	return &GoogleCallbackHandler{
		userStore: params.UserStore,
		sessionStore: params.SessionStore,
        googleOauthConfig: params.GoogleOauthConfig,
	}
}

func (h *GoogleCallbackHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    // Parse the user information
    var userInfo struct {
        ID        string `json:"id"`
        Email     string `json:"email"`
        FirstName string `json:"given_name"`
        LastName  string `json:"family_name"`
    }
    var sessionID string

    // Retrieve the state and code from the query parameters
    state := r.URL.Query().Get("state")
    code := r.URL.Query().Get("code")

    // Validate the state parameter
    cookie, err := r.Cookie("oauthstate")
    if err != nil {
        http.Error(w, "State cookie not found", http.StatusBadRequest)
        return
    }
    if state != cookie.Value {
        http.Error(w, "Invalid state parameter", http.StatusBadRequest)
        return
    }

    // Validate the code parameter
    if code == "" {
        http.Error(w, "Code parameter missing", http.StatusBadRequest)
        return
    }

    // URL-decode the authorization code
    decodedCode, err := url.QueryUnescape(code)
    if err != nil {
        http.Error(w, "Failed to decode authorization code: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Exchange the authorization code for an access token
    token, err := h.googleOauthConfig.Exchange(ctx, decodedCode)
    if err != nil {
        http.Error(w, "Token exchange error: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Create an HTTP client using the access token
    client := h.googleOauthConfig.Client(ctx, token)

    // Retrieve user information from Google's userinfo endpoint
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        http.Error(w, "Failed to fetch user info: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        http.Error(w, "Failed to parse user info", http.StatusInternalServerError)
        return
    }
    
    // Check if the user exists in the database
    user, err := h.userStore.GetUser(userInfo.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            // User does not exist, create a new user
            newUser := models.User{
                Username:     userInfo.Email,
                FirstName:    userInfo.FirstName,
                LastName:     userInfo.LastName,
                Email:        userInfo.Email,
                GoogleID:     userInfo.ID,
                AuthProvider: "google",
            }
            newUserID, err := h.userStore.CreateUser(newUser)
            if err != nil {
                http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
                return
            }
            user = models.User{ID: newUserID}
        } else {
            http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
            return
        }
    }

    sessionID, err = h.sessionStore.CreateSession(strconv.Itoa(user.ID))
    if err != nil {
        http.Error(w, "Failed to create session: "+err.Error(), http.StatusInternalServerError)
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

    fmt.Println("Google User Registered, session cookie set.")

    // Redirect the user to the dashboard or home page
    http.Redirect(w, r, "/", http.StatusSeeOther)
}