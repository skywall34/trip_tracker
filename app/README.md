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
npx expo build:ios
npx expo build:android
```

## 🔧 Configuration

### Environment Variables

The app uses environment variables configured in `app.json`:

```json
{
  "expo": {
    "extra": {
      "apiUrl": "http://localhost:3000"
    }
  }
}
```

**Development vs Production:**

- **Development**: `http://localhost:3000` or `http://YOUR_IP:3000`
- **Production**: `https://your-production-domain.com`

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
- ✅ `POST /api/v1/mobile/auth/refresh` - Token refresh

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

**Next Immediate Steps:**

0. **Fix Issues with network**

   - Emulator or phone is unable to communicate with backend server running at port 3000 on localhost
   - Google OAuth and Email logins fail

1. **Add Trip Screen**

   - Form validation for trip data
   - Airport selection with autocomplete
   - Date/time pickers for departure/arrival
   - Integration with backend POST endpoint

2. **Edit Trip Screen**

   - Pre-populate form with existing trip data
   - Partial update functionality
   - Optimistic UI updates

3. **Trip Detail Screen**
   - Full trip information display
   - Action buttons (edit, delete)
   - Flight status integration

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

**Testing Setup:**

- Physical device testing via Expo Go
- Network configuration for WSL2/Windows environments
- API endpoint testing with proper JWT tokens

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
