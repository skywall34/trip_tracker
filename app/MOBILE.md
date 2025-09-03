# Mia's Trips - Expo Mobile App Integration Plan

## Overview

This document outlines the complete plan for building the Mia's Trips mobile app using Expo with TypeScript, integrated with your existing Go backend. The mobile app will maintain the same dark, modern aesthetic as the website while leveraging Expo's capabilities for cross-platform development.

## UI/UX Design Consistency

### Website Theme to Maintain

Based on your existing website's dark, modern theme:

**Color Palette:**

```typescript
// src/utils/theme.ts
export const colors = {
  // Dark backgrounds (matching website)
  background: "#0f0f23", // Main dark background
  surface: "#1a1a2e", // Card/surface background
  elevated: "#16213e", // Elevated surfaces

  // Accent colors
  primary: "#0ea5e9", // Sky blue (matching website buttons)
  secondary: "#22d3ee", // Cyan for hover states
  accent: "#10b981", // Green for success states

  // Text colors
  text: {
    primary: "#f8fafc", // White text
    secondary: "#94a3b8", // Gray text
    muted: "#64748b", // Muted text
  },

  // Semantic colors
  error: "#ef4444",
  warning: "#f59e0b",
  success: "#10b981",

  // Borders
  border: "#334155",
  borderLight: "#475569",
};

// Typography to match website
export const typography = {
  fontFamily: {
    regular: "Inter-Regular",
    medium: "Inter-Medium",
    semiBold: "Inter-SemiBold",
    bold: "Inter-Bold",
  },
  fontSize: {
    xs: 12,
    sm: 14,
    base: 16,
    lg: 18,
    xl: 20,
    "2xl": 24,
    "3xl": 30,
    "4xl": 36,
  },
};
```

### Component Styling Examples

```typescript
// src/components/common/Card.tsx
import React from "react";
import { View, StyleSheet } from "react-native";
import { colors } from "../../utils/theme";

export const Card: React.FC<{ children: React.ReactNode }> = ({ children }) => (
  <View style={styles.card}>{children}</View>
);

const styles = StyleSheet.create({
  card: {
    backgroundColor: colors.surface,
    borderRadius: 12,
    padding: 16,
    marginVertical: 8,
    marginHorizontal: 16,
    borderWidth: 1,
    borderColor: colors.border,
    // Shadow for iOS
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.3,
    shadowRadius: 4,
    // Elevation for Android
    elevation: 4,
  },
});

// src/components/common/Button.tsx
import React from "react";
import { TouchableOpacity, Text, StyleSheet } from "react-native";
import { colors, typography } from "../../utils/theme";

interface ButtonProps {
  title: string;
  onPress: () => void;
  variant?: "primary" | "secondary" | "outline";
}

export const Button: React.FC<ButtonProps> = ({
  title,
  onPress,
  variant = "primary",
}) => (
  <TouchableOpacity
    style={[styles.button, styles[variant]]}
    onPress={onPress}
    activeOpacity={0.8}
  >
    <Text style={[styles.text, styles[`${variant}Text`]]}>{title}</Text>
  </TouchableOpacity>
);

const styles = StyleSheet.create({
  button: {
    paddingVertical: 12,
    paddingHorizontal: 24,
    borderRadius: 8,
    alignItems: "center",
    justifyContent: "center",
  },
  primary: {
    backgroundColor: colors.primary,
  },
  secondary: {
    backgroundColor: colors.surface,
    borderWidth: 1,
    borderColor: colors.primary,
  },
  outline: {
    backgroundColor: "transparent",
    borderWidth: 1,
    borderColor: colors.border,
  },
  text: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.semiBold,
  },
  primaryText: {
    color: colors.text.primary,
  },
  secondaryText: {
    color: colors.primary,
  },
  outlineText: {
    color: colors.text.secondary,
  },
});
```

## Project Structure with Expo

