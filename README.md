# Trip Tracker Project (Mia's Trips)

**Go Version**: 1.23.5
**Type**: Dual-platform application - Web (PWA) + React Native Mobile App

This project is a comprehensive travel tracking application with both web and mobile interfaces. The Go backend serves a responsive Progressive Web Application (PWA) built with HTMX and Templ, while also providing JWT-based REST API endpoints for a dedicated React Native mobile app. This dual-platform approach allows users to create, edit, delete, and visualize trips on a world map across any device with native mobile capabilities.

**Web Experience**: Full desktop/mobile browser experience with PWA features
**Mobile Experience**: Native React Native app for iOS and Android
**Single Backend**: One Go server serves both platforms with different authentication methods

## Features

### Web Application (PWA)
- **Progressive Web App**: Installable on any device, works offline
- **Responsive Design**: Adapts to desktop, tablet, and mobile screens
- **Real-time Updates**: HTMX-powered dynamic content without page reloads
- **Session Authentication**: Google OAuth with server-side sessions
- **PWA Features**: Service worker caching, offline support, app shortcuts
- **Desktop Navigation**: Traditional top navigation bar for larger screens
- **Mobile Navigation**: Bottom tab navigation optimized for touch

### Mobile Application (React Native)
- **Cross-Platform**: Single codebase for iOS and Android using Expo
- **TypeScript**: Type-safe development with full TypeScript support
- **JWT Authentication**: Secure token-based authentication with refresh tokens
- **Redux State Management**: Centralized state with Redux Toolkit
- **Native Features**: Platform-specific UI components and interactions
- **Offline Capability**: Data synchronization when connection is restored
- **Maps Integration**: Interactive maps to visualize trip locations

### Shared Features
- **Trip Management**: Create, edit, delete, and organize travel trips
- **World Map Visualization**: Interactive maps showing visited locations and flight paths
- **Airport Database**: Comprehensive airport data with IATA codes
- **Flight Information**: Real-time flight data via AviationStack API
- **Statistics Dashboard**: Trip analytics, countries visited, miles flown
- **Google OAuth**: Unified authentication across both platforms

## Tech Stack

### Backend (Shared)
- **Language**: Go 1.23.5
- **HTTP Server**: Built-in `net/http` package
- **Database**: SQLite with comprehensive schema
- **Authentication**: Dual system - Sessions (web) + JWT (mobile)
- **External APIs**: Google OAuth, AviationStack flight data

### Web Frontend (PWA)
- **Templates**: Templ (Go template engine)
- **Real-time**: HTMX for dynamic updates
- **CSS**: Tailwind CSS with mobile-optimized styles
- **JavaScript**: Leaflet.js for maps, custom PWA features
- **PWA**: Service worker, web manifest, offline functionality

### Mobile App (React Native)
- **Framework**: React Native with Expo
- **Language**: TypeScript
- **State Management**: Redux Toolkit + React Redux
- **Navigation**: React Navigation v6
- **API Client**: Axios with JWT interceptors
- **Storage**: Expo SecureStore for sensitive data
- **Maps**: React Native Maps
- **UI**: Custom components matching web theme

## Project Structure

