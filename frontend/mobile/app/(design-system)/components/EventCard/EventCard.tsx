import { Tag } from '@generatesac/lib';

import { EventCardBig } from './EventCardBig';
import { EventCardCalendar } from './EventCardCalendar';
import { EventCardClub } from './EventCardClub';
import { EventCardSmall } from './EventCardSmall';

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
                <EventCardBig
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                />
            );
        case 'small':
            return (
                <EventCardSmall
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                />
            );
        case 'club':
            return (
                <EventCardClub
                    event={event}
                    club={club}
                    startTime={startTime}
                    endTime={endTime}
                    image={image}
                    logo={logo}
                    tags={tags ? tags : []}
                />
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
