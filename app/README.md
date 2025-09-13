# Mia's Trips - Mobile App (React Native with Expo)

A React Native mobile application built with Expo for tracking travel trips. This app works alongside the existing Go backend web application, providing a native mobile experience for iOS and Android.

## 📱 Features

- **Cross-Platform**: Single codebase for iOS and Android
- **TypeScript**: Type-safe development with full TypeScript support
- **Redux State Management**: Centralized state management with Redux Toolkit
- **Secure Authentication**: JWT-based authentication with secure token storage
- **Offline Capability**: Works offline with data synchronization
- **Maps Integration**: Interactive maps to visualize trip locations
- **Native UI**: Platform-specific UI components and interactions

## 🏗️ Architecture

### Technology Stack

- **Framework**: React Native with Expo
- **Language**: TypeScript
- **State Management**: Redux Toolkit + React Redux
- **Navigation**: React Navigation v6
- **API Client**: Axios with interceptors
- **Storage**: Expo SecureStore for sensitive data
- **Maps**: React Native Maps
- **Icons**: Expo Vector Icons

### Project Structure

```
app/
├── App.tsx                          # Main app entry point
├── app.json                         # Expo configuration
├── package.json                     # Dependencies
├── tsconfig.json                    # TypeScript configuration
│
├── src/
│   ├── api/                         # API client and services
│   │   ├── client.ts                # Base API client with interceptors
│   │   ├── auth.ts                  # Authentication API
│   │   ├── trips.ts                 # Trips CRUD API
│   │   ├── airports.ts              # Airport data API
│   │   ├── types.ts                 # TypeScript interfaces
│   │   └── index.ts                 # API exports
│   │
│   ├── components/                  # Reusable components
│   │   ├── common/                  # Shared UI components
│   │   ├── trips/                   # Trip-specific components
│   │   └── map/                     # Map components
│   │
│   ├── screens/                     # Screen components
│   │   ├── auth/                    # Authentication screens
│   │   │   ├── LoginScreen.tsx
│   │   │   └── SplashScreen.tsx
│   │   ├── trips/                   # Trip management screens
│   │   │   ├── TripsListScreen.tsx
│   │   │   ├── TripDetailScreen.tsx
│   │   │   ├── EditTripScreen.tsx
│   │   │   └── AddTripScreen.tsx
│   │   ├── map/
│   │   │   └── MapScreen.tsx
│   │   └── profile/                 # Profile and settings
│   │       ├── ProfileScreen.tsx
│   │       ├── SettingsScreen.tsx
│   │       └── StatisticsScreen.tsx
│   │
│   ├── navigation/                  # Navigation configuration
│   │   ├── AppNavigator.tsx         # Root navigator
│   │   ├── AuthNavigator.tsx        # Authentication flow
│   │   ├── TabNavigator.tsx         # Bottom tab navigation
│   │   ├── TripsNavigator.tsx       # Trips stack navigation
│   │   ├── ProfileNavigator.tsx     # Profile stack navigation
│   │   ├── types.ts                 # Navigation type definitions
│   │   └── index.ts                 # Navigation exports
│   │
│   ├── store/                       # Redux store configuration
│   │   ├── index.ts                 # Store setup
│   │   ├── hooks.ts                 # Typed Redux hooks
│   │   └── slices/                  # Redux slices
│   │       ├── authSlice.ts         # Authentication state
│   │       ├── tripsSlice.ts        # Trips state
│   │       └── uiSlice.ts           # UI state
│   │
│   ├── services/                    # Native services (future)
│   ├── utils/                       # Utility functions (future)
│   └── assets/                      # Static assets (future)
│
├── assets/                          # Expo assets
│   ├── icon.png                     # App icon
│   ├── splash-icon.png              # Splash screen icon
│   ├── adaptive-icon.png            # Android adaptive icon
│   └── favicon.png                  # Web favicon
│
└── node_modules/                    # Dependencies
```

## 🚀 Getting Started

### Prerequisites

- Node.js 18 or higher
- npm or yarn
- Expo CLI: `npm install -g @expo/cli`
- For iOS development: Xcode (macOS only)
- For Android development: Android Studio

