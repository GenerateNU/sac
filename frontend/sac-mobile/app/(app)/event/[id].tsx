import React, { useRef } from 'react';
import {
    ActivityIndicator,
    Animated,
    RefreshControl,
    ScrollView,
    Text,
    View
} from 'react-native';

import { useAssets } from 'expo-asset';
import { useLocalSearchParams } from 'expo-router';
import { StatusBar } from 'expo-status-bar';

import BottomSheet from '@gorhom/bottom-sheet';

import { TagList } from '@/components/all-tags';
import { Button } from '@/components/button';
import { Description } from '@/components/description';
import { HeaderImage } from '@/components/entity-header';
import { Title } from '@/components/title';
import { useAuthStore } from '@/hooks/use-auth';
import { useEvent, useEventTags } from '@/hooks/use-event';

import { AllHosts } from './_components/all-hosts';
import { EventDate } from './_components/event-date';
import { EventLocation } from './_components/event-location';
import { EventTime } from './_components/event-time';
import { HostNames } from './_components/host-names';
import { LocationView } from './_components/location-view';
import { RSVPBottomSheet } from './_components/rsvp-bottom-sheet';

const randomEventColorData = [
    {
        image: 'https://i.imgur.com/pLtAoSt.png',
        color: '#EDA990',
    },
    {
        image: 'https://i.imgur.com/1Auiabb.png',
        color: '#C5D9A4',
    },
    {
        image: 'https://i.imgur.com/pLtAoSt.png',
        color: '#EA90BB',
    },
    {
        image: 'https://i.imgur.com/pLtAoSt.png',
        color: '#EBC767',
    }
]

// TODO: handle link OR location
const EventPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { user } = useAuthStore();

    const { data: event, isLoading, error, refetch } = useEvent(id);

    const ref = useRef<BottomSheet>(null);
    const pan = useRef(new Animated.ValueXY()).current;
    const randomEventColor = randomEventColorData[Math.round(Math.random() * (randomEventColorData.length - 1))];

    const [asset, assetError] = useAssets(
        require('@/assets/images/sandbox-header.png')
    );
    if (assetError) {
        console.error(assetError);
        return <Text>Error fetching event</Text>;
    }

    if (!asset) {
        return <ActivityIndicator size={'large'} />;
    }

    if (isLoading) {
        return <ActivityIndicator size={'large'} />;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!event) {
        return <Text>Event not found</Text>;
    }

    return (
        <>
            <StatusBar style="light" />
            <ScrollView
                refreshControl={
                    <RefreshControl
                        refreshing={isLoading}
                        onRefresh={refetch}
                    />
                }
                scrollEventThrottle={1}
                showsVerticalScrollIndicator={false}
                onScroll={Animated.event(
                    [{ nativeEvent: { contentOffset: { y: pan.y } } }],
                    {
                        useNativeDriver: false
                    }
                )}
            >
                <View>
                    <HeaderImage
                        image={randomEventColor.image}
                        logo={randomEventColor.image}
                        pan={pan}
                        type="event"
                        ref={ref}
                        userID={user?.id ?? ''}
                    />
                    <View
                        style={{
                            display: 'flex',
                            flexDirection: 'column',
                            padding: '5%',
                            paddingTop: '2%',
                            gap: 30
                        }}
                    >
                        <View
                            style={{
                                display: 'flex',
                                flexDirection: 'column',
                                gap: 20
                            }}
                        >
                            <View
                                style={{
                                    display: 'flex',
                                    flexDirection: 'column',
                                    gap: 10
                                }}
                            >
                                <Title title={event[0].name} />
                                <HostNames eventID={event[0].id} />
                            </View>
                            <TagList
                                color={randomEventColor.color}
                                id={event[0].id}
                                tagQueryResult={useEventTags}
                            />
                        </View>
                        <View>
                            <View className="flex-row justify-between pb-5">
                                <EventDate
                                    startTime={new Date(event[0].start_time)}
                                />
                                <EventTime
                                    startTime={new Date(event[0].start_time)}
                                    endTime={new Date(event[0].end_time)}
                                />
                                <EventLocation location={event[0].location} />
                            </View>
                            <Button onPress={() => ref.current?.snapToIndex(0)}>
                                Register
                            </Button>
                        </View>
                        <Description
                            title={'About Event'}
                            description={event[0].content}
                        />
                        <LocationView location={event[0].location} />
                        <AllHosts eventID={event[0].id} />
                    </View>
                </View>
            </ScrollView>
            <RSVPBottomSheet ref={ref} />
        </>
    );
};

export default EventPage;
