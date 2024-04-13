import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import { Link } from 'expo-router';

import { Club } from '@/types/club';

export const ClubHomepageCard = ({ club }: { club: Club }) => {
    return (
        <Link
            key={club.id}
            href={{
                pathname: '/(app)/club/[id]',
                params: { id: club.id }
            }}
            asChild
        >
            <TouchableOpacity className="bg-gray-200 rounded-2xl flex-row p-[5%] my-[2.5%]">
                <View className="flex-row shrink items-center pt-[1%]">
                    <View className="bg-gray-300 rounded-xl w-16 h-16" />
                    <View className="pl-[4%] flex-col shrink">
                        <Text className="text-lg font-bold leading-5 mb-1">
                            {club.name}
                        </Text>
                        <Text numberOfLines={3} ellipsizeMode="tail">
                            {club.description}
                        </Text>
                    </View>
                </View>
            </TouchableOpacity>
        </Link>
    );
};
