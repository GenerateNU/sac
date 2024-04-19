import React from 'react';
import { Text, View } from 'react-native';

import Location from '@/assets/images/svg/location.svg';

type EventLocationProps = {
    location: string;
    meetingLink?: string;
};

const EventLocation = ({ location, meetingLink }: EventLocationProps) => {
    const title = () => {
        if (location && meetingLink) {
            return ['Hybrid', 'Event In-person and online'];
        } else if (location) {
            return ['In Person', 'Event In-person only'];
        } else {
            return ['Online', 'Event online Only'];
        }
    };

    const information = title();

    return (
        <View
            className="p-3 bg-white rounded-lg shadow-sm col-span-1"
            style={{ width: '32%' }}
        >
            <View className="flex-col items-center justify-center gap-y-2">
                <Location width={24} height={24} />
                <Text className="font-semibold text-center text-black">
                    {information[0]}
                </Text>
                <View className="flex-col items-center">
                    <Text className="text-xs text-center font-normal text-neutral-400">
                        {information[1]}
                    </Text>
                </View>
            </View>
        </View>
    );
};

export { EventLocation };
