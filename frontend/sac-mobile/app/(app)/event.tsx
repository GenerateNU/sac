import React from 'react';
import {
    Linking,
    ScrollView,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { router } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import { Button } from '@/components/button';
import HostedBy from '@/components/hosted-by';
import HostedByText from '@/components/hosted-by-text';
import { useEvent } from '@/hooks/use-event';
import {
    getDayOfWeek,
    getFormattedTime,
    getMonth,
    getOrdinalSuffix
} from '@/lib/date';

const displaySameDayEvent = (startTime: Date, endTime: Date) => {
    return (
        <Text className="ml-1 text-base font-medium">
            {getDayOfWeek(startTime)}, {getMonth(startTime)}{' '}
            {startTime.getDate()}
            {getOrdinalSuffix(startTime.getDate())}
            {'\n'}
            {getFormattedTime(startTime)} - {getFormattedTime(endTime)}
        </Text>
    );
};

const displayDifferentDayEvent = (startTime: Date, endTime: Date) => {
    return (
        <>
            <Text className="ml-1 text-base font-medium">
                {getDayOfWeek(startTime)}, {getMonth(startTime)}{' '}
                {startTime.getDate()}
                {getOrdinalSuffix(startTime.getDate())}
                {' - '}
                {getDayOfWeek(endTime)}, {getMonth(endTime)} {endTime.getDate()}
                {getOrdinalSuffix(endTime.getDate())}{' '}
            </Text>
        </>
    );
};

const displayEventTime = (startTime: Date, endTime: Date) => {
    const onSameDay = startTime.getDate() === endTime.getDate();

    return (
        <TouchableOpacity
            onPress={() => console.log('Navigate to the calendar')}
            className="mt-2 flex-row items-center"
        >
            <MaterialCommunityIcons name="calendar" size={36} color="black" />
            {onSameDay
                ? displaySameDayEvent(startTime, endTime)
                : displayDifferentDayEvent(startTime, endTime)}
        </TouchableOpacity>
    );
};

const displayEventLocation = (
    location: string,
    meetingLink: string | undefined
) => {
    return (
        <View className="mt-2 flex-row items-center">
            <MaterialCommunityIcons name="map-marker" size={36} color="black" />
            <View className="ml-1 flex-col">
                <Text className="text-base font-medium">{location}</Text>
                {meetingLink && (
                    <TouchableOpacity
                        onPress={() =>
                            Linking.openURL(meetingLink).catch((err) =>
                                console.error('An error occurred', err)
                            )
                        }
                    >
                        <Text className="text-blue-500 visited:text-purple-500">
                            {meetingLink}
                        </Text>
                    </TouchableOpacity>
                )}
            </View>
        </View>
    );
};

// TODO: handle link OR location
const Event = () => {
    const {
        data: event,
        isLoading: isEventLoading,
        error: eventError
    } = useEvent('tester');

    // TODO: Handle error once we have error components
    if (eventError) {
        console.error(eventError);
        return <Text>Error fetching event</Text>;
    }

    // TODO: Handle loading state once we have loading components
    if (isEventLoading) {
        return <Text>Loading...</Text>;
    }

    if (!event) {
        return <Text>Event not found</Text>;
    }

    return (
        <>
            <ScrollView showsVerticalScrollIndicator={false}>
                <View className="flex flex-col items-center w-screen">
                    <View className="w-full bg-gray-600 h-[200]" />
                    <Button
                        onPress={() => {
                            console.log('Back Button Pressed');
                            router.back();
                        }}
                        size={'icon'}
                        variant={'icon'}
                        className="flex justify-center items-center absolute top-12 left-6 z-10 bg-gray-300 rounded-full"
                    >
                        <MaterialCommunityIcons
                            name="chevron-left"
                            size={28}
                            color="rgb(75 85 99)"
                        />
                    </Button>
                    <View className="flex-row justify-between w-full -translate-y-12 px-6">
                        <View className="w-24 h-24 rounded-xl bg-gray-300" />
                        <View className="flex-row justify-center items-center space-x-3">
                            <Button
                                onPress={() => console.log('X Icon')}
                                size={'icon'}
                                variant={'secondary'}
                            >
                                <Text>X</Text>
                            </Button>
                            <Button
                                onPress={() => console.log('Y Icon')}
                                size={'icon'}
                                variant={'secondary'}
                            >
                                <Text>Y</Text>
                            </Button>
                            <Button
                                onPress={() => console.log('RSVP')}
                                size={'default'}
                                variant={'default'}
                                className="w-18"
                            >
                                <Text>RSVP</Text>
                            </Button>
                        </View>
                    </View>
                    <View className="px-6 py-3 w-full -translate-y-6">
                        <Text className="text-3xl font-bold">{event.name}</Text>
                        <View className="py-1.5">
                            <HostedByText eventID={event.id} />
                        </View>
                        {displayEventTime(event.startTime, event.endTime)}
                        {displayEventLocation(
                            event.location,
                            event.meetingLink
                        )}
                    </View>
                    {/* TODO: Tag Carousel */}
                    <View className="px-6 py-3 w-full">
                        <Text className="text-2xl font-bold">About Event</Text>
                        {/* TODO: use readmoreable text from Melody */}
                        <Text className="mt-2 text-lg">{event.content}</Text>
                    </View>
                    <View className="px-6 py-3 w-full">
                        <View className="flex-row justify-between items-center">
                            <Text className="text-2xl font-bold">Location</Text>
                            <Button
                                // TODO: Navigate to the native map app
                                onPress={() =>
                                    console.log('Navigate to the map')
                                }
                                size={'default'}
                                variant={'default'}
                            >
                                <Text>Open in Maps</Text>
                            </Button>
                        </View>
                        <Text className="mt-2 text-lg">{event.location}</Text>
                        {event.meetingLink && (
                            <TouchableOpacity
                                onPress={() =>
                                    Linking.openURL(event.meetingLink!).catch(
                                        (err) =>
                                            console.error(
                                                'An error occurred',
                                                err
                                            )
                                    )
                                }
                            >
                                <Text className="text-blue-500 visited:text-purple-500">
                                    {event.meetingLink}
                                </Text>
                            </TouchableOpacity>
                        )}
                        <View className="w-full h-36 rounded-xl bg-gray-400 my-3" />
                    </View>
                    <View className="px-6 py-3 w-full">
                        <HostedBy eventID={event.id} />
                    </View>
                </View>
            </ScrollView>
        </>
    );
};

export default Event;
