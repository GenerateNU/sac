import React from 'react';
import { Image, Text, View } from 'react-native';

type ClubAdminCardProps = {
    photo?: string;
    name?: string;
    title?: string;
};

const ClubAdminCard = ({ photo, name, title }: ClubAdminCardProps) => {
    return (
        <View className="flex-col items-start gap-2 mr-2 relative overflow-hidden">
            {photo && (
                <Image
                    source={{ uri: photo }}
                    className="rounded-lg"
                    resizeMode="cover"
                    width={150}
                    height={160}
                />
            )}
            <View className="flex-col items-start gap-1 absolute bottom-0 left-0 p-2 w-full">
                {title && (
                    <Text className="font-light tracking-tight text-white text-md">
                        {title}
                    </Text>
                )}
                {name && (
                    <Text className="text-base font-light text-white">
                        {name}
                    </Text>
                )}
            </View>
        </View>
    );
};

export { ClubAdminCard };
