import { Text, TouchableOpacity, TouchableOpacityProps, View } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

import { cn } from '@/lib/utils';
import { Button } from './button';

const cardVariants = {
    variant: {
        default: ['bg-card-bg', 'text-white', 'justify-end', 'items-start', 'mr-3']
    },
    size: {
        default: ['rounded-lg', 'min-w-96', 'p-4', 'shadow', 'w-80', 'h-40']
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
    VariantProps<typeof cardStyles> {
        title: string
     }

const Card = ({ children, variant, size, title, ...props }: CardProps) => {
    return (
        <View
            className={cn(
                ...cardVariants.variant[variant ?? 'default'],
                ...cardVariants.size[size ?? 'default'],
                'text-center'
            )}
        >
            <View className='flex-row px-3 items-end space-x-10'>
                <View>
                    <Text className={cn(`text-lg font-bold mb-2`, { fontFamily: 'OpenSans-SemiBold' })}>{title}</Text>
                    {children}
                </View>
                <Button className='bg-[#747474] rounded-lg'>Register</Button>
                
            </View>

        </View>
    );
};

Card.displayName = 'card';

export { Card, cardVariants };