```
trip-tracker/
├── [existing Go backend structure]
└── app/                              # Expo app directory
    ├── package.json                  # Dependencies and scripts
    ├── app.json                      # Expo configuration
    ├── tsconfig.json                 # TypeScript configuration
    ├── babel.config.js               # Babel configuration
    ├── .env                          # Environment variables
    ├── eas.json                      # EAS Build configuration
    │
    ├── App.tsx                       # Main entry point
    ├── app.config.ts                 # Dynamic Expo config
    │
    ├── src/
    │   ├── api/                      # API client layer
    │   │   ├── client.ts             # Axios configuration
    │   │   ├── auth.ts               # Authentication
    │   │   ├── trips.ts              # Trip operations
    │   │   └── types.ts              # TypeScript interfaces
    │   │
    │   ├── components/               # Reusable components
    │   │   ├── common/
    │   │   ├── trips/
    │   │   └── map/
    │   │
    │   ├── screens/                  # Screen components
    │   │   ├── auth/
    │   │   ├── trips/
    │   │   ├── map/
    │   │   └── profile/
    │   │
    │   ├── navigation/               # Navigation setup
    │   │   └── AppNavigator.tsx
    │   │
    │   ├── store/                    # Redux store
    │   │   ├── index.ts
    │   │   └── slices/
    │   │
    │   ├── hooks/                    # Custom hooks
    │   │   ├── useAuth.ts
    │   │   └── useTrips.ts
    │   │
    │   ├── utils/                    # Utilities
    │   │   └── constants.ts
    │   │
    │   └── assets/                   # Images, fonts
    │       ├── images/
    │       ├── fonts/
    │       └── icons/
    │
    └── assets/                       # Expo assets (splash, icon)
        ├── icon.png                  # App icon (1024x1024)
        ├── splash.png                # Splash screen (1284x2778)
        └── adaptive-icon.png         # Android adaptive icon
```

## Phase 1: Expo Setup (Week 1)

### Initial Setup (WSL2/Windows)

```bash
# In Windows PowerShell or WSL2
cd trip-tracker

# Install Expo CLI globally
npm install -g expo-cli eas-cli

# Create new Expo app with TypeScript
npx create-expo-app app --template expo-template-blank-typescript

# Navigate to app
cd app

# Install core dependencies
npx expo install expo-status-bar expo-constants expo-device
npx expo install @react-navigation/native @react-navigation/bottom-tabs @react-navigation/stack
npx expo install react-native-screens react-native-safe-area-context
npx expo install react-native-gesture-handler react-native-reanimated
```

### Configure Expo for Your Project

```typescript
// app.config.ts
import { ExpoConfig, ConfigContext } from "expo/config";

export default ({ config }: ConfigContext): ExpoConfig => ({
  ...config,
  name: "Mia's Trips",
  slug: "mias-trips",
  version: "1.0.0",
  orientation: "portrait",
  icon: "./assets/icon.png",
  userInterfaceStyle: "automatic",
  splash: {
    image: "./assets/splash.png",
    resizeMode: "contain",
    backgroundColor: "#1a1a2e",
  },
  assetBundlePatterns: ["**/*"],
  ios: {
    supportsTablet: true,
    bundleIdentifier: "com.miastrips.app",
    config: {
      googleMapsApiKey: process.env.GOOGLE_MAPS_API_KEY,
    },
  },
  android: {
    adaptiveIcon: {
      foregroundImage: "./assets/adaptive-icon.png",
      backgroundColor: "#1a1a2e",
    },
    package: "com.miastrips.app",
    config: {
      googleMaps: {
        apiKey: process.env.GOOGLE_MAPS_API_KEY,
      },
    },
  },
  web: {
    favicon: "./assets/favicon.png",
  },
  extra: {
    apiUrl: process.env.API_URL || "http://localhost:3000",
    eas: {
      projectId: "your-project-id", // Get from Expo dashboard
    },
  },
  plugins: [
    "expo-secure-store",
    "expo-location",
    "expo-camera",
    [
      "expo-image-picker",
      {
        photosPermission:
          "The app needs access to your photos to add trip memories.",
        cameraPermission:
          "The app needs access to your camera to capture trip moments.",
      },
    ],
  ],
});
```

### Environment Configuration

```bash
# .env.development
API_URL=http://192.168.1.100:3000  # Your machine's IP for physical devices
GOOGLE_MAPS_API_KEY=your-dev-key

# .env.production
API_URL=https://api.miastrips.com
GOOGLE_MAPS_API_KEY=your-prod-key
```

