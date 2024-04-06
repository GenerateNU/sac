import { View, Text, TouchableOpacity } from 'react-native'
import React from 'react'
import { MaterialCommunityIcons } from '@expo/vector-icons';

type EventTimeProps = {
    startTime: Date;
    endTime: Date;
}

const EventTime = ({ startTime, endTime }: EventTimeProps) => {
    const formatDate = (date: Date) => {
        const day = date.toLocaleDateString('en-US', { weekday: 'long' });
        const month = date.toLocaleDateString('en-US', { month: 'long' });
        const dateNum = date.getDate();
        const ordinalSuffix = dateNum === 1 ? 'st' : dateNum === 2 ? 'nd' : dateNum === 3 ? 'rd' : 'th';

        return `${day}, ${month} ${dateNum}${ordinalSuffix}`;
    };

    const formatTime = (date: Date) => {
        return date.toLocaleTimeString([], { hour: 'numeric', minute: '2-digit' });
    };

    return (
        <TouchableOpacity
            onPress={() => console.log('Navigate to the calendar')}
            className="flex flex-row items-center py-2"
        >
            <MaterialCommunityIcons name="calendar" size={36} color="black" />
            <View className="flex flex-col">
                <Text className="px-2 text-base font-medium">
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
                            {formatTime(startTime)} - {formatDate(endTime)} {formatTime(endTime)}
                        </>
                    )}
                </Text>
            </View>
        </TouchableOpacity>
    );
};

export { EventTime };
