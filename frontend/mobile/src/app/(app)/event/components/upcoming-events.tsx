import { Pressable } from 'react-native';
import { FlatList } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import { Event } from '@generatesac/lib';

import { Box, Text } from '@/src/app/(design-system)';
import { EventCard } from '@/src/app/(design-system)/components/EventCard';

interface UpcomingEventsProps {
    events: Event[];
}

export const UpcomingEvent: React.FC<UpcomingEventsProps> = ({ events }) => {
    const renderEventCard = ({ item }: { item: Event }) => {
        return (
            <Pressable onPress={() => router.push(`/event/${item.id}`)}>
                <EventCard
                    clubId={'1'}
                    variant="small"
                    event={item.name}
                    club={item.host}
                    eventId={item.id}
                    startTime={new Date(item.start_time)}
                    endTime={new Date(item.end_time)}
                    image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLF3ord7lnV_5Je-pC2AUgUiesHNPcZlpI7A&s"
                />
            </Pressable>
        );
    };
    return (
        <Box flexDirection="column" alignItems="flex-start">
            <Text paddingLeft="xl" paddingBottom="m" variant="subheader-1">
                Upcoming Events
            </Text>
            <FlatList
                showsHorizontalScrollIndicator={false}
                horizontal
                contentContainerStyle={{
                    gap: 15,
                    paddingHorizontal: 32,
                    paddingBottom: 40
                }}
                data={events}
                renderItem={renderEventCard}
                keyExtractor={(item) => item.id}
            />
        </Box>
    );
};