```
trip-tracker/
├── main.go                          # Application entry point with dual routing
├── .air.toml                        # Air configuration for hot reload
├── docker-compose.yml               # Docker deployment setup
├── Dockerfile                       # Container configuration
├── tailwind.config.js               # Tailwind CSS configuration
│
├── internal/
│   ├── api/
│   │   ├── flights.go               # AviationStack API integration
│   │   └── google_*.go              # Google OAuth handlers (web)
│   │
│   ├── database/                    # Database layer
│   │   ├── *.go                     # Database store files
│   │   ├── schema.sql               # Table definitions
│   │   └── database.db              # SQLite database file
│   │
│   ├── handlers/                    # Web HTTP handlers (HTMX/Session-based)
│   │   ├── get*.go                  # GET request handlers
│   │   ├── post*.go                 # POST request handlers
│   │   ├── edit*.go                 # PUT/PATCH request handlers
│   │   ├── delete*.go               # DELETE request handlers
│   │   ├── pwa.go                   # PWA-specific handlers
│   │   └── mobile/                  # Mobile API handlers (JWT-based)
│   │       ├── auth.go              # Mobile Google OAuth + JWT
│   │       ├── refresh.go           # JWT token refresh
│   │       ├── middleware.go        # JWT authentication middleware
│   │       ├── common.go            # Shared API response structures
│   │       ├── gettrips.go          # GET /api/v1/trips
│   │       ├── posttrips.go         # POST /api/v1/trips
│   │       ├── puttrips.go          # PUT /api/v1/trips/{id}
│   │       ├── deletetrips.go       # DELETE /api/v1/trips/{id}
│   │       └── getprofile.go        # GET /api/v1/profile
│   │
│   ├── models/                      # Data models
│   │   └── *.go                     # User, Trip, Airport models
│   │
│   └── middleware/                  # HTTP middleware
│       └── middleware.go            # Auth, CSP, logging, HTMX headers
│
├── static/                          # Static assets for web
│   ├── manifest.json                # PWA web app manifest
│   ├── css/
│   │   ├── input.css                # Tailwind CSS input
│   │   ├── output.css               # Generated CSS
│   │   └── mobile.css               # Mobile-optimized PWA styles
│   ├── icons/                       # PWA icons (72x72 to 512x512)
│   ├── js/
│   │   ├── htmx.min.js              # HTMX library
│   │   ├── leaflet.js               # Map library
│   │   ├── map.js                   # Map configuration
│   │   ├── pwa-features.js          # PWA functionality
│   │   └── sw.js                    # Service worker
│   └── images/                      # Static images
│
├── templates/                       # Templ template files (web only)
│   ├── layout.templ                 # Base layout with PWA support
│   ├── trips.templ                  # Trip-related components
│   ├── *.templ                      # Other page templates
│   └── *_templ.go                   # Generated Go files
│
└── app/                             # React Native mobile app
    ├── package.json                 # Mobile app dependencies (locked versions)
    ├── app.json                     # Expo configuration
    ├── app.config.js                # Dynamic config with ngrok detection
    ├── tsconfig.json                # TypeScript configuration
    ├── .npmrc                       # Exact version enforcement
    ├── App.tsx                      # Mobile app entry point
    │
    ├── src/
    │   ├── api/                     # API client layer
    │   │   ├── client.ts            # Axios configuration with JWT
    │   │   ├── types.ts             # TypeScript interfaces
    │   │   └── index.ts             # API exports
    │   │
    │   ├── components/              # Reusable mobile components
    │   │   ├── common/              # Shared UI components
    │   │   │   ├── Button.tsx
    │   │   │   ├── Card.tsx
    │   │   │   ├── Input.tsx
    │   │   │   └── Loading.tsx
    │   │   └── trips/               # Trip-specific components
    │   │       └── TripCard.tsx
    │   │
    │   ├── screens/                 # Mobile screen components
    │   │   ├── auth/                # Authentication screens
    │   │   │   ├── LoginScreen.tsx
    │   │   │   └── SplashScreen.tsx
    │   │   ├── trips/               # Trip management screens
    │   │   │   └── TripsListScreen.tsx
    │   │   └── profile/             # Profile screens
    │   │       └── ProfileScreen.tsx
    │   │
    │   ├── navigation/              # Navigation configuration
    │   │   └── AppNavigator.tsx     # Root navigator with auth flow
    │   │
    │   ├── store/                   # Redux store
    │   │   ├── index.ts             # Store setup
    │   │   ├── hooks.ts             # Typed Redux hooks
    │   │   └── slices/              # Redux slices
    │   │       ├── authSlice.ts     # Authentication state
    │   │       ├── tripsSlice.ts    # Trips state
    │   │       └── uiSlice.ts       # UI state
    │   │
    │   └── utils/                   # Utility functions
    │       └── theme.ts             # Dark theme matching website
    │
    ├── scripts/                     # Development scripts
    │   ├── start-dev.sh             # Complete development startup
    │   └── cleanup.sh               # Process cleanup
    │
    └── assets/                      # Expo assets
        ├── icon.png                 # App icon
        └── splash-icon.png          # Splash screen icon
```

## Database Schema

**Database**: SQLite (`internal/database/database.db`)

### Tables
1. **airports**: Static airport data with IATA codes, coordinates
2. **users**: User accounts with Google OAuth data
3. **trips**: Trip records with origin/destination airports
4. **sessions**: User session management (web only)
5. **password_reset_tokens**: 1-hour expiry tokens

