import { ScrollView, Text, View } from 'react-native';

import { useEvents } from '@/hooks/use-event';
import { EventHomePageCardHorizontal } from '../../(tabs)/_components/event-homepage-card';

const Events = ({}: { clubID?: string }) => {
    // const { data: events, isLoading, error } = useEvents(clubID);
    const { data: events, isLoading, error } = useEvents(); // TODO: add limits and filters

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <View className="flex-row items-center justify-between">
                <Text className="text-xl font-light leading-relaxed text-zinc-900">
                    Events
                </Text>
                <Text className="text-neutral-400 text-[15px] font-normal">
                    View All
                </Text>
            </View>
            <ScrollView>
                <View className="flex-col mt-3">
                    <View className="gap-y-3">
                        {events &&
                            events.slice(0, 2).map((event, index) => (
                                <View key={index}>
                                    <EventHomePageCardHorizontal event={event} />
                                </View>
                            ))}
                    </View>
                </View>
            </ScrollView>
        </View>
    );
};

export { Events };
