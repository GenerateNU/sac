import { Button } from '@/components/button';
import { EventCard, EventCardProps } from '@/components/eventCard';
import { FaqCard, FaqCardProps } from '@/components/faqCard';
import { useEvents } from '@/hooks/use-events';
import { ScrollView, Text, View } from 'react-native';

const Events = () => {
    const { data: events, isLoading, error } = useEvents("sac"); // TODO: add limits and filters

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View>
            <View className="flex-row items-center justify-between">
                <Text className="text-lg font-semibold text-black">Events</Text>
                <Text className='text-neutral-400 text-[15px] font-normal'>View All</Text>
            </View>
            <ScrollView className="pt-[2%] pb-[10%]">
                <View className="flex-col">
                    <View className='px-2'>
                        {events && events.slice(0, 2).map((event: EventCardProps, index) => (
                            <>
                                <Text className='text-neutral-600 text-[15px] font-medium'>{event.date}</Text>
                                <EventCard {...event} />
                            </>
                        ))}
                        <Button className='bg-[#f8c45c]' variant="secondary" size="screenwide">View All Events</Button>
                    </View>
                </View>
            </ScrollView>
        </View>

    );
}

export { Events };