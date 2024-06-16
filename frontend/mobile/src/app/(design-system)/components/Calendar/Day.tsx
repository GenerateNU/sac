import { Box } from '@/src/app/(design-system)/components/Box/Box';
import DayTimeSection, {
    EventSection
} from '@/src/app/(design-system)/components/Calendar/DayTimeSection';
import { Text } from '@/src/app/(design-system)/components/Text/Text';
import NoEventsIcon from '@/src/assets/svg/NoSearchResult.svg';
import { formatTime } from '@/src/utils/time';

type DayProps = {
    section: EventSection;
    error?: boolean;
};

function formatSectionHeader(date: string) {
    let addedZero = false;
    const { date: dayOfMonth, dayOfWeek } = formatTime(
        date + 'T00:00:00'
    );
    if (dayOfMonth.length < 2) {
        addedZero = true;
    }

    return `${dayOfWeek} ${addedZero ? '0' : ''}${dayOfMonth}`;
}

export default function Day({ section, error }: DayProps) {
    const keys = Object.keys(section.data).sort();
    return (
        <Box>
            <Text variant="body-1" fontWeight={'bold'} marginTop="l">
                {formatSectionHeader(section.date)}
            </Text>
            {keys.map((key, index) => {
                return (
                    <DayTimeSection
                        key={`${key}${index}`}
                        time={parseInt(key)}
                        data={section.data[key]}
                    />
                );
            })}
            {keys.length === 0 && (
                <Box width="100%" alignItems="center">
                    {!error && <NoEventsIcon height={175} />}
                    <Text textAlign="center" fontWeight="bold" color="darkGray">
                        {error
                            ? 'Oops, something went wrong!'
                            : 'No events today'}
                    </Text>
                    <Text color="darkGray" variant="body-2" marginTop="s">
                        {error
                            ? 'Please try again later'
                            : 'Make sure to check again soon!'}
                    </Text>
                </Box>
            )}
        </Box>
    );
}
