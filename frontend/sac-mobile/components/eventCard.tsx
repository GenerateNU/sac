import { Pressable, Text, TouchableOpacity, TouchableOpacityProps, View } from 'react-native';

import { cn } from '@/lib/utils';
import { Button } from './button';
import MiniLocationPin from '@/assets/images/svg/mini-location-pin.svg';
import Expand from '@/assets/images/svg/expand.svg';
import React from 'react';
import { router } from 'expo-router';

export interface EventCardProps
    {
        title: string
        time: string
        date: string
        location: string
     }

const EventCard = ({title, time, date, location }: EventCardProps) => {
    return (
        <Pressable
            onPress={() => router.push(`/event/eea48981-8c89-4f3a-8b04-c9ec60671640`)}
            className='flex my-3 bg-white border border-gray-300 shadow-sm rounded-2xl'>
            <View className='flex p-4 gap-y-1'>
                <View className='flex-row items-center justify-between gap-x-2'>
                    <Text className='font-normal text-neutral-400 text-md'>{time.toLocaleUpperCase()}</Text>
                    <Expand width={12} height={12} />
                </View>
                <Text className='text-lg font-semibold text-black'>{title}</Text>
                <View className='flex-row items-center gap-x-2'>
                    <MiniLocationPin width={12} height={12} />
                    <Text className='font-normal text-neutral-600 text-md'>{location}</Text>
                </View>
            </View>

        </Pressable>
    );
};

EventCard.displayName = 'Event Card';

export { EventCard };
