package database

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type SessionStore struct {
	db *sql.DB
}

type NewSessionStoreParams struct {
	DB *sql.DB
}

func NewSessionStore(params NewSessionStoreParams) *SessionStore {
	return &SessionStore{db: params.DB}
}


func (s *SessionStore) CreateSession(userId string) (string, error) {
	sessionId := uuid.New().String()

	stmt, err := s.db.Prepare("INSERT INTO sessions (session_id, user_id) VALUES (?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionId, userId)
	if err != nil {
		return "", err
	}
	return sessionId, err
}

func (s *SessionStore) GetUserFromSession(sessionID string) (int, error) {
	var id string
	err := s.db.QueryRow("SELECT user_id FROM sessions WHERE session_id = ?", sessionID).Scan(&id)
	if err != nil {
		return 0, err
	}
	if id == "" {
		log.Fatal("No User Associated with the session")
		return 0, sql.ErrNoRows
	}
	numId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Fatal error converting userID string to int: ", err)
		return 0, err
	}

	return numId, nil
}