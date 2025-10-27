package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/skywall34/trip-tracker/internal/api"
	"github.com/skywall34/trip-tracker/internal/database"
	"github.com/skywall34/trip-tracker/internal/handlers"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
)

func main() {

	dotenvPath := os.Getenv("DOTENV_PATH")
	if dotenvPath == "" {
		dotenvPath = ".env" // default if not set
	}

	if err := godotenv.Load(dotenvPath); err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	db, err := database.InitDB("file:./internal/database/database.db?_enable_math_functions=1")
	if err != nil {
		log.Fatal(err)
	}

	// Load the countries from a JSON file into memory
	if err := models.LoadCountriesFromFile("./static/data/countries.json"); err != nil {
		log.Fatalf("Failed to load countries: %v", err)
	}
	if err := models.LoadAirportTimezonesFromFile("./static/data/airport2timezone.json"); err != nil {
		log.Fatalf("Failed to load airport timezones: %v", err)
	}

	userStore := database.NewUserStore(database.NewUserStoreParams{DB: db})
	tripStore := database.NewTripStore(database.NewTripStoreParams{DB: db})
	sessionStore := database.NewSessionStore(database.NewSessionStoreParams{DB: db})
	passwordResetStore := database.NewPasswordResetStore(database.PasswordResetStoreParams{DB: db})
	placeStore := database.NewPlaceStore(db)
	//TODO: Chaining middleware seems to break css for some reason
	authMiddleware := m.NewAuthMiddleware(sessionStore, "session_id")

	// Google OAuth Initilization to Add the Environemnt Variables
	googleOauthConfig := api.NewGoogleOauthConfig()

	// Email Service for Resetting Passwords
	// Placeholders for now
	gmailUser := os.Getenv("GMAIL_SERVICE_APP_USERNAME")
	gmailPsw := os.Getenv("GMAIL_SERVICE_APP_PASSWORD")
	emailService := models.EmailService{
		SMTPHost: "smtp.gmail.com",
		SMTPPort: 587,
		Username: gmailUser,
		Password: gmailPsw, // use App Password, not real password
		From:     gmailUser,
	}

	//

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// PWA routes
	mux.Handle("/manifest.json", handlers.NewPWAManifestHandler())
	mux.Handle("/sw.js", handlers.NewServiceWorkerHandler())
	mux.Handle("/offline", handlers.NewOfflineHandler())

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

	mux.Handle("PUT /trips",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewEditTripHandler(
							handlers.EditTripHandlerParams{
								TripStore: tripStore}).ServeHTTP)))))

	mux.Handle("DELETE /trips",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewDeleteTripHandler(
							handlers.DeleteTripHandlerParams{
								TripStore: tripStore}).ServeHTTP)))))

	// Places Routes
	mux.Handle("GET /places",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewGetPlacesHandler(
							handlers.GetPlacesHandlerParams{
								PlaceStore: placeStore,
								TripStore:  tripStore,
							}).ServeHTTP)))))

	mux.Handle("POST /places",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewPostPlaceHandler(
							handlers.PostPlaceHandlerParams{
								PlaceStore: placeStore,
							}).ServeHTTP)))))

	mux.Handle("PUT /places",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewPutPlaceHandler(
							handlers.PutPlaceHandlerParams{
								PlaceStore: placeStore,
							}).ServeHTTP)))))

	mux.Handle("DELETE /places",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewDeletePlaceHandler(
							handlers.DeletePlaceHandlerParams{
								PlaceStore: placeStore,
							}).ServeHTTP)))))

	mux.Handle("GET /editplaceform",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewGetEditPlaceFormHandler(
							handlers.GetEditPlaceFormHandlerParams{
								PlaceStore: placeStore,
							}).ServeHTTP)))))

	// Google Places API routes
	mux.Handle("GET /api/places/search",
		authMiddleware.AddUserToContext(
			m.LoggingMiddleware(
				handlers.NewGetPlaceSearchHandler(
					handlers.GetPlaceSearchHandlerParams{}).ServeHTTP)))

	mux.Handle("GET /api/places/details",
		authMiddleware.AddUserToContext(
			m.LoggingMiddleware(
				handlers.NewGetPlaceDetailsHandler(
					handlers.GetPlaceDetailsHandlerParams{}).ServeHTTP)))

	mux.Handle("GET /api/places/modal",
		authMiddleware.AddUserToContext(
			m.LoggingMiddleware(
				handlers.NewGetPlaceModalHandler(
					handlers.GetPlaceModalHandlerParams{}).ServeHTTP)))

	mux.Handle("GET /api/places/modal/close",
		authMiddleware.AddUserToContext(
			m.LoggingMiddleware(
				handlers.NewGetPlaceModalHandler(
					handlers.GetPlaceModalHandlerParams{}).ServeHTTP)))

	mux.Handle("GET /api/places/filter",
		authMiddleware.AddUserToContext(
			m.LoggingMiddleware(
				handlers.NewGetPlaceFilterHandler(
					handlers.GetPlaceFilterHandlerParams{
						PlaceStore: placeStore,
						TripStore:  tripStore,
					}).ServeHTTP)))

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
							UserStore:    userStore,
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
						UserStore:    userStore,
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

	mux.Handle("GET /worldmap3d",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewGetWorldMap3dHandlerHandler().ServeHTTP)))))

	mux.Handle("GET /createtripform",
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewGetCreateTripHandler().ServeHTTP)))))

	mux.Handle("GET /edittripform", // Edit trip
		authMiddleware.AddUserToContext(
			m.CSPMiddleware(
				m.TextHTMLMiddleware(
					m.LoggingMiddleware(
						handlers.NewGetEditTripHandlerParmas(
							handlers.GetEditTripHandlerParams{
								TripStore: tripStore,
							}).ServeHTTP)))))

	mux.Handle("GET /forgot-password",
		m.CSPMiddleware(
			m.TextHTMLMiddleware(
				m.LoggingMiddleware(
					handlers.NewGetForgotPasswordHandler().ServeHTTP))))

	mux.Handle("GET /reset-password",
		m.CSPMiddleware(
			m.TextHTMLMiddleware(
				m.LoggingMiddleware(
					handlers.NewGetResetPasswordHandlerParams().ServeHTTP))))

	// API CALLS TODO: Possibly version
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

	mux.Handle("POST /api/forgot-password",
		m.CSPMiddleware(
			m.LoggingMiddleware(
				handlers.NewPostForgotPasswordHandler(
					handlers.PostForgotPasswordHandlerParams{
						UserStore:          userStore,
						PasswordResetStore: passwordResetStore,
						EmailService:       emailService,
					}).ServeHTTP)))

	mux.Handle("POST /api/reset-password",
		m.CSPMiddleware(
			m.LoggingMiddleware(
				handlers.NewPostResetPasswordHandler(
					handlers.PostResetPasswordHandlerParams{
						UserStore:          userStore,
						PasswordResetStore: passwordResetStore,
					}).ServeHTTP)))

	// Google Auth
	mux.HandleFunc("/auth/google/login",
		api.NewGoogleLoginHandlerParams(
			api.GoogleLoginHandlerParams{
				GoogleOauthConfig: googleOauthConfig,
			}).ServeHTTP)

	mux.HandleFunc("/auth/google/callback",
		api.NewGoogleCallbackHandlerParams(
			api.GoogleCallbackHandlerParams{
				UserStore:         userStore,
				SessionStore:      sessionStore,
				GoogleOauthConfig: googleOauthConfig,
			}).ServeHTTP)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", appPort),
		Handler: mux,
	}

	fmt.Printf("Server running on port :%s\n", appPort)
	server.ListenAndServe()
}
