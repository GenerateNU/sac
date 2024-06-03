import { Box } from "../../Box/Box";
import { Button } from "../../Button/Button";
import { EventCard } from "../../EventCard/EventCard";
import { Text } from "../../Text/Text";
import { EventSectionData } from "../mockData"

export type AgendaTimeSectionProps = {
    time: number;
    data: EventSectionData;
}

function convertNumToTime(num: number) {
    let time = num;
    let meridiem = num == 23 ? 'AM' : 'PM';
    let minutes = num % 100;

    if (num != 0 && (num < 1000 || num > 2300)) {
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
        minute: minutes == 0 ? undefined : minutes.toString(),
        meridiem
    }
}

function TimeHeader ({ time } : { time: number }) {
    const { hour, minute, meridiem } = convertNumToTime(time);
    return (
        <Box flexDirection="row" gap="s" alignItems="center" marginVertical="s">
            <Text variant='caption-1'>
                {hour}{minute ? `:${minute}` : ''} {meridiem}
            </Text>
            <Box borderBottomColor='gray' borderBottomWidth={1} flex={1} />
        </Box>
    )
}

export default function AgendaTimeSection({ time, data }: AgendaTimeSectionProps ) {
    return (
        <Box flex={1}>
            <Box flex={1}>
                <TimeHeader time={time} />
            </Box>
            <Box flex={1} alignItems="center">
                {data.events.map((event, index) => {
                    return (
                        <Box marginVertical="xxs" key={`${event.title}${index}`}>
                            <EventCard
                                event={event.title}
                                variant="calendar"
                                tags={[]}
                                club={event.host}
                                startTime={new Date()}
                                endTime={new Date()}
                                logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                                image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                            />
                        </Box>
                    )
                })}
                {data.overflow && (
                    <Button 
                        variant='standardButton'
                        color='white'
                        size='md'
                        onPress={() => {}} // TODO: LINK TO EVENT LIST PAGE
                    >
                        See More
                    </Button>
                )}
            </Box>
        </Box>
    )
}