### Key Relationships
- `trips.user_id` → `users.id`
- `trips.origin` → `airports.iata_code`
- `trips.destination` → `airports.iata_code`
- `sessions.user_id` → `users.id`

## Development Setup

### Prerequisites
- **Go 1.23+**
- **Node.js 18+**
- **SQLite3**
- **Expo CLI**: `npm install -g @expo/cli`
- **Air** (Go hot reload): `go install github.com/cosmtrek/air@latest`
- **Tailwind CLI binary**
- **Templ CLI**: `go install github.com/a-h/templ/cmd/templ@latest`

### Environment Variables

**Backend (.env in root):**
```bash
APP_PORT=3000                                    # Optional, defaults to 3000
JWT_SECRET=your-super-secure-jwt-secret         # Optional, auto-generated if not set
GOOGLE_OAUTH_CLIENT_ID=your-google-client-id
GOOGLE_OAUTH_CLIENT_SECRET=your-google-client-secret
```

**Mobile App (app/.env):**
```bash
EXPO_PUBLIC_DEV_API_URL=http://0.0.0.0:3000  # Your computer's IP for physical devices
EXPO_PUBLIC_PRODUCTION_API_URL=https://your-domain.com
```

### Quick Start

#### 1. Backend Setup (One-time)
```bash
# Install Go dependencies
go mod download

# Create database
sqlite3 internal/database/database.db < internal/database/schema.sql

# Generate CSS and templates
./tailwindcss -i ./static/css/input.css -o ./static/css/output.css
templ generate
```

#### 2. Start Backend Development
```bash
# Hot reload development server
air
# Backend will run on http://localhost:3000
```

#### 3. Mobile App Development
```bash
# Navigate to mobile app
cd app

# Install dependencies (first time)
npm install

# Start complete development environment
npm run dev
```

This will:
1. Start ngrok tunnel for public HTTPS access
2. Start Go backend with hot reload
3. Start Expo development server
4. Auto-configure API URL for mobile app

### Development Commands

**Backend:**
```bash
go run main.go                    # Basic server start
air                              # Hot reload development
./tailwindcss -i ./static/css/input.css -o ./static/css/output.css  # Build CSS
templ generate                   # Generate templates
```

**Mobile App:**
```bash
npm run dev                      # Complete development environment
npm run cleanup                 # Stop all processes
npm start                       # Expo only
npm run android                 # Android emulator
npm run ios                     # iOS simulator (macOS only)
npm run web                     # Web browser
```

### Testing

**Web Application:**
1. **Desktop**: `http://localhost:3000`
2. **PWA Features**: Use ngrok for HTTPS testing
3. **Mobile Browser**: Test responsive design

**Mobile App:**
1. **Expo Go**: Scan QR code on physical device
2. **Emulator**: Press 'a' (Android) or 'i' (iOS)
3. **Physical Device**: Ensure same WiFi network

**API Testing:**
```bash
# Test mobile endpoints
curl -X POST http://localhost:3000/api/v1/mobile/auth/google \
  -H "Content-Type: application/json" \
  -d '{"google_token": "mock-google-token-development"}'

# Test JWT protected endpoint
curl -X GET http://localhost:3000/api/v1/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Architecture Patterns

### Dual Authentication System

**Web Application (Session-based):**
- Google OAuth flow with server-side sessions
- Session cookies for authentication
- HTMX responses with HTML fragments
- CSRF protection via Content Security Policy

**Mobile Application (JWT-based):**
- Google OAuth with JWT token exchange
- Access tokens (15 min) + Refresh tokens (7 days)
- JSON API responses
- Automatic token refresh via interceptors

### Handler Organization

**Naming Convention**: `{method}{resource}.go`
- `gethome.go` - GET request for home page
- `posttrip.go` - POST request to create trip
- `mobile/gettrips.go` - GET /api/v1/trips (JSON)
- `mobile/auth.go` - Mobile authentication

**Response Patterns:**
```go
// Web handler (returns HTML)
err := templates.ComponentName(data).Render(r.Context(), w)

