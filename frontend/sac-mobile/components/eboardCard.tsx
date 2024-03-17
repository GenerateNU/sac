import React from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Image } from 'react-native';
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

export interface EBoardCardProps extends TouchableOpacityProps, VariantProps<typeof cardStyles> {
    photo?: string;
    name?: string;
    title?: string;
}

const EBoardCard = ({ photo, name, title, variant, size, ...props }: EBoardCardProps) => {
    return (
        <TouchableOpacity {...props}>
            <View className="bg-white p-4 w-full h-48 text-center items-center pb-12" >
                {photo ? (
                    <Image source={{ uri: photo }} className="aspect-square rounded-lg w-full h-full mb-2" />
                ) : (
                    <View className="aspect-square rounded-lg w-full h-full bg-card-bg mb-2" />
                )}
                {name && <Text className="text-sm text-gray-500 mb-2">{name}</Text>}
                {title && <Text className="text-sm">{title}</Text>}
            </View>
        </TouchableOpacity>
    );
};


EBoardCard.displayName = 'eboardCard';

export { EBoardCard, cardVariants };
