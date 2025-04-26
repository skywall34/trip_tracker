package handlers

import (
	"fmt"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/models"
)


type ForgotPasswordHandler struct {
	userStore *db.UserStore
	passwordResetStore *db.PasswordResetStore
	emailService m.EmailService
}

type ForgotPasswordHandlerParams struct {
	UserStore *db.UserStore
	PasswordResetStore *db.PasswordResetStore
	EmailService m.EmailService
}

func NewForgotPasswordHandler(params ForgotPasswordHandlerParams) (*ForgotPasswordHandler) {
	return &ForgotPasswordHandler{
		userStore: params.UserStore,
		passwordResetStore: params.PasswordResetStore,
		emailService: params.EmailService,
	}
}


func (h *ForgotPasswordHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	user, err := h.userStore.GetUserGivenEmail(email) // TODO: Implement
	if err != nil {
		// Always pretend we succeeded
		w.Write([]byte(`<div class="text-center text-green-600 font-semibold">If the email exists, we sent a reset link.</div>`))
		return
	}

	token, err := h.passwordResetStore.GenerateResetToken(user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the email TODO: Implement
	resetLink := fmt.Sprintf("http://localhost/reset-password?token=%s", token)
	err = h.emailService.SendPassswordResetEmail(user.Email, resetLink)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Send success message
	w.Write([]byte(`<div class="text-center text-green-600 font-semibold">If the email exists, we sent a reset link.</div>`))
}