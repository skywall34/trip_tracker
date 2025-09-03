import React from 'react';
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { Card } from '../common';
import { colors, typography, spacing } from '../../utils/theme';
import { Trip } from '../../api/types';

interface TripCardProps {
  trip: Trip;
  onPress?: () => void;
  onEdit?: () => void;
  onDelete?: () => void;
  showActions?: boolean;
}

export const TripCard: React.FC<TripCardProps> = ({
  trip,
  onPress,
  onEdit,
  onDelete,
  showActions = false,
}) => {
  const departureDate = new Date(trip.departure_time * 1000);
  const arrivalDate = new Date(trip.arrival_time * 1000);
  
  const formatDate = (date: Date) => {
    return date.toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    });
  };

  const formatTime = (date: Date) => {
    return date.toLocaleTimeString('en-US', {
      hour: 'numeric',
      minute: '2-digit',
      hour12: true,
    });
  };

  const isPastTrip = departureDate < new Date();

  return (
    <Card
      variant="elevated"
      padding="none"
      onPress={onPress}
      style={StyleSheet.flatten([styles.card, isPastTrip && styles.pastTrip])}
    >
      {/* Header with route */}
      <View style={styles.header}>
        <View style={styles.route}>
          <Text style={styles.airport}>{trip.origin}</Text>
          <Ionicons name="airplane" size={16} color={colors.primary} />
          <Text style={styles.airport}>{trip.destination}</Text>
        </View>
        
        {showActions && (
          <View style={styles.actions}>
            {onEdit && (
              <TouchableOpacity onPress={onEdit} style={styles.actionButton}>
                <Ionicons name="pencil" size={16} color={colors.text.secondary} />
              </TouchableOpacity>
            )}
            {onDelete && (
              <TouchableOpacity onPress={onDelete} style={styles.actionButton}>
                <Ionicons name="trash" size={16} color={colors.error} />
              </TouchableOpacity>
            )}
          </View>
        )}
      </View>

      {/* Flight details */}
      <View style={styles.content}>
        <View style={styles.flightInfo}>
          <Text style={styles.airline}>
            {trip.airline} {trip.flight_number}
          </Text>
          {trip.gate && (
            <Text style={styles.detail}>Gate {trip.gate}</Text>
          )}
          {trip.terminal && (
            <Text style={styles.detail}>Terminal {trip.terminal}</Text>
          )}
        </View>

        {/* Times */}
        <View style={styles.timeContainer}>
          <View style={styles.timeSection}>
            <Text style={styles.timeLabel}>Departure</Text>
            <Text style={styles.time}>{formatTime(departureDate)}</Text>
            <Text style={styles.date}>{formatDate(departureDate)}</Text>
          </View>
          
          <View style={styles.flightDuration}>
            <Ionicons name="time" size={14} color={colors.text.muted} />
          </View>
          
          <View style={styles.timeSection}>
            <Text style={styles.timeLabel}>Arrival</Text>
            <Text style={styles.time}>{formatTime(arrivalDate)}</Text>
            <Text style={styles.date}>{formatDate(arrivalDate)}</Text>
          </View>
        </View>

        {/* Additional info */}
        <View style={styles.footer}>
          {trip.reservation && (
            <Text style={styles.reservation}>
              Reservation: {trip.reservation}
            </Text>
          )}
          
          <View style={styles.statusContainer}>
            <View style={[styles.statusDot, isPastTrip ? styles.completedDot : styles.upcomingDot]} />
            <Text style={styles.statusText}>
              {isPastTrip ? 'Completed' : 'Upcoming'}
            </Text>
          </View>
        </View>
      </View>
    </Card>
  );
};

const styles = StyleSheet.create({
  card: {
    marginVertical: spacing.sm,
    marginHorizontal: spacing.md,
  },

  pastTrip: {
    opacity: 0.8,
  },

  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: spacing.md,
    paddingBottom: spacing.sm,
    borderBottomWidth: 1,
    borderBottomColor: colors.border,
  },

  route: {
    flexDirection: 'row',
    alignItems: 'center',
    flex: 1,
    gap: spacing.md,
  },

  airport: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.bold,
    color: colors.text.primary,
  },

  actions: {
    flexDirection: 'row',
    gap: spacing.sm,
  },

  actionButton: {
    padding: spacing.sm,
  },

  content: {
    padding: spacing.md,
    gap: spacing.md,
  },

  flightInfo: {
    gap: spacing.xs,
  },

  airline: {
    fontSize: typography.fontSize.base,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.primary,
  },

  detail: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
  },

  timeContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },

  timeSection: {
    alignItems: 'center',
    flex: 1,
  },

  timeLabel: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.medium,
    color: colors.text.muted,
    textTransform: 'uppercase',
    marginBottom: spacing.xs,
  },

  time: {
    fontSize: typography.fontSize.lg,
    fontFamily: typography.fontFamily.bold,
    color: colors.text.primary,
  },

  date: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    marginTop: spacing.xs,
  },

  flightDuration: {
    alignItems: 'center',
    justifyContent: 'center',
    paddingHorizontal: spacing.md,
  },

  footer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },

  reservation: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.muted,
  },

  statusContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: spacing.xs,
  },

  statusDot: {
    width: 8,
    height: 8,
    borderRadius: 4,
  },

  completedDot: {
    backgroundColor: colors.success,
  },

  upcomingDot: {
    backgroundColor: colors.primary,
  },

  statusText: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.medium,
    color: colors.text.secondary,
    textTransform: 'uppercase',
  },
});

export default TripCard;