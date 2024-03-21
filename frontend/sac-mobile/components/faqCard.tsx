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

export interface FaqCardProps extends TouchableOpacityProps, VariantProps<typeof cardStyles> {
    question?: string;
    answer?: string;
}

const FaqCard = ({ question, answer, variant, size, ...props }: FaqCardProps) => {
    return (
        <TouchableOpacity {...props}>
            <View className="bg-white p-4 rounded-lg w-60 h-48 pb-12 border">
                {question && <Text className="text-sm mb-2 font-semibold">{question}</Text>}
                {answer && <Text className="text-sm">{answer}</Text>}
            </View>
        </TouchableOpacity>
    );
};


FaqCard.displayName = 'faqCard';

export { FaqCard, cardVariants };
