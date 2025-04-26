package models

type Session struct {
	ID        int  `json:"id"`
	SessionID string `json:"session_id"`
	UserID    int   `json:"user_id"`
}