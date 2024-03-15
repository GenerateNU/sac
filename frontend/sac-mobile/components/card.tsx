import { Text, TouchableOpacity, TouchableOpacityProps, View } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

import { cn } from '@/lib/utils';

const cardVariants = {
    variant: {
        default: ['bg-card-bg', 'text-white', 'justify-end', 'items-start']
    },
    size: {
        default: ['rounded-lg', 'min-w-96', 'p-4', 'shadow', 'w-80', 'h-48']
    }
};

const cardStyles = cva(
    ['flex'],
    {
        variants: cardVariants,
        defaultVariants: {
            variant: 'default',
            size: 'default'
        }
    }
);

export interface CardProps
    extends TouchableOpacityProps,
        VariantProps<typeof cardStyles> {}

const Card = ({ children, variant, size, ...props }: CardProps) => {
    return (
        <View
            className={cn(
                ...cardVariants.variant[variant ?? 'default'],
                ...cardVariants.size[size ?? 'default'],
                'text-center'
            )}
        >
            <View>
                <Text className={cn(`text-lg font-bold mb-2`, { fontFamily: 'OpenSans-SemiBold' })}>{"Orientation"}</Text>
                <Text className={cn(`text-gray-600`)}>{"08:00 - 09:30 PM"}</Text>
                <Text className={cn(`text-gray-600`)}>{"April 23, 2024"}</Text>
                <Text className={cn(`text-gray-600`)}>{"Ryder Hall, Room 294"}</Text>
            </View>
            
        </View>
    );
};

Card.displayName = 'card';

export { Card, cardVariants };
