import React from 'react';
import { Dimensions } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import { Image } from '@rneui/base';

import { Box, Text } from '@/src/app/(design-system)';
import { createOptions, eventTime } from '@/src/utils/time';

interface EventCardSmallProps {
    event: string;
    club: string;
    clubId: string;
    startTime: Date;
    endTime: Date;
    image: string;
}

export const EventCardSmall: React.FC<EventCardSmallProps> = ({
    event,
    club,
    clubId,
    startTime,
    endTime,
    image
}) => {
    const screenWidth = Dimensions.get('window').width;
    const boxWidth = screenWidth * 0.4;

    return (
        <Box width={boxWidth}>
            <TouchableOpacity
                onPress={() => router.navigate(`/event/${clubId}`)}
            >
                <Box flexDirection="column" gap="xs">
                    <Image
                        containerStyle={{
                            aspectRatio: 1,
                            width: '100%',
                            borderRadius: 15
                        }}
                        source={{ uri: image }}
                    />
                    <Text fontWeight="500">{event}</Text>
                    <Text variant="caption-3" color="darkGray">
                        {eventTime(
                            startTime,
                            endTime,
                            createOptions('dayOfWeek', 'monthAndDate')
                        )}
                    </Text>
                    <Text variant="caption-3" color="darkGray">
                        {club}
                    </Text>
                </Box>
            </TouchableOpacity>
        </Box>
    );
};
