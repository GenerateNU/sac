import { Dimensions } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { Tag } from '@generatesac/lib';

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
    image?: string;
    logo?: string;
    variant: 'big' | 'small' | 'club' | 'calendar';
}

export const EventCard = ({
    event,
    club,
    startTime,
    endTime,
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
                    <Box
                        aspectRatio={1}
                        width={'100%'}
                        borderRadius={15}
                        backgroundColor="darkGray"
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
                    <Box
                        aspectRatio={1}
                        borderRadius={15}
                        backgroundColor="darkGray"
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

    const horizontalVariant = () => (
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
                        <Box
                            flex={1.15}
                            borderRadius={8}
                            aspectRatio={1}
                            backgroundColor="darkGray"
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
                                <Box
                                    width={14}
                                    aspectRatio={1}
                                    backgroundColor="darkGray"
                                    borderRadius={50}
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
                        <Box
                            flex={1}
                            borderRadius={12}
                            backgroundColor="darkGray"
                        />
                        <Box flex={2.5} flexDirection="column" gap="xxs">
                            <Text variant="body-1">{event}</Text>
                            <Box
                                flexDirection="row"
                                alignItems="center"
                                gap="xxs"
                            >
                                <Text
                                    variant="caption-1"
                                    color="darkGray"
                                >{`${eventTime(startTime, endTime, createOptions('start', 'end'))} •`}</Text>
                                <Box
                                    width={14}
                                    aspectRatio={1}
                                    backgroundColor="darkGray"
                                    borderRadius={50}
                                />
                                <Text color="darkGray" variant="caption-1">
                                    {club}
                                </Text>
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
                return horizontalVariant();
            case 'calendar':
                return calendarVariant();
            default:
                return bigVariant();
        }
    };

    return renderVariant();
};
