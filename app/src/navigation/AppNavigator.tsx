import React from 'react';
import { View, Text } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { Ionicons } from '@expo/vector-icons';
import { useSelector } from 'react-redux';

// Import screens
import { LoginScreen, SplashScreen } from '../screens/auth';
import { HomeScreen } from '../screens/home';
import { TripsListScreen } from '../screens/trips';
// import { MapScreen } from '../screens/map';
// import { ProfileScreen } from '../screens/profile';

import { colors, typography } from '../utils/theme';
import { RootState } from '../store';

const Stack = createStackNavigator();
const Tab = createBottomTabNavigator();

// Temporary placeholder screens
const MapScreen = () => (
  <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', backgroundColor: colors.background }}>
    <Text style={{ color: colors.text.primary, fontSize: 18 }}>Map Screen - Coming Soon</Text>
  </View>
);

const ProfileScreen = () => (
  <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', backgroundColor: colors.background }}>
    <Text style={{ color: colors.text.primary, fontSize: 18 }}>Profile Screen - Coming Soon</Text>
  </View>
);

// Tab Navigator for authenticated users
const TabNavigator = () => (
  <Tab.Navigator
    screenOptions={({ route }) => ({
      tabBarIcon: ({ focused, color, size }) => {
        let iconName: keyof typeof Ionicons.glyphMap;

        switch (route.name) {
          case 'Trips':
            iconName = focused ? 'airplane' : 'airplane-outline';
            break;
          case 'Map':
            iconName = focused ? 'map' : 'map-outline';
            break;
          case 'Profile':
            iconName = focused ? 'person' : 'person-outline';
            break;
          default:
            iconName = 'help-outline';
        }

        return <Ionicons name={iconName} size={size} color={color} />;
      },
      tabBarActiveTintColor: colors.primary,
      tabBarInactiveTintColor: colors.text.muted,
      tabBarStyle: {
        backgroundColor: colors.surface,
        borderTopColor: colors.border,
        borderTopWidth: 1,
        paddingBottom: 8,
        paddingTop: 8,
        height: 65,
      },
      tabBarLabelStyle: {
        fontSize: 12,
        fontFamily: typography.fontFamily.medium,
        marginTop: 4,
      },
      headerStyle: {
        backgroundColor: colors.surface,
        borderBottomColor: colors.border,
        borderBottomWidth: 1,
      },
      headerTitleStyle: {
        color: colors.text.primary,
        fontSize: typography.fontSize.lg,
        fontFamily: typography.fontFamily.semiBold,
      },
      headerTintColor: colors.text.primary,
    })}
  >
    <Tab.Screen 
      name="Trips" 
      component={TripsListScreen}
      options={{ title: 'My Trips' }}
    />
    <Tab.Screen 
      name="Map" 
      component={MapScreen}
      options={{ title: 'World Map' }}
    />
    <Tab.Screen 
      name="Profile" 
      component={ProfileScreen}
      options={{ title: 'Profile' }}
    />
  </Tab.Navigator>
);

// Main App Stack (for authenticated users)
const MainAppStack = () => (
  <Stack.Navigator
    screenOptions={{
      headerStyle: {
        backgroundColor: colors.surface,
        borderBottomColor: colors.border,
        borderBottomWidth: 1,
      },
      headerTitleStyle: {
        color: colors.text.primary,
        fontSize: typography.fontSize.lg,
        fontFamily: typography.fontFamily.semiBold,
      },
      headerTintColor: colors.text.primary,
      cardStyle: { backgroundColor: colors.background },
    }}
  >
    <Stack.Screen 
      name="MainTabs" 
      component={TabNavigator}
      options={{ headerShown: false }}
    />
    <Stack.Screen 
      name="TripDetail" 
      component={MapScreen} // Placeholder, will be replaced
      options={{ title: 'Trip Details' }}
    />
    <Stack.Screen 
      name="AddTrip" 
      component={MapScreen} // Placeholder, will be replaced
      options={{ 
        title: 'Add Trip',
        presentation: 'modal',
      }}
    />
    <Stack.Screen 
      name="EditTrip" 
      component={MapScreen} // Placeholder, will be replaced
      options={{ 
        title: 'Edit Trip',
        presentation: 'modal',
      }}
    />
  </Stack.Navigator>
);

// Auth Stack (for non-authenticated users)
const AuthStack = () => (
  <Stack.Navigator
    screenOptions={{
      headerShown: false,
      cardStyle: { backgroundColor: colors.background },
    }}
  >
    <Stack.Screen name="Splash" component={SplashScreen} />
    <Stack.Screen name="Home" component={HomeScreen} />
    <Stack.Screen name="Login" component={LoginScreen} />
  </Stack.Navigator>
);

// Main App Navigator
const AppNavigator: React.FC = () => {
  const { isAuthenticated, isInitializing } = useSelector(
    (state: RootState) => state.auth
  );

  return (
    <NavigationContainer>
      {isInitializing || (!isAuthenticated && !isInitializing) ? (
        <AuthStack />
      ) : (
        <MainAppStack />
      )}
    </NavigationContainer>
  );
};

export default AppNavigator;