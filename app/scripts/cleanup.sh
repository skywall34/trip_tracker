#!/bin/bash
# Cleanup script to terminate all development processes

echo "🧹 Cleaning up development processes..."

# Kill ngrok processes
echo "Stopping ngrok..."
pkill -f ngrok
sleep 1

# Kill air processes (Go backend)
echo "Stopping air (Go backend)..."
pkill -f air
sleep 1

# Kill Go processes
echo "Stopping Go processes..."
pkill -f "go run main.go"
pkill -f main.go
sleep 1

# Kill Expo processes
echo "Stopping Expo..."
pkill -f "expo start"
pkill -f "@expo/cli"
sleep 1

# Kill any remaining Node processes related to Expo
echo "Stopping remaining Node/Expo processes..."
pkill -f "expo-cli"
pkill -f "react-native"
sleep 1

# Clean up any remaining processes on common ports
echo "Checking for processes on common ports..."
lsof -ti:3000 | xargs kill -9 2>/dev/null || true
lsof -ti:4040 | xargs kill -9 2>/dev/null || true
lsof -ti:8081 | xargs kill -9 2>/dev/null || true

echo "✅ Cleanup complete!"
echo ""
echo "All development processes have been terminated."
echo "You can now run 'npm run dev' to start fresh."