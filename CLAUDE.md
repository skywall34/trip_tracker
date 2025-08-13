# Mia's Trips - Claude Code Context

This document provides comprehensive context for Claude Code to assist with development of the Mia's Trips travel tracking application.

## Project Overview

**Name**: Trip Tracker Project (Mia's Trips)  
**Go Version**: 1.23.5  
**Type**: Progressive Web Application (PWA) - unified website and mobile app  
**Primary User**: Single user (Mia) with potential for expansion

**Description**: A unified Progressive Web App for managing travel trips, built as a learning project to understand production-ready Go applications with HTMX and modern PWA features. The same application serves as both a responsive website and an installable mobile app, allowing users to create, edit, delete, and visualize trips on a world map across any device.

## Tech Stack

### Backend

- **Language**: Go 1.23.5
- **HTTP Server**: Built-in `net/http` package
- **Database**: SQLite
- **Templates**: Templ (Go template engine)
- **Real-time Updates**: HTMX

### Frontend

- **CSS Framework**: Tailwind CSS + mobile-optimized CSS
- **JavaScript**: HTMX + custom JS files + PWA features
- **Maps**: Leaflet.js for world map visualization
- **UI Pattern**: Server-side rendering with HTMX for dynamic updates
- **PWA**: Service worker, web manifest, offline functionality
- **Responsive Design**: Unified experience for website and mobile app

### External APIs

- **Flight Data**: AviationStack API (100 calls/month free tier)
- **Authentication**: Google OAuth

### Development Tools

- **Hot Reload**: Air for Go code auto-reload
- **CSS Build**: Tailwind CLI for CSS compilation
- **Template Generation**: Templ CLI

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
|   |   └── google_*.go              # Google Auth handlers
│   ├── database/                    # Database layer
│   │   ├── *.go                     # Database store files
│   │   ├── schema.sql               # Table definitions
|   |   └── database.db              # SQLite database file
│   ├── handlers/                    # HTTP handlers
│   │   ├── get*.go                  # GET request handlers
│   │   ├── post*.go                 # POST request handlers
│   │   ├── update*.go               # PUT/PATCH request handlers
│   │   ├── delete*.go               # DELETE request handlers
│   │   └── pwa.go                   # PWA-specific handlers (manifest, service worker, offline)
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
└── templates/                       # Templ template files
    ├── layout.templ                 # Base layout template (includes PWA meta tags, mobile nav)
    ├── trips.templ                  # Trip-related components (TripsPage removed)
    ├── *.templ                      # Other page templates
    └── *_templ.go                   # Generated Go files (don't edit, generate by templ generate)
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
- **Examples**:
  - `gethome.go` - GET request for home page
  - `posttrip.go` - POST request to create trip
  - `updatetrip.go` - PUT request to update trip
  - `deletetrip.go` - DELETE request to remove trip
  - `pwa.go` - PWA-specific handlers (manifest, service worker, offline page)

### PWA Implementation Notes

- **Unified Experience**: Same application serves both website and mobile app
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

- `PORT` (optional): Server port, defaults to 3000

### Development Setup

1. **Install Go 1.23+**
2. **Install SQLite3**
3. **Install Tailwind CLI binary**
4. **Install Templ CLI**
5. **Install Air for hot reload**

### Development Commands

During initial setup only, no need to run if database.db already created:

```bash
sqlite3 database.db < database/schema.sql
```

Run every time you want to start the server:

```bash
# Generate Tailwind CSS (includes mobile styles)
./tailwindcss -i ./static/css/input.css -o ./static/css/output.css

# Generate templ files (includes PWA templates)
templ generate

# Start development server with hot reload
air
```

### PWA Testing Commands

```bash
# Test as website (desktop)
open http://localhost:3000

# Test as mobile website (find your IP first)
ip addr show | grep "inet " | grep -v 127.0.0.1
# Then open http://YOUR_IP:3000 on mobile

# Test as full PWA (requires HTTPS)
air &
ngrok http 3000
# Use the https://xxx.ngrok.io URL for full PWA features
```

## Common Patterns

### Error Handling

```go
if err != nil {
    http.Error(w, "Error message", http.StatusInternalServerError)
    return
}
```

### Context Usage

```go
ctx := r.Context()
userID, ok := ctx.Value(m.UserKey).(int)
if !ok {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
}
```

### Template Rendering

```go
err := templates.ComponentName(data).Render(r.Context(), w)
if err != nil {
    http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
    return
}
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

### Potential Enhancements

1. Multi-user support
2. Mobile app development (PWA or native)
3. Enhanced statistics and analytics
4. Trip sharing capabilities
5. Integration with more flight APIs

## File Importance Levels

### Critical Files (Don't Delete/Break)

- `main.go` - Application entry point
- `internal/database/schema.sql` - Database structure
- `templates/layout.templ` - Base template layout
- `.air.toml` - Development hot reload config
- `internal/*` - main internal files such as handlers which handle the logic of the project

### Generated Files (Don't Edit Manually)

- `*_templ.go` - Generated from `.templ` files
- `static/css/output.css` - Generated from Tailwind
- `database.db` - SQLite database file

### Configuration Files

- `tailwind.config.js` - CSS framework config
- `docker-compose.yml` - Deployment setup
- `Dockerfile` - Container definition
