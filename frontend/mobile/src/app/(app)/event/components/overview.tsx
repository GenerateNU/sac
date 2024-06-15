import {
    faCalendarDays,
    faGlobe,
    faLocationDot
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';
import { EventType } from '@generatesac/lib';
import { Avatar } from '@rneui/base';

import { Box, Colors, SACColors, Text } from '@/src/app/(design-system)';
import { firstLetterUppercase } from '@/src/utils/string';
import { createOptions, eventTime } from '@/src/utils/time';

interface OverviewProps {
    logo: string;
    eventName: string;
    club: string;
    clubId: string;
    startTime: Date;
    endTime: Date;
    location: string;
    type: EventType;
    color: SACColors;
}

export const Overview: React.FC<OverviewProps> = ({
    logo,
    // clubId,
    eventName,
    color,
    club,
    startTime,
    endTime,
    location,
    type
}) => {
    return (
        <Box
            flexWrap="wrap"
            gap="m"
            flexDirection="column"
            alignItems="flex-start"
        >
            <Box>
                <Text variant="header-1">{eventName}</Text>
                <Box
                    marginTop="xs"
                    flexDirection="row"
                    gap="xxs"
                    alignItems="center"
                >
                    <Text color="darkGray">Hosted by {club}</Text>
                    <Avatar rounded source={{ uri: logo }} size={20} />
                </Box>
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
                    <Text>{firstLetterUppercase(type)}</Text>
                </Box>
            </Box>
        </Box>
    );
};
