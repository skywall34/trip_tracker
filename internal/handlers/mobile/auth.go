package mobile

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// JWT Claims structure
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Token response structure
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// User response structure
type UserResponse struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture,omitempty"`
}

// Login response structure
type LoginResponse struct {
	Success bool           `json:"success"`
	Data    *LoginData     `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type LoginData struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int64       `json:"expires_in"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Google login request
type GoogleLoginRequest struct {
	GoogleToken string `json:"google_token"`
}

// Refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// MobileAuthHandler handles mobile authentication similar to website
type MobileAuthHandler struct {
	userStore         *db.UserStore
	jwtSecret         []byte
	googleOauthConfig *oauth2.Config
}

// MobileAuthHandlerParams constructor parameters
type MobileAuthHandlerParams struct {
	UserStore *db.UserStore
}

// NewMobileAuthHandler creates a new MobileAuthHandler
func NewMobileAuthHandler(params MobileAuthHandlerParams) *MobileAuthHandler {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "fallback-secret-key-change-in-production"
	}

	// Set up Google OAuth config (mimic website)
	googleConfig := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		// For mobile, we use a different redirect approach
		RedirectURL: "postmessage", // Special value for mobile apps
		Scopes:      []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:    google.Endpoint,
	}

	return &MobileAuthHandler{
		userStore:         params.UserStore,
		jwtSecret:         []byte(secret),
		googleOauthConfig: googleConfig,
	}
}

// Generate JWT tokens
func (h *MobileAuthHandler) generateTokens(userID int, email string) (*TokenResponse, error) {
	now := time.Now()

	// Access token (15 minutes)
	accessClaims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "miastrips-mobile",
			Subject:   fmt.Sprintf("user:%d", userID),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(h.jwtSecret)
	if err != nil {
		return nil, err
	}

	// Refresh token (7 days)
	refreshClaims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "miastrips-mobile",
			Subject:   fmt.Sprintf("refresh:%d", userID),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(h.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    15 * 60, // 15 minutes in seconds
	}, nil
}

// Verify JWT token
func (h *MobileAuthHandler) verifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return h.jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ServeHTTP handles the Google OAuth login process
func (h *MobileAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GoogleLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "INVALID_REQUEST",
				Message: "Invalid request body",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Handle mock token for development
	if req.GoogleToken == "mock-google-token-development" {
		mockUser := m.User{
			ID:           1,
			Email:        "dev@example.com",
			FirstName:    "Development",
			LastName:     "User",
			GoogleID:     "mock-google-id",
			AuthProvider: "google",
		}

		tokens, err := h.generateTokens(mockUser.ID, mockUser.Email)
		if err != nil {
			response := LoginResponse{
				Success: false,
				Error: &ErrorResponse{
					Code:    "TOKEN_GENERATION_FAILED",
					Message: "Failed to generate authentication tokens",
				},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := LoginResponse{
			Success: true,
			Data: &LoginData{
				User: UserResponse{
					ID:      mockUser.ID,
					Email:   mockUser.Email,
					Name:    fmt.Sprintf("%s %s", mockUser.FirstName, mockUser.LastName),
					Picture: "",
				},
				AccessToken:  tokens.AccessToken,
				RefreshToken: tokens.RefreshToken,
				ExpiresIn:    tokens.ExpiresIn,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Exchange the authorization code for a token
	ctx := context.Background()

	// Create an OAuth2 token from the provided access token
	token := &oauth2.Token{
		AccessToken: req.GoogleToken,
		TokenType:   "Bearer",
	}

	// Create an HTTP client using the access token
	client := h.googleOauthConfig.Client(ctx, token)

	// Retrieve user information from Google's userinfo endpoint
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "NETWORK_ERROR",
				Message: "Failed to fetch user info: " + err.Error(),
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer resp.Body.Close()

	// Parse user information
	var userInfo struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "PARSE_ERROR",
				Message: "Failed to parse user info",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if the user exists in the database
	user, err := h.userStore.GetUserGivenEmail(userInfo.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// User does not exist, create a new user
			newUser := m.User{
				Username:     userInfo.Email,
				FirstName:    userInfo.FirstName,
				LastName:     userInfo.LastName,
				Email:        userInfo.Email,
				GoogleID:     userInfo.ID,
				AuthProvider: "google",
			}
			newUserID, err := h.userStore.CreateUser(newUser)
			if err != nil {
				response := LoginResponse{
					Success: false,
					Error: &ErrorResponse{
						Code:    "USER_CREATION_FAILED",
						Message: "Failed to create user: " + err.Error(),
					},
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
				return
			}
			user = m.User{ID: newUserID}
		} else {
			response := LoginResponse{
				Success: false,
				Error: &ErrorResponse{
					Code:    "DATABASE_ERROR",
					Message: "Database error: " + err.Error(),
				},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// Generate JWT tokens
	tokens, err := h.generateTokens(user.ID, userInfo.Email)
	if err != nil {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "TOKEN_GENERATION_FAILED",
				Message: "Failed to generate authentication tokens",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return successful authentication
	response := LoginResponse{
		Success: true,
		Data: &LoginData{
			User: UserResponse{
				ID:      user.ID,
				Email:   userInfo.Email,
				Name:    fmt.Sprintf("%s %s", userInfo.FirstName, userInfo.LastName),
				Picture: "",
			},
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
			ExpiresIn:    tokens.ExpiresIn,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
