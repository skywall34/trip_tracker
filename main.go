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
    sessionStore := database.NewSessionStore(database.NewSessionStoreParams{DB: db})
    authMiddleware := m.NewAuthMiddleware(sessionStore, "session_id")

    mux := http.NewServeMux()

    fs:= http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))

    // Main
    mux.Handle("/", authMiddleware.AddUserToContext(
        m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetHomeHandler().ServeHTTP))))

    // Page Routes
    mux.Handle("GET /trips",  authMiddleware.AddUserToContext(m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetTripHandler(
        handlers.GetTripHandlerParams{
            TripStore: tripStore,
        }).ServeHTTP))))
    mux.Handle("POST /trips",  authMiddleware.AddUserToContext(m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewPostTripHandler(
        handlers.PostTripHandlerParams{
            TripStore: tripStore,
        }).ServeHTTP))))
    mux.Handle("GET /login",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetLoginHandler().ServeHTTP)))
    mux.Handle("POST /login",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewPostLoginHandler(
        handlers.PostLoginHandlerParams{
            UserStore: userStore,
            SessionStore: sessionStore,
        }).ServeHTTP)))
    mux.Handle("POST /logout", m.CSPMiddleware(handlers.NewPostLogoutHandler(
        handlers.PostLogoutHandlerParams{
            SessionCookieName: "session_id",
        }).ServeHTTP))
    mux.Handle("GET /register",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewGetRegisterHandler().ServeHTTP)))
    mux.Handle("POST /register",  m.CSPMiddleware(m.TextHTMLMiddleware(handlers.NewPostRegisterHandler(
        handlers.PostRegisterHandlerParams{
            UserStore: userStore,
            SessionStore: sessionStore,
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