### Installation

1. **Navigate to the app directory:**

   ```bash
   cd app
   ```

2. **Install dependencies:**

   ```bash
   npm install
   ```

3. **Start the development server:**
   ```bash
   npx expo start
   ```

## 📱 Running the App

### Development with Expo Go

1. **Install Expo Go** on your mobile device:

   - [iOS App Store](https://apps.apple.com/app/expo-go/id982107779)
   - [Google Play Store](https://play.google.com/store/apps/details?id=host.exp.exponent)

2. **Start the development server:**

   ```bash
   npx expo start
   ```

3. **Scan the QR code** with Expo Go (Android) or Camera app (iOS)

### Using Emulators

#### iOS Simulator (macOS only)

1. **Install Xcode** from the Mac App Store
2. **Open iOS Simulator:**
   ```bash
   npx expo start --ios
   ```

#### Android Emulator

1. **Install Android Studio** and set up an Android Virtual Device (AVD)
2. **Start the Android emulator:**
   ```bash
   npx expo start --android
   ```

### Development Scripts

**New Automated Development Workflow:**
```bash
# One-command development startup (recommended)
npm run dev
# Automatically starts: ngrok + backend + Expo

# Clean shutdown of all processes
npm run cleanup

# Individual components
npm run backend      # Start Go backend with air
npm run ngrok       # Start ngrok tunnel only
npm run check-ngrok # Display current ngrok URL
```

**Standard Expo Commands:**
```bash
# Start development server
npm start
# or
npx expo start

# Start with specific platform
npx expo start --ios
npx expo start --android
npx expo start --web

# Clear cache and restart
npx expo start --clear

# Build for production
eas build --profile production --platform ios
eas build --profile production --platform android
```

## 🔧 Configuration

### Environment Variables

The app uses environment variables for secure configuration management:

**Setup (Required):**
```bash
# Copy environment template
cp .env.example .env

# Edit .env with your actual values
```

**Available Environment Variables:**
```bash
# Google OAuth Client IDs (Get from Google Cloud Console)
EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_IOS=your-ios-client-id.apps.googleusercontent.com
EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_ANDROID=your-android-client-id.apps.googleusercontent.com  
EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_WEB=your-web-client-id.apps.googleusercontent.com

# Production API URL
EXPO_PUBLIC_PRODUCTION_API_URL=https://yourdomain.com

# Optional: Development API URL override (bypasses ngrok auto-detection)
EXPO_PUBLIC_DEV_API_URL=https://your-custom-dev-url.com

# Google Maps API Key (for React Native Maps)
EXPO_PUBLIC_GOOGLE_MAPS_API_KEY=your-google-maps-api-key
```

**Dynamic Configuration:**
- **Development**: Automatic ngrok URL detection via `app.config.js`
- **Production**: Uses `EXPO_PUBLIC_PRODUCTION_API_URL`
- **Override**: Can force specific URL via `EXPO_PUBLIC_DEV_API_URL`

**Security:**
- All sensitive values stored in `.env` (git-ignored)
- No hardcoded credentials in source code
- Environment-based OAuth client IDs

### Backend Integration

**✅ Complete Backend Integration:**

The mobile app connects to the existing Go backend server with dedicated mobile handlers.

**Backend Requirements:**

1. ✅ **Backend server running** on the specified URL
2. ✅ **JWT endpoints implemented** for mobile authentication
3. ✅ **Separated mobile handlers** following code conventions
4. ✅ **JSON API responses** (no HTMX headers)

**Available API Endpoints:**

**Authentication:**

- ✅ `POST /api/v1/mobile/auth/google` - Google OAuth login
- ⚠️ `POST /api/v1/mobile/auth/login` - Email/password login (needs backend implementation)
- ✅ `POST /api/v1/mobile/auth/refresh` - Token refresh
- ✅ `POST /api/v1/mobile/auth/logout` - Logout

**Trip Management:**

- ✅ `GET /api/v1/trips` - Get trips list (JSON response)
- ✅ `POST /api/v1/trips` - Create trip (JSON request/response)
- ✅ `PUT /api/v1/trips/{id}` - Update trip (JSON request/response)
- ✅ `DELETE /api/v1/trips/{id}` - Delete trip (JSON response)

**Backend Architecture:**

- **Web Handlers**: HTMX-based with session authentication
- **Mobile Handlers**: RESTful JSON API with JWT authentication
- **Clean Separation**: No conflicts between web and mobile functionality

## 🧪 Testing

### Running Tests

```bash
# Run all tests
npm test

# Run tests in watch mode
npm run test:watch

# Run tests with coverage
npm run test:coverage
```

### Testing on Devices

1. **Physical Device Testing:**

   - Use Expo Go for development builds
   - Create development builds for testing native features

2. **Emulator Testing:**
   - iOS Simulator for iOS testing
   - Android Emulator for Android testing

## 📦 Building for Production

### Development Builds

For testing native features:

```bash
# Create development build
npx expo install expo-dev-client
eas build --profile development --platform ios
eas build --profile development --platform android
```

### Production Builds

```bash
# iOS App Store build
eas build --profile production --platform ios

# Android Play Store build
eas build --profile production --platform android
```

## 🚀 Deployment

### App Store Deployment

1. **iOS App Store:**

   ```bash
   eas submit --platform ios
   ```

2. **Google Play Store:**
   ```bash
   eas submit --platform android
   ```

### Over-the-Air Updates

Use Expo Updates for seamless app updates:

```bash
# Publish update
eas update --branch main --message "Bug fixes and improvements"
```

## 🔐 Security

### Token Storage

- **Secure Storage**: Uses Expo SecureStore for JWT tokens
- **Biometric Authentication**: Can be extended with biometric authentication
- **Certificate Pinning**: Can be implemented for API calls

### Best Practices

- All sensitive data stored in SecureStore
- API calls use HTTPS only
- JWT tokens have expiration and refresh mechanism
- Input validation on all user inputs

## 🛠️ Development

### Code Style

- **ESLint**: Code linting with React Native rules
- **Prettier**: Code formatting
- **TypeScript**: Strict type checking

### Debugging

1. **Expo DevTools**: Built-in debugging tools
2. **React Native Debugger**: Standalone debugging app
3. **Flipper**: Facebook's debugging platform (for development builds)

### Hot Reloading

Expo provides fast refresh for instant code updates during development.

## 📱 Features Implementation Status

### ✅ Completed Features

**Foundation & Setup:**

- [x] Project setup with Expo and TypeScript
- [x] Core dependencies and navigation libraries installed
- [x] Project structure following React Native best practices

**Authentication System:**

- [x] Google OAuth integration with expo-auth-session
- [x] JWT token management with automatic refresh
- [x] Secure token storage using Expo SecureStore
- [x] Authentication screens (Login, Splash)
- [x] Authentication flow with navigation guards

**State Management:**

- [x] Redux Toolkit store configuration
- [x] Authentication slice with async thunks
- [x] Trips slice with CRUD operations
- [x] UI slice for app state management
- [x] Typed Redux hooks

**API Integration:**

- [x] Axios API client with JWT interceptors
- [x] Automatic token refresh mechanism
- [x] Error handling and response processing
- [x] TypeScript interfaces for API responses

**UI System:**

- [x] Dark theme system matching website aesthetic
- [x] Common UI components (Button, Card, Input, Loading)
- [x] Trip-specific components (TripCard)
- [x] Consistent styling with website colors
- [x] Theme utilities and constants

**Navigation:**

- [x] Navigation structure (Auth, Main, Tabs)
- [x] Bottom tab navigation with proper icons
- [x] Stack navigation for modal screens
- [x] Authentication flow routing

**Backend Integration:**

- [x] Mobile-specific JWT API endpoints
- [x] Separated mobile handlers in Go backend
- [x] RESTful API design with JSON responses
- [x] Proper authentication middleware for mobile

### 🚧 In Progress

**Trip Management:**

- [x] Trip list screen with pull-to-refresh
- [ ] Add trip screen with form validation
- [ ] Edit trip screen with pre-filled data
- [ ] Trip detail screen with full information
- [ ] Delete confirmation dialogs

### 📋 Planned Features

**Core Functionality:**

- [ ] Maps integration with React Native Maps
- [ ] Trip visualization on world map
- [ ] Profile and settings screens
- [ ] Statistics and analytics

**Advanced Features:**

- [ ] Offline data synchronization
- [ ] Push notifications
- [ ] Camera integration for trip photos
- [ ] Biometric authentication
- [ ] Advanced trip statistics
- [ ] Trip sharing capabilities

## 🤝 Contributing

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Commit changes**: `git commit -m 'Add amazing feature'`
4. **Push to branch**: `git push origin feature/amazing-feature`
5. **Open a Pull Request**

## 🆘 Support

### Common Issues

1. **Metro bundler cache issues:**

   ```bash
   npx expo start --clear --tunnel
   ```

2. **iOS Simulator not opening:**

   ```bash
   sudo xcode-select --switch /Applications/Xcode.app
   ```

3. **Android emulator connection issues:**
   ```bash
   adb reverse tcp:3000 tcp:3000
   ```

### Getting Help

- **Expo Documentation**: [docs.expo.dev](https://docs.expo.dev/)
- **React Native Documentation**: [reactnative.dev](https://reactnative.dev/)
- **GitHub Issues**: Report bugs and request features

## 🔄 Integration with Web App

This mobile app is designed to work alongside the existing Go web application:

- **Shared Backend**: Uses the same Go server and database
- **Separate Authentication**: JWT tokens for mobile, sessions for web
- **API Versioning**: Mobile uses `/api/v1/` endpoints
- **Data Synchronization**: Real-time sync between web and mobile

The web PWA features will be removed as outlined in the project plan, with the mobile app becoming the primary mobile experience.

## 🚀 Development Progress & Next Steps

### ✅ Completed Development Milestones

**Phase 1: Foundation (✅ Complete)**

- Expo project setup with TypeScript
- Core dependencies installation
- Project structure organization
- Redux store configuration
- API client implementation

**Phase 2: Authentication System (✅ Complete)**

- Google OAuth integration with expo-auth-session
- JWT token management and secure storage
- Backend mobile API endpoints (`/api/v1/mobile/auth/`)
- Authentication screens and navigation flow
- Token refresh mechanism

**Phase 3: UI System (✅ Complete)**

- Dark theme matching website aesthetic
- Reusable UI components (Button, Card, Input, Loading)
- Trip-specific components (TripCard)
- Consistent styling system

**Phase 4: Backend Integration (✅ Complete)**

- Separated mobile handlers following code conventions:
  - `internal/handlers/mobile/gettrips.go`
  - `internal/handlers/mobile/posttrips.go`
  - `internal/handlers/mobile/puttrips.go`
  - `internal/handlers/mobile/deletetrips.go`
- JWT authentication middleware
- RESTful API design with JSON responses

### 🎯 Current Development Phase

**Phase 5: Trip Management Screens (🚧 In Progress)**

**Completed:**

- ✅ Trip list screen with pull-to-refresh functionality
- ✅ Redux integration for trip state management
- ✅ Error handling and loading states

## 📋 Comprehensive Development Todo List

Based on the comprehensive codebase audit, here's the complete roadmap for bringing the mobile app to feature parity and production readiness:

### 🚨 Critical Issues (Fix First)

#### Backend API Integration
- [ ] **Add Email Login Backend Endpoint** - `POST /api/v1/mobile/auth/login` missing
  - Location: Create `internal/handlers/mobile/login.go` (exists but route not registered)
  - Frontend calls this in `authSlice.ts:51-54`
  - Backend only has Google OAuth route

#### API Response Structure Mismatch
- [ ] **Fix API Response Handling in Trips Slice**
  - Location: `/src/store/slices/tripsSlice.ts:36`
  - Issue: `apiClient.get()` expects different structure than backend provides
  - Backend returns: `{ success: boolean, data: Trip[], error?: ErrorResponse }`
  - Frontend expects: Direct array or `response.data as Trip[]`

### 🎯 High Priority Features (Phase 1)

#### Core Trip Management Screens
- [ ] **Add Trip Screen Implementation**
  - Location: Replace placeholder in `AppNavigator.tsx:133-139`
  - Features needed:
    - Form validation with React Hook Form
    - Airport autocomplete/selection
    - Date/time pickers for departure/arrival
    - Airline and flight number inputs
    - Integration with `POST /api/v1/trips`
    - Error handling and loading states

- [ ] **Edit Trip Screen Implementation**
  - Location: Replace placeholder in `AppNavigator.tsx:141-147`
  - Features needed:
    - Pre-filled form with existing trip data
    - Same form components as Add Trip
    - Integration with `PUT /api/v1/trips/{id}`
    - Optimistic UI updates
    - Partial update support

- [ ] **Trip Detail Screen Implementation**
  - Location: Replace placeholder in `AppNavigator.tsx:128-131`
  - Features needed:
    - Full trip information display
    - Action buttons (edit, delete)
    - Delete confirmation dialog
    - Flight status integration (if available)
    - Share trip functionality

#### Maps Integration
- [ ] **World Map Screen Implementation**
  - Location: Replace placeholder in `AppNavigator.tsx:23-27`
  - Dependencies: Install `react-native-maps`
  - Features needed:
    - Interactive world map with trip markers
    - Flight route visualization
    - Marker clustering for dense areas
    - Trip filtering and legend
    - Integration with Google Maps API key from environment

### 🌟 Medium Priority Features (Phase 2)

#### User Profile & Settings
- [ ] **Profile Screen Implementation**
  - Location: Replace placeholder in `AppNavigator.tsx:29-33`
  - Features needed:
    - User information display
    - Profile picture upload
    - Account settings
    - Logout functionality
    - Delete account option

- [ ] **Settings Screen Implementation**
  - Features needed:
    - App preferences (theme, notifications)
    - Privacy settings
    - Data management (export, delete)
    - About section with app version

#### Enhanced Functionality
- [ ] **Flight Search Implementation**
  - Location: Complete TODO in `HomeScreen.tsx:24-27`
  - Integration with AviationStack API
  - Flight lookup and auto-population
  - Real-time flight status

- [ ] **Statistics Dashboard**
  - Features needed:
    - Total trips, miles flown, countries visited
    - Travel patterns and insights
    - Data visualizations (charts/graphs)
    - Year-over-year comparisons

- [ ] **Offline Support Implementation**
  - Redux Persist configuration
  - Offline queue for API calls
  - Sync mechanism when online
  - Offline indicators and messaging

### 🔧 Technical Improvements

#### Code Quality & Architecture
- [ ] **TypeScript Coverage Improvement**
  - Replace `any` types with proper interfaces
  - Add missing prop interface definitions
  - Strict type checking compliance

- [ ] **Error Boundaries Implementation**
  - React error boundaries for crash prevention
  - Error logging and reporting
  - Graceful fallback UI components

- [ ] **Performance Optimization**
  - FlatList optimization for large trip lists
  - React.memo and useCallback optimization
  - Image loading optimization
  - Bundle size analysis and reduction

#### Testing Framework
- [ ] **Unit Testing Setup**
  - Jest configuration for React Native
  - Component testing with React Native Testing Library
  - Redux store testing
  - API client testing with mocks

- [ ] **Integration Testing**
  - E2E testing with Detox
  - Navigation flow testing
  - Authentication flow testing
  - API integration testing

#### Enhanced UI/UX
- [ ] **Loading States & Skeletons**
  - Skeleton components for all screens
  - Better loading indicators
  - Progressive loading for images

- [ ] **Toast Notification System**
  - Success/error message toasts
  - Action confirmation toasts
  - Offline status notifications

- [ ] **Accessibility Improvements**
  - Screen reader support
  - Accessibility labels and hints
  - Focus management
  - High contrast support

### 🚀 Production Readiness

#### App Store Preparation
- [ ] **iOS App Store Configuration**
  - App icons (all required sizes)
  - Splash screens and launch images
  - App Store metadata and screenshots
  - TestFlight beta testing setup

- [ ] **Android Play Store Configuration**
  - Adaptive icons and splash screens
  - Play Store metadata and screenshots
  - Internal testing track setup
  - Signed APK/AAB generation

#### Security & Monitoring
- [ ] **Enhanced Security Implementation**
  - Input validation on all forms
  - XSS prevention measures
  - Certificate pinning for API calls
  - Biometric authentication option

- [ ] **Analytics & Monitoring Setup**
  - Crash reporting (Sentry integration)
  - Usage analytics (Firebase/Amplitude)
  - Performance monitoring
  - Error tracking and alerting

#### Advanced Features
- [ ] **Push Notifications**
  - Expo Notifications setup
  - Trip reminder notifications
  - Flight status update notifications
  - Permission handling

- [ ] **Camera Integration**
  - Trip photo capture
  - Photo gallery integration
  - Image upload to backend
  - Photo management

### 📱 Nice-to-Have Features (Phase 3)

- [ ] **Advanced Search & Filtering**
  - Trip search functionality
  - Filter by airline, date range, destinations
  - Sorting options

- [ ] **Data Export & Backup**
  - Export trips to CSV/PDF
  - Backup/restore functionality
  - Data synchronization across devices

- [ ] **Social Features**
  - Trip sharing with friends
  - Travel recommendations
  - Social login options

- [ ] **Travel Planning Tools**
  - Trip planning assistance
  - Weather integration
  - Currency converter
  - Time zone calculator

### 🏁 Current Status Summary
- **Foundation**: ✅ Complete (100%)
- **Authentication**: ✅ Complete (95% - missing email login backend)
- **Basic UI**: ✅ Complete (90%)
- **Trip Listing**: ✅ Complete (100%)
- **Trip Management**: 🚧 In Progress (20% - only listing implemented)
- **Maps**: ❌ Not Started (0%)
- **Profile/Settings**: ❌ Not Started (0%)
- **Production Ready**: ❌ Not Started (10%)

**Overall Completion**: ~40% of planned features
**Estimated Time to MVP**: 3-4 weeks for Phase 1
**Estimated Time to Production**: 6-8 weeks total

### 🗺️ Upcoming Major Phases

**Phase 6: Maps Integration (📋 Planned)**

- React Native Maps implementation
- Trip visualization on world map
- Flight route display
- Location markers with clustering

**Phase 7: Advanced Features (📋 Planned)**

- Profile and settings screens
- Statistics and analytics
- Offline data synchronization
- Push notifications

### 🛠️ Current Development Workflow

**Backend Development:**

```bash
# Terminal 1: Backend server
cd /path/to/trip-tracker
air  # Hot reload enabled
# Server runs on http://localhost:3000
```

**Mobile Development:**

```bash
# Terminal 2: Mobile app
cd app
npx expo start --tunnel
# Choose platform: Android (a), iOS (i), or Web (w)
```

**Automated Development Setup:**

```bash
# Quick start (recommended)
npm run dev
# Automatically handles: ngrok tunneling, backend startup, Expo server

# Manual setup (if needed)
# Terminal 1: Backend server
cd /home/mshin/trip-tracker
air  # Hot reload enabled

# Terminal 2: Mobile app  
cd app
npx expo start --clear
```

**Testing Setup:**

- Physical device testing via Expo Go (for most features)
- Development builds required for OAuth testing (`npx expo run:android`)
- Network configuration automated via ngrok tunneling
- Environment variable management via `.env` files

### 📋 Development Checklist

**Completed:**

- ✅ Redux store implementation
- ✅ Authentication flow setup
- ✅ Backend JWT endpoints
- ✅ UI component system
- ✅ Trip listing functionality

**In Progress:**

- 🚧 Trip management screens (CRUD operations)

**Upcoming:**

- ⏳ Maps functionality integration
- ⏳ Testing framework setup
- ⏳ Production build configuration

---

**Note**: This mobile app is part of the Mia's Trips project migration from PWA to native mobile application. It provides enhanced native capabilities while maintaining compatibility with the existing Go backend.

### Resources

https://medium.com/@akbarimo/developing-react-native-with-expo-android-emulators-on-wsl2-linux-subsystem-ad5a8b0fa23c
