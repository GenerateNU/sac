import React from 'react';
import { ActivityIndicator, Pressable, Text, View } from 'react-native';

import { router } from 'expo-router';

import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';
import Pin from '@/assets/images/svg/pin.svg';
import Time from '@/assets/images/svg/time.svg';
import { useClub } from '@/hooks/use-club';
import { Event } from '@/types/event';
import { useEventTags } from '@/hooks/use-event';

const EventHomePageCard = ({ event }: { event: Event }) => {
    const { data: host, isLoading, error } = useClub(event.host);

    if (isLoading) {
        // skeleton loading
        <ActivityIndicator size={'large'} />;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    const isToday = (date: Date) => {
        const today = new Date();
        const eventDate = new Date(date);

        return today.toDateString() === eventDate.toDateString();
    };

    return (
        <Pressable
            className='flex-1 w-[80%] bg-white rounded-xl shadow-sm'
            onPress={() => router.push(`/event/${event.id}`)}
        >
            <View>
                <View className="px-[5%] pt-[5%] pb-[4%]">
                    <View className="flex-row items-start pb-5">
                        <EventDemoDay
                            width={98}
                            height={98}
                            className="rounded-xl"
                        />
                        <View className="pl-[5%] flex-col gap-y-2 shrink pr-[3%]">
                            {isToday(event.start_time) && (
                                <Text className="text-md font-bold leading-5 text-[#609F80]">
                                    Happening Today
                                </Text>
                            )}
                            <Text className="text-base font-light text-zinc-900">
                                {event.name}
                            </Text>
                            <View className="flex-row items-center">
                                <Text className="text-xs text-neutral-600 font-extralight">
                                    Hosted by {host?.name ?? 'No Hosts'}
                                </Text>
                            </View>
                            <View className="flex-row items-center">
                                <Pin />
                                <Text className="text-sm pl-[2%] leading-5">
                                    {event.location.split(',')[0]}
                                </Text>
                            </View>
                        </View>
                    </View>
                    <HomePageTagList id={event.id} />
                </View>
            </View>
        </Pressable>
    );
};

type HomePageTagListProps = {
    id: string;
};

const HomePageTagList = ({ id }: HomePageTagListProps) => {
    const { data: tags, isLoading, error } = useEventTags(id);

    if (isLoading) {
        return <ActivityIndicator size={'small'} />;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View className="flex flex-row flex-wrap">
            {tags && tags.map((tag) =>
                <View className='flex flex-row px-3 py-1 mr-2 border border-gray-200 rounded-xl' key={tag.id}>
                    <Text className='text-xs text-zinc-900 font-extralight' >
                        {tag.name}
                    </Text>
                </View>
            )}
        </View>
    )
}


export { EventHomePageCard };
