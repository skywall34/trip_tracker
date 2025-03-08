package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/skywall34/trip-tracker/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashedBytes, err  := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedBytes), nil
}

func ComparePasswords(password, hashedPassword string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return false, err
    }
    return true, nil
}


// If the variable starts with a capital letter, it is a constant and public
var Users = map[int]models.User{
    1: {Username: "alexj", Password: "password123", FirstName: "Alex", LastName: "Johnson", Email: "alex.johnson@example.com"},
    2: {Username: "sarah28", Password: "securepass", FirstName: "Sarah", LastName: "Thompson", Email: "sarah.thompson@example.com"},
    3: {Username: "mike45", Password: "1234secure", FirstName: "Michael", LastName: "Smith", Email: "mike.smith@example.com"},
    4: {Username: "emilyc", Password: "mypassword", FirstName: "Emily", LastName: "Carter", Email: "emily.carter@example.com"},
    5: {Username: "johnD", Password: "john123", FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
    6: {Username: "anna_b", Password: "annapass", FirstName: "Anna", LastName: "Brown", Email: "anna.brown@example.com"},
    7: {Username: "peterp", Password: "passPeter", FirstName: "Peter", LastName: "Parker", Email: "peter.parker@example.com"},
    8: {Username: "tony.s", Password: "starkrules", FirstName: "Tony", LastName: "Stark", Email: "tony.stark@example.com"},
    9: {Username: "natasha_r", Password: "redspy", FirstName: "Natasha", LastName: "Romanoff", Email: "natasha.romanoff@example.com"},
    10: {Username: "bruce.b", Password: "hulkSmash", FirstName: "Bruce", LastName: "Banner", Email: "bruce.banner@example.com"},
}

// Internal Function to Get User
func GetUser(email string) (models.User, error) {
	for _, user := range Users {
        if user.Email == email {
            return user, nil
        }
    }
    return models.User{}, errors.New("user not found")
}

// Internal Function to Create User (Used by register)
func CreateUser(email string, password string, firstname string, lastname string) (models.User, error) {

    // Check if the user already exists
    // TODO: We'd replace this with store SQL query
    for _, user := range Users {
        if user.Email == email {
            return models.User{}, errors.New("user already exists")
        }
    }

    // Hash the password
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return models.User{}, err
    }

    newUser := models.User{
        Username: email,
        Password: hashedPassword,
        FirstName: firstname,
        LastName: lastname,
        Email: email,
    }

    userId := len(Users) + 1
    Users[userId] = newUser

    return newUser, nil
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    // get the id of the user to be deleted if it exists
    userId := r.URL.Query().Get("id")
    if userId != "" { 
        userNum, err := strconv.Atoi(userId)
        if err != nil {
            http.Error(w, "Invalid user ID, user ID must be number!", http.StatusBadRequest)
            return
        }

		user, exists := Users[userNum]
		if !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
        json.NewEncoder(w).Encode(user)    
    } else {
        json.NewEncoder(w).Encode(Users)
    }
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement POST handler
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse and validate the request body
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validation logic 
    if newUser.Username == "" || newUser.Password == "" || newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" { 
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        marshaled, err := json.MarshalIndent(newUser, "", "   ")
        if err != nil {
            log.Fatalf("marshaling error: %s", err)
        }
        fmt.Println(string(marshaled))
        return
    }

    // Validate the request body and add the new user to the Users slice
    userId := len(Users) + 1
    Users[userId] = newUser

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func EditUsersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // get the id of the user to be deleted
    userId := r.URL.Query().Get("id")
    if userId == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }

    userNum, err := strconv.Atoi(userId)
    if err != nil {
        http.Error(w, "Invalid user ID, user ID must be number!", http.StatusBadRequest)
        return
    }

    // Check if the user exists
	if _, exists := Users[userNum]; !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

    // Parse and validate the incoming JSON
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}


    // Validation logic
	if updatedUser.Username == "" || updatedUser.Password == "" || updatedUser.FirstName == "" || updatedUser.LastName == "" || updatedUser.Email == "" {
		http.Error(w, "User missing required fields", http.StatusBadRequest)
		return
	}

	// Update the user
	Users[userNum] = updatedUser

	// Respond with the updated user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)

}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // get the id of the user to be deleted
    userId := r.URL.Query().Get("id")
    if userId == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }

    userNum, err := strconv.Atoi(userId)
    if err != nil {
        http.Error(w, "Invalid user ID, user ID must be number!", http.StatusBadRequest)
        return
    }

    // Check if the user exists
	if _, exists := Users[userNum]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete the user
	delete(Users, userNum)

	// Respond with success
	w.WriteHeader(http.StatusNoContent)
}