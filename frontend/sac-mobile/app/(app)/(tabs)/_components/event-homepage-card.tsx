import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import { Link } from 'expo-router';

import Location from '@/assets/icons/location';
import Time from '@/assets/icons/time';
import { useEventHosts } from '@/hooks/use-event';
import { Event } from '@/types/event';

export const EventHomePageCard = ({ event }: { event: Event }) => {
    const getHost = () => {
        const { data: hosts } = useEventHosts(event.id);
        if (hosts !== undefined && hosts.length > 0) {
            return hosts[0].name;
        }
        return null;
    };

    return (
        <Link
            key={event.id}
            href={{
                pathname: '/(app)/event/[id]',
                params: { id: event.id }
            }}
            asChild
        >
            <TouchableOpacity className="bg-gray-200 rounded-2xl my-[2.5%]">
                <View>
                    <View className="px-[5%] pt-[5%] pb-[4%]">
                        <View className="flex-row items-center">
                            <View className="bg-gray-300 h-16 w-16 rounded-xl" />
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
            </TouchableOpacity>
        </Link>
    );
};
