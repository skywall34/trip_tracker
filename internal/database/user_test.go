package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	m "github.com/skywall34/trip-tracker/internal/models"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	createUsersTable := `
    CREATE TABLE users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        password TEXT,
        first_name TEXT,
        last_name TEXT,
        email TEXT UNIQUE NOT NULL
    );`
	if _, err := db.Exec(createUsersTable); err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}
	return db
}

func TestGetUserGivenEmail(t *testing.T) {
	sqlDB := setupTestDB(t)
	store := NewUserStore(NewUserStoreParams{DB: sqlDB})

	user := m.User{
		Username:  "testuser",
		Password:  "secret",
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@example.com",
	}
	if _, err := store.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	got, err := store.GetUserGivenEmail("test@example.com")
	if err != nil {
		t.Fatalf("GetUserGivenEmail returned error: %v", err)
	}
	if got.Email != user.Email || got.Username != user.Username {
		t.Fatalf("expected user %v, got %v", user, got)
	}
}
