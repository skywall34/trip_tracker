import React, { useState } from 'react';
import {
  View,
  Text,
  StyleSheet,
  SafeAreaView,
  StatusBar,
  TextInput,
  TouchableOpacity,
  ScrollView,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';

import { Button, Card } from '../../components/common';
import { colors, typography, spacing, borderRadius } from '../../utils/theme';

interface HomeScreenProps {
  navigation: any;
}

export const HomeScreen: React.FC<HomeScreenProps> = ({ navigation }) => {
  const [flightCode, setFlightCode] = useState('');

  const handleSearchFlight = () => {
    // TODO: Implement flight search when backend integration is ready
    console.log('Searching for flight:', flightCode);
  };

  const handleManualAdd = () => {
    // Navigate to login first since user needs to be authenticated to add trips
    navigation.navigate('Login');
  };

  const handleGetStarted = () => {
    navigation.navigate('Login');
  };

  return (
    <SafeAreaView style={styles.container}>
      <StatusBar barStyle="light-content" backgroundColor={colors.background} />
      
      <ScrollView 
        style={styles.scrollContainer}
        contentContainerStyle={styles.scrollContent}
        showsVerticalScrollIndicator={false}
      >
        {/* Header */}
        <View style={styles.header}>
          <View style={styles.logoContainer}>
            <Ionicons name="airplane" size={32} color={colors.primary} />
          </View>
          <Text style={styles.appTitle}>Mia's Trips</Text>
        </View>

        {/* Hero Section */}
        <Card variant="glass" style={styles.heroCard}>
          <View style={styles.heroContent}>
            <Text style={styles.heroTitle}>
              Track your adventures like a pro producer.
            </Text>
            <Text style={styles.heroSubtitle}>
              Search flights, add trips, and see what's next—right from here.
            </Text>

            {/* Search Form */}
            <View style={styles.searchForm}>
              <View style={styles.searchInputContainer}>
                <TextInput
                  style={styles.searchInput}
                  placeholder="Flight IATA code (e.g. UA100)"
                  placeholderTextColor={colors.text.muted}
                  value={flightCode}
                  onChangeText={setFlightCode}
                  autoCapitalize="characters"
                />
              </View>
              
              <TouchableOpacity
                style={styles.searchButton}
                onPress={handleSearchFlight}
                disabled={!flightCode.trim()}
              >
                <Text style={styles.searchButtonText}>Search</Text>
              </TouchableOpacity>
              
              <TouchableOpacity
                style={styles.manualButton}
                onPress={handleManualAdd}
              >
                <Text style={styles.manualButtonText}>Manual Add</Text>
              </TouchableOpacity>
            </View>
          </View>
        </Card>

        {/* Features Section */}
        <View style={styles.featuresSection}>
          <Text style={styles.sectionTitle}>Why Choose Mia's Trips?</Text>
          
          <View style={styles.featuresGrid}>
            <Card variant="default" padding="lg" style={styles.featureCard}>
              <Ionicons name="airplane-outline" size={32} color={colors.primary} style={styles.featureIcon} />
              <Text style={styles.featureTitle}>Flight Tracking</Text>
              <Text style={styles.featureDescription}>
                Automatically track your flights with real-time updates
              </Text>
            </Card>

            <Card variant="default" padding="lg" style={styles.featureCard}>
              <Ionicons name="map-outline" size={32} color={colors.primary} style={styles.featureIcon} />
              <Text style={styles.featureTitle}>World Map</Text>
              <Text style={styles.featureDescription}>
                Visualize all your trips on an interactive world map
              </Text>
            </Card>

            <Card variant="default" padding="lg" style={styles.featureCard}>
              <Ionicons name="stats-chart-outline" size={32} color={colors.primary} style={styles.featureIcon} />
              <Text style={styles.featureTitle}>Travel Stats</Text>
              <Text style={styles.featureDescription}>
                See your travel statistics and achievements
              </Text>
            </Card>

            <Card variant="default" padding="lg" style={styles.featureCard}>
              <Ionicons name="cloud-outline" size={32} color={colors.primary} style={styles.featureIcon} />
              <Text style={styles.featureTitle}>Sync Everywhere</Text>
              <Text style={styles.featureDescription}>
                Access your trips on web and mobile seamlessly
              </Text>
            </Card>
          </View>
        </View>

        {/* Call to Action */}
        <Card variant="elevated" padding="lg" style={styles.ctaCard}>
          <Text style={styles.ctaTitle}>Ready to start tracking?</Text>
          <Text style={styles.ctaSubtitle}>
            Sign in with Google to begin your journey
          </Text>
          
          <Button
            title="Get Started"
            onPress={handleGetStarted}
            variant="primary"
            size="lg"
            fullWidth
            leftIcon={
              <Ionicons name="logo-google" size={20} color={colors.text.primary} />
            }
            style={styles.ctaButton}
          />
        </Card>
      </ScrollView>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.background,
  },

  scrollContainer: {
    flex: 1,
  },

  scrollContent: {
    paddingBottom: spacing.xl,
  },

  header: {
    alignItems: 'center',
    paddingTop: spacing.lg,
    paddingHorizontal: spacing.lg,
    marginBottom: spacing.xl,
  },

  logoContainer: {
    width: 64,
    height: 64,
    borderRadius: 32,
    backgroundColor: colors.surface,
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: spacing.md,
    borderWidth: 1,
    borderColor: colors.border,
  },

  appTitle: {
    fontSize: typography.fontSize['2xl'],
    fontFamily: typography.fontFamily.bold,
    color: colors.text.primary,
    textAlign: 'center',
  },

  heroCard: {
    marginHorizontal: spacing.lg,
    marginBottom: spacing.xl,
  },

  heroContent: {
    padding: spacing.lg,
  },

  heroTitle: {
    fontSize: typography.fontSize['2xl'],
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
    textAlign: 'center',
    marginBottom: spacing.sm,
    lineHeight: typography.lineHeight['2xl'],
  },

  heroSubtitle: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: 'center',
    marginBottom: spacing.xl,
    lineHeight: typography.lineHeight.base,
  },

  searchForm: {
    gap: spacing.sm,
  },

  searchInputContainer: {
    position: 'relative',
  },

  searchInput: {
    backgroundColor: colors.elevated,
    borderWidth: 1,
    borderColor: colors.border,
    borderRadius: borderRadius.lg,
    paddingHorizontal: spacing.md,
    paddingVertical: spacing.md,
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.primary,
  },

  searchButton: {
    backgroundColor: colors.primary,
    borderRadius: borderRadius.lg,
    paddingVertical: spacing.md,
    alignItems: 'center',
    justifyContent: 'center',
  },

  searchButtonText: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
  },

  manualButton: {
    borderWidth: 1,
    borderColor: colors.border,
    borderRadius: borderRadius.lg,
    paddingVertical: spacing.md,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: 'transparent',
  },

  manualButtonText: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.medium,
    color: colors.text.secondary,
  },

  featuresSection: {
    paddingHorizontal: spacing.lg,
    marginBottom: spacing.xl,
  },

  sectionTitle: {
    fontSize: typography.fontSize.xl,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
    textAlign: 'center',
    marginBottom: spacing.xl,
  },

  featuresGrid: {
    gap: spacing.lg,
  },

  featureCard: {
    alignItems: 'center',
  },

  featureIcon: {
    marginBottom: spacing.md,
  },

  featureTitle: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
    textAlign: 'center',
    marginBottom: spacing.sm,
  },

  featureDescription: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: 'center',
    lineHeight: typography.lineHeight.sm + 2,
  },

  ctaCard: {
    marginHorizontal: spacing.lg,
    alignItems: 'center',
  },

  ctaTitle: {
    fontSize: typography.fontSize.xl,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
    textAlign: 'center',
    marginBottom: spacing.sm,
  },

  ctaSubtitle: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: 'center',
    marginBottom: spacing.xl,
  },

  ctaButton: {
    minWidth: 200,
  },
});