import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import * as SecureStore from "expo-secure-store";
import apiClient from "../../api/client";
import { User } from "../../api/types";

interface AuthState {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
  isInitializing: boolean;
}

const initialState: AuthState = {
  user: null,
  accessToken: null,
  refreshToken: null,
  isAuthenticated: false,
  isLoading: false,
  error: null,
  isInitializing: true,
};

// Async thunks for authentication
export const initializeAuth = createAsyncThunk("auth/initialize", async () => {
  try {
    const accessToken = await SecureStore.getItemAsync("access_token");
    const refreshToken = await SecureStore.getItemAsync("refresh_token");
    const userData = await SecureStore.getItemAsync("user_data");

    if (accessToken && refreshToken && userData) {
      const user = JSON.parse(userData);
      return { user, accessToken, refreshToken };
    }
    return null;
  } catch (error) {
    console.error("Error initializing auth:", error);
    return null;
  }
});

export const loginWithEmail = createAsyncThunk(
  "auth/loginWithEmail",
  async (
    { email, password }: { email: string; password: string },
    { rejectWithValue }
  ) => {
    try {
      const response = await apiClient.post("/api/v1/mobile/auth/login", {
        email,
        password,
      });

      console.log("Login response:", JSON.stringify(response.data, null, 2));

      const response_data = response.data as {
        user: User;
        access_token: string;
        refresh_token: string;
        expires_in?: number;
      };
      
      // Check if we have the required fields
      if (!response_data.user || !response_data.access_token || !response_data.refresh_token) {
        console.log("Login failed - missing required fields:", response_data);
        return rejectWithValue("Login failed - missing authentication data");
      }

      const { user, access_token, refresh_token } = response_data;

      // Store tokens securely
      await SecureStore.setItemAsync("access_token", access_token);
      await SecureStore.setItemAsync("refresh_token", refresh_token);
      await SecureStore.setItemAsync("user_data", JSON.stringify(user));

      return { user, accessToken: access_token, refreshToken: refresh_token };
    } catch (error: any) {
      // Extract serializable error message
      const errorMessage = error.response?.data?.error?.message || error.message || "An unknown error occurred";
      return rejectWithValue(errorMessage);
    }
  }
);

export const loginWithGoogle = createAsyncThunk(
  "auth/loginWithGoogle",
  async (googleToken: string, { rejectWithValue }) => {
    try {
      const response = await apiClient.post("/api/v1/mobile/auth/google", {
        google_token: googleToken,
      });

      const data = response.data as {
        user: User;
        access_token: string;
        refresh_token: string;
      };
      const { user, access_token, refresh_token } = data;

      // Store tokens securely
      await SecureStore.setItemAsync("access_token", access_token);
      await SecureStore.setItemAsync("refresh_token", refresh_token);
      await SecureStore.setItemAsync("user_data", JSON.stringify(user));

      return { user, accessToken: access_token, refreshToken: refresh_token };
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || "Login failed");
    }
  }
);

export const refreshAccessToken = createAsyncThunk(
  "auth/refreshToken",
  async (_, { getState, rejectWithValue }) => {
    try {
      const refreshToken = await SecureStore.getItemAsync("refresh_token");

      if (!refreshToken) {
        throw new Error("No refresh token available");
      }

      const response = await apiClient.post("/api/v1/mobile/auth/refresh", {
        refresh_token: refreshToken,
      });

      const data = response.data as {
        access_token: string;
        refresh_token: string;
      };
      const { access_token, refresh_token: newRefreshToken } = data;

      // Update stored tokens
      await SecureStore.setItemAsync("access_token", access_token);
      await SecureStore.setItemAsync("refresh_token", newRefreshToken);

      return { accessToken: access_token, refreshToken: newRefreshToken };
    } catch (error: any) {
      return rejectWithValue(
        error.response?.data?.message || "Token refresh failed"
      );
    }
  }
);

export const logout = createAsyncThunk("auth/logout", async () => {
  try {
    // Call logout endpoint if needed
    await apiClient.post("/api/v1/mobile/auth/logout");
  } catch (error) {
    console.error("Logout API error:", error);
  } finally {
    // Always clear local storage
    await SecureStore.deleteItemAsync("access_token");
    await SecureStore.deleteItemAsync("refresh_token");
    await SecureStore.deleteItemAsync("user_data");
  }
});

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
    setTokens: (
      state,
      action: PayloadAction<{ accessToken: string; refreshToken: string }>
    ) => {
      state.accessToken = action.payload.accessToken;
      state.refreshToken = action.payload.refreshToken;
    },
  },
  extraReducers: (builder) => {
    builder
      // Initialize auth
      .addCase(initializeAuth.pending, (state) => {
        state.isInitializing = true;
      })
      .addCase(initializeAuth.fulfilled, (state, action) => {
        state.isInitializing = false;
        if (action.payload) {
          state.user = action.payload.user;
          state.accessToken = action.payload.accessToken;
          state.refreshToken = action.payload.refreshToken;
          state.isAuthenticated = true;
        } else {
          state.isAuthenticated = false;
        }
      })
      .addCase(initializeAuth.rejected, (state) => {
        state.isInitializing = false;
        state.isAuthenticated = false;
      })

      // Email login
      .addCase(loginWithEmail.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(loginWithEmail.fulfilled, (state, action) => {
        state.isLoading = false;
        state.user = action.payload.user;
        state.accessToken = action.payload.accessToken;
        state.refreshToken = action.payload.refreshToken;
        state.isAuthenticated = true;
        state.error = null;
      })
      .addCase(loginWithEmail.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload as string;
        state.isAuthenticated = false;
      })

      // Google login
      .addCase(loginWithGoogle.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(loginWithGoogle.fulfilled, (state, action) => {
        state.isLoading = false;
        state.user = action.payload.user;
        state.accessToken = action.payload.accessToken;
        state.refreshToken = action.payload.refreshToken;
        state.isAuthenticated = true;
        state.error = null;
      })
      .addCase(loginWithGoogle.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload as string;
        state.isAuthenticated = false;
      })

      // Token refresh
      .addCase(refreshAccessToken.fulfilled, (state, action) => {
        state.accessToken = action.payload.accessToken;
        state.refreshToken = action.payload.refreshToken;
      })
      .addCase(refreshAccessToken.rejected, (state) => {
        // If refresh fails, logout user
        state.user = null;
        state.accessToken = null;
        state.refreshToken = null;
        state.isAuthenticated = false;
      })

      // Logout
      .addCase(logout.fulfilled, (state) => {
        state.user = null;
        state.accessToken = null;
        state.refreshToken = null;
        state.isAuthenticated = false;
        state.error = null;
      });
  },
});

export const { clearError, setTokens } = authSlice.actions;
export default authSlice.reducer;
