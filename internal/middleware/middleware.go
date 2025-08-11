package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
)

type key string

var NonceKey key = "nonces"

type Nonces struct {
	Htmx            string
	ResponseTargets string
	Tw              string
	Modal			string
	TabsJS		    string
	MapJS           string
	Map3dJS         string
	ThreeJS         string
	Leaflet         string
	HtmxCSSHash     string
	ConvertTS       string
}


func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

// CSPMiddleware sets a Content Security Policy (CSP) header on the response
// and attaches a set of nonce values to the request context for use in templates
// and inline scripts/styles. A new set of nonces is generated for each request.
func CSPMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a new set of nonce values for this request.
		// These nonces allow specific inline scripts and styles to run.
		nonceSet := Nonces{
			// Nonce for inline HTMX scripts.
			Htmx: generateRandomString(16),
			ResponseTargets: generateRandomString(16),
			// Nonce for replacing content instead of hx-on
			Modal: generateRandomString(16),
			TabsJS: generateRandomString(16), // For tab functionality in templ UI
			// Mapping JS and CSS (Leaflet.js and leaflet.css)
			Leaflet: generateRandomString(16),
			MapJS: generateRandomString(16),
			Map3dJS: generateRandomString(16),
			ThreeJS: generateRandomString(16),
			// Nonce for convertTimes.js
			ConvertTS: generateRandomString(16),
		}

		// Store the nonce set in the request context so other parts of the application
		// (e.g., templates) can access it.
		ctx := context.WithValue(r.Context(), NonceKey, nonceSet)

		// Build the CSP header using the generated nonces.
		cspHeader := fmt.Sprintf(
			"default-src 'self'; " +
			"base-uri 'self'; " +
			"object-src 'none'; " +
			"frame-ancestors 'none'; " +
			"form-action 'self' https://accounts.google.com; " +
			"script-src 'self' 'strict-dynamic' " +
				"'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s' " +
				"https://cdn.jsdelivr.net; " +
			"style-src 'self' 'nonce-%s' 'nonce-%s' https://fonts.googleapis.com; " +
			"img-src 'self' data: https://*.tile.openstreetmap.org https://*.basemaps.cartocdn.com; " +
			"connect-src 'self' https://*.tile.openstreetmap.org https://*.basemaps.cartocdn.com https://accounts.google.com https://oauth2.googleapis.com; " +
			"font-src 'self' https://fonts.gstatic.com; ",
			nonceSet.Htmx,
			nonceSet.ResponseTargets,
			nonceSet.ConvertTS,
			nonceSet.Modal,
			nonceSet.TabsJS,
			nonceSet.Leaflet,
			nonceSet.MapJS,
			nonceSet.Map3dJS,
			nonceSet.ThreeJS,
			nonceSet.Tw,
			nonceSet.Leaflet,
		)
		w.Header().Set("Content-Security-Policy", cspHeader)


		// Call the next handler in the chain, passing the updated context.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


// Return the header text/html
func TextHTMLMiddleware(next http.HandlerFunc) http.HandlerFunc  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func ApplicationJsonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// get the Nonce from the context, it is a struct called Nonces,
// so we can get the nonce we need by the key, i.e. HtmxNonce
func GetNonces(ctx context.Context) Nonces {
	nonceSet := ctx.Value(NonceKey)
	if nonceSet == nil {
		log.Fatal("error getting nonce set - is nil")
	}

	nonces, ok := nonceSet.(Nonces)

	if !ok {
		log.Fatal("error getting nonce set - not ok")
	}

	return nonces
}

func GetHtmxNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.Htmx
}

func GetResponseTargetsNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.ResponseTargets
}

func GetTwNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.Tw
}

func GetConvertTSNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.ConvertTS
}

func GetModalNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.Modal
}

func GetLeafletNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.Leaflet
}

func GetMapJSNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.MapJS
}

func GetMap3DJSNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.ThreeJS
}

func GetThreeJSNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.ThreeJS
}

func GetTabsJSNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.TabsJS
}
/***********************************Auth Middleware**********************************************/

type AuthMiddleware struct {
	sessionStore *db.SessionStore
	sessionCookieName string
}

func NewAuthMiddleware(sessionStore *db.SessionStore, sessionCookieName string) *AuthMiddleware {
	return &AuthMiddleware{
		sessionStore: sessionStore,
		sessionCookieName: sessionCookieName,
	}
}

type UserContextKey string

var UserKey UserContextKey = "user"


func (m *AuthMiddleware) AddUserToContext(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated
		cookie, err := r.Cookie("session_id")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		sessionID := cookie.Value
		if sessionID == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Get user from session
		userId, err := m.sessionStore.GetUserFromSession(sessionID)

		if err != nil {
			// if the request came from /home, redirect just return
			if r.URL.Path == "/" {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, userId)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// TODO: Once using store change false return to nil and return userStore
func GetUserUsingContext(ctx context.Context) int {
	userId, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userId
}

/***********************************Logging Middleware**********************************************/

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}