// Mobile handler (returns JSON)
response := mobile.ApiResponse{
    Success: true,
    Data:    result,
}
json.NewEncoder(w).Encode(response)
```

### PWA Implementation

**Progressive Enhancement:**
- Works as website without PWA features
- Enhanced experience when installed
- Offline-first with service worker caching
- Mobile navigation appears only on small screens

**Key Files:**
- `static/manifest.json` - PWA manifest
- `static/js/sw.js` - Service worker
- `handlers/pwa.go` - PWA-specific endpoints
- `templates/layout.templ` - PWA meta tags

## Google OAuth Setup

### Web Application
1. **Google Cloud Console** → Create OAuth 2.0 Client
2. **Application Type**: Web application
3. **Authorized Redirect URIs**: `http://localhost:3000/auth/google/callback`
4. **Environment Variables**: Set `GOOGLE_OAUTH_CLIENT_ID` and `GOOGLE_OAUTH_CLIENT_SECRET`

### Mobile Application
1. **iOS Client ID**: For iOS app builds
2. **Android Client ID**: For Android app builds
3. **Web Client ID**: For Expo Go development
4. **App Configuration**: Update `app.config.js` with platform-specific client IDs

**Development Mode:**
The mobile app includes mock authentication that bypasses Google OAuth:
- Use token: `"mock-google-token-development"`
- Creates user: `dev@example.com`
- Returns real JWT tokens for testing

## Deployment

### Production Build

**Backend:**
```bash
# Build production binary
go build -o trip-tracker main.go

# Docker build and deploy
./docker_build.sh YOUR_GITHUB_PAT
```

**Mobile App:**
```bash
# Build for app stores
cd app
eas build --platform all

# Preview build
eas build --profile preview

# Development build (for testing)
eas build --profile development
```

### Environment Configuration

**Development:**
- Backend: `http://localhost:3000`
- Mobile: Local IP for physical device testing
- Database: Local SQLite file

**Production:**
- Backend: HTTPS domain with SSL termination
- Mobile: Production API URL in app config
- Database: Production SQLite or external database

## Development Progress & Roadmap

### Completed Features
- **Foundation**: Go backend with dual routing architecture
- **Web PWA**: Complete responsive web application with offline support
- **Mobile App**: React Native app with Expo, Redux state management
- **Authentication**: Dual auth system (sessions + JWT) with Google OAuth
- **API Integration**: RESTful JSON API for mobile consumption
- **Database**: SQLite schema with proper relationships
- **Trip Management**: Full CRUD operations for both platforms
- **Profile System**: User profile display and management
- **Security**: Dependency version locking, proper error handling
- **Development Workflow**: Automated development environment setup

### Next Priorities

**Phase 1: Core Mobile Features**
1. Trip creation and editing on mobile
2. Maps integration with React Native Maps
3. Image upload and trip photos
4. Push notifications for trip reminders
5. Offline synchronization improvements

**Phase 2: Advanced Features**
1. Trip sharing capabilities
2. Statistics and analytics dashboard
3. Calendar integration
4. Flight status tracking
5. Travel document storage

**Phase 3: Production Readiness**
1. App store deployment (iOS/Android)
2. Production OAuth configuration
3. Performance optimization
4. Error tracking and analytics
5. User feedback system

**Phase 4: Enhancements**
1. Multi-user support
2. Trip collaboration features
3. Social sharing integration
4. Advanced reporting
5. API rate limiting and caching

### Technical Debt & Improvements
1. **Middleware Chaining**: Fix middleware integration issues
2. **Test Coverage**: Add comprehensive test suites
3. **Documentation**: API documentation with OpenAPI/Swagger
4. **Monitoring**: Production logging and metrics
5. **Backup**: Automated database backup system

## Code Conventions

### File Organization
- **One responsibility per file**: Each handler handles one HTTP endpoint
- **Clear naming**: Files named after HTTP method + resource
- **Separation of concerns**: Web handlers vs mobile API handlers
- **Template co-location**: Related Templ components in same file

### Error Handling
```go
// Web handler error response
if err != nil {
    http.Error(w, "Error message", http.StatusInternalServerError)
    return
}

// Mobile API error response
response := mobile.ApiResponse{
    Success: false,
    Error: &mobile.ErrorResponse{
        Code:    "ERROR_CODE",
        Message: "Human readable message",
    },
}
```