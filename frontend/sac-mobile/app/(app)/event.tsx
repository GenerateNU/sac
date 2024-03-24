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
import { useEvent } from '@/hooks/use-event';
import {
    getDayOfWeek,
    getFormattedTime,
    getMonth,
    getOrdinalSuffix
} from '@/lib/date';

const displayHosts = (hosts: string[]) => {
    if (hosts.length === 1) {
        return (
            <Text>
                Hosted by
                <Text
                    onPress={() => console.log(`${hosts[0]} pressed`)}
                    className="font-bold"
                >
                    {hosts[0]}
                </Text>
            </Text>
        );
    } else {
        return (
            <Text>
                Hosted by{' '}
                {hosts.slice(0, -1).map((host, index) => (
                    <Text
                        key={index}
                        onPress={() => console.log(`${host} pressed`)}
                        className="font-bold"
                    >
                        {host}
                        {index === hosts.length - 2 ? '' : ', '}
                    </Text>
                ))}{' '}
                &{' '}
                <Text
                    onPress={() =>
                        console.log(`${hosts[hosts.length - 1]} pressed`)
                    }
                    className="font-bold"
                >
                    {hosts[hosts.length - 1]}
                </Text>
            </Text>
        );
    }
};

const displaySameDayEvent = (startTime: Date, endTime: Date) => {
    return (
        <>
            <Text className="ml-1 text-base font-medium">
                {getDayOfWeek(startTime)}, {getMonth(startTime)}{' '}
                {startTime.getDate()}
                {getOrdinalSuffix(startTime.getDate())}
                {'\n'}
                {getFormattedTime(startTime)} - {getFormattedTime(endTime)}
            </Text>
        </>
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

const Event = () => {
    const { data: event, isLoading, error } = useEvent('tester');

    // TODO: Handle error once we have error components
    if (error) {
        console.error(error);
    }

    // TODO: Handle loading state once we have loading components
    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (!event) {
        return <Text>Event not found</Text>;
    }

    return (
        <>
            <ScrollView showsVerticalScrollIndicator={false}>
                <View className="w-screen bg-gray-600 h-[200]" />
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
                <View className="flex-row justify-between">
                    <View className="w-24 h-24 rounded-xl bg-gray-300 ml-6 -translate-y-12" />
                    <View className="flex-row justify-center items-center w-full gap-3">
                        <Button
                            onPress={() => console.log('X Icon')}
                            size={'icon'}
                            variant={'secondary'}
                            className="-translate-y-12"
                        >
                            <Text>X</Text>
                        </Button>
                        <Button
                            onPress={() => console.log('Y Icon')}
                            size={'icon'}
                            variant={'secondary'}
                            className="-translate-y-12"
                        >
                            <Text>Y</Text>
                        </Button>
                        <Button
                            onPress={() => console.log('RSVP')}
                            size={'default'}
                            variant={'default'}
                            className="w-18 -translate-y-12 mr-6"
                        >
                            <Text>RSVP</Text>
                        </Button>
                    </View>
                </View>
                <View className="ml-6 -translate-y-6">
                    <Text className="text-3xl font-bold">{event.name}</Text>
                    <Text className="mt-3 text-lg">
                        {displayHosts(event.hosts)}
                    </Text>
                    {displayEventTime(event.startTime, event.endTime)}
                    {displayEventLocation(event.location, event.meetingLink)}
                </View>
            </ScrollView>
        </>
    );
};

export default Event;
