import React, { useEffect, useState } from "react";
import {
  View,
  Text,
  StyleSheet,
  SafeAreaView,
  ScrollView,
  Alert,
  RefreshControl,
} from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useDispatch, useSelector } from "react-redux";

import { Button, Card } from "../../components/common";
import { colors, typography, spacing } from "../../utils/theme";
import { logout } from "../../store/slices/authSlice";
import { RootState, AppDispatch } from "../../store";
import apiClient from "../../api/client";

interface ProfileScreenProps {
  navigation: any;
}

export const ProfileScreen: React.FC<ProfileScreenProps> = ({ navigation }) => {
  const dispatch = useDispatch<AppDispatch>();
  const { user, isLoading } = useSelector((state: RootState) => state.auth);
  const [profileData, setProfileData] = useState<any>(null);
  const [refreshing, setRefreshing] = useState(false);
  const [loadingProfile, setLoadingProfile] = useState(false);

  const fetchProfile = async () => {
    try {
      setLoadingProfile(true);
      const response = await apiClient.getProfile();
      if (response.success && response.data) {
        setProfileData(response.data);
      }
    } catch (error) {
      console.error("Failed to fetch profile:", error);
      Alert.alert("Error", "Failed to load profile data");
    } finally {
      setLoadingProfile(false);
    }
  };

  const onRefresh = async () => {
    setRefreshing(true);
    await fetchProfile();
    setRefreshing(false);
  };

  useEffect(() => {
    fetchProfile();
  }, []);

  const handleLogout = () => {
    Alert.alert(
      "Logout",
      "Are you sure you want to logout?",
      [
        {
          text: "Cancel",
          style: "cancel",
        },
        {
          text: "Logout",
          style: "destructive",
          onPress: async () => {
            try {
              await dispatch(logout()).unwrap();
              // Navigation will be handled by auth state change
            } catch (error) {
              console.error("Logout error:", error);
              Alert.alert("Error", "Failed to logout. Please try again.");
            }
          },
        },
      ]
    );
  };

  const navigateToTrips = () => {
    navigation.navigate("Trips");
  };

  // Use profile data if available, fallback to user data from auth
  const displayData = profileData || user;

  if (!displayData) {
    return (
      <SafeAreaView style={styles.container}>
        <View style={styles.errorContainer}>
          <Ionicons name="person-circle-outline" size={64} color={colors.text.muted} />
          <Text style={styles.errorText}>No user data available</Text>
        </View>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView style={styles.container}>
      <ScrollView
        style={styles.scrollView}
        contentContainerStyle={styles.scrollContent}
        showsVerticalScrollIndicator={false}
        refreshControl={
          <RefreshControl
            refreshing={refreshing}
            onRefresh={onRefresh}
            tintColor={colors.primary}
            colors={[colors.primary]}
          />
        }
      >
        {/* Header */}
        <View style={styles.header}>
          <View style={styles.avatarContainer}>
            {displayData.picture ? (
              <Ionicons name="person" size={48} color={colors.primary} />
            ) : (
              <Ionicons name="person-circle" size={64} color={colors.primary} />
            )}
          </View>
          <Text style={styles.welcomeText}>Welcome back!</Text>
          <Text style={styles.nameText}>
            {displayData.name ||
             (displayData.first_name && displayData.last_name
               ? `${displayData.first_name} ${displayData.last_name}`
               : displayData.first_name || displayData.last_name || "User")}
          </Text>
        </View>

        {/* User Information Card */}
        <Card variant="glass" padding="lg" style={styles.infoCard}>
          <View style={styles.cardHeader}>
            <Ionicons name="person-outline" size={24} color={colors.primary} />
            <Text style={styles.cardTitle}>Account Information</Text>
          </View>

          <View style={styles.infoSection}>
            {/* Username */}
            {displayData.username && (
              <View style={styles.infoRow}>
                <Text style={styles.infoLabel}>Username</Text>
                <Text style={styles.infoValue}>{displayData.username}</Text>
              </View>
            )}

            {/* First Name */}
            {displayData.first_name && (
              <View style={styles.infoRow}>
                <Text style={styles.infoLabel}>First Name</Text>
                <Text style={styles.infoValue}>{displayData.first_name}</Text>
              </View>
            )}

            {/* Last Name */}
            {displayData.last_name && (
              <View style={styles.infoRow}>
                <Text style={styles.infoLabel}>Last Name</Text>
                <Text style={styles.infoValue}>{displayData.last_name}</Text>
              </View>
            )}

            {/* Email */}
            <View style={styles.infoRow}>
              <Text style={styles.infoLabel}>Email</Text>
              <Text style={styles.infoValue}>{displayData.email}</Text>
            </View>

            {/* Account Type */}
            <View style={styles.infoRow}>
              <Text style={styles.infoLabel}>Account Type</Text>
              <View style={styles.accountType}>
                {displayData.auth_provider === "google" ? (
                  <>
                    <Ionicons name="logo-google" size={16} color={colors.primary} />
                    <Text style={styles.infoValue}>Google Account</Text>
                  </>
                ) : (
                  <>
                    <Ionicons name="mail" size={16} color={colors.primary} />
                    <Text style={styles.infoValue}>Email Account</Text>
                  </>
                )}
              </View>
            </View>

            {/* Member Since */}
            {displayData.created_at && (
              <View style={styles.infoRow}>
                <Text style={styles.infoLabel}>Member Since</Text>
                <Text style={styles.infoValue}>
                  {new Date(displayData.created_at).toLocaleDateString()}
                </Text>
              </View>
            )}
          </View>
        </Card>

        {/* Quick Actions Card */}
        <Card variant="subtle" padding="lg" style={styles.actionsCard}>
          <View style={styles.cardHeader}>
            <Ionicons name="settings-outline" size={24} color={colors.primary} />
            <Text style={styles.cardTitle}>Quick Actions</Text>
          </View>

          <View style={styles.actionsSection}>
            <Button
              title="View My Trips"
              onPress={navigateToTrips}
              variant="outline"
              size="lg"
              fullWidth
              leftIcon={
                <Ionicons name="airplane-outline" size={20} color={colors.primary} />
              }
              style={styles.actionButton}
            />

            <Button
              title="Logout"
              onPress={handleLogout}
              variant="secondary"
              size="lg"
              fullWidth
              loading={isLoading}
              leftIcon={
                <Ionicons name="log-out-outline" size={20} color={colors.error} />
              }
              style={[styles.actionButton, styles.logoutButton]}
            />
          </View>
        </Card>

        {/* App Info Card */}
        <Card variant="subtle" padding="md" style={styles.appInfoCard}>
          <Text style={styles.appInfoTitle}>Mia's Trips</Text>
          <Text style={styles.appInfoSubtitle}>
            Track your adventures like a pro producer
          </Text>
          <Text style={styles.appInfoVersion}>Version 1.0.0</Text>
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

  scrollView: {
    flex: 1,
  },

  scrollContent: {
    padding: spacing.lg,
    paddingBottom: spacing["2xl"],
  },

  header: {
    alignItems: "center",
    marginBottom: spacing["2xl"],
  },

  avatarContainer: {
    width: 80,
    height: 80,
    borderRadius: 40,
    backgroundColor: colors.surface,
    alignItems: "center",
    justifyContent: "center",
    marginBottom: spacing.lg,
    borderWidth: 2,
    borderColor: colors.primary + "20",
  },

  welcomeText: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    marginBottom: spacing.xs,
  },

  nameText: {
    fontSize: typography.fontSize["2xl"],
    fontFamily: typography.fontFamily.bold,
    color: colors.text.primary,
  },

  infoCard: {
    marginBottom: spacing.lg,
  },

  cardHeader: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: spacing.lg,
  },

  cardTitle: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
    marginLeft: spacing.sm,
  },

  infoSection: {
    gap: spacing.md,
  },

  infoRow: {
    paddingVertical: spacing.sm,
    borderBottomWidth: 1,
    borderBottomColor: colors.border,
  },

  infoLabel: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.medium,
    color: colors.text.secondary,
    marginBottom: spacing.xs,
  },

  infoValue: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.primary,
  },

  accountType: {
    flexDirection: "row",
    alignItems: "center",
    gap: spacing.xs,
  },

  actionsCard: {
    marginBottom: spacing.lg,
  },

  actionsSection: {
    gap: spacing.md,
  },

  actionButton: {
    marginBottom: spacing.xs,
  },

  logoutButton: {
    borderColor: colors.error,
  },

  appInfoCard: {
    alignItems: "center",
    backgroundColor: colors.surface + "50",
  },

  appInfoTitle: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.bold,
    color: colors.text.primary,
    marginBottom: spacing.xs,
  },

  appInfoSubtitle: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: "center",
    marginBottom: spacing.sm,
  },

  appInfoVersion: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.muted,
  },

  errorContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    padding: spacing.lg,
  },

  errorText: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.muted,
    marginTop: spacing.lg,
  },
});