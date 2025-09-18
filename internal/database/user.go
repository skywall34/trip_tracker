package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	m "github.com/skywall34/trip-tracker/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Password Hashing
type UserStore struct {
	db *sql.DB
}

type NewUserStoreParams struct {
	DB *sql.DB
}

func NewUserStore(params NewUserStoreParams) *UserStore {
	return &UserStore{db: params.DB}
}

func (u *UserStore) CreateUser(user m.User) (int, error) {
	stmt, err := u.db.Prepare(`INSERT INTO users
		(username, password, first_name, last_name, email, google_id, auth_provider)
		VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Password, user.FirstName, user.LastName, user.Email, user.GoogleID, user.AuthProvider)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (u *UserStore) GetUser(username string) (m.User, error) {
	var user m.User
	err := u.db.QueryRow(`SELECT
							id,
							username,
							password,
							first_name,
							last_name,
							email,
							google_id,
							auth_provider,
							created_at
						FROM users
						WHERE username = ?`, username).Scan(
							&user.ID,
							&user.Username,
							&user.Password,
							&user.FirstName,
							&user.LastName,
							&user.Email,
							&user.GoogleID,
							&user.AuthProvider,
							&user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserStore) GetUsers(username string) ([]m.User, error) {
	var users []m.User
	rows, err := u.db.Query(`SELECT
							id,
							username,
							password,
							first_name,
							last_name,
							email,
							google_id,
							auth_provider,
							created_at
						FROM users
						WHERE username = ?`, username)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user m.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.GoogleID,
			&user.AuthProvider,
			&user.CreatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserStore) GetUserGivenID(id int) (m.User, error) {
	var user m.User
	var username, password, firstName, lastName, googleID sql.NullString

	err := u.db.QueryRow(`SELECT
							id,
							username,
							password,
							first_name,
							last_name,
							email,
							google_id,
							auth_provider,
							created_at
						FROM
							users
						WHERE id = ?`, id).Scan(
							&user.ID,
							&username,
							&password,
							&firstName,
							&lastName,
							&user.Email,
							&googleID,
							&user.AuthProvider,
							&user.CreatedAt)

	if err != nil {
		return user, err
	}

	// Convert NullString to regular strings
	if username.Valid {
		user.Username = username.String
	}
	if password.Valid {
		user.Password = password.String
	}
	if firstName.Valid {
		user.FirstName = firstName.String
	}
	if lastName.Valid {
		user.LastName = lastName.String
	}
	if googleID.Valid {
		user.GoogleID = googleID.String
	}

	return user, nil
}


func (u *UserStore) GetUserGivenEmail(email string) (m.User, error) {
	var user m.User
	err := u.db.QueryRow(`SELECT
							id,
							username,
							password,
							first_name,
							last_name,
							email,
							google_id,
							auth_provider,
							created_at
						FROM
							users
						WHERE email = ?`, email).Scan(
							&user.ID,
							&user.Username,
							&user.Password,
							&user.FirstName,
							&user.LastName,
							&user.Email,
							&user.GoogleID,
							&user.AuthProvider,
							&user.CreatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}


func (u *UserStore) UpdatePassword(userID int, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	q := `
		UPDATE users SET password = ? WHERE id = ?
	`

	stmt, err := u.db.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(hashedPassword, userID)
	if err != nil {
		return err
	}
	return nil
}

