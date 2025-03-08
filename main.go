package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/skywall34/trip-tracker/internal/handlers"
	m "github.com/skywall34/trip-tracker/internal/middleware"
)

func main() {

    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    // Initialize User Data
    // Hashing the passwords for the users
    for id, user := range handlers.Users {
        hashedPassword, err := handlers.HashPassword(user.Password)
        if err != nil {
            fmt.Printf("Error hashing password for user %d: %s\n", id, err)
            return
        }
        user.Password = hashedPassword
        handlers.Users[id] = user
    }

    mux := http.NewServeMux()

    fs:= http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))

    // Main
    mux.HandleFunc("/", m.AuthMiddleware(m.CSPMiddleware(m.TextHTMLMiddleware(handlers.HomeHandler))))

    // Page Routes
    mux.Handle("GET /trips",  m.AuthMiddleware(m.CSPMiddleware(m.TextHTMLMiddleware(handlers.HtmxTripsHandler))))
    mux.Handle("GET /login",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.HtmxLoginHandler)))
    mux.Handle("POST /login",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.LoginHandler)))
    mux.Handle("GET /register",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.HtmxRegisterHandler)))
    mux.Handle("POST /register",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.RegisterHandler)))

    // API Routes
    mux.Handle("GET /api/trips", m.ApplicationJsonMiddleware(handlers.GetTripsHandler))
    mux.Handle("POST /api/trips", m.ApplicationJsonMiddleware(handlers.PostTripsHandler))
    mux.Handle("PUT /api/trips", m.ApplicationJsonMiddleware(handlers.EditTripsHandler))
    mux.HandleFunc("DELETE /api/trips", handlers.DeleteTripsHandler)

    // User Routes
    mux.Handle("GET /api/users", m.ApplicationJsonMiddleware(handlers.GetUsersHandler))
    mux.Handle("POST /api/users", m.ApplicationJsonMiddleware(handlers.PostUsersHandler))
    mux.Handle("PUT /api/users", m.ApplicationJsonMiddleware(handlers.EditUsersHandler))
    mux.HandleFunc("DELETE /api/users", handlers.DeleteUsersHandler)

    // User Filter Routes
    mux.Handle("GET /api/user/trips", m.ApplicationJsonMiddleware(handlers.GetTripsForUserHandler))

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    fmt.Printf("Server running on port %s\n", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
