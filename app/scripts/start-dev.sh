#!/bin/bash
# Development startup script

echo "🚀 Starting Mia's Trips development environment..."
echo ""

# Check if we're in the right directory
if [ ! -f "../main.go" ]; then
    echo "❌ Error: This script must be run from the app/ directory"
    echo "Current directory: $(pwd)"
    exit 1
fi

# Function to check if a port is in use
check_port() {
    local port=$1
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null ; then
        echo "⚠️  Port $port is already in use"
        return 1
    fi
    return 0
}

# Check required ports
echo "🔍 Checking ports..."
if ! check_port 3000; then
    echo "Backend port 3000 is busy. Run 'npm run cleanup' first."
    exit 1
fi

if ! check_port 4040; then
    echo "ngrok API port 4040 is busy. Run 'npm run cleanup' first."
    exit 1
fi

# Start ngrok tunnel
echo "🌐 Starting ngrok tunnel..."
cd ..
ngrok http 3000 > /dev/null 2>&1 &
NGROK_PID=$!
cd app

# Wait for ngrok to start
echo "⏳ Waiting for ngrok to initialize..."
sleep 3

# Verify ngrok is running
if ! curl -s http://localhost:4040/api/tunnels > /dev/null; then
    echo "❌ Failed to start ngrok tunnel"
    exit 1
fi

# Start backend with air
echo "⚙️  Starting Go backend with air..."
cd ..
air > /dev/null 2>&1 &
BACKEND_PID=$!
cd app

# Wait for backend to start
echo "⏳ Waiting for backend to start..."
sleep 5

# Verify backend is running
if ! curl -s http://localhost:3000 > /dev/null; then
    echo "❌ Failed to start backend"
    exit 1
fi

# Get ngrok URL for display
NGROK_URL=$(curl -s http://localhost:4040/api/tunnels | jq -r '.tunnels[] | select(.config.addr == "http://localhost:3000" and .proto == "https") | .public_url')

echo ""
echo "✅ Development environment ready!"
echo "🌐 ngrok URL: $NGROK_URL"
echo "⚙️  Backend: http://localhost:3000"
echo ""
echo "🎯 Starting Expo..."

# Start Expo (this will run in foreground)
npx expo start --clear