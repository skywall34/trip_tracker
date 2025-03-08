package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skywall34/trip-tracker/internal/models"
	m "github.com/skywall34/trip-tracker/internal/models"
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

func (u *UserStore) CreateUser(user models.User) (int64, error) {
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

	return id, nil
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

