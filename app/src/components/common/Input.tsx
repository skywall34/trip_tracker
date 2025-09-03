import React, { useState, forwardRef } from 'react';
import {
  TextInput,
  View,
  Text,
  StyleSheet,
  TextInputProps,
  ViewStyle,
  TouchableOpacity,
} from 'react-native';
import { colors, typography, borderRadius, spacing } from '../../utils/theme';

interface InputProps extends Omit<TextInputProps, 'style'> {
  label?: string;
  error?: string;
  helperText?: string;
  leftIcon?: React.ReactNode;
  rightIcon?: React.ReactNode;
  onRightIconPress?: () => void;
  variant?: 'default' | 'filled' | 'outline';
  size?: 'sm' | 'md' | 'lg';
  containerStyle?: ViewStyle;
  inputStyle?: ViewStyle;
}

export const Input = forwardRef<TextInput, InputProps>(
  (
    {
      label,
      error,
      helperText,
      leftIcon,
      rightIcon,
      onRightIconPress,
      variant = 'default',
      size = 'md',
      containerStyle,
      inputStyle,
      ...textInputProps
    },
    ref
  ) => {
    const [isFocused, setIsFocused] = useState(false);

    const containerStyles = [
      styles.container,
      containerStyle,
    ];

    const inputContainerStyles = [
      styles.inputContainer,
      styles[variant],
      styles[size],
      isFocused && styles.focused,
      error && styles.error,
    ];

    const textInputStyles = [
      styles.input,
      styles[`${size}Input`],
      inputStyle,
    ];

    return (
      <View style={containerStyles}>
        {label && (
          <Text style={styles.label}>{label}</Text>
        )}
        
        <View style={inputContainerStyles}>
          {leftIcon && (
            <View style={styles.leftIcon}>
              {leftIcon}
            </View>
          )}
          
          <TextInput
            ref={ref}
            style={textInputStyles}
            placeholderTextColor={colors.text.muted}
            selectionColor={colors.primary}
            onFocus={() => setIsFocused(true)}
            onBlur={() => setIsFocused(false)}
            {...textInputProps}
          />
          
          {rightIcon && (
            <TouchableOpacity
              style={styles.rightIcon}
              onPress={onRightIconPress}
              disabled={!onRightIconPress}
            >
              {rightIcon}
            </TouchableOpacity>
          )}
        </View>
        
        {(error || helperText) && (
          <Text style={[styles.helperText, error && styles.errorText]}>
            {error || helperText}
          </Text>
        )}
      </View>
    );
  }
);

Input.displayName = 'Input';

const styles = StyleSheet.create({
  container: {
    marginVertical: spacing.sm,
  },

  label: {
    fontSize: typography.fontSize.sm,
    fontFamily: typography.fontFamily.semiBold,
    color: colors.text.tertiary,
    marginBottom: spacing.sm,
  },

  inputContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    borderRadius: borderRadius.lg,
    borderWidth: 1,
  },

  // Variants
  default: {
    backgroundColor: colors.elevated,
    borderColor: colors.border,
  },

  filled: {
    backgroundColor: colors.surface,
    borderColor: 'transparent',
  },

  outline: {
    backgroundColor: 'transparent',
    borderColor: colors.border,
  },

  // Sizes
  sm: {
    minHeight: 36,
    paddingHorizontal: spacing.md,
  },

  md: {
    minHeight: 44,
    paddingHorizontal: spacing.md,
  },

  lg: {
    minHeight: 52,
    paddingHorizontal: spacing.lg,
  },

  // States
  focused: {
    borderColor: colors.primary,
    borderWidth: 2,
    // Add glow effect
    shadowColor: colors.primary,
    shadowOffset: { width: 0, height: 0 },
    shadowOpacity: 0.3,
    shadowRadius: 4,
    elevation: 4,
  },

  error: {
    borderColor: colors.error,
    borderWidth: 2,
  },

  // Input text
  input: {
    flex: 1,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.secondary,
    paddingVertical: 0, // Remove default padding
  },

  smInput: {
    fontSize: typography.fontSize.sm,
    lineHeight: typography.lineHeight.sm,
  },

  mdInput: {
    fontSize: typography.fontSize.base,
    lineHeight: typography.lineHeight.base,
  },

  lgInput: {
    fontSize: typography.fontSize.lg,
    lineHeight: typography.lineHeight.lg,
  },

  // Icons
  leftIcon: {
    marginRight: spacing.sm,
    alignItems: 'center',
    justifyContent: 'center',
  },

  rightIcon: {
    marginLeft: spacing.sm,
    alignItems: 'center',
    justifyContent: 'center',
    padding: spacing.xs,
  },

  // Helper text
  helperText: {
    fontSize: typography.fontSize.xs,
    fontFamily: typography.fontFamily.regular,
    color: colors.text.muted,
    marginTop: spacing.xs,
    marginLeft: spacing.xs,
  },

  errorText: {
    color: colors.error,
  },
});

export default Input;