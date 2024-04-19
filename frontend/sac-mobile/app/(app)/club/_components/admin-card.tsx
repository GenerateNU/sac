import React from 'react';
import { Image, Text, View } from 'react-native';

type ClubAdminCardProps = {
    photo?: string;
    name?: string;
    title?: string;
};

const ClubAdminCard = ({ photo, name, title }: ClubAdminCardProps) => {
    return (
        <View className="flex-col items-start gap-2 mr-2">
            {photo && (
                <Image
                    source={{ uri: photo }}
                    className="rounded-lg aspect-square"
                    resizeMode="cover"
                    width={145}
                    height={132}
                />
            )}
            <View className="flex-col items-start gap-1">
                {name && (
                    <Text className="text-sm font-normal text-zinc-600">
                        {name}
                    </Text>
                )}
                {title && (
                    <Text className="font-medium tracking-tight text-black text-md">
                        {title}
                    </Text>
                )}
            </View>
        </View>
    );
};

export { ClubAdminCard };
