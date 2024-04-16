import React from 'react';
import { Pressable, Text, View } from 'react-native';

import { router } from 'expo-router';

import Location from '@/assets/icons/location';
import Time from '@/assets/icons/time';
import { useEventHosts } from '@/hooks/use-event';
import { Event } from '@/types/event';

const EventHomePageCard = ({ event }: { event: Event }) => {
    const getHost = () => {
        const { data: hosts } = useEventHosts(event.id);
        if (hosts !== undefined && hosts.length > 0) {
            return hosts[0].name;
        }
        return null;
    };

    return (
        <Pressable
            className="bg-gray-200 rounded-2xl my-[2.5%]"
            onPress={() => router.push(`/event/${event.id}`)}
        >
            <View>
                <View className="px-[5%] pt-[5%] pb-[4%]">
                    <View className="flex-row items-center">
                        <View className="w-16 h-16 bg-gray-300 rounded-xl" />
                        <View className="pl-[5%] flex-col shrink pr-[3%]">
                            <Text className="text-lg font-bold leading-5">
                                {getHost()}
                            </Text>
                            <Text className="text-sm font-semibold leading-5">
                                {event.name}
                            </Text>
                            <View className="flex-row items-center">
                                <Location />
                                <Text className="text-sm pl-[2%] leading-5">
                                    {event.location}
                                </Text>
                            </View>
                            <View className="flex-row items-center">
                                <Time />
                                <Text className="text-sm pl-[2%]">
                                    {String(event.start_time)}
                                </Text>
                            </View>
                        </View>
                    </View>
                    <Text
                        className="pt-[2%]"
                        numberOfLines={4}
                        ellipsizeMode="tail"
                    >
                        {event.preview}
                    </Text>
                </View>
                <View className="bg-gray-300 h-60 rounded-2xl" />
            </View>
        </Pressable>
    );
};

export { EventHomePageCard };
