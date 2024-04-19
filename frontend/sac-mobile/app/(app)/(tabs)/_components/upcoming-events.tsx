import React from 'react';
import { Text, View } from 'react-native';
import { FlatList } from 'react-native-gesture-handler';

import { useEvents } from '@/hooks/use-event';
import { isToday } from '@/lib/utils';
import { Event } from '@/types/event';

import { EventHomePageCard } from './event-homepage-card';

const UpcomingEvents = () => {
    const { data: events, isLoading, error } = useEvents();

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    const renderEvents = ({ item }: { item: Event }) => {
        if (!isToday(item.start_time)) {
            return <EventHomePageCard event={item} />;
        }

        return null;
    };

    return (
        <View style={{ paddingLeft: 20 }}>
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

export { UpcomingEvents };
