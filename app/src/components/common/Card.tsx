import React from 'react';
import {
  View,
  StyleSheet,
  ViewStyle,
  TouchableOpacity,
  TouchableOpacityProps,
} from 'react-native';
import { colors, borderRadius, spacing, shadows } from '../../utils/theme';

interface CardProps {
  children: React.ReactNode;
  style?: ViewStyle;
  variant?: 'default' | 'elevated' | 'glass' | 'outline' | 'subtle';
  padding?: 'none' | 'sm' | 'md' | 'lg';
  onPress?: TouchableOpacityProps['onPress'];
  disabled?: boolean;
}

export const Card: React.FC<CardProps> = ({
  children,
  style,
  variant = 'default',
  padding = 'md',
  onPress,
  disabled = false,
}) => {
  const paddingKey = `padding${padding.charAt(0).toUpperCase() + padding.slice(1)}` as keyof typeof styles;
  const cardStyle = [
    styles.card,
    styles[variant],
    styles[paddingKey],
    style,
  ];

  if (onPress) {
    return (
      <TouchableOpacity
        style={cardStyle}
        onPress={onPress}
        disabled={disabled}
        activeOpacity={0.8}
      >
        {children}
      </TouchableOpacity>
    );
  }

  return <View style={cardStyle}>{children}</View>;
};

const styles = StyleSheet.create({
  card: {
    borderRadius: borderRadius.lg,
    overflow: 'hidden',
  },

  // Variants
  default: {
    backgroundColor: colors.surface,
    borderWidth: 1,
    borderColor: colors.border,
    ...shadows.sm,
  },

  elevated: {
    backgroundColor: colors.elevated,
    borderWidth: 1,
    borderColor: colors.borderLight,
    ...shadows.md,
  },

  glass: {
    backgroundColor: colors.glass,
    borderWidth: 1,
    borderColor: colors.borderLight,
    ...shadows.lg,
    // Blur effect would need a blur view library
  },

  outline: {
    backgroundColor: 'transparent',
    borderWidth: 1,
    borderColor: colors.border,
  },

  subtle: {
    backgroundColor: colors.surface + '50', // Semi-transparent surface
    borderWidth: 1,
    borderColor: colors.border + '30', // Subtle border
    ...shadows.sm,
  },

  // Padding variants
  paddingNone: {
    padding: 0,
  },

  paddingSm: {
    padding: spacing.sm,
  },

  paddingMd: {
    padding: spacing.md,
  },

  paddingLg: {
    padding: spacing.lg,
  },
});

export default Card;