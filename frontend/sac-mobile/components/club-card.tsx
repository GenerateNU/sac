import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import { Club } from '@/types/item';

const ClubCard = ({ club }: { club: Club }) => {
    return (
        <TouchableOpacity className="bg-gray-200 rounded-2xl flex-row p-[5%] my-[2.5%]">
            <View className="flex-row shrink items-center">
                <View className="bg-gray-300 rounded-xl w-16 h-16"></View>
                <View className="pl-[4%] flex-col shrink">
                    <Text className="text-lg font-bold">{club.name}</Text>
                    <Text numberOfLines={3} ellipsizeMode="tail">
                        {club.description}
                    </Text>
                </View>
            </View>
        </TouchableOpacity>
    );
};

export default ClubCard;
