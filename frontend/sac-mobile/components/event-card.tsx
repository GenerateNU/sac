import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import { Event } from '@/types/item';

const EventCard = ({ event }: { event: Event }) => {
    return (
        <TouchableOpacity className="bg-gray-200 rounded-2xl my-[2.5%]">
            <View className="px-[5%] pt-[5%] pb-[4%]">
                <View className="flex-row items-center">
                    <View className="bg-gray-300 h-16 w-16 rounded-xl"></View>
                    <View className="pl-[5%] flex-col shrink">
                        <Text className="text-lg font-bold leading-5">
                            {event.clubName}
                        </Text>
                        <Text className="text-sm font-bold">
                            {event.eventName}
                        </Text>
                        <Text className="text-xs">{event.location}</Text>
                        <Text className="text-xs">{event.time}</Text>
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
