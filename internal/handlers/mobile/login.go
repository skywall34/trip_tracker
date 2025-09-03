package mobile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// EmailLoginHandler handles email/password login for mobile
type EmailLoginHandler struct {
	userStore    *db.UserStore
	authHandler *MobileAuthHandler
}

// EmailLoginHandlerParams constructor parameters
type EmailLoginHandlerParams struct {
	AuthHandler *MobileAuthHandler
	UserStore   *db.UserStore
}

// NewEmailLoginHandler creates a new EmailLoginHandler
func NewEmailLoginHandler(params EmailLoginHandlerParams) *EmailLoginHandler {
	return &EmailLoginHandler{
		userStore:    params.UserStore,
		authHandler: params.AuthHandler,
	}
}

// EmailLoginRequest represents the login request body
type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// comparePasswords compares a plain text password with a hashed password
func comparePasswords(password, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

// No need for generateTokens method as we use the one from MobileAuthHandler

// ServeHTTP handles POST /api/v1/mobile/auth/login requests
func (h *EmailLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("AUTH CALLED")

	// Creating a duration to prevent timing attacks
	const duration = 1 * time.Second
	startTime := time.Now()

	var req EmailLoginRequest
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

	// Validate input
	if req.Email == "" || req.Password == "" {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "MISSING_CREDENTIALS",
				Message: "Email and password are required",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get user by email
	user, err := h.userStore.GetUserGivenEmail(req.Email)
	if err != nil {
		// Ensure consistent timing even on user not found
		if time.Since(startTime) < duration {
			time.Sleep(duration - time.Since(startTime))
		}
		
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "INVALID_CREDENTIALS",
				Message: "Invalid email or password",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Compare passwords
	passwordIsValid, err := comparePasswords(req.Password, user.Password)
	if err != nil || !passwordIsValid {
		// Ensure consistent timing on password failure
		if time.Since(startTime) < duration {
			time.Sleep(duration - time.Since(startTime))
		}
		
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "INVALID_CREDENTIALS",
				Message: "Invalid email or password",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate JWT tokens using the auth handler
	tokens, err := h.authHandler.generateTokens(user.ID, user.Email)
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

	// Ensure consistent timing
	if time.Since(startTime) < duration {
		time.Sleep(duration - time.Since(startTime))
	}

	// Return successful authentication
	response := LoginResponse{
		Success: true,
		Data: &LoginData{
			User: UserResponse{
				ID:      user.ID,
				Email:   user.Email,
				Name:    user.FirstName + " " + user.LastName,
				Picture: "", // Email login doesn't have profile picture
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