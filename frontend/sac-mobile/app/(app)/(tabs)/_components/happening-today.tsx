import React from 'react';
import { Text, View } from 'react-native';
import { FlatList } from 'react-native-gesture-handler';

import { useEvents } from '@/hooks/use-event';
import { isToday } from '@/lib/utils';
import { Event } from '@/types/event';

import { EventHomePageCardHorizontal } from './event-homepage-card';

const HappeningToday = () => {
    const { data: events, isLoading, error } = useEvents();

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    const renderEvents = ({ item }: { item: Event }) => {
        if (isToday(item.start_time)) {
            return (
                <View className='flex-grow mr-3'>
                    <EventHomePageCardHorizontal event={item} />
                </View>
            );
        }

        return null;
    };

    return (
        <View className='w-full'>
            <FlatList
                data={events}
                renderItem={renderEvents}
                keyExtractor={(event) => event.id}
                horizontal
                showsHorizontalScrollIndicator={false}
            />
        </View>
    );
};

export { HappeningToday };
