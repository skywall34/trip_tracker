package handlers

import (
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
)


type PostResetPasswordHandler struct {
	userStore *db.UserStore
	passwordResetStore *db.PasswordResetStore
}

type PostResetPasswordHandlerParams struct {
	UserStore *db.UserStore
	PasswordResetStore *db.PasswordResetStore

}

func NewPostResetPasswordHandler(params PostResetPasswordHandlerParams) (*PostResetPasswordHandler) {
	return &PostResetPasswordHandler{
		userStore: params.UserStore,
		passwordResetStore: params.PasswordResetStore,
	}
}


func (h *PostResetPasswordHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	token := r.FormValue("token")

	if password != confirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	user, err := h.passwordResetStore.ValidateResetToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	err = h.userStore.UpdatePassword(user.ID, password)
	if err != nil {
		http.Error(w, "Failed to reset password", http.StatusInternalServerError)
		return
	}

	// TODO: Implement
	err = h.passwordResetStore.MarkTokenUsed(token)
	if err != nil {
		http.Error(w, "Failed to mark token as used", http.StatusInternalServerError)
		return
	}

	// Success Message 
	w.Write([]byte(`
        <div class="text-center text-green-600 font-semibold">
            Your password has been reset! <a href="/login" class="text-blue-500 underline">Login</a>.
        </div>
    `))

}