## Phase 2: Core Dependencies & Navigation (Week 2)

### Install Essential Packages

```bash
# State Management & API
npm install @reduxjs/toolkit react-redux redux-persist
npm install axios react-query @tanstack/react-query

# UI Components
npx expo install expo-font @expo-google-fonts/inter
npx expo install react-native-paper react-native-vector-icons
npx expo install react-native-maps

# Utilities
npx expo install expo-secure-store expo-location expo-camera
npx expo install expo-image-picker expo-notifications
npx expo install expo-auth-session expo-crypto expo-web-browser
npx expo install @react-native-async-storage/async-storage

# Development
npm install --save-dev @types/react @types/react-native
npm install --save-dev eslint prettier @typescript-eslint/parser
```

### Navigation Setup

```typescript
// src/navigation/AppNavigator.tsx
import React from "react";
import { NavigationContainer } from "@react-navigation/native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createStackNavigator } from "@react-navigation/stack";
import { Ionicons } from "@expo/vector-icons";

// Import screens
import HomeScreen from "../screens/HomeScreen";
import TripsScreen from "../screens/trips/TripsScreen";
import MapScreen from "../screens/map/MapScreen";
import ProfileScreen from "../screens/profile/ProfileScreen";
import LoginScreen from "../screens/auth/LoginScreen";

const Tab = createBottomTabNavigator();
const Stack = createStackNavigator();

const TabNavigator = () => (
  <Tab.Navigator
    screenOptions={({ route }) => ({
      tabBarIcon: ({ focused, color, size }) => {
        let iconName: keyof typeof Ionicons.glyphMap = "home";

        if (route.name === "Home") iconName = focused ? "home" : "home-outline";
        else if (route.name === "Trips")
          iconName = focused ? "airplane" : "airplane-outline";
        else if (route.name === "Map")
          iconName = focused ? "map" : "map-outline";
        else if (route.name === "Profile")
          iconName = focused ? "person" : "person-outline";

        return <Ionicons name={iconName} size={size} color={color} />;
      },
      tabBarActiveTintColor: "#4A90E2",
      tabBarInactiveTintColor: "gray",
    })}
  >
    <Tab.Screen name="Home" component={HomeScreen} />
    <Tab.Screen name="Trips" component={TripsScreen} />
    <Tab.Screen name="Map" component={MapScreen} />
    <Tab.Screen name="Profile" component={ProfileScreen} />
  </Tab.Navigator>
);

export default function AppNavigator() {
  const isAuthenticated = false; // Replace with auth check

  return (
    <NavigationContainer>
      <Stack.Navigator screenOptions={{ headerShown: false }}>
        {isAuthenticated ? (
          <Stack.Screen name="Main" component={TabNavigator} />
        ) : (
          <Stack.Screen name="Login" component={LoginScreen} />
        )}
      </Stack.Navigator>
    </NavigationContainer>
  );
}
```

## Phase 3: API Integration (Week 3)

### API Client with Expo

```typescript
// src/api/client.ts
import axios, { AxiosInstance } from "axios";
import * as SecureStore from "expo-secure-store";
import Constants from "expo-constants";

class ApiClient {
  private client: AxiosInstance;

  constructor() {
    const baseURL =
      Constants.expoConfig?.extra?.apiUrl || "http://localhost:3000";

    this.client = axios.create({
      baseURL,
      timeout: 10000,
      headers: {
        "Content-Type": "application/json",
      },
    });

    // Request interceptor
    this.client.interceptors.request.use(
      async (config) => {
        const token = await SecureStore.getItemAsync("access_token");
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    // Response interceptor for token refresh
    this.client.interceptors.response.use(
      (response) => response,
      async (error) => {
        if (error.response?.status === 401) {
          await this.refreshToken();
          return this.client.request(error.config);
        }
        return Promise.reject(error);
      }
    );
  }

  private async refreshToken() {
    const refreshToken = await SecureStore.getItemAsync("refresh_token");
    const response = await this.client.post("/api/v1/mobile/auth/refresh", {
      refresh_token: refreshToken,
    });

    await SecureStore.setItemAsync(
      "access_token",
      response.data.data.access_token
    );
    await SecureStore.setItemAsync(
      "refresh_token",
      response.data.data.refresh_token
    );
  }

  // API methods
  async get<T>(url: string, params?: any) {
    const response = await this.client.get<T>(url, { params });
    return response.data;
  }

  async post<T>(url: string, data?: any) {
    const response = await this.client.post<T>(url, data);
    return response.data;
  }

  async put<T>(url: string, data?: any) {
    const response = await this.client.put<T>(url, data);
    return response.data;
  }

  async delete<T>(url: string) {
    const response = await this.client.delete<T>(url);
    return response.data;
  }
}

export default new ApiClient();
```

