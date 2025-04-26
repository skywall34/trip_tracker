package models

import "time"

type ResetToken struct {
	ID string `json:"id"`
	UserId int64 `json:"user_id"`
	TokenHash string `json:"token_hash"`
	ExpiresAt time.Time `json:"expires_at"`
	Used bool `json:"used"`
}