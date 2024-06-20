// EventCardBig.tsx
import React from 'react';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import { Image } from '@rneui/base';

import { Box, Text } from '@/src/app/design-system';
import { createOptions, eventTime } from '@/src/utils/time';

interface EventCardBigProps {
    event: string;
    club: string;
    eventId: string;
    startTime: string;
    endTime: string;
    image: string;
}

export const EventCardBig: React.FC<EventCardBigProps> = ({
    event,
    club,
    eventId,
    startTime,
    endTime,
    image
}) => {
    return (
        <Box>
            <TouchableOpacity
                onPress={() => router.navigate(`/app/event/${eventId}`)}
            >
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
