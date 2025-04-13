package api

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type GoogleLoginHandler struct {
	googleOauthConfig *oauth2.Config
}

type GoogleLoginHandlerParams struct {
	GoogleOauthConfig *oauth2.Config
}

func NewGoogleLoginHandlerParams(params GoogleLoginHandlerParams) *GoogleLoginHandler {
	return &GoogleLoginHandler{
		googleOauthConfig: params.GoogleOauthConfig,
	}
}

func generateState() string {
    b := make([]byte, 16)
    if _, err := rand.Read(b); err != nil {
        log.Fatalf("Unable to generate state: %v", err)
    }
    return base64.URLEncoding.EncodeToString(b)
}


func (h *GoogleLoginHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	// TODO: Replace with a securely generated state string and save it in a session
	state := generateState()

    // Store the state in a secure cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "oauthstate",
        Value:    state,
        Expires:  time.Now().Add(10 * time.Minute),
        HttpOnly: true,
        Secure:   true, // Set to true if using HTTPS
        Path:     "/",
    })

	// Generate the OAuth URL with the state parameter
    url := h.googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}