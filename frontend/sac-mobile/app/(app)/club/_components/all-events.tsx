import { Button } from '@/components/button';
import { EventCard, EventCardProps } from '@/components/eventCard';
import { FaqCard, FaqCardProps } from '@/components/faqCard';
import { useEvents } from '@/hooks/use-events';
import { ScrollView, Text, View } from 'react-native';

const Events = () => {
    const { data: events, isLoading, error } = useEvents("sac");

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <Text className="text-black font-bold">Events</Text>
            <ScrollView horizontal={true} className="pt-[2%] pb-[10%]">
                <View className="flex-row">
                    {events && events.map((event: EventCardProps) => (
                        <EventCard title={event.title} time={event.time} date={event.date} location={event.location} />
                    ))}
                </View>
            </ScrollView>
        </View>

    );
}

export { Events };