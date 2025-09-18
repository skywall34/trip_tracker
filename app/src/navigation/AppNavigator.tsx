import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { Ionicons } from '@expo/vector-icons';
import { useSelector } from 'react-redux';
import { useSafeAreaInsets } from 'react-native-safe-area-context';

// Import screens
import { LoginScreen, SplashScreen } from '../screens/auth';
import { HomeScreen } from '../screens/home';
import { TripsListScreen } from '../screens/trips';
import { ProfileScreen } from '../screens/profile';

import { colors, typography } from '../utils/theme';
import { RootState } from '../store';
import { PlaceholderScreen } from '../components/common';

const Stack = createStackNavigator();
const Tab = createBottomTabNavigator();

// Placeholder screens for features under development
const MapScreen = () => <PlaceholderScreen title="World Map" subtitle="Interactive map visualization coming soon!" icon="map-outline" />;
const TripDetailScreen = () => <PlaceholderScreen title="Trip Details" subtitle="Detailed trip view coming soon!" icon="airplane-outline" />;
const AddTripScreen = () => <PlaceholderScreen title="Add Trip" subtitle="Trip creation form coming soon!" icon="add-circle-outline" />;
const EditTripScreen = () => <PlaceholderScreen title="Edit Trip" subtitle="Trip editing form coming soon!" icon="create-outline" />;


// Tab Navigator for authenticated users
const TabNavigator = () => {
  const insets = useSafeAreaInsets();

  return (
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
        paddingBottom: Math.max(insets.bottom, 8),
        paddingTop: 8,
        height: 65 + Math.max(insets.bottom - 8, 0),
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
};

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
      component={TripDetailScreen}
      options={{ title: 'Trip Details' }}
    />
    <Stack.Screen
      name="AddTrip"
      component={AddTripScreen}
      options={{
        title: 'Add Trip',
        presentation: 'modal',
      }}
    />
    <Stack.Screen
      name="EditTrip"
      component={EditTripScreen}
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