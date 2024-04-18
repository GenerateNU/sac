import React from 'react';
import { Text, View } from 'react-native';

import TimeIcon from '@/assets/images/svg/time.svg';

type EventTimeProps = {
    startTime: Date;
    endTime: Date;
};

const EventTime = ({ startTime, endTime }: EventTimeProps) => {
    const formatLength = () => {
        const length =
            (endTime.getTime() - startTime.getTime()) / (1000 * 60 * 60);

        if (length < 1) {
            return `${length} minutes`;
        } else if (length === 1) {
            return `${length} hour`;
        } else {
            return `${length} hours`;
        }
    };

    return (
        <View
            className="p-3 bg-white rounded-lg shadow-sm col-span-1"
            style={{ width: '32%' }}
        >
            <View className="flex-col items-center justify-center gap-y-2">
                <TimeIcon width={24} height={24} />
                <Text className="font-semibold text-center text-black">
                    {formatLength()}
                </Text>
                <View className="flex-col items-center">
                    <Text className="text-xs text-center font-normal text-neutral-400">
                        Length of the Event
                    </Text>
                </View>
            </View>
        </View>
    );
};

export { EventTime };
