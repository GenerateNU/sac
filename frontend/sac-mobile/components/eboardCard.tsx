import React from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Image } from 'react-native';
import { VariantProps, cva } from 'class-variance-authority';
import { cn } from '@/lib/utils';

export interface EBoardCardProps extends TouchableOpacityProps{
    photo?: string;
    name?: string;
    title?: string;
}

const EBoardCard = ({ photo, name, title, ...props }: EBoardCardProps) => {
    return (
        <TouchableOpacity {...props}>
            <View className="bg-white pt-[3%] pr-[5%] w-full h-48 text-center items-center pb-12" >
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

export { EBoardCard };
