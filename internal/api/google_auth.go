package api

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func NewGoogleOauthConfig() *oauth2.Config {
    return &oauth2.Config{
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        RedirectURL:  "http://localhost:3000/auth/google/callback",
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint:     google.Endpoint,
    }
}