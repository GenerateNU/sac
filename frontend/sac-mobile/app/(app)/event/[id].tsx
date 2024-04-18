import React, { useRef } from 'react';
import {
    ActivityIndicator,
    Animated,
    Platform,
    ScrollView,
    Text,
    View
} from 'react-native';

import { Stack, useLocalSearchParams } from 'expo-router';
import { StatusBar } from 'expo-status-bar';

import BottomSheet from '@gorhom/bottom-sheet';

import { AllHosts } from '@/app/(app)/event/_components/all-hosts';
import { TagList } from '@/components/all-tags';
import { Description } from '@/components/description';
import { Title } from '@/components/title';
import { useEvent } from '@/hooks/use-event';

import { HeaderImage } from '../../../components/entity-header';
import { HostNames } from './_components/host-names';
import { LocationView } from './_components/location-view';
import { RSVPBottomSheet } from './_components/rsvp-bottom-sheet';
import LocationIcon from '@/assets/images/svg/location.svg'
import PersonIcon from '@/assets/images/svg/person.svg'
import TimeIcon from '@/assets/images/svg/time.svg'
import { EventDate } from './_components/event-date';
import { EventTime } from './_components/event-time';
import { EventLocation } from './_components/event-location';
import { Button } from '@/components/button';
import { MenuView } from '@react-native-menu/menu';

// TODO: handle link OR location
const EventPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();

    const { data: event, isLoading, error } = useEvent(id);

    const ref = useRef<BottomSheet>(null);
    const pan = useRef(new Animated.ValueXY()).current;

    if (error) {
        console.error(error);
        return <Text>Error fetching event</Text>;
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
                        image="https://blenderartists.org/uploads/default/optimized/4X/7/e/2/7e2d7bea4ac21388c4a96e1371f375c4ce00094b_2_1024x576.jpg"
                        pan={pan}
                        type="event"
                        ref={ref}
                    />
                    <View
                        style={{ display: 'flex', flexDirection: 'column', padding: '5%', paddingTop: '2%', gap: 30, }}
                    >
                        <View style={{ display: 'flex', flexDirection: 'column', gap: 20, }}>
                            <View style={{ display: 'flex', flexDirection: 'column', gap: 10, }}>
                                <Title
                                    title={event[0].name}
                                />
                                <HostNames eventID={event[0].id} />
                            </View>
                            <TagList id={event[0].id} />
                        </View>
                        <View>
                            <View className='flex-row justify-between pb-5'>
                                <EventDate 
                                    startTime={new Date(event[0].start_time)}
                                />
                                <EventTime 
                                    startTime={new Date(event[0].start_time)}
                                    endTime={new Date(event[0].end_time)}
                                />
                                <EventLocation location={event[0].location} />
                            </View>
                            <Button
                                onPress={() => ref.current?.snapToIndex(0)}
                            >
                                Register
                            </Button>
                        </View>
                        <Description
                            title={'About Event'}
                            description={event[0].content}
                        />
                        <LocationView
                            location={event[0].location}
                            meetingLink={event[0].meeting_link}
                        />
                        <AllHosts eventID={event[0].id} />
                    </View>
                </View>
            </ScrollView>
            <RSVPBottomSheet ref={ref} />
        </>
    );
};

export default EventPage;
