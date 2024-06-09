import { formatTime } from '@/utils/time';
import { Text } from '../Text/Text';
import { Box } from '@/app/(design-system)/components/Box/Box';
import DayTimeSection, { EventSection } from '@/app/(design-system)/components/Calendar/DayTimeSection';

type DayProps = {
    section: EventSection;
    error?: boolean;
};

function formatSectionHeader(date: string) {
    let addedZero = false;
    const { date: dayOfMonth, dayOfWeek } = formatTime(
        new Date(date + 'T00:00:00')
    );
    if (dayOfMonth.length < 2) {
        addedZero = true;
    }

    return `${dayOfWeek} ${addedZero ? '0' : ''}${dayOfMonth}`;
}

export default function Day({ section, error }: DayProps) {
    const keys = Object.keys(section.data).sort();

    console.log(error);

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
                <Box>
                    <Text textAlign="center" marginVertical="s">
                        {error ? "Oops, something went wrong!" : "No Events Today" }
                    </Text>
                </Box>
            )}
        </Box>
    );
};
