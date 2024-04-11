import React from 'react';
import { Alert, Platform, ScrollView, Share, Text, View } from 'react-native';

import { Stack, useLocalSearchParams } from 'expo-router';

import { MenuView } from '@react-native-menu/menu';

import { AllHosts } from '@/app/(app)/event/_components/all-hosts';
import ShareIcon from '@/assets/images/svg/share.svg';
import { TagList } from '@/components/all-tags';
import { Button } from '@/components/button';
import { Description } from '@/components/description';
import { useAuthStore } from '@/hooks/use-auth';
import { useEvent } from '@/hooks/use-event';

import { Title } from '../_components/title';
import { EventTime } from './_components/event-time';
import { HostNames } from './_components/host-names';
import { LocationView } from './_components/location-view';
import { EventLocation } from './_components/event-location';
import { EventHeader } from './_components/event-header';

// TODO: handle link OR location
const EventPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { user } = useAuthStore();

    const { data: event, isLoading, error } = useEvent(id);

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
            <Stack.Screen
                options={{ title: `${event.name} - ${user?.first_name}` }}
            />
            <ScrollView showsVerticalScrollIndicator={false}>
                <View className="flex-col justify-between flex-1 w-full">
                    <View className="w-full bg-gray-600 h-[100]" />
                    <View className="w-full px-6 pb-5">
                        <EventHeader />
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
        </>
    );
};

export default EventPage