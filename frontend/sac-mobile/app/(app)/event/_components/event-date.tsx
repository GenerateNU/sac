import React from 'react';
import { Text, View } from 'react-native';

import CalendarIcon from '@/assets/images/svg/calendar.svg';

type EventDateProps = {
    startTime: Date;
};

const EventDate = ({ startTime }: EventDateProps) => {
    const formatTime = (date: Date) => {
        return date.toLocaleTimeString([], {
            hour: 'numeric',
            minute: '2-digit'
        });
    };

    const date = startTime.toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric'
    });
    const weekday = startTime.toLocaleDateString('en-US', { weekday: 'long' });
    const start = formatTime(startTime);

    return (
        <View
            className="pt-3 pb-3 bg-white rounded-lg shadow-sm col-span-1"
            style={{ width: '32%' }}
        >
            <View className="flex-col items-center justify-center gap-y-2">
                <CalendarIcon width={24} height={24} />
                <Text className="font-semibold text-center text-black">
                    {date}
                </Text>
                <View className="flex-col items-center">
                    <Text className="text-xs font-normal text-neutral-400">
                        {weekday}
                    </Text>
                    <Text className="text-xs font-normal text-neutral-400">
                        {start}
                    </Text>
                </View>
            </View>
        </View>
    );
};

export { EventDate };