### Google Authentication with Expo

```typescript
// src/api/auth.ts
import * as WebBrowser from "expo-web-browser";
import * as Google from "expo-auth-session/providers/google";
import * as SecureStore from "expo-secure-store";
import { useEffect } from "react";
import apiClient from "./client";

WebBrowser.maybeCompleteAuthSession();

export const useGoogleAuth = () => {
  const [request, response, promptAsync] = Google.useAuthRequest({
    expoClientId: "YOUR_EXPO_CLIENT_ID.apps.googleusercontent.com",
    iosClientId: "YOUR_IOS_CLIENT_ID.apps.googleusercontent.com",
    androidClientId: "YOUR_ANDROID_CLIENT_ID.apps.googleusercontent.com",
    webClientId: "YOUR_WEB_CLIENT_ID.apps.googleusercontent.com",
  });

  useEffect(() => {
    if (response?.type === "success") {
      const { authentication } = response;
      handleGoogleLogin(authentication!.accessToken);
    }
  }, [response]);

  const handleGoogleLogin = async (googleToken: string) => {
    try {
      // Send Google token to your backend
      const response = await apiClient.post("/api/v1/mobile/auth/google", {
        google_token: googleToken,
      });

      // Store JWT tokens
      await SecureStore.setItemAsync(
        "access_token",
        response.data.access_token
      );
      await SecureStore.setItemAsync(
        "refresh_token",
        response.data.refresh_token
      );

      // Navigate to main app
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  return { request, promptAsync };
};
```

## Phase 4: Core Features Implementation (Weeks 4-6)

### Trip List Screen

```typescript
// src/screens/trips/TripsScreen.tsx
import React, { useState, useEffect } from "react";
import {
  View,
  FlatList,
  StyleSheet,
  RefreshControl,
  TouchableOpacity,
} from "react-native";
import { Card, Title, Paragraph, FAB } from "react-native-paper";
import { useQuery } from "@tanstack/react-query";
import { tripApi } from "../../api/trips";
import { Trip } from "../../api/types";

export default function TripsScreen({ navigation }: any) {
  const [refreshing, setRefreshing] = useState(false);

  const { data: trips, refetch } = useQuery({
    queryKey: ["trips"],
    queryFn: () => tripApi.getTrips(),
  });

  const onRefresh = async () => {
    setRefreshing(true);
    await refetch();
    setRefreshing(false);
  };

  const renderTrip = ({ item }: { item: Trip }) => (
    <TouchableOpacity
      onPress={() => navigation.navigate("TripDetail", { tripId: item.id })}
    >
      <Card style={styles.card}>
        <Card.Content>
          <Title>{`${item.origin} → ${item.destination}`}</Title>
          <Paragraph>{`${item.airline} ${item.flight_number}`}</Paragraph>
          <Paragraph>
            {new Date(item.departure_time * 1000).toLocaleDateString()}
          </Paragraph>
        </Card.Content>
      </Card>
    </TouchableOpacity>
  );

  return (
    <View style={styles.container}>
      <FlatList
        data={trips?.data || []}
        renderItem={renderTrip}
        keyExtractor={(item) => item.id.toString()}
        refreshControl={
          <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />
        }
      />
      <FAB
        style={styles.fab}
        icon="plus"
        onPress={() => navigation.navigate("AddTrip")}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#f5f5f5",
  },
  card: {
    margin: 10,
  },
  fab: {
    position: "absolute",
    margin: 16,
    right: 0,
    bottom: 0,
    backgroundColor: "#4A90E2",
  },
});
```

### Map Integration with Expo

