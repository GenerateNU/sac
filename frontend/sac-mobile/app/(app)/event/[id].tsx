import React, { useRef } from 'react';
import {
    ActivityIndicator,
    Animated,
    RefreshControl,
    ScrollView,
    Text,
    View
} from 'react-native';

import { Stack, useLocalSearchParams } from 'expo-router';
import { StatusBar } from 'expo-status-bar';

import BottomSheet from '@gorhom/bottom-sheet';

import { MenuView } from '@react-native-menu/menu';

// import { HostList } from '@/app/(app)/event/_components/host-list';
import ShareIcon from '@/assets/images/svg/share.svg';
import { TagList } from '@/components/all-tags';
import { Button } from '@/components/button';
import { Description } from '@/components/description';
import { Title } from '@/components/title';
import { useEvent, useEventTags } from '@/hooks/use-event';

import { EventLocation } from './_components/event-location';
import { EventTime } from './_components/event-time';
import { HostNames } from './_components/host-names';
import { LocationView } from './_components/location-view';
import { AllHosts } from './_components/all-hosts';
import { RSVPBottomSheet } from './_components/rsvp-bottom-sheet';
import { HeaderImage } from '@/components/entity-header';
import { useAssets } from 'expo-asset';

// TODO: handle link OR location
const EventPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();

    const { data: event, isLoading, error, refetch } = useEvent(id);

    const ref = useRef<BottomSheet>(null);
    const pan = useRef(new Animated.ValueXY()).current;

    const [asset, assetError] = useAssets(require('@/assets/images/sandbox-header.png'));
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

    if (!event) {
        return <Text>Event not found</Text>;
    }

    return (
        <>
            <Stack.Screen options={{ title: `${event[0].name}` }} />
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
                <View className="flex-col">
                    <HeaderImage
                        image="https://blenderartists.org/uploads/default/optimized/4X/7/e/2/7e2d7bea4ac21388c4a96e1371f375c4ce00094b_2_1024x576.jpg"
                        pan={pan}
                        type="event"
                        ref={ref}
                    />
                    <View className="px-5 pb-8">
                        <View className="flex gap-y-2">
                            <Title title={event[0].name} />
                            <HostNames eventID={event[0].id} />
                            <View className="flex py-1 m-0 gap-y-2">
                                <EventTime
                                    startTime={new Date(event[0].start_time)}
                                    endTime={new Date(event[0].end_time)}
                                />
                                <EventLocation
                                    location={event[0].location}
                                    meetingLink={event[0].meeting_link}
                                />
                            </View>
                            <TagList tagQueryResult={useEventTags} id={event[0].id} />
                            <Description
                                title={'About Event'}
                                description={event[0].content}
                            />
                            <LocationView
                                location={event[0].location}
                                meetingLink={event[0].meeting_link}
                            />
                            <AllHosts eventID={event[0].id} />
                            <Button
                                onPress={() => console.log('RSVP')}
                                size={'default'}
                                variant={'default'}
                                className="w-20 rounded-xl"
                            >
                                <Text>Report</Text>
                            </Button>
                        </View>
                    </View>
                </View>
            </ScrollView>
            <RSVPBottomSheet ref={ref} />
        </>
    );
};

export default EventPage;
