import React from 'react';
import { ActivityIndicator, Image, Pressable, Text, View } from 'react-native';

import { useAssets } from 'expo-asset';
import { router } from 'expo-router';

import { UseQueryResult } from '@tanstack/react-query';

import BluePin from '@/assets/images/svg/blue-pin.svg';
import EventDemoDay from '@/assets/images/svg/event-demo-day.svg';
import Pin from '@/assets/images/svg/pin.svg';
import { useClub } from '@/hooks/use-club';
import { useEventTags } from '@/hooks/use-event';
import { isToday } from '@/lib/utils';
import { Event } from '@/types/event';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';
import { randomEventColorData } from '@/lib/utils';

const EventHomePageCardHorizontal = ({ event }: { event: Event }) => {
    const { data: host, isLoading, error } = useClub(event.host);

    if (isLoading) {
        // skeleton loading
        <ActivityIndicator size={'large'} />;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    const randomEventColor =
        randomEventColorData[
        Math.round(Math.random() * (randomEventColorData.length - 1))
        ];

    return (
        <Pressable
            className="flex-1 w-full mx-2 bg-white shadow-sm rounded-xl"
            onPress={() => router.push(`/event/${event.id}`)}
        >
            <View className="flex-row">
                <View className="px-3 my-5">
                    <View className="flex-row items-start pb-2">
                        <Image
                            source={{ uri: randomEventColor.image }}
                            style={{ width: 98, height: 98 }}
                            className="rounded-xl"
                        />
                        <View className="pl-[5%] flex-col gap-y-1 shrink pr-[3%]">
                            {isToday(event.start_time) && (
                                <Text className="text-md font-bold leading-5 text-[#609F80]">
                                    Happening Today
                                </Text>
                            )}
                            <Text className="leading-5 text-md text-zinc-900 font-extralight">
                                {event.name}
                            </Text>
                            <View className="flex-row items-center">
                                <Text className="text-xs text-neutral-600 font-extralight">
                                    Hosted by {host?.name ?? 'No Hosts'}
                                </Text>
                            </View>
                            <View className="flex-row items-center gap-x-2">
                                <BluePin />
                                <Text className="text-xs text-zinc-900 font-extralight">
                                    {event.location.split(',')[0]}
                                </Text>
                            </View>
                        </View>
                    </View>
                    <HomePageTagList
                        id={event.id}
                        tagQueryResult={useEventTags}
                    />
                </View>
            </View>
        </Pressable>
    );
};

const EventHomePageCard = ({ event }: { event: Event }) => {
    const { data: host, isLoading, error } = useClub(event.host);

    // const [asset, assetError] = useAssets(
    //     require('@/assets/images/showcase.png')
    // );

    // if (assetError) {
    //     return <Text>Error: {assetError.message}</Text>;
    // }

    if (isLoading) {
        // skeleton loading
        <ActivityIndicator size={'large'} />;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }


    const randomEventColor =
        randomEventColorData[
        Math.round(Math.random() * (randomEventColorData.length - 1))
        ];

    return (
        <Pressable
            className="flex-1 mr-3 bg-white shadow-sm rounded-xl"
            style={{ width: 168, height: 225 }}
            onPress={() => router.push(`/event/${event.id}`)}
        >
            <View className="flex-col p-4 gap-y-2">
                    <Image
                    source={{ uri: randomEventColor.image }}
                        style={{ width: 142, height: 64 }}
                        resizeMode="cover"
                        className="rounded-xl"
                />
                <View className=" flex-col gap-y-1 shrink pr-[3%]">
                    <Text
                        className="text-xs text-zinc-900 font-extralight"
                        numberOfLines={1}
                    >
                        {host?.name ?? 'No Hosts'}
                    </Text>
                    <Text
                        className="text-lg leading-5 text-zinc-900 font-extralight"
                        numberOfLines={1}
                    >
                        {event.name}
                    </Text>
                </View>
                <View className="flex-col items-start pt-2 gap-x-2 gap-y-1">
                    <Text className="text-md text-zinc-900 font-extralight">
                        {new Intl.DateTimeFormat('en-US', {
                            hour: 'numeric',
                            minute: 'numeric'
                        }).format(new Date(event.start_time))}{' '}
                        -{' '}
                        {new Intl.DateTimeFormat('en-US', {
                            hour: 'numeric',
                            minute: 'numeric'
                        }).format(new Date(event.end_time))}
                    </Text>
                    <Text className="text-md text-zinc-900 font-extralight">
                        {new Intl.DateTimeFormat('en-US', {
                            month: 'long',
                            day: 'numeric',
                            year: 'numeric'
                        }).format(new Date(event.start_time))}
                    </Text>
                    <View className="flex-row items-center ">
                        <Pin />
                        <Text className="pl-2 text-xs text-zinc-900 font-extralight">
                            {event.location.split(',')[0]}
                        </Text>
                    </View>
                </View>
            </View>
        </Pressable>
    );
};

type HomePageTagListProps = {
    id: string;
    tagQueryResult: (id: uuid) => UseQueryResult<Tag[], Error>;
};

const HomePageTagList = ({ id, tagQueryResult }: HomePageTagListProps) => {
    const { data: tags, isLoading, error } = tagQueryResult(id);

    if (isLoading) {
        return <ActivityIndicator size={'small'} />;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View className="flex flex-row flex-wrap gap-y-2">
            {tags &&
                tags.slice(0, 2).map((tag) => (
                    <View
                        className="flex flex-row px-3 py-1 mr-2 border border-gray-200 rounded-xl"
                        key={tag.id}
                    >
                        <Text className="text-xs text-zinc-900 font-extralight">
                            {tag.name}
                        </Text>
                    </View>
                ))}
        </View>
    );
};

export { EventHomePageCardHorizontal, EventHomePageCard, HomePageTagList };
