import { TouchableOpacity } from 'react-native-gesture-handler';

import {
    faCalendarDays,
    faGlobe,
    faLocationDot
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';
import { EventType } from '@generatesac/lib';
import { Avatar } from '@rneui/base';

import { Box, Colors, SACColors, Text } from '@/src/app/design-system';
import { setClubId, setClubShouldPreview } from '@/src/store/slices/clubSlice';
import { useAppDispatch } from '@/src/store/store';
import { firstLetterUppercase } from '@/src/utils/string';
import { createOptions, eventTime } from '@/src/utils/time';

const eventTypeMappings: Record<EventType, string> = {
    'in_person': 'In Person',
    'virtual': 'Virtual',
    'hybrid': 'Hybrid'
}

interface OverviewProps {
    logo: string;
    eventName: string;
    club: string;
    clubId: string;
    startTime: string;
    endTime: string;
    location: string;
    type: EventType;
    color: SACColors;
}

export const Overview: React.FC<OverviewProps> = ({
    logo,
    clubId,
    eventName,
    color,
    club,
    startTime,
    endTime,
    location,
    type
}) => {
    const dispatch = useAppDispatch();

    return (
        <Box gap="m" flexDirection="column">
            <Box>
                <Text variant="header-1">{eventName}</Text>
                <TouchableOpacity
                    onPress={() => {
                        dispatch(setClubId(clubId));
                        dispatch(setClubShouldPreview(true));
                    }}
                >
                    <Box
                        marginTop="xs"
                        flexWrap="wrap"
                        flexDirection="row"
                        gap="xxs"
                        alignItems="center"
                        width="70%"
                    >
                        <Text color="darkGray">Hosted by {club}</Text>
                        <Avatar rounded source={{ uri: logo }} size={20} />
                    </Box>
                </TouchableOpacity>
            </Box>
            <Box flexDirection="column" gap="xs" width="90%">
                <Box flexDirection="row" gap="s" alignItems="center">
                    <FontAwesomeIcon
                        size={23}
                        color={Colors[color]}
                        icon={faCalendarDays}
                    />
                    <Text>
                        {eventTime(
                            startTime,
                            endTime,
                            createOptions(
                                'dayOfWeek',
                                'monthAndDate',
                                'start',
                                'end'
                            )
                        )}
                    </Text>
                </Box>
                <Box flexDirection="row" gap="s" alignItems="center">
                    <FontAwesomeIcon
                        size={23}
                        color={Colors[color]}
                        icon={faLocationDot}
                    />
                    <Text>{location}</Text>
                </Box>
                <Box flexDirection="row" gap="s" alignItems="center">
                    <FontAwesomeIcon
                        size={23}
                        color={Colors[color]}
                        icon={faGlobe}
                    />
                    <Text>{eventTypeMappings[type]}</Text>
                </Box>
            </Box>
        </Box>
    );
};
