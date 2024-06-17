import { FlatList, Pressable } from 'react-native';

import { router } from 'expo-router';

import { Event } from '@generatesac/lib';

import { EventCard } from './EventCard';

interface EventCardListProps {
    events: Event[];
    club?: string;
}

export const EventCardList: React.FC<EventCardListProps> = ({
    events,
    club
}) => {
    const renderEventCard = ({ item }: { item: Event }) => {
        return (
            <Pressable onPress={() => router.push(`/event/${item.id}`)}>
                <EventCard
                    variant="small"
                    event={item.name}
                    club={club || item.host}
                    eventId={item.id}
                    startTime={item.start_time}
                    endTime={item.end_time}
                    image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLF3ord7lnV_5Je-pC2AUgUiesHNPcZlpI7A&s"
                    logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                />
            </Pressable>
        );
    };
    return (
        <FlatList
            showsHorizontalScrollIndicator={false}
            horizontal
            contentContainerStyle={{
                gap: 15,
                paddingBottom: 40
            }}
            data={events}
            renderItem={renderEventCard}
            keyExtractor={(item) => item.id}
        />
    );
};
