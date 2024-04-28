import React from 'react';
import { ActivityIndicator, Text, View } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

import { cn } from '@/lib/utils';

const spinnerVariants = {
    size: {
        default: 'w-6 h-6',
        large: 'w-8 h-8',
        small: 'w-4 h-4'
    },
    color: {
        default: 'text-gray-500',
        primary: 'text-blue-500',
        secondary: 'text-red-500'
    }
};

const spinnerStyles = cva(['items-center', 'justify-center'], {
    variants: spinnerVariants,
    defaultVariants: {
        size: 'default',
        color: 'default'
    }
});

export interface SpinnerProps extends VariantProps<typeof spinnerStyles> {
    text?: string;
}

const Spinner = ({ size, color, text }: SpinnerProps) => {
    return (
        <View className={cn(spinnerStyles({ size, color }))}>
            <ActivityIndicator />
            {text && <Text>{text}</Text>}
        </View>
    );
};

Spinner.displayName = 'Spinner';

export { Spinner, spinnerVariants };
