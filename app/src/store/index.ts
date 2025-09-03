import { configureStore } from '@reduxjs/toolkit';
import authReducer from './slices/authSlice';
import tripsReducer from './slices/tripsSlice';
import uiReducer from './slices/uiSlice';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    trips: tripsReducer,
    ui: uiReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [
          'persist/PERSIST',
          'persist/REHYDRATE',
          // Ignore Set objects in UI slice
          'ui/addActiveRequest',
          'ui/removeActiveRequest',
        ],
        ignoredActionsPaths: ['payload.activeRequests'],
        ignoredPaths: ['ui.activeRequests'],
      },
    }),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;