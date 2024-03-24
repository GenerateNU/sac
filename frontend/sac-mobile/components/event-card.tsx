import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import Location from '@/assets/icons/location';
import Time from '@/assets/icons/time';
import { Event } from '@/types/item';

const EventCard = ({ event }: { event: Event }) => {
    return (
        <TouchableOpacity className="bg-gray-200 rounded-2xl my-[2.5%]">
            <View className="px-[5%] pt-[5%] pb-[4%]">
                <View className="flex-row items-center">
                    <View className="bg-gray-300 h-16 w-16 rounded-xl"></View>
                    <View className="pl-[5%] flex-col shrink pr-[3%]">
                        <Text className="text-lg font-bold leading-5">
                            {event.clubName}
                        </Text>
                        <Text className="text-sm font-semibold leading-5">
                            {event.eventName}
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
                                {event.time}
                            </Text>
                        </View>
                    </View>
                </View>
                <Text
                    className="pt-[2%]"
                    numberOfLines={4}
                    ellipsizeMode="tail"
                >
                    {event.description}
                </Text>
            </View>
            <View className="bg-gray-300 h-60 rounded-2xl"></View>
        </TouchableOpacity>
    );
};

export default EventCard;
