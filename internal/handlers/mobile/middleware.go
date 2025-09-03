package mobile

import (
	"context"
	"net/http"
	"strings"
)

// Context keys for JWT claims
type contextKey string

const UserIDKey contextKey = "user_id"
const UserEmailKey contextKey = "user_email"

// JWT authentication middleware for mobile endpoints
type JWTMiddleware struct {
	authHandler *MobileAuthHandler
}

func NewJWTMiddleware(authHandler *MobileAuthHandler) *JWTMiddleware {
	return &JWTMiddleware{
		authHandler: authHandler,
	}
}

func (m *JWTMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check for Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := tokenParts[1]

		// Verify JWT token
		claims, err := m.authHandler.verifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)

		// Continue to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper function to get user ID from context
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(UserIDKey).(int)
	return userID, ok
}

// Helper function to get user email from context
func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(UserEmailKey).(string)
	return email, ok
}