import React from 'react';
import { ActivityIndicator, Pressable, Text, View } from 'react-native';

import { router } from 'expo-router';

import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';
import Pin from '@/assets/images/svg/pin.svg';
import Time from '@/assets/images/svg/time.svg';
import { useClub } from '@/hooks/use-club';
import { Event } from '@/types/event';

const EventHomePageCard = ({ event }: { event: Event }) => {
    const { data: host, isLoading, error } = useClub(event.host);

    if (isLoading) {
        // skeleton loading
        <ActivityIndicator size={'large'} />;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <Pressable
            className="bg-gray-200 rounded-2xl my-[2.5%]"
            onPress={() => router.push(`/event/${event.id}`)}
        >
            <View>
                <View className="px-[5%] pt-[5%] pb-[4%]">
                    <View className="flex-row items-center">
                        <EventDemoDay
                            width={64}
                            height={64}
                            className="rounded-xl"
                        />
                        <View className="pl-[5%] flex-col shrink pr-[3%]">
                            <Text className="text-lg font-bold leading-5">
                                {host?.name ?? 'No Hosts'}
                            </Text>
                            <Text className="text-sm font-semibold leading-5">
                                {event.name}
                            </Text>
                            <View className="flex-row items-center">
                                <Pin />
                                <Text className="text-sm pl-[2%] leading-5">
                                    {event.location}
                                </Text>
                            </View>
                            <View className="flex-row items-center">
                                <Time />
                                <Text className="text-sm pl-[2%]">
                                    {new Date(
                                        event.start_time
                                    ).toLocaleDateString('en-US', {
                                        month: 'long',
                                        day: 'numeric',
                                        hour: 'numeric',
                                        minute: 'numeric'
                                    })}
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
