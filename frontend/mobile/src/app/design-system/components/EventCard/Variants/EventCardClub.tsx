import React from 'react';
import { TouchableOpacity } from 'react-native';

import { router } from 'expo-router';

import { Tag } from '@generatesac/lib';
import { Avatar, Image } from '@rneui/base';

import { Box, Text, createStyles } from '@/src/app/design-system';
import { createOptions, eventTime } from '@/src/utils/time';

import { EventTags } from '../EventCardTags/EventCardTags';
import { sharedStyles } from '../shared/styles';

interface EventCardClubProps {
    event: string;
    club: string;
    eventId: string;
    startTime: string;
    endTime: string;
    image: string;
    logo: string;
    tags?: Tag[];
}

export const EventCardClub: React.FC<EventCardClubProps> = ({
    event,
    club,
    eventId,
    startTime,
    endTime,
    image,
    logo,
    tags
}) => {
    return (
        <Box>
            <TouchableOpacity
                onPress={() => router.navigate(`/event/${eventId}`)}
            >
                <Box {...styles.cardContainer}>
                    <Box {...styles.cardSubContainer}>
                        <Image
                            containerStyle={{
                                flexGrow: 0,
                                borderRadius: 8,
                                aspectRatio: 1
                            }}
                            source={{ uri: image }}
                        />
                        <Box flex={3} flexDirection="column" gap="xxs">
                            <Text color="darkGray" variant="caption-1">
                                {eventTime(
                                    startTime,
                                    endTime,
                                    createOptions(
                                        'monthAndDate',
                                        'year',
                                        'start',
                                        'end'
                                    )
                                )}
                            </Text>
                            <Text variant="body-1">{event}</Text>
                            <Box {...sharedStyles.clubInfoContainer}>
                                <Text color="darkGray" variant="caption-1">
                                    Hosted by {club}
                                </Text>
                                <Avatar
                                    size={15}
                                    rounded
                                    source={{ uri: logo }}
                                />
                            </Box>
                            {tags && <EventTags tags={tags} />}
                        </Box>
                    </Box>
                </Box>
            </TouchableOpacity>
        </Box>
    );
};

const styles = createStyles({
    cardContainer: {
        shadowColor: 'black',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.1,
        shadowRadius: 2,
        backgroundColor: 'white',
        borderRadius: 'md'
    },
    cardSubContainer: {
        width: '100%',
        borderRadius: 'md',
        padding: 'm',
        backgroundColor: 'white',
        flexDirection: 'row',
        justifyContent: 'space-between',
        gap: 's'
    }
});
