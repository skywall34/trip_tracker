package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type key string

var NonceKey key = "nonces"

type Nonces struct {
	Htmx            string
	ResponseTargets string
	Tw              string
	HtmxCSSHash     string
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
			// Nonce for inline Tailwind CSS (or similar).
			Tw: generateRandomString(16),
			// Precomputed hash for HTMX CSS.
			HtmxCSSHash: "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg=",
		}

		// Store the nonce set in the request context so other parts of the application
		// (e.g., templates) can access it.
		ctx := context.WithValue(r.Context(), NonceKey, nonceSet)

		// Build the CSP header using the generated nonces.
		cspHeader := fmt.Sprintf(
			"default-src 'self'; script-src 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' '%s';",
			nonceSet.Htmx,
			nonceSet.ResponseTargets,
			nonceSet.Tw,
			nonceSet.HtmxCSSHash,
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


/***********************************Auth Middleware**********************************************/

type UserContextKey string

var UserKey UserContextKey = "user"

// validateSession checks if there's a "session_id" cookie, and uses that ID to
// look up the associated user. In a real app, you'd query a database or cache.
// TODO: Replace with actual session validation logic. The return will be the whole User
func validateSession(r *http.Request) (int, error) {
    // Check if we have a session_id cookie
    cookie, err := r.Cookie("session_id")
    if err != nil {
        // No session_id cookie found
        return 0, errors.New("session cookie not found")
    }

    // In a real implementation, you'd look up the session in your session store.
    // For demonstration, we'll pretend any non-empty cookie value is valid.
    sessionID := cookie.Value
    if sessionID == "" {
        return 0, errors.New("invalid or empty session ID")
    }

    // Mock user ID retrieval from the session store.
    // For example, maybe session ID 'abc123' = user ID 42, etc.
    // We'll just return a static userID for demonstration.
    mockUserID := 1

    // If everything checks out, return the user ID
    return mockUserID, nil
}


func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated
		userId, err := validateSession(r)
		if err != nil {
			// if the request came from /home, redirect just retunr
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