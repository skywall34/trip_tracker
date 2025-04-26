package models

import "time"

type PasswordReset struct {
	ID        string `json:"id"`
	UserId    int `json:"user_id"`
	TokenHash string `json:"token_hash"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      uint8 `json:"used"`
}