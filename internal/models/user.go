package models

type User struct {
    ID int64 `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"` // TODO: Email validator
}