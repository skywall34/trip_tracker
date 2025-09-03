import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface Toast {
  id: string;
  type: 'success' | 'error' | 'warning' | 'info';
  title: string;
  message?: string;
  duration?: number;
}

export interface Modal {
  id: string;
  type: 'confirm' | 'alert' | 'custom';
  title: string;
  message?: string;
  component?: string;
  onConfirm?: string; // Action type to dispatch
  onCancel?: string; // Action type to dispatch
}

interface UIState {
  // Loading states
  isAppLoading: boolean;
  activeRequests: Set<string>;
  
  // Toast notifications
  toasts: Toast[];
  
  // Modal system
  activeModal: Modal | null;
  
  // Navigation state
  currentTab: string;
  navigationHistory: string[];
  
  // Pull to refresh
  isRefreshing: boolean;
  
  // Network state
  isOnline: boolean;
  
  // App state
  isAppActive: boolean;
  
  // Search and filters
  searchQuery: string;
  activeFilters: Record<string, any>;
}

const initialState: UIState = {
  isAppLoading: false,
  activeRequests: new Set(),
  toasts: [],
  activeModal: null,
  currentTab: 'Home',
  navigationHistory: [],
  isRefreshing: false,
  isOnline: true,
  isAppActive: true,
  searchQuery: '',
  activeFilters: {},
};

const uiSlice = createSlice({
  name: 'ui',
  initialState,
  reducers: {
    // Loading states
    setAppLoading: (state, action: PayloadAction<boolean>) => {
      state.isAppLoading = action.payload;
    },
    
    addActiveRequest: (state, action: PayloadAction<string>) => {
      state.activeRequests.add(action.payload);
    },
    
    removeActiveRequest: (state, action: PayloadAction<string>) => {
      state.activeRequests.delete(action.payload);
    },
    
    // Toast management
    showToast: (state, action: PayloadAction<Omit<Toast, 'id'>>) => {
      const toast: Toast = {
        id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
        duration: 4000, // Default 4 seconds
        ...action.payload,
      };
      state.toasts.push(toast);
    },
    
    hideToast: (state, action: PayloadAction<string>) => {
      state.toasts = state.toasts.filter(toast => toast.id !== action.payload);
    },
    
    clearAllToasts: (state) => {
      state.toasts = [];
    },
    
    // Modal management
    showModal: (state, action: PayloadAction<Omit<Modal, 'id'>>) => {
      state.activeModal = {
        id: Date.now().toString(),
        ...action.payload,
      };
    },
    
    hideModal: (state) => {
      state.activeModal = null;
    },
    
    // Navigation
    setCurrentTab: (state, action: PayloadAction<string>) => {
      if (state.currentTab !== action.payload) {
        state.navigationHistory.push(state.currentTab);
      }
      state.currentTab = action.payload;
    },
    
    navigateBack: (state) => {
      if (state.navigationHistory.length > 0) {
        state.currentTab = state.navigationHistory.pop() || 'Home';
      }
    },
    
    clearNavigationHistory: (state) => {
      state.navigationHistory = [];
    },
    
    // Refresh state
    setRefreshing: (state, action: PayloadAction<boolean>) => {
      state.isRefreshing = action.payload;
    },
    
    // Network state
    setOnlineStatus: (state, action: PayloadAction<boolean>) => {
      state.isOnline = action.payload;
    },
    
    // App state
    setAppActive: (state, action: PayloadAction<boolean>) => {
      state.isAppActive = action.payload;
    },
    
    // Search and filters
    setSearchQuery: (state, action: PayloadAction<string>) => {
      state.searchQuery = action.payload;
    },
    
    setFilter: (state, action: PayloadAction<{ key: string; value: any }>) => {
      state.activeFilters[action.payload.key] = action.payload.value;
    },
    
    removeFilter: (state, action: PayloadAction<string>) => {
      delete state.activeFilters[action.payload];
    },
    
    clearFilters: (state) => {
      state.activeFilters = {};
    },
    
    // Reset UI state
    resetUI: (state) => {
      state.toasts = [];
      state.activeModal = null;
      state.searchQuery = '';
      state.activeFilters = {};
      state.isRefreshing = false;
    },
  },
});

// Selectors
export const selectActiveRequestsCount = (state: { ui: UIState }) => 
  state.ui.activeRequests.size;

export const selectIsLoading = (state: { ui: UIState }) => 
  state.ui.isAppLoading || state.ui.activeRequests.size > 0;

export const selectToasts = (state: { ui: UIState }) => state.ui.toasts;

export const selectActiveModal = (state: { ui: UIState }) => state.ui.activeModal;

export const selectCurrentTab = (state: { ui: UIState }) => state.ui.currentTab;

export const selectIsRefreshing = (state: { ui: UIState }) => state.ui.isRefreshing;

export const selectIsOnline = (state: { ui: UIState }) => state.ui.isOnline;

export const selectSearchQuery = (state: { ui: UIState }) => state.ui.searchQuery;

export const selectActiveFilters = (state: { ui: UIState }) => state.ui.activeFilters;

export const {
  setAppLoading,
  addActiveRequest,
  removeActiveRequest,
  showToast,
  hideToast,
  clearAllToasts,
  showModal,
  hideModal,
  setCurrentTab,
  navigateBack,
  clearNavigationHistory,
  setRefreshing,
  setOnlineStatus,
  setAppActive,
  setSearchQuery,
  setFilter,
  removeFilter,
  clearFilters,
  resetUI,
} = uiSlice.actions;

export default uiSlice.reducer;