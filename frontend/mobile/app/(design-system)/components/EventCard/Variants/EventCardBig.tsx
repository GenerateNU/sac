// EventCardBig.tsx
import React from 'react';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { Image } from '@rneui/base';

import { Box, Text } from '@/app/(design-system)';
import { calculateDuration, createOptions, eventTime } from '@/utils/time';
import { router } from 'expo-router';

interface EventCardBigProps {
    event: string;
    club: string;
    clubId: string;
    startTime: Date;
    endTime: Date;
    image: string;
}

export const EventCardBig: React.FC<EventCardBigProps> = ({
    event,
    club,
    clubId,
    startTime,
    endTime,
    image
}) => {
    return (
        <Box>
            <TouchableOpacity onPress={() => router.navigate(`/event/${clubId}`)}>
                <Box gap="xs">
                    <Image
                        containerStyle={{
                            aspectRatio: 1,
                            width: '100%',
                            borderRadius: 15
                        }}
                        source={{ uri: image }}
                    />
                    <Box
                        flexDirection="row"
                        justifyContent="space-between"
                        width="100%"
                    >
                        <Text variant="subheader-2">{event}</Text>
                        <Text />
                        <Text>{calculateDuration(startTime, endTime)}</Text>
                    </Box>
                    <Text color="darkGray">{club}</Text>
                    <Box flexDirection="row" gap="xs">
                        <Text color="darkGray">
                            {eventTime(
                                startTime,
                                endTime,
                                createOptions('dayOfWeek', 'monthAndDate')
                            )}
                        </Text>
                        <Text fontWeight="400" color="black">
                            {eventTime(
                                startTime,
                                endTime,
                                createOptions('start')
                            )}
                        </Text>
                    </Box>
                </Box>
            </TouchableOpacity>
        </Box>
    );
};
