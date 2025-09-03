import React, { useEffect } from 'react';
import {
  View,
  Text,
  StyleSheet,
  SafeAreaView,
  StatusBar,
  Animated,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { useDispatch, useSelector } from 'react-redux';

import { Loading } from '../../components/common';
import { colors, typography, spacing } from '../../utils/theme';
import { initializeAuth } from '../../store/slices/authSlice';
import { RootState, AppDispatch } from '../../store';

interface SplashScreenProps {
  navigation: any;
}

export const SplashScreen: React.FC<SplashScreenProps> = ({ navigation }) => {
  const dispatch = useDispatch<AppDispatch>();
  const { isInitializing, isAuthenticated } = useSelector(
    (state: RootState) => state.auth
  );

  const fadeAnim = React.useRef(new Animated.Value(0)).current;
  const scaleAnim = React.useRef(new Animated.Value(0.8)).current;

  useEffect(() => {
    // Start animations
    Animated.parallel([
      Animated.timing(fadeAnim, {
        toValue: 1,
        duration: 800,
        useNativeDriver: true,
      }),
      Animated.spring(scaleAnim, {
        toValue: 1,
        friction: 4,
        useNativeDriver: true,
      }),
    ]).start();

    // Initialize authentication
    dispatch(initializeAuth());
  }, [dispatch, fadeAnim, scaleAnim]);

  useEffect(() => {
    if (!isInitializing) {
      // Add a small delay for better UX
      const timer = setTimeout(() => {
        if (isAuthenticated) {
          navigation.replace('MainApp');
        } else {
          navigation.replace('Home');
        }
      }, 500);

      return () => clearTimeout(timer);
    }
  }, [isInitializing, isAuthenticated, navigation]);

  return (
    <SafeAreaView style={styles.container}>
      <StatusBar barStyle="light-content" backgroundColor={colors.background} />
      
      <View style={styles.content}>
        <Animated.View
          style={[
            styles.logoContainer,
            {
              opacity: fadeAnim,
              transform: [{ scale: scaleAnim }],
            },
          ]}
        >
          <View style={styles.iconWrapper}>
            <Ionicons name="airplane" size={60} color={colors.primary} />
          </View>
          
          <Text style={styles.title}>Mia's Trips</Text>
          <Text style={styles.subtitle}>Track your adventures</Text>
        </Animated.View>

        <View style={styles.loadingContainer}>
          <Loading
            size="large"
            color={colors.primary}
            message="Loading your trips..."
          />
        </View>
      </View>

      <View style={styles.footer}>
        <Text style={styles.versionText}>Version 1.0.0</Text>
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.background,
  },

  content: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: spacing.lg,
  },

  logoContainer: {
    alignItems: 'center',
    marginBottom: spacing['3xl'],
  },

  iconWrapper: {
    width: 120,
    height: 120,
    borderRadius: 60,
    backgroundColor: colors.surface,
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: spacing.xl,
    borderWidth: 2,
    borderColor: colors.border,
    // Add glow effect
    shadowColor: colors.primary,
    shadowOffset: { width: 0, height: 0 },
    shadowOpacity: 0.3,
    shadowRadius: 20,
    elevation: 10,
  },

  title: {
    fontSize: typography.fontSize['4xl'],
    fontFamily: typography.fontFamily.bold,
    color: colors.text.primary,
    textAlign: 'center',
    marginBottom: spacing.sm,
  },

  subtitle: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: 'center',
  },

  loadingContainer: {
    position: 'absolute',
    bottom: spacing['2xl'],
    width: '100%',
  },

  footer: {
    alignItems: 'center',
    paddingBottom: spacing.lg,
  },

  versionText: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.muted,
  },
});

