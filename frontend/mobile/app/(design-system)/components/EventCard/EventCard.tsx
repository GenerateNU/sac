import { Tag } from '@generatesac/lib';

import { EventCardBig } from './EventCardVariants/EventCardBig';
import { EventCardCalendar } from './EventCardVariants/EventCardCalendar';
import { EventCardClub } from './EventCardVariants/EventCardClub';
import { EventCardSmall } from './EventCardVariants/EventCardSmall';
import { Box } from '../Box/Box';

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
    switch (variant) {
        case 'big':
            return (
                <Box marginVertical={'m'}>
                <EventCardBig
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                />
                </Box>
            );
        case 'small':
            return (
                <Box marginVertical={'m'}>
                <EventCardSmall
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                />
                </Box>
            );
        case 'club':
            return (
                <Box marginVertical={'m'}>
                <EventCardClub
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                    logo={logo}
                    tags={tags ? tags : []}
                />
                </Box>
            );
        case 'calendar':
            return (
                <EventCardCalendar
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                    logo={logo}
                    tags={tags ? tags : []}
                />
            );
        default:
            return (
                <EventCardBig
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                />
            );
    }
};
