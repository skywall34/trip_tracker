# Google OAuth Setup for Mobile App - Step by Step Guide

This guide shows how to properly set up Google OAuth for the mobile app to work exactly like the website.

## Step 1: Google Cloud Console Setup

### 1.1 Create/Configure Google Cloud Project
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Select your existing project or create a new one
3. Enable the Google+ API and Google OAuth2 API

### 1.2 Configure OAuth Consent Screen
1. Go to "APIs & Services" → "OAuth consent screen"
2. Choose "External" user type
3. Fill in required information:
   - App name: "Mia's Trips"
   - User support email: your email
   - Developer contact information: your email
4. Add scopes:
   - `../auth/userinfo.email`
   - `../auth/userinfo.profile`
5. Add test users if needed

### 1.3 Create OAuth 2.0 Client IDs

#### For Development (Expo):
1. Go to "APIs & Services" → "Credentials"
2. Click "Create Credentials" → "OAuth 2.0 Client ID"
3. Choose "Web application"
4. Name: "Mia's Trips - Expo Dev"
5. Add authorized redirect URIs:
   ```
   https://auth.expo.io/@your-expo-username/mias-trips
   ```
6. Copy the Client ID - this goes in your mobile app

#### For Production (Standalone):
1. Create another OAuth client for each platform:

**iOS:**
- Application type: "iOS"
- Bundle ID: `com.miasname.trips` (from app.json)

**Android:**
- Application type: "Android"
- Package name: `com.miasname.trips` (from app.json)
- SHA-1 certificate fingerprint: (get from Expo build)

**Web (for testing):**
- Application type: "Web application"
- Authorized redirect URIs: `http://localhost:3000/auth/google/callback`

## Step 2: Environment Variables

Update your `.env` file with the Google OAuth credentials:

```bash
# Google OAuth (same as website)
GOOGLE_CLIENT_ID=your-web-client-id.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=your-client-secret
GOOGLE_CALLBACK_URL=http://localhost:3000/auth/google/callback

# JWT Secret for mobile tokens
JWT_SECRET=your-super-secure-jwt-secret-key
```

## Step 3: Mobile App Configuration

### 3.1 Update app.json
```json
{
  "expo": {
    "scheme": "mias-trips",
    "ios": {
      "bundleIdentifier": "com.miasname.trips",
      "googleServicesFile": "./GoogleService-Info.plist"
    },
    "android": {
      "package": "com.miasname.trips",
      "googleServicesFile": "./google-services.json"
    }
  }
}
```

### 3.2 Update LoginScreen.tsx with proper client IDs

Replace the placeholder client IDs in `LoginScreen.tsx`:

```typescript
const [request, response, promptAsync] = Google.useAuthRequest({
  clientId: Platform.select({
    ios: 'YOUR_IOS_CLIENT_ID.apps.googleusercontent.com',
    android: 'YOUR_ANDROID_CLIENT_ID.apps.googleusercontent.com', 
    default: 'YOUR_WEB_CLIENT_ID.apps.googleusercontent.com', // For Expo Go
  }),
  scopes: ['openid', 'profile', 'email'],
});
```

## Step 4: Backend Integration Points

The mobile authentication now properly mimics the website:

1. **Google OAuth Flow**: Uses same Google API endpoints as website
2. **User Creation**: Same logic as `google_callback.go`
3. **JWT Tokens**: Mobile gets JWT tokens instead of session cookies
4. **User Management**: Same user store methods as website

### API Endpoints:
- `POST /api/v1/mobile/auth/login` - Email/password login
- `POST /api/v1/mobile/auth/google` - Google OAuth login  
- `POST /api/v1/mobile/auth/refresh` - Token refresh

## Step 5: Development vs Production

### Development (Current Mock):
```typescript
// For immediate testing without OAuth setup
const mockGoogleToken = 'mock-google-token-development';
dispatch(loginWithGoogle(mockGoogleToken));
```

### Production (Real OAuth):
```typescript
// Real Google OAuth flow
const [request, response, promptAsync] = Google.useAuthRequest({
  clientId: 'your-real-client-id.apps.googleusercontent.com',
  scopes: ['openid', 'profile', 'email'],
});
```

## Step 6: Testing the Integration

### Test with Expo Go:
1. Use the web client ID in the default case
2. Test on physical device with Expo Go app
3. OAuth redirect will work through Expo's auth proxy

### Test Production Build:
1. Create development build: `eas build --profile development`
2. Use platform-specific client IDs
3. Test on device or simulator

## Step 7: Troubleshooting

### Common Issues:

1. **"Invalid client" error**:
   - Check client ID matches exactly
   - Verify bundle ID/package name matches Google Console

2. **"Redirect URI mismatch"**:
   - For Expo Go: Use `https://auth.expo.io/@username/project-slug`
   - For standalone: Use custom scheme from app.json

3. **Token verification fails**:
   - Check Google APIs are enabled
   - Verify scopes match between frontend and backend

### Debug Mode:
The backend includes mock authentication that bypasses Google OAuth for development:
- Use token: `"mock-google-token-development"`
- Creates user: `dev@example.com`
- Returns real JWT tokens for testing

## Step 8: Security Considerations

1. **Environment Variables**: Never commit real credentials to git
2. **JWT Secrets**: Use strong, unique secrets in production
3. **Token Expiry**: Access tokens expire in 15 minutes
4. **Refresh Flow**: Implement proper token refresh on mobile
5. **HTTPS**: Production requires HTTPS for OAuth callbacks

## Migration Path

1. **Phase 1**: Use mock authentication for development
2. **Phase 2**: Set up Google OAuth with test credentials  
3. **Phase 3**: Configure production OAuth clients
4. **Phase 4**: Deploy with real credentials and HTTPS

This setup ensures the mobile app authentication works exactly like the website, just with JWT tokens instead of session cookies!