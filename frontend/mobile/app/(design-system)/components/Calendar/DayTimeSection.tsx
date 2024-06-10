import { Tag } from '@generatesac/lib';

import { Box } from '@/app/(design-system)/components/Box/Box';
import { Button } from '@/app/(design-system)/components/Button/Button';
import { EventCard } from '@/app/(design-system)/components/EventCard/EventCard';
import { Text } from '@/app/(design-system)/components/Text/Text';

export type EventSectionTimes = {
    [key: string]: EventSectionData;
};

export type EventSection = {
    date: string;
    data: EventSectionTimes;
};

export type EventPreview = {
    title: string;
    startTime: string;
    endTime: string;
    host: string;
    tags: Tag[];
};

export type EventSectionData = {
    overflow: boolean;
    events: EventPreview[];
};

export type DayTimeSectionProps = {
    time: number;
    data: EventSectionData;
};

function convertNumToTime(num: number) {
    let time = num;
    let meridiem = num === 23 ? 'AM' : 'PM';
    let minutes = num % 100;

    if (num !== 0 && (num < 1000 || num > 2300)) {
        throw new Error('Invalid time');
    }

    time = time / 100;

    if (num < 12) {
        time += 12;
        meridiem = 'AM';
    }
    if (num > 12) {
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
                                club={'Generate'}
                                clubId={event.host}
                                startTime={new Date(event.startTime)}
                                endTime={new Date(event.endTime)}
                                logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                                image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
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
