import React from 'react';
import { Text, View } from 'react-native';

import TimeIcon from '@/assets/images/svg/time.svg';

type EventTimeProps = {
    startTime: Date;
    endTime: Date;
};

const EventTime = ({ startTime, endTime }: EventTimeProps) => {
    const formatTime = (date: Date) => {
        return date.toLocaleTimeString([], {
            hour: 'numeric',
            minute: '2-digit'
        });
    };

    const formatLength = () => {
        const length = (endTime.getTime() - startTime.getTime()) / (1000 * 60 * 60);

        if (length < 1) {
            return `${length} minutes`
        } else if (length === 1) {
            return `${length} hour`
        } else {
            return `${length} hours`
        }
    }

    const weekday = startTime.toLocaleDateString('en-US', { weekday: 'long' });
    const start = formatTime(startTime);

    return (
        <View className='p-3 bg-white rounded-lg shadow-sm col-span-1' style={{ width: '32%' }}>
            <View className='flex-col items-center justify-center gap-y-2'>
                <TimeIcon width={24} height={24} />
                <Text className='font-semibold text-center text-black'>
                    {formatLength()}
                </Text>
                <View className='flex-col items-center'>
                    <Text className='text-xs text-center font-normal text-neutral-400'>
                        Length of the Event
                    </Text>
                </View>
            </View>
        </View>
    );
};

export { EventTime };
