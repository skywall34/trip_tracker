package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/skywall34/trip-tracker/internal/api"
	"github.com/skywall34/trip-tracker/internal/database"
	"github.com/skywall34/trip-tracker/internal/handlers"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
)

func main() {

    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    db, err := database.InitDB("file:./internal/database/database.db?_enable_math_functions=1")
    if err != nil {
        log.Fatal(err)
    }

    // Load the countries from a JSON file into memory
    if err := models.LoadCountriesFromFile("./static/data/countries.json"); err != nil {
        log.Fatalf("Failed to load countries: %v", err)
    }

    userStore := database.NewUserStore(database.NewUserStoreParams{DB: db})
    tripStore := database.NewTripStore(database.NewTripStoreParams{DB: db})
    sessionStore := database.NewSessionStore(database.NewSessionStoreParams{DB: db})
    //TODO: Chaining middleware seems to break css for some reason
    authMiddleware := m.NewAuthMiddleware(sessionStore, "session_id")
    
    // Google OAuth Initilization to Add the Environemnt Variables
    googleOauthConfig := api.NewGoogleOauthConfig()

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
                        handlers.NewGetStatisticsPageHandler(
                            handlers.GetStatisticsPageHandlerParams{
                                UserStore: userStore,
                                TripStore: tripStore,
                            }).ServeHTTP)))))

    mux.Handle("GET /worldmap",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                        handlers.NewGetWorldMapHandler(
                            handlers.GetWorldMapHandlerParams{
                                TripStore: tripStore,
                            }).ServeHTTP)))))

    mux.Handle("GET /createtrip",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                        handlers.NewGetCreateTripHandler().ServeHTTP)))))     
                    

    // API CALLS
    mux.Handle("GET /api/flights",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.TextHTMLMiddleware(
                    m.LoggingMiddleware(
                        handlers.NewGetFlightHandler().ServeHTTP)))))                         


    mux.Handle("GET /api/trips",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.LoggingMiddleware(
                    handlers.NewGetTripMapApiHandler(
                        handlers.GetTripMapApiHandlerParams{
                            TripStore: tripStore}).ServeHTTP))))

    mux.Handle("GET /api/statistics",  
        authMiddleware.AddUserToContext(
            m.CSPMiddleware(
                m.LoggingMiddleware(
                    handlers.NewGetStatisticsHandlerParams(
                        handlers.GetStatisticsHandlerParams{
                            TripStore: tripStore}).ServeHTTP))))

    // Google Auth
    mux.HandleFunc("/auth/google/login", 
        api.NewGoogleLoginHandlerParams(
            api.GoogleLoginHandlerParams{
                GoogleOauthConfig: googleOauthConfig,
            }).ServeHTTP)

    mux.HandleFunc("/auth/google/callback", 
        api.NewGoogleCallbackHandlerParams(
            api.GoogleCallbackHandlerParams{
                UserStore: userStore,
                SessionStore: sessionStore,
                GoogleOauthConfig: googleOauthConfig,
            }).ServeHTTP)

    server := http.Server {
        Addr: ":3000",
        Handler: mux,
    }

    fmt.Println("Server running on port :3000")
    server.ListenAndServe()
}