# Mia's Trips - Claude Code Context

This document provides comprehensive context for Claude Code to assist with development of the Mia's Trips travel tracking application.

## Project Overview

**Name**: Trip Tracker Project (Mia's Trips)  
**Go Version**: 1.23.5  
**Type**: Dual-platform application - Web application + React Native mobile app  
**Primary User**: Single user (Mia) with potential for expansion

**Description**: A comprehensive travel tracking application with both web and mobile interfaces. The Go backend serves a responsive web application built with HTMX and Templ, while also providing JWT-based REST API endpoints for a dedicated React Native mobile app. This dual-platform approach allows users to create, edit, delete, and visualize trips on a world map across any device with native mobile capabilities.

## Tech Stack

### Backend

- **Language**: Go 1.23.5
- **HTTP Server**: Built-in `net/http` package
- **Database**: SQLite
- **Templates**: Templ (Go template engine)
- **Real-time Updates**: HTMX (web only)
- **Authentication**: Dual system - Session-based (web) + JWT (mobile)
- **API Architecture**: REST endpoints with JSON responses for mobile

### Web Frontend

- **CSS Framework**: Tailwind CSS + mobile-optimized CSS
- **JavaScript**: HTMX + custom JS files + PWA features
- **Maps**: Leaflet.js for world map visualization
- **UI Pattern**: Server-side rendering with HTMX for dynamic updates
- **PWA**: Service worker, web manifest, offline functionality
- **Authentication**: Session-based with Google OAuth

### Mobile App (React Native + Expo)

- **Framework**: React Native with Expo
- **Language**: TypeScript
- **State Management**: Redux Toolkit + React Redux
- **Navigation**: React Navigation v6
- **API Client**: Axios with JWT interceptors
- **Storage**: Expo SecureStore for sensitive data
- **Maps**: React Native Maps
- **Authentication**: JWT-based with Google OAuth
- **UI Components**: Custom components matching web theme

### External APIs

- **Flight Data**: AviationStack API (100 calls/month free tier)
- **Authentication**: Google OAuth

### Development Tools

**Backend:**
- **Hot Reload**: Air for Go code auto-reload
- **CSS Build**: Tailwind CLI for CSS compilation
- **Template Generation**: Templ CLI

**Mobile:**
- **Development**: Expo CLI for React Native development
- **Testing**: Expo Go for physical device testing
- **Build**: EAS Build for production builds

## Project Structure

```
trip-tracker/
├── main.go                          # Application entry point
├── docker_build.sh                  # Script to build and load the docker image to gcr
├── .air.toml                        # Air configuration for hot reload
├── docker-compose.yml               # Docker deployment setup
├── Dockerfile                       # Container configuration
├── tailwind.config.js               # Tailwind CSS configuration
├── test-instructions.md             # PWA mobile testing guide
├── internal/
|   ├── api/
|   |   |── flights.go               # Aviation Flight API functions
|   |   └── google_*.go              # Google Auth handlers (web)
│   ├── database/                    # Database layer
│   │   ├── *.go                     # Database store files
│   │   ├── schema.sql               # Table definitions
|   |   └── database.db              # SQLite database file
│   ├── handlers/                    # Web HTTP handlers (HTMX/Session-based)
│   │   ├── get*.go                  # GET request handlers
│   │   ├── post*.go                 # POST request handlers
│   │   ├── edit*.go                 # PUT/PATCH request handlers
│   │   ├── delete*.go               # DELETE request handlers
│   │   ├── pwa.go                   # PWA-specific handlers (manifest, service worker, offline)
│   │   └── mobile/                  # Mobile API handlers (JWT-based)
│   │       ├── auth.go              # Mobile Google OAuth + JWT
│   │       ├── refresh.go           # JWT token refresh
│   │       ├── middleware.go        # JWT authentication middleware
│   │       ├── common.go           # Shared API response structures
│   │       ├── gettrips.go         # GET /api/v1/trips
│   │       ├── posttrips.go        # POST /api/v1/trips
│   │       ├── puttrips.go         # PUT /api/v1/trips/{id}
│   │       └── deletetrips.go      # DELETE /api/v1/trips/{id}
│   │
│   └── middleware/                  # HTTP middleware
│       └── middleware.go            # Authentication, CSP, logging, HTMX headers + PWA nonces
├── static/                          # Static assets
│   ├── manifest.json                # PWA web app manifest
│   ├── css/
│   │   ├── input.css                # Tailwind CSS input
|   |   ├── leaflet.css              # Leaflet CSS input
│   │   ├── output.css               # Generated CSS
│   │   └── mobile.css               # Mobile-optimized PWA styles
│   ├── icons/                       # PWA icons (72x72 to 512x512)
│   │   ├── icon-*.png               # App icons for all sizes
│   │   ├── apple-touch-icon.png     # iOS home screen icon
│   │   └── favicon-*.png            # Browser favicons
│   ├── js/
│   │   ├── htmx.min.js              # HTMX library
│   │   ├── convertTimes.js          # UTC time conversion
│   │   ├── leaflet.js               # Map library
│   │   ├── map.js                   # Map configuration
│   │   ├── modal.js                 # Modal show/hide logic
│   │   ├── response-targets.js      # HTMX response targeting
│   │   ├── tabs.js                  # Tab sliding animations
│   │   ├── pwa-features.js          # PWA functionality (geolocation, camera, pull-to-refresh)
│   │   └── sw.js                    # Service worker for caching and offline support
│   └── images/                      # Static images and assets
├── templates/                       # Templ template files (web only)
│   ├── layout.templ                 # Base layout template (includes PWA meta tags, mobile nav)
│   ├── trips.templ                  # Trip-related components (TripsPage removed)
│   ├── *.templ                      # Other page templates
│   └── *_templ.go                   # Generated Go files (don't edit, generate by templ generate)
└── app/                             # React Native mobile app
    ├── package.json                 # Mobile app dependencies
    ├── app.json                     # Expo configuration
    ├── tsconfig.json                # TypeScript configuration
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
    │   │   └── trips/               # Trip management screens
    │   │       └── TripsListScreen.tsx
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
    └── assets/                      # Expo assets
        ├── icon.png                 # App icon
        └── splash-icon.png          # Splash screen icon
```

## Database Schema

**Database**: SQLite (`database/database.db`)

### Tables

1. **airports**: Static airport data from CSV files
2. **users**: User account information
3. **trips**: Trip records with origin/destination airports
4. **sessions**: User session management
5. **password_reset_tokens**: 1-hour expiry tokens for password reset

### Key Relationships

- `trips.user_id` → `users.id`
- `trips.origin` → `airports.iata_code`
- `trips.destination` → `airports.iata_code`
- `sessions.user_id` → `users.id`

## Code Conventions

### Handler Naming Convention

- **Pattern**: `{method}{resource}.go`
- **Web Examples**:
  - `gethome.go` - GET request for home page
  - `posttrip.go` - POST request to create trip
  - `edittrip.go` - PUT request to update trip
  - `deletetrip.go` - DELETE request to remove trip
  - `pwa.go` - PWA-specific handlers (manifest, service worker, offline page)
- **Mobile API Examples**:
  - `gettrips.go` - GET /api/v1/trips (JSON response)
  - `posttrips.go` - POST /api/v1/trips (JSON request/response)
  - `puttrips.go` - PUT /api/v1/trips/{id} (JSON request/response)
  - `deletetrips.go` - DELETE /api/v1/trips/{id} (JSON response)

### Dual Architecture: Web + Mobile API

#### Web Application (HTMX + Session Auth)
- **Authentication**: Session-based with Google OAuth
- **Response Format**: HTML templates with HTMX headers
- **URL Pattern**: `/trips`, `/login`, etc.
- **State Management**: Server-side sessions
- **Real-time Updates**: HTMX for dynamic content

#### Mobile API (REST + JWT Auth)  
- **Authentication**: JWT tokens with Google OAuth
- **Response Format**: JSON with standardized `ApiResponse` structure
- **URL Pattern**: `/api/v1/trips`, `/api/v1/mobile/auth/google`, etc.
- **State Management**: Redux store with secure token storage
- **Handler Separation**: Dedicated mobile handlers in `internal/handlers/mobile/`

### PWA Implementation Notes (Web Only)

- **Progressive Enhancement**: Features gracefully degrade without PWA support
- **Mobile Navigation**: Bottom nav appears only on mobile devices (`md:hidden`)
- **Responsive Layout**: Desktop uses top nav, mobile uses bottom nav
- **Offline-First**: Service worker caches key pages and assets
- **HTTPS Required**: PWA installation only works over HTTPS in production

### Handler Structure Template

```go
// Handler struct with dependencies
type [Action][Resource]Handler struct {
    [resource]Store *db.[Resource]Store
}

// Constructor parameters struct
type [Action][Resource]HandlerParams struct {
    [Resource]Store *db.[Resource]Store
}

// Constructor function
func New[Action][Resource]Handler(params [Action][Resource]HandlerParams) *[Action][Resource]Handler {
    return &[Action][Resource]Handler{
        [resource]Store: params.[Resource]Store,
    }
}

// HTTP handler method
func (h *[Action][Resource]Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Handler logic here
}
```

### Template Organization

- **One file per page/feature**: All related components in same `.templ` file
- **Example**: `trips.templ` contains `RenderTrips`, `RenderPastTrips`, `EditTripForm`, etc.
- **PWA Layout**: `layout.templ` includes PWA meta tags, mobile navigation, and service worker registration
- **Layout Usage**:
  ```go
  c := templates.HomePage()
  templates.Layout(c, "Page Title").Render(r.Context(), w)
  ```
- **HTMX Fragments**: Many handlers now return partial templates for dynamic updates
- **Mobile Responsive**: Templates adapt to device size with CSS classes

## Middleware Stack

### Current Middleware (applied individually, not chained)

1. **Auth**: Session-based authentication
2. **TextHTML**: Sets HTML content type for templ responses  
3. **CSP**: Content Security Policy for XSS protection (includes PWA nonces)
4. **Logging**: Request/response logging

### CSP Enhancement for PWA

The CSP middleware now includes nonces for PWA-specific scripts:
- PWA installation and service worker registration
- Mobile-specific JavaScript features
- All nonces are dynamically generated per request

### Known Issue

Middleware chaining breaks CSP and TextHTML middleware functionality.

## Key Features

### Unified Web + Mobile Experience

- **Responsive Design**: Adapts to desktop, tablet, and mobile screens
- **Progressive Enhancement**: Works as website, enhanced as mobile app
- **Cross-Platform**: Single codebase serves all devices
- **Offline Capability**: Core functionality available without internet

### Authentication

- **Google OAuth**: Primary authentication method
- **Session-based**: Server-side session management  
- **Files**: `google_auth.go`, `google_login.go`, `google_callback.go`

### Trip Management

- **CRUD Operations**: Create, read, update, delete trips
- **Airport Integration**: IATA codes for origin/destination
- **Flight Data**: AviationStack API for real-time flight information
- **Mobile Optimized**: Touch-friendly forms and interactions

### Visualization

- **World Map**: Leaflet.js integration showing visited locations
- **Statistics**: Trip counts, countries visited, miles flown
- **Responsive Charts**: Adapt to screen size and orientation

### HTMX Integration

- **Dynamic Updates**: Page sections update without full reload
- **Form Handling**: HTMX-powered form submissions
- **Response Targeting**: Custom JS for flexible response handling
- **Mobile Navigation**: HTMX loads content in mobile-optimized layout

### PWA Features

- **Service Worker**: Caches assets and enables offline functionality
- **Web Manifest**: Enables installation on mobile devices
- **Background Sync**: Queues actions when offline, syncs when online
- **Native Features**: Geolocation, camera access, pull-to-refresh
- **App Shortcuts**: Quick actions from installed app icon

## Environment Configuration

### Required Environment Variables

**Backend:**
- `APP_PORT` (optional): Server port, defaults to 3000
- `JWT_SECRET` (optional): JWT signing secret, auto-generated if not set
- `GOOGLE_OAUTH_CLIENT_ID`: Google OAuth client ID
- `GOOGLE_OAUTH_CLIENT_SECRET`: Google OAuth client secret

**Mobile App:**
- `API_URL`: Backend server URL (e.g., `http://192.168.1.100:3000` for development)
- `GOOGLE_MAPS_API_KEY`: Google Maps API key for React Native Maps

### Development Setup

**Backend:**
1. **Install Go 1.23+**
2. **Install SQLite3**
3. **Install Tailwind CLI binary**
4. **Install Templ CLI**
5. **Install Air for hot reload**

**Mobile App:**
1. **Install Node.js 18+**
2. **Install Expo CLI**: `npm install -g @expo/cli`
3. **Install EAS CLI**: `npm install -g eas-cli` (for builds)
4. **iOS Development**: Xcode (macOS only)
5. **Android Development**: Android Studio with AVD

### Development Commands

**Backend Setup (One-time):**
```bash
# Create database if it doesn't exist
sqlite3 database.db < internal/database/schema.sql
```

**Backend Development:**
```bash
# Generate Tailwind CSS (includes mobile styles)
./tailwindcss -i ./static/css/input.css -o ./static/css/output.css

# Generate templ files (includes PWA templates)
templ generate

# Start development server with hot reload
air
# Backend will run on http://localhost:3000
```

**Mobile App Development:**
```bash
# Navigate to mobile app directory
cd app

# Install dependencies (first time)
npm install

# Start Expo development server
npx expo start

# Development options:
# - Press 'a' for Android emulator
# - Press 'i' for iOS simulator (macOS only)
# - Press 'w' for web browser
# - Scan QR code with Expo Go app on physical device
```

**Mobile Testing on Physical Devices:**
```bash
# 1. Find your computer's IP address
ip addr show | grep "inet " | grep -v 127.0.0.1
# Example result: 192.168.1.100

# 2. Update app/app.json with your IP:
# "extra": { "apiUrl": "http://192.168.1.100:3000" }

# 3. Start both backend and mobile app
# Terminal 1: (project root)
air

# Terminal 2: (in app/ directory)  
npx expo start

# 4. Install Expo Go on your phone and scan QR code
```

**Testing Commands:**
```bash
# Web application
open http://localhost:3000

# Mobile API endpoints (for testing)
curl http://localhost:3000/api/v1/trips -H "Authorization: Bearer <JWT_TOKEN>"

# Full PWA testing (requires HTTPS)
air &
ngrok http 3000
# Use the https://xxx.ngrok.io URL for full PWA features
```

## Common Patterns

### Web Handler Patterns

**Error Handling:**
```go
if err != nil {
    http.Error(w, "Error message", http.StatusInternalServerError)
    return
}
```

**Context Usage (Session Auth):**
```go
ctx := r.Context()
userID, ok := ctx.Value(m.UserKey).(int)
if !ok {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
}
```

**Template Rendering:**
```go
err := templates.ComponentName(data).Render(r.Context(), w)
if err != nil {
    http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
    return
}
```

### Mobile API Patterns

**JWT Authentication:**
```go
userID, ok := mobile.GetUserIDFromContext(r.Context())
if !ok {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
}
```

**JSON Response:**
```go
response := mobile.ApiResponse{
    Success: true,
    Data:    result,
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(response)
```

**JSON Error Response:**
```go
response := mobile.ApiResponse{
    Success: false,
    Error: &mobile.ErrorResponse{
        Code:    "ERROR_CODE",
        Message: "Human readable error message",
    },
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusBadRequest) // or appropriate status
json.NewEncoder(w).Encode(response)
```

### Mobile App (React Native) Patterns

**Redux Usage:**
```typescript
const dispatch = useDispatch<AppDispatch>();
const { trips, isLoading, error } = useSelector((state: RootState) => state.trips);

// Async action
await dispatch(fetchTrips({ page: 1, refresh: true })).unwrap();
```

**API Client:**
```typescript
// Automatic JWT handling
const response = await apiClient.get('/api/v1/trips');
// JWT tokens are automatically added via interceptors
```

**Theme Usage:**
```typescript
import { colors, typography, spacing } from '../utils/theme';

const styles = StyleSheet.create({
  container: {
    backgroundColor: colors.background,
    padding: spacing.md,
  },
  text: {
    color: colors.text.primary,
    fontSize: typography.fontSize.base,
  },
});
```

## Security Considerations

### Content Security Policy

- Prevents XSS attacks
- Restricts script sources
- Uses nonces for inline scripts
- Critical for HTMX applications

### HTTPS Requirements

- Required for Google OAuth
- Required for production deployment
- Traefik handles SSL termination

## Deployment

### Docker Setup

- **Multi-stage build**: Go compilation in builder stage
- **Production image**: Alpine Linux base
- **Reverse Proxy**: Traefik with Let's Encrypt
- **Auto-updates**: Watchtower for rolling updates

### Hosting

- **Provider**: Hostinger VPS
- **OS**: Ubuntu
- **Container Registry**: GitHub Container Registry (ghcr.io)

## Current Limitations & Future Improvements

### Known Issues

1. Middleware chaining breaks CSP and TextHTML middleware
2. Single-user focus (designed for Mia)
3. Limited flight API calls (100/month)
4. Mobile Google OAuth requires proper client IDs configuration

### Recent Improvements

1. ✅ **Dual-platform architecture**: Web app + React Native mobile app
2. ✅ **JWT Authentication**: Secure mobile API with token refresh
3. ✅ **Redux State Management**: Centralized state for mobile app
4. ✅ **Theme Consistency**: Mobile app matches website's dark theme
5. ✅ **Separated API Handlers**: Clean separation between web and mobile endpoints

### Potential Enhancements

1. Multi-user support (requires authentication overhaul)
2. Enhanced statistics and analytics
3. Trip sharing capabilities
4. Integration with more flight APIs
5. Push notifications for mobile app
6. Offline data synchronization
7. Camera integration for trip photos

## File Importance Levels

### Critical Files (Don't Delete/Break)

**Backend:**
- `main.go` - Application entry point with dual web/mobile routing
- `internal/database/schema.sql` - Database structure
- `templates/layout.templ` - Base template layout
- `.air.toml` - Development hot reload config
- `internal/handlers/` - Web handlers (HTMX/session-based)
- `internal/handlers/mobile/` - Mobile API handlers (JWT-based)

**Mobile App:**
- `app/App.tsx` - Mobile app entry point
- `app/src/store/` - Redux store configuration
- `app/src/navigation/AppNavigator.tsx` - Navigation with auth flow
- `app/src/utils/theme.ts` - Theme configuration
- `app/app.json` - Expo configuration

### Generated Files (Don't Edit Manually)

**Backend:**
- `*_templ.go` - Generated from `.templ` files
- `static/css/output.css` - Generated from Tailwind
- `internal/database/database.db` - SQLite database file

**Mobile App:**
- `app/node_modules/` - Dependencies
- Expo build artifacts

### Configuration Files

**Backend:**
- `tailwind.config.js` - CSS framework config
- `docker-compose.yml` - Deployment setup
- `Dockerfile` - Container definition

**Mobile App:**
- `app/package.json` - Dependencies and scripts
- `app/tsconfig.json` - TypeScript configuration
