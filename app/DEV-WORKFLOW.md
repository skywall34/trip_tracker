# Development Workflow Guide

## Setup (First Time Only)

1. **Environment Variables** (Optional but recommended):
```bash
cd app
cp .env.example .env
# Edit .env with your Google OAuth client IDs and other settings
```

## Quick Start

```bash
cd app
npm run dev
```

This single command will:
1. Start ngrok tunnel for public HTTPS access
2. Start Go backend with `air` (hot reload)
3. Start Expo development server
4. Automatically configure API URL via app.config.js

## Available Commands

### Main Commands
- `npm run dev` - Start complete development environment
- `npm run cleanup` - Stop all development processes

### Individual Components  
- `npm run backend` - Start only Go backend (air)
- `npm run ngrok` - Start only ngrok tunnel
- `npm run check-ngrok` - Display current ngrok URL

### Standard Expo Commands
- `npm start` - Start Expo with tunnel
- `npm run android` - Start for Android
- `npm run ios` - Start for iOS  
- `npm run web` - Start for web

## Development Process

### Starting Development
1. **Clean start**: `npm run cleanup && npm run dev`
2. **Quick restart**: `npm run dev` (if no processes running)

### Stopping Development
- `npm run cleanup` - Stops all processes cleanly
- `Ctrl+C` in terminal - Stops current process only

## How It Works

### Environment-Based Configuration
- **Environment Variables**: All sensitive values in `.env` file (git-ignored)
- **Google OAuth**: Separate client IDs for iOS/Android/Web platforms
- **API URL Priority**:
  1. `EXPO_PUBLIC_DEV_API_URL` (manual override)
  2. ngrok auto-detection (dynamic)
  3. `http://localhost:3000` (fallback)
  4. `EXPO_PUBLIC_PRODUCTION_API_URL` (production)

### Automatic ngrok Detection
- `app.config.js` calls ngrok API at `http://localhost:4040/api/tunnels`
- Finds HTTPS tunnel for `localhost:3000`
- Automatically configures `extra.apiUrl` for mobile app
- Falls back to `localhost:3000` if ngrok not running

### Process Management
- **ngrok**: Runs in background, provides public HTTPS URL
- **air**: Runs in background, auto-reloads Go backend on changes
- **expo**: Runs in foreground, shows QR code and logs

## Troubleshooting

### Port Already in Use
```bash
npm run cleanup
# Wait a few seconds, then:
npm run dev
```

### ngrok Not Working
```bash
# Check if ngrok is running
curl http://localhost:4040/api/tunnels

# Manual ngrok start
npm run ngrok
# In another terminal:
npm start
```

### Backend Not Starting
```bash
# Check from project root
cd ..
air
# Should see "Server running on :3000"
```

### App.config.js Issues
```bash
# Check if axios can reach ngrok API
node -e "const axios = require('axios'); axios.get('http://localhost:4040/api/tunnels').then(r => console.log(r.data)).catch(e => console.log('Error:', e.message))"
```

## File Structure
```
app/
├── scripts/
│   ├── cleanup.sh       # Cleanup script
│   └── start-dev.sh     # Development startup script
├── app.config.js        # Dynamic Expo config with ngrok detection
├── package.json         # Updated with new npm scripts
└── DEV-WORKFLOW.md      # This file
```