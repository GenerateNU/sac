import { Tag } from '@generatesac/lib';

import { Box } from '@/src/app/design-system/components/Box/Box';
import { Button } from '@/src/app/design-system/components/Button/Button';
import { EventCard } from '@/src/app/design-system/components/EventCard/EventCard';
import { Text } from '@/src/app/design-system/components/Text/Text';

export type EventSectionTimes = {
    [key: string]: EventSectionData;
};

export type EventSection = {
    date: string;
    data: EventSectionTimes;
};

export type EventCalendarPreview = {
    id: string;
    title: string;
    startTime: string;
    endTime: string;
    host_name: string;
    host_logo: string;
    image: string;
    tags: Tag[];
};

export type EventSectionData = {
    overflow: boolean;
    events: EventCalendarPreview[];
};

export type DayTimeSectionProps = {
    time: number;
    data: EventSectionData;
};

function convertNumToTime(num: number) {
    let time = num;
    let meridiem = num === 23 ? 'AM' : 'PM';
    let minutes = num % 100;

    if (num !== 0 && (num < 0 || num > 2359)) {
        throw new Error('Invalid time');
    }

    time = time / 100;
    if (time < 12) {
        meridiem = 'AM';
    }
    if (time > 12) {
        time -= 12;
    }

    return {
        hour: Math.floor(time),
        minute: minutes === 0 ? undefined : minutes.toString(),
        meridiem
    };
}

function TimeHeader({ time }: { time: number }) {
    const { hour, minute, meridiem } = convertNumToTime(time);
    return (
        <Box flexDirection="row" gap="s" alignItems="center" marginVertical="s">
            <Text variant="caption-1">
                {hour}
                {minute ? `:${minute}` : ''} {meridiem}
            </Text>
            <Box flex={1} backgroundColor="gray" height={0.75} />
        </Box>
    );
}

export default function DayTimeSection({ time, data }: DayTimeSectionProps) {
    return (
        <Box>
            <Box>
                <TimeHeader time={time} />
            </Box>
            <Box alignItems="center" flexDirection="column" gap="s">
                {data.events.map((event, index) => {
                    return (
                        <Box
                            marginVertical="xxs"
                            marginHorizontal="s"
                            key={`${event.title}${index}`}
                        >
                            <EventCard
                                event={event.title}
                                variant="calendar"
                                tags={event.tags}
                                club={event.host_name}
                                logo={event.host_logo}
                                eventId={event.id}
                                startTime={event.startTime}
                                endTime={event.endTime}
                                image={event.image}
                            />
                        </Box>
                    );
                })}
                {data.overflow && (
                    <Button
                        variant="standardButton"
                        color="transparent"
                        size="sm"
                        onPress={() => {}} // TODO: LINK TO EVENT LIST PAGE
                    >
                        See More
                    </Button>
                )}
            </Box>
        </Box>
    );
}
