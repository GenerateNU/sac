import { Dimensions } from 'react-native';
import { StyleSheet } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { Tag } from '@generatesac/lib';
import { Image } from '@rneui/base';
import { Avatar } from '@rneui/base';

import { Box, Text } from '@/app/(design-system)';
import {
    calculateDuration,
    createOptions,
    eventTime,
    happeningNow
} from '@/utils/time';

import { EventTags } from './event-tag';

interface EventCardProps {
    event: string;
    club: string;
    startTime: Date;
    endTime: Date;
    tags?: Tag[];
    image: string;
    logo: string;
    variant: 'big' | 'small' | 'club' | 'calendar';
}

export const EventCard = ({
    event,
    club,
    startTime,
    endTime,
    image,
    logo,
    tags,
    variant
}: EventCardProps) => {
    const isHappening = happeningNow(startTime, endTime);
    const screenWidth = Dimensions.get('window').width;
    const boxWidth = screenWidth * 0.4;

    const bigVariant = () => (
        <Box margin="xl">
            <TouchableOpacity>
                <Box gap="xs">
                    <Image
                        containerStyle={styles.image}
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

    const smallVariant = () => (
        <Box width={boxWidth}>
            <TouchableOpacity>
                <Box flexDirection="column" gap="xs">
                    <Image
                        containerStyle={styles.image}
                        source={{ uri: image }}
                    />
                    <Text fontWeight="500">{event}</Text>
                    <Text variant="caption-1" color="darkGray">
                        {eventTime(
                            startTime,
                            endTime,
                            createOptions('dayOfWeek', 'monthAndDate')
                        )}
                    </Text>
                    <Text variant="caption-1" color="darkGray">
                        {club}
                    </Text>
                </Box>
            </TouchableOpacity>
        </Box>
    );

    const clubVariant = () => (
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

    const calendarVariant = () => (
        <Box margin="xl">
            <TouchableOpacity>
                {isHappening && (
                    <Box
                        zIndex={1}
                        marginRight="xl"
                        top={16}
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
                            borderRadius={4}
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
                    <Box width="100%" borderRadius={12} />
                    <Box
                        width="100%"
                        borderRadius={12}
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
                                    variant="caption-1"
                                    color="darkGray"
                                >{`${eventTime(startTime, endTime, createOptions('start', 'end'))} •`}</Text>
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
                                    <Text color="darkGray" variant="caption-1">
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

    const renderVariant = () => {
        switch (variant) {
            case 'big':
                return bigVariant();
            case 'small':
                return smallVariant();
            case 'club':
                return clubVariant();
            case 'calendar':
                return calendarVariant();
            default:
                return bigVariant();
        }
    };

    return renderVariant();
};

const styles = StyleSheet.create({
    image: {
        aspectRatio: 1,
        width: '100%',
        borderRadius: 15
    },
    clubImage: {
        flex: 1.15,
        borderRadius: 8,
        aspectRatio: 1
    },
    calendarImage: {
        flex: 1,
        borderRadius: 12
    }
});
