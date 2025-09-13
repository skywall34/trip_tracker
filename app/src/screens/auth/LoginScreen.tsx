import React, { useEffect, useState } from "react";
import {
  View,
  Text,
  StyleSheet,
  SafeAreaView,
  StatusBar,
  Alert,
  Platform,
  KeyboardAvoidingView,
  ScrollView,
} from "react-native";
import { Ionicons } from "@expo/vector-icons";
import * as WebBrowser from "expo-web-browser";
import * as Google from "expo-auth-session/providers/google";
import { useDispatch, useSelector } from "react-redux";

import { Button, Card, Input } from "../../components/common";
import { colors, typography, spacing } from "../../utils/theme";
import { loginWithEmail, loginWithGoogle } from "../../store/slices/authSlice";
import { RootState, AppDispatch } from "../../store";
import apiClient from "../../api/client";

// Complete auth session for Google OAuth
WebBrowser.maybeCompleteAuthSession();

interface LoginScreenProps {
  navigation: any;
}

export const LoginScreen: React.FC<LoginScreenProps> = ({ navigation }) => {
  const dispatch = useDispatch<AppDispatch>();
  const { isLoading, error } = useSelector((state: RootState) => state.auth);

  // Form state
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [emailError, setEmailError] = useState("");
  const [passwordError, setPasswordError] = useState("");
  const [connectionStatus, setConnectionStatus] = useState<{
    tested: boolean;
    success: boolean;
    message: string;
    url: string;
  } | null>(null);

  // Configure Google OAuth (you'll need to set up proper client IDs)
  const [request, response, promptAsync] = Google.useAuthRequest({
    clientId: Platform.select({
      // You'll need to replace these with your actual Google OAuth client IDs
      ios: "731083279268-udc3jitbjbn40k281oa1ppourlk9ikuj.apps.googleusercontent.com",
      android:
        "731083279268-udc3jitbjbn40k281oa1ppourlk9ikuj.apps.googleusercontent.com",
      default:
        "731083279268-udc3jitbjbn40k281oa1ppourlk9ikuj.apps.googleusercontent.com",
    }),
    scopes: ["openid", "profile", "email"],
  });

  // Handle Google OAuth response
  useEffect(() => {
    if (response?.type === "success") {
      const { accessToken } = response.authentication || {};
      if (accessToken) {
        dispatch(loginWithGoogle(accessToken));
      }
    } else if (response?.type === "error") {
      Alert.alert(
        "Authentication Error",
        "Failed to sign in with Google. Please try again.",
        [{ text: "OK" }]
      );
    }
  }, [response, dispatch]);

  // Show error alert if login fails
  useEffect(() => {
    if (error) {
      Alert.alert("Login Failed", error);
    }
  }, [error]);

  // Validate email
  const validateEmail = (email: string): boolean => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!email) {
      setEmailError("Email is required");
      return false;
    } else if (!emailRegex.test(email)) {
      setEmailError("Please enter a valid email address");
      return false;
    } else {
      setEmailError("");
      return true;
    }
  };

  // Validate password
  const validatePassword = (password: string): boolean => {
    if (!password) {
      setPasswordError("Password is required");
      return false;
    } else if (password.length < 6) {
      setPasswordError("Password must be at least 6 characters");
      return false;
    } else {
      setPasswordError("");
      return true;
    }
  };

  // Handle email/password login
  const handleEmailLogin = async () => {
    const isEmailValid = validateEmail(email);
    const isPasswordValid = validatePassword(password);

    if (!isEmailValid || !isPasswordValid) {
      return;
    }

    try {
      await dispatch(loginWithEmail({ email, password })).unwrap();
      // Navigation will be handled by auth state change
    } catch (error: any) {
      // Error is already handled by Redux and displayed via useEffect above
      console.error(`${email}, ${password} Failed!`);
      console.error("Login error:", error);
    }
  };

  const handleGoogleLogin = () => {
    promptAsync();
  };

  // Test connection to backend
  const testConnection = async () => {
    try {
      const result = await apiClient.testConnection();
      setConnectionStatus({
        tested: true,
        success: result.success,
        message: result.message,
        url: result.url,
      });
      
      if (result.success) {
        Alert.alert(
          "Connection Test",
          `✅ ${result.message}\n\nServer: ${result.url}`,
          [{ text: "OK" }]
        );
      } else {
        Alert.alert(
          "Connection Failed",
          `❌ ${result.message}\n\nTrying to reach: ${result.url}\n\nMake sure:\n• Your phone and computer are on the same WiFi\n• Backend server is running\n• IP address is correct`,
          [{ text: "OK" }]
        );
      }
    } catch (error) {
      Alert.alert(
        "Connection Test Failed",
        `Error testing connection: ${error}`,
        [{ text: "OK" }]
      );
    }
  };

  return (
    <SafeAreaView style={styles.container}>
      <StatusBar barStyle="light-content" backgroundColor={colors.background} />

      <KeyboardAvoidingView
        style={styles.keyboardView}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
      >
        <ScrollView
          style={styles.scrollView}
          contentContainerStyle={styles.scrollContent}
          keyboardShouldPersistTaps="handled"
        >
          {/* Header */}
          <View style={styles.header}>
            <View style={styles.logoContainer}>
              <Ionicons name="airplane" size={48} color={colors.primary} />
            </View>
            <Text style={styles.title}>Welcome Back</Text>
            <Text style={styles.subtitle}>Login to continue your journey</Text>
          </View>

          {/* Login Card */}
          <Card variant="glass" padding="lg" style={styles.loginCard}>
            <Text style={styles.cardTitle}>Sign In</Text>
            <Text style={styles.cardSubtitle}>
              Enter your credentials to access your account
            </Text>

            {/* Email Input */}
            <View style={styles.inputGroup}>
              <Text style={styles.label}>Email</Text>
              <Input
                value={email}
                onChangeText={setEmail}
                placeholder="Enter your email"
                keyboardType="email-address"
                autoCapitalize="none"
                autoCorrect={false}
                error={emailError}
                onBlur={() => validateEmail(email)}
              />
            </View>

            {/* Password Input */}
            <View style={styles.inputGroup}>
              <Text style={styles.label}>Password</Text>
              <Input
                value={password}
                onChangeText={setPassword}
                placeholder="Enter your password"
                secureTextEntry
                error={passwordError}
                onBlur={() => validatePassword(password)}
              />
            </View>

            {/* Login Button */}
            <Button
              title="Sign In"
              onPress={handleEmailLogin}
              loading={isLoading}
              disabled={isLoading}
              variant="primary"
              size="lg"
              fullWidth
              style={styles.loginButton}
            />

            {/* Links */}
            <View style={styles.linksContainer}>
              <Button
                title="Register"
                onPress={() => {
                  /* Navigate to register */
                }}
                variant="ghost"
                size="sm"
              />
              <Button
                title="Forgot Password?"
                onPress={() => {
                  /* Navigate to forgot password */
                }}
                variant="ghost"
                size="sm"
              />
            </View>

            {/* Divider */}
            <View style={styles.divider}>
              <View style={styles.dividerLine} />
              <Text style={styles.dividerText}>OR</Text>
              <View style={styles.dividerLine} />
            </View>

            {/* Google Login */}
            <Button
              title="Continue with Google"
              onPress={handleGoogleLogin}
              loading={isLoading}
              disabled={isLoading || !request}
              variant="secondary"
              size="lg"
              fullWidth
              leftIcon={
                <Ionicons
                  name="logo-google"
                  size={20}
                  color={colors.text.primary}
                />
              }
              style={styles.googleButton}
            />

            <Text style={styles.termsText}>
              By signing in, you agree to our Terms of Service and Privacy
              Policy
            </Text>
          </Card>

          {/* Connection Test Button - Development Only */}
          {__DEV__ && (
            <Card variant="subtle" padding="md" style={styles.debugCard}>
              <Text style={styles.debugTitle}>Development Tools</Text>
              <Button
                title="Test Backend Connection"
                onPress={testConnection}
                variant="outline"
                size="sm"
                fullWidth
                style={styles.testButton}
              />
              {connectionStatus && (
                <View style={styles.connectionStatus}>
                  <Text
                    style={[
                      styles.connectionText,
                      {
                        color: connectionStatus.success
                          ? colors.success
                          : colors.error,
                      },
                    ]}
                  >
                    {connectionStatus.success ? "✅" : "❌"}{" "}
                    {connectionStatus.message}
                  </Text>
                  <Text style={styles.connectionUrl}>
                    {connectionStatus.url}
                  </Text>
                </View>
              )}
            </Card>
          )}
        </ScrollView>
      </KeyboardAvoidingView>
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
    justifyContent: "center",
    paddingHorizontal: spacing.lg,
    paddingVertical: spacing.xl,
  },

  header: {
    alignItems: "center",
    marginBottom: spacing["2xl"],
  },

  logoContainer: {
    width: 80,
    height: 80,
    borderRadius: 40,
    backgroundColor: colors.surface,
    alignItems: "center",
    justifyContent: "center",
    marginBottom: spacing.lg,
    borderWidth: 1,
    borderColor: colors.border,
  },

  title: {
    fontSize: typography.fontSize["4xl"],
    fontFamily: typography.fontFamily.bold,
    color: colors.text.primary,
    textAlign: "center",
    marginBottom: spacing.sm,
  },

  subtitle: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: "center",
  },

  loginCard: {
    marginBottom: spacing.xl,
  },

  cardTitle: {
    fontSize: typography.fontSize["2xl"],
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
    textAlign: "center",
    marginBottom: spacing.sm,
  },

  cardSubtitle: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: "center",
    marginBottom: spacing.xl,
  },

  googleButton: {
    marginBottom: spacing.lg,
  },

  termsText: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.muted,
    textAlign: "center",
    lineHeight: typography.lineHeight.xs + 2,
  },

  footer: {
    alignItems: "center",
  },

  footerText: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
  },

  footerLink: {
    color: colors.primary,
    fontFamily: typography.fontFamily.semiBold,
  },
  keyboardView: {
    flex: 1,
  },
  scrollView: {
    flex: 1,
  },
  scrollContent: {
    flexGrow: 1,
    justifyContent: "center",
    padding: spacing.md,
  },
  inputGroup: {
    marginBottom: spacing.md,
  },
  label: {
    color: colors.text.secondary,
    fontFamily: typography.fontFamily.regular,
    fontSize: typography.fontSize.sm,
    marginBottom: spacing.xs,
  },
  loginButton: {
    marginTop: spacing.md,
  },
  linksContainer: {
    flexDirection: "row",
    justifyContent: "space-between",
    marginTop: spacing.sm,
  },
  divider: {
    flexDirection: "row",
    alignItems: "center",
    marginVertical: spacing.lg,
  },
  dividerLine: {
    flex: 1,
    height: 1,
    backgroundColor: colors.border,
  },
  dividerText: {
    marginHorizontal: spacing.sm,
    color: colors.text.secondary,
    fontFamily: typography.fontFamily.regular,
    fontSize: typography.fontSize.sm,
  },
  debugCard: {
    marginTop: spacing.lg,
    backgroundColor: colors.surface,
    borderColor: colors.border,
    borderWidth: 1,
  },
  debugTitle: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.secondary,
    textAlign: "center",
    marginBottom: spacing.sm,
  },
  testButton: {
    marginBottom: spacing.sm,
  },
  connectionStatus: {
    padding: spacing.sm,
    backgroundColor: colors.background,
    borderRadius: 8,
  },
  connectionText: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.medium,
    textAlign: "center",
    marginBottom: spacing.xs,
  },
  connectionUrl: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.muted,
    textAlign: "center",
  },
});
