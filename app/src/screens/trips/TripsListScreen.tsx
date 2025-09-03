import React, { useEffect, useState, useCallback } from 'react';
import {
  View,
  Text,
  StyleSheet,
  FlatList,
  RefreshControl,
  Alert,
  TouchableOpacity,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { useDispatch, useSelector } from 'react-redux';
import { useFocusEffect } from '@react-navigation/native';

import { Loading, TripCardSkeleton, Button } from '../../components/common';
import { TripCard } from '../../components/trips';
import { colors, spacing } from '../../utils/theme';
import { RootState, AppDispatch } from '../../store';
import { fetchTrips, deleteTrip, setCurrentTrip } from '../../store/slices/tripsSlice';
import { showToast } from '../../store/slices/uiSlice';
import { Trip } from '../../api/types';

interface TripsListScreenProps {
  navigation: any;
  route: any;
}

export const TripsListScreen: React.FC<TripsListScreenProps> = ({ 
  navigation, 
  route 
}) => {
  const dispatch = useDispatch<AppDispatch>();
  const { trips, isLoading, error, hasMore, page } = useSelector(
    (state: RootState) => state.trips
  );
  const { isRefreshing } = useSelector((state: RootState) => state.ui);
  
  const [refreshing, setRefreshing] = useState(false);
  const [loadingMore, setLoadingMore] = useState(false);

  // Load trips when screen comes into focus
  useFocusEffect(
    useCallback(() => {
      if (trips.length === 0) {
        dispatch(fetchTrips({ page: 1, refresh: true }));
      }
    }, [dispatch, trips.length])
  );

  // Handle pull-to-refresh
  const onRefresh = useCallback(async () => {
    setRefreshing(true);
    try {
      await dispatch(fetchTrips({ page: 1, refresh: true })).unwrap();
    } catch (error) {
      console.error('Failed to refresh trips:', error);
    } finally {
      setRefreshing(false);
    }
  }, [dispatch]);

  // Handle load more (pagination)
  const loadMore = useCallback(async () => {
    if (loadingMore || !hasMore || isLoading) return;
    
    setLoadingMore(true);
    try {
      await dispatch(fetchTrips({ page: page + 1 })).unwrap();
    } catch (error) {
      console.error('Failed to load more trips:', error);
    } finally {
      setLoadingMore(false);
    }
  }, [dispatch, loadingMore, hasMore, isLoading, page]);

  // Handle trip press (view details)
  const handleTripPress = useCallback((trip: Trip) => {
    dispatch(setCurrentTrip(trip));
    navigation.navigate('TripDetail', { tripId: trip.id });
  }, [dispatch, navigation]);

  // Handle edit trip
  const handleEditTrip = useCallback((trip: Trip) => {
    dispatch(setCurrentTrip(trip));
    navigation.navigate('EditTrip', { tripId: trip.id });
  }, [dispatch, navigation]);

  // Handle delete trip
  const handleDeleteTrip = useCallback((trip: Trip) => {
    Alert.alert(
      'Delete Trip',
      `Are you sure you want to delete your trip from ${trip.origin} to ${trip.destination}?`,
      [
        { text: 'Cancel', style: 'cancel' },
        {
          text: 'Delete',
          style: 'destructive',
          onPress: async () => {
            try {
              await dispatch(deleteTrip(trip.id)).unwrap();
              dispatch(showToast({
                type: 'success',
                title: 'Trip deleted',
                message: 'Your trip has been successfully deleted.',
              }));
            } catch (error: any) {
              dispatch(showToast({
                type: 'error',
                title: 'Delete failed',
                message: error.message || 'Failed to delete trip. Please try again.',
              }));
            }
          },
        },
      ]
    );
  }, [dispatch]);

  // Handle add new trip
  const handleAddTrip = useCallback(() => {
    navigation.navigate('AddTrip');
  }, [navigation]);

  // Render trip item
  const renderTrip = useCallback(({ item }: { item: Trip }) => (
    <TripCard
      trip={item}
      onPress={() => handleTripPress(item)}
      onEdit={() => handleEditTrip(item)}
      onDelete={() => handleDeleteTrip(item)}
      showActions={true}
    />
  ), [handleTripPress, handleEditTrip, handleDeleteTrip]);

  // Render loading skeleton
  const renderSkeleton = useCallback(() => (
    <TripCardSkeleton />
  ), []);

  // Render footer loading indicator
  const renderFooter = useCallback(() => {
    if (!loadingMore) return null;
    return (
      <View style={styles.footerLoading}>
        <Loading size="small" />
      </View>
    );
  }, [loadingMore]);

  // Show error state
  if (error && trips.length === 0) {
    return (
      <View style={styles.errorContainer}>
        <Ionicons name="alert-circle" size={48} color={colors.error} />
        <Text style={styles.errorText}>Failed to load trips</Text>
        <Button
          title="Try Again"
          onPress={() => dispatch(fetchTrips({ page: 1, refresh: true }))}
          variant="outline"
        />
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <FlatList
        data={trips}
        renderItem={renderTrip}
        keyExtractor={(item) => item.id.toString()}
        refreshControl={
          <RefreshControl
            refreshing={refreshing}
            onRefresh={onRefresh}
            tintColor={colors.primary}
            colors={[colors.primary]}
          />
        }
        onEndReached={loadMore}
        onEndReachedThreshold={0.3}
        ListFooterComponent={renderFooter}
        contentContainerStyle={[
          styles.listContent,
          ...(trips.length === 0 ? [styles.emptyContent] : [])
        ]}
        showsVerticalScrollIndicator={false}
        // Show loading skeletons while initial load
        ListEmptyComponent={
          isLoading ? (
            <View>
              {Array.from({ length: 5 }, (_, index) => (
                <TripCardSkeleton key={index} />
              ))}
            </View>
          ) : (
            <View style={styles.emptyState}>
              <Ionicons name="airplane" size={64} color={colors.text.muted} />
              <Text style={styles.emptyTitle}>No trips yet</Text>
              <Text style={styles.emptyMessage}>
                Start tracking your adventures by adding your first trip!
              </Text>
              <Button
                title="Add Your First Trip"
                onPress={handleAddTrip}
                variant="primary"
                style={styles.emptyButton}
              />
            </View>
          )
        }
      />

      {/* Floating Add Button */}
      <TouchableOpacity
        style={styles.fab}
        onPress={handleAddTrip}
        activeOpacity={0.8}
      >
        <Ionicons name="add" size={24} color={colors.text.primary} />
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.background,
  },

  listContent: {
    paddingBottom: 100, // Space for FAB
  },

  emptyContent: {
    flexGrow: 1,
  },

  emptyState: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: spacing.xl,
    paddingVertical: spacing['3xl'],
  },

  emptyTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    color: colors.text.primary,
    marginTop: spacing.lg,
    marginBottom: spacing.sm,
  },

  emptyMessage: {
    fontSize: 16,
    color: colors.text.secondary,
    textAlign: 'center',
    lineHeight: 24,
    marginBottom: spacing.xl,
  },

  emptyButton: {
    minWidth: 200,
  },

  errorContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: spacing.xl,
    backgroundColor: colors.background,
  },

  errorText: {
    fontSize: 18,
    color: colors.text.primary,
    marginVertical: spacing.lg,
    textAlign: 'center',
  },

  footerLoading: {
    paddingVertical: spacing.lg,
    alignItems: 'center',
  },

  fab: {
    position: 'absolute',
    bottom: spacing.xl,
    right: spacing.xl,
    width: 56,
    height: 56,
    borderRadius: 28,
    backgroundColor: colors.primary,
    alignItems: 'center',
    justifyContent: 'center',
    elevation: 8,
    shadowColor: colors.primary,
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.3,
    shadowRadius: 8,
  },
});

