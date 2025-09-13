import axios, { AxiosInstance, AxiosError, InternalAxiosRequestConfig } from 'axios';
import * as SecureStore from 'expo-secure-store';
import Constants from 'expo-constants';
import { ApiResponse } from './types';

// Extend the Axios request config to include our retry flag
interface ExtendedAxiosRequestConfig extends InternalAxiosRequestConfig {
  _retry?: boolean;
}

// Get API URL based on environment and platform
const getApiUrl = (): string => {
  const configUrl = Constants.expoConfig?.extra?.apiUrl;
  
  if (configUrl) {
    console.log('Using configured API URL:', configUrl);
    return configUrl;
  }
  
  // Fallback logic for different environments
  if (__DEV__) {
    // Development environment
    const Platform = require('react-native').Platform;
    if (Platform.OS === 'android') {
      return 'http://10.0.2.2:3000'; // Android emulator
    } else if (Platform.OS === 'ios') {
      return 'http://localhost:3000'; // iOS simulator
    }
  }
  
  // Production fallback
  return Constants.expoConfig?.extra?.productionApiUrl || 'https://api.miastrips.com';
};

const API_BASE_URL = getApiUrl();

class ApiClient {
  private client: AxiosInstance;
  private refreshPromise: Promise<string> | null = null;

  constructor() {
    this.client = axios.create({
      baseURL: API_BASE_URL,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor to add auth token
    this.client.interceptors.request.use(
      async (config) => {
        try {
          const token = await SecureStore.getItemAsync('access_token');
          if (token) {
            config.headers.Authorization = `Bearer ${token}`;
          }
        } catch (error) {
          console.warn('Failed to get access token:', error);
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    // Response interceptor for token refresh and error handling
    this.client.interceptors.response.use(
      (response) => response,
      async (error: AxiosError) => {
        const originalRequest = error.config as ExtendedAxiosRequestConfig;

        if (
          error.response?.status === 401 &&
          originalRequest &&
          !originalRequest._retry
        ) {
          originalRequest._retry = true;

          try {
            const newToken = await this.refreshToken();
            if (originalRequest.headers) {
              originalRequest.headers.Authorization = `Bearer ${newToken}`;
            }
            return this.client(originalRequest);
          } catch (refreshError) {
            // Clear all tokens and redirect to login
            await this.clearAuth();
            return Promise.reject(refreshError);
          }
        }

        return Promise.reject(error);
      }
    );
  }

  private async refreshToken(): Promise<string> {
    if (this.refreshPromise) {
      return this.refreshPromise;
    }

    this.refreshPromise = (async () => {
      try {
        const refreshToken = await SecureStore.getItemAsync('refresh_token');
        if (!refreshToken) {
          throw new Error('No refresh token available');
        }

        const response = await axios.post(`${API_BASE_URL}/api/v1/mobile/auth/refresh`, {
          refresh_token: refreshToken,
        });

        const { access_token, refresh_token: newRefreshToken } = response.data.data;

        await SecureStore.setItemAsync('access_token', access_token);
        await SecureStore.setItemAsync('refresh_token', newRefreshToken);

        this.refreshPromise = null;
        return access_token;
      } catch (error) {
        this.refreshPromise = null;
        throw error;
      }
    })();

    return this.refreshPromise;
  }

  async clearAuth(): Promise<void> {
    try {
      await SecureStore.deleteItemAsync('access_token');
      await SecureStore.deleteItemAsync('refresh_token');
      await SecureStore.deleteItemAsync('user');
    } catch (error) {
      console.warn('Failed to clear auth tokens:', error);
    }
  }

  async get<T>(url: string, params?: any): Promise<ApiResponse<T>> {
    const response = await this.client.get(url, { params });
    return response.data;
  }

  async post<T>(url: string, data?: any): Promise<ApiResponse<T>> {
    const response = await this.client.post(url, data);
    return response.data;
  }

  async put<T>(url: string, data?: any): Promise<ApiResponse<T>> {
    const response = await this.client.put(url, data);
    return response.data;
  }

  async delete<T>(url: string): Promise<ApiResponse<T>> {
    const response = await this.client.delete(url);
    return response.data;
  }

  // Utility method to check if user is authenticated
  async isAuthenticated(): Promise<boolean> {
    try {
      const token = await SecureStore.getItemAsync('access_token');
      return !!token;
    } catch {
      return false;
    }
  }

  // Get stored user data
  async getStoredUser(): Promise<any> {
    try {
      const userStr = await SecureStore.getItemAsync('user');
      return userStr ? JSON.parse(userStr) : null;
    } catch {
      return null;
    }
  }

  // Test connectivity to backend
  async testConnection(): Promise<{ success: boolean; message: string; url: string }> {
    try {
      console.log('Testing connection to:', API_BASE_URL);
      
      // Test a simple endpoint (401 response is actually success - means server is reachable)
      const response = await this.client.get('/api/v1/trips', { 
        timeout: 5000,
        // Don't include authorization header for this test
        headers: {} 
      });
      
      return {
        success: true,
        message: 'Connected successfully',
        url: API_BASE_URL
      };
    } catch (error: any) {
      console.log('Connection test error:', error.message);
      console.log('Error details:', error.response?.status, error.code);
      
      // 401 Unauthorized means server is reachable but needs auth - this is actually success!
      if (error.response?.status === 401) {
        return {
          success: true,
          message: 'Connected successfully (auth required)',
          url: API_BASE_URL
        };
      }
      
      // Network errors
      if (error.code === 'NETWORK_ERROR' || error.message.includes('Network Error')) {
        return {
          success: false,
          message: `Cannot reach server at ${API_BASE_URL}. Check network connection.`,
          url: API_BASE_URL
        };
      }
      
      // Timeout errors
      if (error.code === 'ECONNABORTED') {
        return {
          success: false,
          message: `Connection timeout to ${API_BASE_URL}`,
          url: API_BASE_URL
        };
      }
      
      return {
        success: false,
        message: `Connection failed: ${error.message}`,
        url: API_BASE_URL
      };
    }
  }
}

export default new ApiClient();