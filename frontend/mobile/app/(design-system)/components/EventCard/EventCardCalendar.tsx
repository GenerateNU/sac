import React from 'react';
import { StyleSheet, TouchableOpacity } from 'react-native';

import { Tag } from '@generatesac/lib';
import { Avatar, Image } from '@rneui/base';

import { createOptions, eventTime, happeningNow } from '@/utils/time';

import { Box } from '../Box/Box';
import { Text } from '../Text/Text';
import { EventTags } from './EventTag';

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

    return (
        <Box margin="xl">
            <TouchableOpacity>
                {isHappening && (
                    <Box
                        zIndex={1}
                        position="absolute"
                        marginRight="xl"
                        top={-16}
                        right={6}
                        flexDirection="row"
                        flex={1}
                        justifyContent="flex-end"
                    >
                        <Box
                            width={60}
                            height={30}
                            flexDirection="column"
                            justifyContent="center"
                            alignItems="center"
                            borderRadius="small"
                            backgroundColor="darkRed"
                            shadowColor="black"
                            shadowOffset={{ width: 0, height: 1 }}
                            shadowOpacity={0.3}
                            shadowRadius={2}
                        >
                            <Text color="white" variant="body-1">
                                NOW
                            </Text>
                        </Box>
                    </Box>
                )}
                <Box
                    shadowColor="black"
                    shadowOffset={{ width: 0, height: 1 }}
                    shadowOpacity={0.1}
                    shadowRadius={2}
                >
                    <Box width="100%" borderRadius="medium" />
                    <Box
                        width="100%"
                        borderRadius="medium"
                        padding="m"
                        backgroundColor="white"
                        flexDirection="row"
                        justifyContent="space-between"
                        gap="s"
                    >
                        <Image
                            containerStyle={styles.calendarImage}
                            source={{ uri: image }}
                        />
                        <Box flex={2.5} flexDirection="column" gap="xxs">
                            <Text variant="body-1">{event}</Text>
                            <Box
                                flexDirection="row"
                                alignItems="center"
                                gap="xxs"
                                flexWrap="wrap"
                            >
                                <Text
                                    variant="caption-3"
                                    color="darkGray"
                                >{`${eventTime(
                                    startTime,
                                    endTime,
                                    createOptions('start', 'end')
                                )} •`}</Text>
                                <Box
                                    flexDirection="row"
                                    alignItems="center"
                                    gap="xxs"
                                >
                                    <Avatar
                                        size={15}
                                        rounded
                                        source={{ uri: logo }}
                                    />
                                    <Text color="darkGray" variant="caption-3">
                                        {club}
                                    </Text>
                                </Box>
                            </Box>
                            {tags && <EventTags tags={tags} variant="plus" />}
                        </Box>
                    </Box>
                </Box>
            </TouchableOpacity>
        </Box>
    );
};

const styles = StyleSheet.create({
    calendarImage: {
        flex: 1,
        borderRadius: 12
    }
});
