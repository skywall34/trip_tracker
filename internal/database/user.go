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

// TODO: Check to make sure user has not already been created with email
func (u *UserStore) CreateUser(user m.User) (int, error) {
	stmt, err := u.db.Prepare("INSERT INTO users (username, password, first_name, last_name, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Password, user.FirstName, user.LastName, user.Email)
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
	err := u.db.QueryRow("SELECT id, username, password, first_name, last_name, email FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserStore) GetUsers(username string) ([]m.User, error) {
	var users []m.User
	rows, err := u.db.Query("SELECT id, username, password, first_name, last_name, email FROM users WHERE username = ?")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user m.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserStore) GetUserGivenID(id int) (m.User, error) {
	var user m.User
	err := u.db.QueryRow(`SELECT 
							id, 
							username, 
							password, 
							first_name, 
							last_name, 
							email 
						FROM 
							users 
						WHERE id = ?`, id).Scan(
							&user.ID, 
							&user.Username, 
							&user.Password, 
							&user.FirstName, 
							&user.LastName, 
							&user.Email)

	if err != nil {
		return user, err
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
							email 
						FROM 
							users 
						WHERE email = ?`, email).Scan(
							&user.ID, 
							&user.Username, 
							&user.Password, 
							&user.FirstName, 
							&user.LastName, 
							&user.Email)

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

