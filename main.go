package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/skywall34/trip-tracker/internal/database"
	"github.com/skywall34/trip-tracker/internal/handlers"
	m "github.com/skywall34/trip-tracker/internal/middleware"
)

func main() {

    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    db, err := database.InitDB("./internal/database/database.db")
    if err != nil {
        log.Fatal(err)
    }

    userStore := database.NewUserStore(database.NewUserStoreParams{DB: db})
    tripStore := database.NewTripStore(database.NewTripStoreParams{DB: db})

    mux := http.NewServeMux()

    fs:= http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))

    // Main
    mux.Handle("/", m.AuthMiddleware(m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetHomeHandler().ServeHTTP))))

    // Page Routes
    mux.Handle("GET /trips",  m.AuthMiddleware(m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetTripHandler(
        handlers.GetTripHandlerParams{
            TripStore: tripStore,
    }).ServeHTTP))))
    mux.Handle("GET /login",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetLoginHandler().ServeHTTP)))
    mux.Handle("POST /login",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewPostLoginHandler(
        handlers.PostLoginHandlerParams{
            UserStore: userStore,
    }).ServeHTTP)))
    mux.Handle("GET /register",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetRegisterHandler().ServeHTTP)))
    mux.Handle("POST /register",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewPostRegisterHandler(
        handlers.PostRegisterHandlerParams{
            UserStore: userStore,
    }).ServeHTTP)))

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    fmt.Printf("Server running on port %s\n", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
