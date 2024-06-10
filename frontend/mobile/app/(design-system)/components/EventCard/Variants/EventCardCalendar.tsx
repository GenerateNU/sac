import React from 'react';
import { TouchableOpacity } from 'react-native';

import { Tag } from '@generatesac/lib';
import { Avatar, Image } from '@rneui/base';

import { Box, Text, createStyles } from '@/app/(design-system)';
import { createOptions, eventTime, happeningNow } from '@/utils/time';

import { EventTags } from '../EventCardTags/EventCardTags';
import { sharedStyles } from '../shared/styles';

interface EventCardCalendarProps {
    event: string;
    club: string;
    startTime: Date;
    endTime: Date;
    image: string;
    logo: string;
    tags?: Tag[];
}

export const EventCardCalendar: React.FC<EventCardCalendarProps> = ({
    event,
    club,
    startTime,
    endTime,
    image,
    logo,
    tags
}) => {
    const isHappening = happeningNow(startTime, endTime);
    const isPast = endTime.getTime() < new Date().getTime();

    return (
        <Box position="relative">
            <TouchableOpacity>
                {isHappening && (
                    <Box {...styles.isHappeningContainer}>
                        <Box {...styles.isHappeningContent}>
                            <Text color="white" variant="body-1">
                                NOW
                            </Text>
                        </Box>
                    </Box>
                )}
                {isPast && <Box {...styles.pastEvent} />}
                <Box {...styles.cardContainer}>
                    <Box {...styles.cardSubContainer}>
                        <Image
                            containerStyle={{
                                flexGrow: 0,
                                borderRadius: 4,
                                aspectRatio: 1
                            }}
                            source={{ uri: image }}
                        />
                        <Box flex={2.5} flexDirection="column" gap="xxs">
                            <Text variant="body-1">{event}</Text>
                            <Box {...styles.cardContentContainer}>
                                <Text variant="caption-1" color="darkGray">
                                    {`${eventTime(
                                        startTime,
                                        endTime,
                                        createOptions('start', 'end')
                                    )} •`}
                                </Text>
                                <Box {...sharedStyles.clubInfoContainer}>
                                    <Avatar
                                        size={15}
                                        rounded
                                        source={{ uri: logo }}
                                    />
                                    <Text color="darkGray" variant="caption-1">
                                        {club}
                                    </Text>
                                </Box>
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
    calendarImage: {
        flex: 1,
        borderRadius: 'sm'
    },
    isHappeningContainer: {
        zIndex: 1,
        position: 'absolute',
        marginRight: 'xl',
        top: '-15%',
        right: '-5%',
        flexDirection: 'row',
        justifyContent: 'flex-end'
    },
    isHappeningContent: {
        paddingHorizontal: 's',
        paddingVertical: 'xxs',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 'sm',
        backgroundColor: 'darkRed',
        shadowColor: 'black',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.3,
        shadowRadius: 2
    },
    cardContainer: {
        borderRadius: 'md',
        shadowColor: 'black',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.2,
        shadowRadius: 8,
        backgroundColor: 'white'
    },
    cardSubContainer: {
        width: '100%',
        borderRadius: 'md',
        padding: 'm',
        backgroundColor: 'white',
        flexDirection: 'row',
        justifyContent: 'space-between',
        gap: 's'
    },
    cardContentContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        gap: 'xxs',
        flexWrap: 'wrap'
    },
    pastEvent: {
        position: 'absolute',
        width: '100%',
        height: '100%',
        backgroundColor: 'black',
        zIndex: 2,
        borderRadius: 'md',
        opacity: 0.15
    }
});
