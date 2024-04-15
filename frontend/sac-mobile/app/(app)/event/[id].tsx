import React, { useRef } from 'react';
import { ScrollView, Text, View } from 'react-native';

import { Stack, useLocalSearchParams } from 'expo-router';

import BottomSheet from '@gorhom/bottom-sheet';

import { AllHosts } from '@/app/(app)/event/_components/all-hosts';
import { TagList } from '@/components/all-tags';
import { Button } from '@/components/button';
import { Description } from '@/components/description';
import { Title } from '@/components/title';
import { useEvent } from '@/hooks/use-event';

import { EventHeader } from './_components/event-header';
import { EventLocation } from './_components/event-location';
import { EventTime } from './_components/event-time';
import { HostNames } from './_components/host-names';
import { LocationView } from './_components/location-view';
import RSVPBottomSheet from './_components/rsvp-bottom-sheet';

// TODO: handle link OR location
const EventPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();

    const { data: event, isLoading, error } = useEvent(id);

    const ref = useRef<BottomSheet>(null);

    if (error) {
        console.error(error);
        return <Text>Error fetching event</Text>;
    }

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (!event) {
        return <Text>Event not found</Text>;
    }

    console.log('[event]', event);

    return (
        <>
            <Stack.Screen options={{ title: `${event.name}` }} />
            <ScrollView showsVerticalScrollIndicator={false}>
                <View className="flex-col justify-between flex-1 w-full">
                    <View className="w-full bg-gray-600 h-[100]" />
                    <View className="w-full px-6 pb-5">
                        <EventHeader ref={ref} />
                        <View className="flex gap-y-2">
                            <Title title={event.name} />
                            <HostNames eventID={event.id} />
                            <View className="flex py-1 m-0 gap-y-2">
                                <EventTime
                                    startTime={event.start_time}
                                    endTime={event.end_time}
                                />
                                <EventLocation
                                    location={event.location}
                                    meetingLink={event.meeting_link}
                                />
                            </View>
                            <TagList id={event.id} />
                            <Description
                                title={'About Event'}
                                description={event.content}
                            />
                            <LocationView
                                location={event.location}
                                meetingLink={event.meeting_link}
                            />
                            <AllHosts eventID={event.id} />
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
