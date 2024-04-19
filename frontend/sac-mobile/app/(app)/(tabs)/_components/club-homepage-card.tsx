import React from 'react';
import { Pressable, Text, View } from 'react-native';

import { router } from 'expo-router';

import { Club } from '@/types/club';
import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';

export const ClubHomePageCard = ({ club }: { club: Club }) => {
    return (
        <Pressable
            className="bg-gray-200 rounded-2xl flex-row p-[5%] my-[2.5%]"
            onPress={() => router.push(`/club/${club.id}`)}
        >
            <View className="flex-row shrink items-center pt-[1%]">
                <EventDemoDay
                    width={64}
                    height={64}
                    className="rounded-xl"
                />
                <View className="pl-[4%] flex-col shrink">
                    <Text className="mb-1 text-lg font-bold leading-5">
                        {club.name}
                    </Text>
                    <Text numberOfLines={3} ellipsizeMode="tail">
                        {club.description}
                    </Text>
                </View>
            </View>
        </Pressable>
    );
};
