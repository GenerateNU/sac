import React from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Image } from 'react-native';
import { VariantProps, cva } from 'class-variance-authority';
import { cn } from '@/lib/utils';
import { useAssets } from 'expo-asset';

export type ClubAdminCardProps = {
    photo?: string;
    name?: string;
    title?: string;
}

const ClubAdminCard = ({ photo, name, title }: ClubAdminCardProps) => {
    const [asset, error] = useAssets(
        require('@/assets/images/guy-1.jpeg')
    );

    if (error) {
        console.error(error);
    }

    return (
        <View className="flex-col items-start gap-2 mr-2">
            {asset && (
                <Image
                    source={{ uri: asset[0].uri }}
                    className="rounded-lg aspect-square"
                    resizeMode="cover"
                    width={145}
                    height={132}
                />
            )}
            <View className='flex-col items-start gap-1'>
                {name && <Text className="text-sm font-normal text-zinc-600">{name}</Text>}
                {title && <Text className="font-medium tracking-tight text-black text-md">{title}</Text>}
            </View>
        </View>
    );
};


export { ClubAdminCard };
