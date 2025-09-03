import React from 'react';
import {
  View,
  ActivityIndicator,
  StyleSheet,
  Text,
  ViewStyle,
  DimensionValue,
} from 'react-native';
import { colors, typography, spacing } from '../../utils/theme';

interface LoadingProps {
  size?: 'small' | 'large';
  color?: string;
  message?: string;
  fullScreen?: boolean;
  style?: ViewStyle;
}

export const Loading: React.FC<LoadingProps> = ({
  size = 'large',
  color = colors.primary,
  message,
  fullScreen = false,
  style,
}) => {
  const containerStyle = [
    styles.container,
    fullScreen && styles.fullScreen,
    style,
  ];

  return (
    <View style={containerStyle}>
      <ActivityIndicator size={size} color={color} />
      {message && (
        <Text style={styles.message}>{message}</Text>
      )}
    </View>
  );
};

// Skeleton loading component for list items
interface SkeletonProps {
  height?: number;
  width?: DimensionValue;
  borderRadius?: number;
  style?: ViewStyle;
}

export const Skeleton: React.FC<SkeletonProps> = ({
  height = 20,
  width = '100%',
  borderRadius = 4,
  style,
}) => {
  const skeletonStyle = [
    styles.skeleton,
    { height, width, borderRadius },
    style,
  ];

  return <View style={skeletonStyle} />;
};

// Card skeleton for trip cards
export const TripCardSkeleton: React.FC = () => (
  <View style={styles.tripCardSkeleton}>
    <View style={styles.skeletonRow}>
      <Skeleton width={120} height={16} />
      <Skeleton width={80} height={16} />
    </View>
    <Skeleton width={200} height={14} style={{ marginTop: spacing.sm }} />
    <Skeleton width={150} height={14} style={{ marginTop: spacing.xs }} />
    <View style={[styles.skeletonRow, { marginTop: spacing.md }]}>
      <Skeleton width={60} height={12} />
      <Skeleton width={60} height={12} />
    </View>
  </View>
);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: spacing.lg,
  },

  fullScreen: {
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: colors.background,
    zIndex: 1000,
  },

  message: {
    marginTop: spacing.md,
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    textAlign: 'center',
  },

  skeleton: {
    backgroundColor: colors.border,
    opacity: 0.3,
  },

  tripCardSkeleton: {
    backgroundColor: colors.surface,
    borderRadius: 12,
    padding: spacing.md,
    margin: spacing.sm,
    borderWidth: 1,
    borderColor: colors.border,
  },

  skeletonRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
});

export default Loading;