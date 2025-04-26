package database

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"time"

	// Removed import to avoid import cycle
	m "github.com/skywall34/trip-tracker/internal/models"
)

// Handles the functions accessinng table password_reset_tokens
type PasswordResetStore struct {
	db *sql.DB
}

type PasswordResetStoreParams struct {
	DB *sql.DB
}

func NewPasswordResetStore(params PasswordResetStoreParams) *PasswordResetStore {
	return &PasswordResetStore{db: params.DB}
}



func (rs *PasswordResetStore) GenerateResetToken(userID int) (string, error) {
	rawToken := make([]byte, 32)
	_, err := rand.Read(rawToken)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(rawToken)

	// Hash the token for storage
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	expiry := time.Now().Add(1 * time.Hour)

	q := `INSERT INTO password_reset_tokens (user_id, token_hash, expires_at) VALUES (?, ?, ?)`

	// Insert into `password_reset_tokens` table
	stmt, err := rs.db.Prepare(q)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, tokenHash, expiry)
	if err != nil {
		return "", err
	}

	return token, nil

}


func (rs *PasswordResetStore) ValidateResetToken(token string) (*m.User, error) {
	var passwordReset m.PasswordReset	
	var user m.User

	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	tokenQ := `
		SELECT id, user_id, token_hash, expires_at, used 
		FROM password_reset_tokens
		WHERE token_hash = ? and used = false and expires_at > NOW()
	`

	// Look up the token hash
	stmt, err := rs.db.Prepare(tokenQ)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(tokenHash).Scan(&passwordReset.ID, &passwordReset.UserId, &passwordReset.TokenHash, &passwordReset.ExpiresAt, &passwordReset.Used)
	if err != nil {
		return nil, err
	}

	userQ := `SELECT id, username, password, first_name, last_name, email FROM users WHERE username = ?`
	
	// Fetch the user
	// Can't import from user because of circular dependency
	err = rs.db.QueryRow(userQ, passwordReset.UserId).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (rs *PasswordResetStore) MarkTokenUsed(token string) error {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	q := `
		UPDATE password_reset_tokens SET user = 1 WHERE token_hash = ?
	`

	stmt, err := rs.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(q, tokenHash)
	if err != nil {
		return err
	}
	return nil
}