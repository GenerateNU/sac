import React from 'react';
import { StyleSheet, TouchableOpacity } from 'react-native';

import { Tag } from '@generatesac/lib';
import { Avatar, Image } from '@rneui/base';

import { Box, Text } from '@/app/(design-system)';
import { createOptions, eventTime } from '@/utils/time';

import { EventTags } from './event-tag';

interface EventCardClubProps {
    event: string;
    club: string;
    startTime: Date;
    endTime: Date;
    image: string;
    logo: string;
    tags?: Tag[];
}

export const EventCardClub: React.FC<EventCardClubProps> = ({
    event,
    club,
    startTime,
    endTime,
    image,
    logo,
    tags
}) => {
    return (
        <Box margin="xl">
            <TouchableOpacity>
                <Box
                    shadowColor="black"
                    shadowOffset={{ width: 0, height: 1 }}
                    shadowOpacity={0.1}
                    shadowRadius={2}
                >
                    <Box
                        width="100%"
                        borderRadius={8}
                        padding="m"
                        backgroundColor="white"
                        flexDirection="row"
                        justifyContent="space-between"
                        gap="s"
                    >
                        <Image
                            containerStyle={styles.clubImage}
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
                            <Box
                                flexDirection="row"
                                alignItems="center"
                                gap="xxs"
                            >
                                <Text color="darkGray" variant="caption-1">
                                    Hosted by {club}
                                </Text>
                                <Avatar
                                    size={15}
                                    rounded
                                    source={{ uri: logo }}
                                />
                            </Box>
                            {tags && (
                                <EventTags tags={tags} variant="limited" />
                            )}
                        </Box>
                    </Box>
                </Box>
            </TouchableOpacity>
        </Box>
    );
};

const styles = StyleSheet.create({
    clubImage: {
        flex: 1.15,
        borderRadius: 8,
        aspectRatio: 1
    }
});
