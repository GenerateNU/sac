import React from 'react';
import { Text, View } from 'react-native';
import { FlatList } from 'react-native-gesture-handler';

import { useEvents } from '@/hooks/use-event';
import { isToday } from '@/lib/utils';
import { Event } from '@/types/event';

import { EventHomePageCard } from './event-homepage-card';

const UpcomingEvents = () => {
    const { data: events, isLoading, error } = useEventsQuery();

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error</Text>;
    }

    if (!events) {
        return <Text>No events found</Text>;
    }

    const sortedEvents = [...events].sort(
        (a, b) =>
            new Date(a.start_time).getTime() - new Date(b.start_time).getTime()
    );

    const renderEvents = ({ item }: { item: Event }) => {
        if (!isToday(item.start_time)) {
            return <EventHomePageCard event={item} />;
        }

        return null;
    };

    return (
        <View>
            <FlatList
                data={sortedEvents}
                renderItem={renderEvents}
                keyExtractor={(event) => event.id}
                horizontal
                showsHorizontalScrollIndicator={false}
            />
        </View>
    );
};

export { UpcomingEvents };
