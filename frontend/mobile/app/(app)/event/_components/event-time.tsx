import React from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

import CalendarIcon from '@/assets/images/svg/calendar.svg';

type EventTimeProps = {
    startTime: Date;
    endTime: Date;
};

const EventTime = ({ startTime, endTime }: EventTimeProps) => {
    const formatDate = (date: Date) => {
        const day = date.toLocaleDateString('en-US', { weekday: 'long' });
        const month = date.toLocaleDateString('en-US', { month: 'long' });
        const dateNum = date.getDate();
        const ordinalSuffix =
            dateNum === 1
                ? 'st'
                : dateNum === 2
                  ? 'nd'
                  : dateNum === 3
                    ? 'rd'
                    : 'th';

        return `${day}, ${month} ${dateNum}${ordinalSuffix}`;
    };

    const formatTime = (date: Date) => {
        return date.toLocaleTimeString([], {
            hour: 'numeric',
            minute: '2-digit'
        });
    };

    return (
        <TouchableOpacity
            onPress={() => console.log('Navigate to the calendar')}
            className="flex flex-row"
        >
            <CalendarIcon className="self-center" />
            <View className="flex-col pl-2">
                <Text className="text-black">
                    {startTime.getDate() === endTime.getDate() ? (
                        <>
                            {formatDate(startTime)}
                            {'\n'}
                            {formatTime(startTime)} - {formatTime(endTime)}
                        </>
                    ) : (
                        <>
                            {formatDate(startTime)}
                            {'\n'}
                            {formatTime(startTime)} - {formatDate(endTime)}{' '}
                            {formatTime(endTime)}
                        </>
                    )}
                </Text>
            </View>
        </TouchableOpacity>
    );
};

export { EventTime };
