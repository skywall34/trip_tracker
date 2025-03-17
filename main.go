package main

import (
	"fmt"
	"log"
	"net/http"

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
    //TODO: Chaining middleware seems to break css for some reason
    authMiddleware := m.NewAuthMiddleware(sessionStore, "session_id")

    mux := http.NewServeMux()

    fs:= http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))

    // Main
    mux.Handle("/", 
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                        handlers.NewGetHomeHandler().ServeHTTP)))))

    // Page Routes
    mux.Handle("GET /trips",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                    handlers.NewGetTripHandler(
                        handlers.GetTripHandlerParams{
                            TripStore: tripStore}).ServeHTTP)))))

    mux.Handle("POST /trips",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                        handlers.NewPostTripHandler(
                            handlers.PostTripHandlerParams{
                                TripStore: tripStore}).ServeHTTP)))))

    mux.Handle("DELETE /trips",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                        handlers.NewDeleteTripHandler(
                            handlers.DeleteTripHandlerParams{
                                TripStore: tripStore}).ServeHTTP)))))

    mux.Handle("GET /login",  
        m.CSPMiddleware(
            m.TextHTMLMiddleware(
                m.LoggingMiddleware(
                    handlers.NewGetLoginHandler().ServeHTTP))))

    mux.Handle("POST /login",  
        m.CSPMiddleware(
            m.TextHTMLMiddleware(
                m.LoggingMiddleware(
                handlers.NewPostLoginHandler(
                    handlers.PostLoginHandlerParams{
                        UserStore: userStore,
                        SessionStore: sessionStore}).ServeHTTP))))

    mux.Handle("POST /logout", 
        m.CSPMiddleware(
            m.LoggingMiddleware(handlers.NewPostLogoutHandler(
                handlers.PostLogoutHandlerParams{
                    SessionCookieName: "session_id",
                }).ServeHTTP)))

    mux.Handle("GET /register",  
        m.CSPMiddleware(
            m.TextHTMLMiddleware(
                m.LoggingMiddleware(
                    handlers.NewGetRegisterHandler().ServeHTTP))))

    mux.Handle("POST /register",  
        m.CSPMiddleware(
            m.TextHTMLMiddleware(
                m.LoggingMiddleware(handlers.NewPostRegisterHandler(
                    handlers.PostRegisterHandlerParams{
                        UserStore: userStore,
                        SessionStore: sessionStore,
                    }).ServeHTTP))))

    mux.Handle("GET /statistics",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                        handlers.NewGetTripMapHandler(
                            handlers.GetTripMapHandlerParams{
                                UserStore: userStore,
                            }).ServeHTTP)))))

    // API CALLS
    mux.Handle("GET /api/trips",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.LoggingMiddleware(
                    handlers.NewGetTripMapApiHandler(
                        handlers.GetTripMapApiHandlerParams{
                            TripStore: tripStore}).ServeHTTP))))

    server := http.Server {
        Addr: ":3000",
        Handler: mux,
    }

    fmt.Println("Server running on port :3000")
    server.ListenAndServe()
}