```typescript
// src/screens/map/MapScreen.tsx
import React, { useEffect, useState } from "react";
import { StyleSheet, View } from "react-native";
import MapView, { Marker, Polyline, PROVIDER_GOOGLE } from "react-native-maps";
import * as Location from "expo-location";
import { useQuery } from "@tanstack/react-query";
import { tripApi } from "../../api/trips";

export default function MapScreen() {
  const [location, setLocation] = useState<Location.LocationObject | null>(
    null
  );

  const { data: trips } = useQuery({
    queryKey: ["trips-map"],
    queryFn: () => tripApi.getTrips(),
  });

  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== "granted") return;

      const currentLocation = await Location.getCurrentPositionAsync({});
      setLocation(currentLocation);
    })();
  }, []);

  return (
    <View style={styles.container}>
      <MapView
        style={styles.map}
        provider={PROVIDER_GOOGLE}
        initialRegion={{
          latitude: location?.coords.latitude || 37.78825,
          longitude: location?.coords.longitude || -122.4324,
          latitudeDelta: 0.0922,
          longitudeDelta: 0.0421,
        }}
      >
        {trips?.data?.map((trip) => (
          <React.Fragment key={trip.id}>
            <Marker
              coordinate={{
                latitude: trip.origin_lat || 0,
                longitude: trip.origin_lng || 0,
              }}
              title={trip.origin}
            />
            <Marker
              coordinate={{
                latitude: trip.destination_lat || 0,
                longitude: trip.destination_lng || 0,
              }}
              title={trip.destination}
            />
            <Polyline
              coordinates={[
                {
                  latitude: trip.origin_lat || 0,
                  longitude: trip.origin_lng || 0,
                },
                {
                  latitude: trip.destination_lat || 0,
                  longitude: trip.destination_lng || 0,
                },
              ]}
              strokeColor="#4A90E2"
              strokeWidth={2}
            />
          </React.Fragment>
        ))}
      </MapView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  map: {
    width: "100%",
    height: "100%",
  },
});
```

## Phase 5: Testing with Expo on WSL2/Windows (Week 7)

### Testing Setup for WSL2

```bash
# 1. Start your Go backend in WSL2
wsl
cd ~/trip-tracker
air  # or go run main.go

# 2. Get your network IP (in Windows PowerShell)
ipconfig
# Look for IPv4 Address under your active network adapter
# Example: 192.168.1.100

# 3. Configure port forwarding (Windows PowerShell as Admin)
netsh interface portproxy add v4tov4 `
  listenport=3000 `
  listenaddress=0.0.0.0 `
  connectport=3000 `
  connectaddress=(wsl hostname -I)

# 4. Update your Expo app config with your IP
# In app/.env.development
API_URL=http://192.168.1.100:3000
```

### Running Expo from Windows

```bash
# In Windows PowerShell (NOT WSL2)
cd C:\Projects\trip-tracker\app

# Start Expo development server
npx expo start

# Options will appear:
# › Press a to open Android emulator
# › Press i to open iOS simulator (Mac only)
# › Press w to open web browser
# › Scan QR code with Expo Go app
```

### Testing on Physical Devices

```bash
# 1. Ensure phone is on same WiFi network
# 2. Install Expo Go app
# 3. Scan QR code from terminal or browser

# For Android physical device
# Enable Developer Mode and USB Debugging
adb devices  # Verify device is connected
npx expo start --tunnel  # Use for tunneling with wsl2

# The app will connect to your backend at:
# http://192.168.1.100:3000 (your computer's IP)
```

### Expo Go Development Benefits

1. **Instant Updates**: Changes appear immediately without rebuilding
2. **Cross-Platform Testing**: Test iOS from Windows using Expo Go
3. **No Build Configuration**: Skip Android Studio/Xcode setup initially
4. **Developer Menu**: Shake device for debugging options

### Common WSL2 + Expo Issues

```bash
# Issue: Expo can't find Metro bundler
# Solution: Run Expo on Windows, not WSL2

# Issue: Mobile can't connect to WSL2 backend
# Solution: Use your Windows IP, not localhost or WSL2 IP

# Issue: Port 19000 (Expo) blocked by firewall
# Solution: Allow Node.js through Windows Firewall

# Clear Expo cache if having issues
npx expo start --clear --tunnel
```
