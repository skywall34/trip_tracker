import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import apiClient from '../../api/client';
import { Trip, CreateTripRequest, UpdateTripRequest } from '../../api/types';

interface TripsState {
  trips: Trip[];
  currentTrip: Trip | null;
  isLoading: boolean;
  isCreating: boolean;
  isUpdating: boolean;
  isDeleting: boolean;
  error: string | null;
  lastFetch: number | null;
  hasMore: boolean;
  page: number;
}

const initialState: TripsState = {
  trips: [],
  currentTrip: null,
  isLoading: false,
  isCreating: false,
  isUpdating: false,
  isDeleting: false,
  error: null,
  lastFetch: null,
  hasMore: true,
  page: 1,
};

// Async thunks for trip operations
export const fetchTrips = createAsyncThunk(
  'trips/fetchTrips',
  async ({ page = 1, refresh = false }: { page?: number; refresh?: boolean } = {}, { rejectWithValue }) => {
    try {
      const response = await apiClient.get('/api/v1/trips', { page, limit: 20 });
      const trips = (response.data as Trip[]) || [];
      return { 
        trips, 
        page, 
        refresh,
        hasMore: trips.length === 20 
      };
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to fetch trips');
    }
  }
);

export const fetchTripById = createAsyncThunk(
  'trips/fetchTripById',
  async (tripId: number, { rejectWithValue }) => {
    try {
      const response = await apiClient.get(`/api/v1/trips/${tripId}`);
      return response.data as Trip;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to fetch trip');
    }
  }
);

export const createTrip = createAsyncThunk(
  'trips/createTrip',
  async (tripData: CreateTripRequest, { rejectWithValue }) => {
    try {
      const response = await apiClient.post('/api/v1/trips', tripData);
      return response.data as Trip;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to create trip');
    }
  }
);

export const updateTrip = createAsyncThunk(
  'trips/updateTrip',
  async (tripData: UpdateTripRequest, { rejectWithValue }) => {
    try {
      const { id, ...updateData } = tripData;
      const response = await apiClient.put(`/api/v1/trips/${id}`, updateData);
      return response.data as Trip;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to update trip');
    }
  }
);

export const deleteTrip = createAsyncThunk(
  'trips/deleteTrip',
  async (tripId: number, { rejectWithValue }) => {
    try {
      await apiClient.delete(`/api/v1/trips/${tripId}`);
      return tripId;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to delete trip');
    }
  }
);

const tripsSlice = createSlice({
  name: 'trips',
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
    setCurrentTrip: (state, action: PayloadAction<Trip | null>) => {
      state.currentTrip = action.payload;
    },
    resetTrips: (state) => {
      state.trips = [];
      state.page = 1;
      state.hasMore = true;
      state.lastFetch = null;
    },
  },
  extraReducers: (builder) => {
    builder
      // Fetch trips
      .addCase(fetchTrips.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(fetchTrips.fulfilled, (state, action) => {
        state.isLoading = false;
        const { trips, page, refresh, hasMore } = action.payload;
        
        if (refresh || page === 1) {
          state.trips = trips;
        } else {
          // Append new trips for pagination
          state.trips.push(...trips);
        }
        
        state.page = page;
        state.hasMore = hasMore;
        state.lastFetch = Date.now();
        state.error = null;
      })
      .addCase(fetchTrips.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload as string;
      })
      
      // Fetch trip by ID
      .addCase(fetchTripById.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(fetchTripById.fulfilled, (state, action) => {
        state.isLoading = false;
        state.currentTrip = action.payload as Trip;
        state.error = null;
      })
      .addCase(fetchTripById.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload as string;
      })
      
      // Create trip
      .addCase(createTrip.pending, (state) => {
        state.isCreating = true;
        state.error = null;
      })
      .addCase(createTrip.fulfilled, (state, action) => {
        state.isCreating = false;
        state.trips.unshift(action.payload as Trip); // Add new trip to beginning
        state.error = null;
      })
      .addCase(createTrip.rejected, (state, action) => {
        state.isCreating = false;
        state.error = action.payload as string;
      })
      
      // Update trip
      .addCase(updateTrip.pending, (state) => {
        state.isUpdating = true;
        state.error = null;
      })
      .addCase(updateTrip.fulfilled, (state, action) => {
        state.isUpdating = false;
        const updatedTrip = action.payload as Trip;
        const index = state.trips.findIndex(trip => trip.id === updatedTrip.id);
        if (index !== -1) {
          state.trips[index] = updatedTrip;
        }
        if (state.currentTrip?.id === updatedTrip.id) {
          state.currentTrip = updatedTrip;
        }
        state.error = null;
      })
      .addCase(updateTrip.rejected, (state, action) => {
        state.isUpdating = false;
        state.error = action.payload as string;
      })
      
      // Delete trip
      .addCase(deleteTrip.pending, (state) => {
        state.isDeleting = true;
        state.error = null;
      })
      .addCase(deleteTrip.fulfilled, (state, action) => {
        state.isDeleting = false;
        const tripId = action.payload;
        state.trips = state.trips.filter(trip => trip.id !== tripId);
        if (state.currentTrip?.id === tripId) {
          state.currentTrip = null;
        }
        state.error = null;
      })
      .addCase(deleteTrip.rejected, (state, action) => {
        state.isDeleting = false;
        state.error = action.payload as string;
      });
  },
});

export const { clearError, setCurrentTrip, resetTrips } = tripsSlice.actions;
export default tripsSlice.reducer;