import { useEffect, useRef, useState } from 'react';
import { Dimensions } from 'react-native';
import Animated, {
    interpolate,
    useAnimatedRef,
    useAnimatedStyle,
    useScrollViewOffset
} from 'react-native-reanimated';

import { Stack, useLocalSearchParams } from 'expo-router';

import { EventType, clubApi } from '@generatesac/lib';
import { eventApi } from '@generatesac/lib';
import BottomSheet from '@gorhom/bottom-sheet';

import { Arrow, Box, KebabMenu } from '@/src/app/(design-system)';
import { SACColors } from '@/src/app/(design-system)';
import { Button } from '@/src/app/(design-system)/components/Button/Button';
import { description, events, tags } from '@/src/consts/event-page';

import { AboutEvent } from './components/about';
import { Description } from './components/description';
import EventPageError from './components/error';
import { Location } from './components/location';
import { Overview } from './components/overview';
import { RegisterBottomSheet } from './components/register';
import { ShareEventBottomSheet } from './components/share-event';
import EventPageSkeleton from './components/skeleton';

const MockEvent = {
    eventImage:
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQMm5aVXtgL-4e7X_Ed3N7P6TxqLHBNNvabYA&s',
    clubId: 'generate', // uuid
    logo: 'data:image/png;base64,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',
    eventName: 'Generate Spring Showcase 2024',
    color: 'aqua' as SACColors,
    club: 'Generate',
    startTime: new Date('2024-06-01T10:00:00'),
    endTime: new Date('2024-06-01T12:00:00'),
    location: 'ISEC',
    eventType: 'hybrid' as EventType,
    zoomLink: 'https://google.com', // zoom link if event is hybrid or virtual
    description: description,
    tags: tags,
    events: events, // upcoming events
    link: 'https://generatenu.com' // link of this event page
};

const EventPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();
    const { width } = Dimensions.get('window');
    const IMG_HEIGHT = width;
    const shareEvent = useRef<BottomSheet>(null);
    const register = useRef<BottomSheet>(null);
    const bottomSheet = useRef<BottomSheet>(null);
    const scrollRef = useAnimatedRef<Animated.ScrollView>();
    const scrollOffset = useScrollViewOffset(scrollRef);

    const [retriggerFetch, setRetriggerFetch] = useState(false);

    const [
        getEvent,
        { isLoading: eventLoading, error: eventError, data: event }
    ] = eventApi.useLazyEventQuery();
    const [getClub, { isLoading: clubLoading, error: clubError, data: club }] =
        clubApi.useLazyClubQuery();
    const [
        getEventTags,
        { isLoading: tagsLoading, error: tagsError, data: tags }
    ] = eventApi.useLazyEventTagsQuery();

    const imageAnimatedStyle = useAnimatedStyle(() => {
        return {
            transform: [
                {
                    translateY: interpolate(
                        scrollOffset.value,
                        [-IMG_HEIGHT, 0, IMG_HEIGHT],
                        [-IMG_HEIGHT / 2, 0, IMG_HEIGHT * 0.75]
                    )
                },
                {
                    scale: interpolate(
                        scrollOffset.value,
                        [-IMG_HEIGHT, 0, IMG_HEIGHT],
                        [2, 1, 1]
                    )
                }
            ]
        };
    });

    const headerAnimatedStyle = useAnimatedStyle(() => {
        return {
            transform: [
                {
                    translateY: interpolate(
                        scrollOffset.value,
                        [-IMG_HEIGHT, 0, IMG_HEIGHT],
                        [0, 0, -IMG_HEIGHT * 0.75]
                    )
                }
            ]
        };
    });

    useEffect(() => {
        // Fetch events
        getEvent(id as string).then(({ data: eventData }) => {
            if (eventData) {
                // Fetch club
                getClub(eventData.host as string);
                // Fetch tags:
                getEventTags(eventData.id);
            }
        });
    }, [retriggerFetch, id, getClub, getEvent, getEventTags]);

    const apiLoading = eventLoading || clubLoading || tagsLoading;
    const apiError = eventError || clubError || tagsError;
    const allData = event && tags && club;

    return (
        <Box backgroundColor="white" height={'100%'}>
            <Stack.Screen
                options={{
                    headerTitle: '',
                    headerTransparent: true,
                    headerShown: true,
                    headerTitleStyle: {
                        color: 'white'
                    },
                    headerLeft: () => (
                        <Animated.View style={headerAnimatedStyle}>
                            <Arrow color={apiError ? 'black' : 'white'} />
                        </Animated.View>
                    ),
                    headerRight: !apiError
                        ? () => (
                              <Animated.View style={headerAnimatedStyle}>
                                  <KebabMenu
                                      onPress={() =>
                                          shareEvent.current?.snapToIndex(0)
                                      }
                                      color="white"
                                  />
                              </Animated.View>
                          )
                        : () => <></>
                }}
            />
            <Animated.ScrollView
                showsVerticalScrollIndicator={false}
                scrollEventThrottle={16}
                ref={scrollRef}
            >
                {apiLoading ? (
                    <EventPageSkeleton />
                ) : apiError ? (
                    <EventPageError refetch={setRetriggerFetch} />
                ) : (
                    allData && (
                        <>
                            <Box>
                                <Animated.Image
                                    source={{
                                        uri: MockEvent.eventImage
                                    }}
                                    style={[
                                        imageAnimatedStyle,
                                        { width: '100%', aspectRatio: 1 }
                                    ]}
                                />
                            </Box>
                            <Box backgroundColor="white">
                                <Box margin="xl" gap="xxl">
                                    <Box gap="l">
                                        <Overview
                                            clubId={event.host}
                                            logo={club.logo ?? ''}
                                            eventName={event.name}
                                            color={MockEvent.color}
                                            club={club?.name}
                                            startTime={
                                                new Date(event.start_time)
                                            }
                                            endTime={new Date(event.end_time)}
                                            location={
                                                event.location ||
                                                'ISEC, Room 204'
                                            }
                                            type={event.event_type}
                                        />
                                        <Button
                                            variant="standardButton"
                                            color={MockEvent.color}
                                            onPress={() =>
                                                register.current?.snapToIndex(0)
                                            }
                                        >
                                            Register
                                        </Button>
                                    </Box>
                                    <AboutEvent
                                        color={MockEvent.color}
                                        zoomLink={event.link}
                                        description={event.description}
                                        tags={tags}
                                        onPress={() =>
                                            bottomSheet.current?.snapToIndex(0)
                                        }
                                    />
                                    <Location
                                        location={
                                            event.location || 'ISEC, Room 204'
                                        }
                                    />
                                </Box>
                                {/* <UpcomingEvent events={MockEvent.events} /> */}
                            </Box>
                        </>
                    )
                )}
            </Animated.ScrollView>
            <ShareEventBottomSheet ref={shareEvent} link={MockEvent.link} />
            <RegisterBottomSheet
                startTime={new Date(event?.start_time as string)}
                endTime={new Date(event?.end_time as string)}
                eventDetail={event?.description as string}
                eventName={event?.name as string}
                location={event?.location as string}
                eventType={event?.event_type as EventType}
                ref={register}
            />
            <Description
                description={event?.description as string}
                ref={bottomSheet}
            />
        </Box>
    );
};

export default EventPage;
