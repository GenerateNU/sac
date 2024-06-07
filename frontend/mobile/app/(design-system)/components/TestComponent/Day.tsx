import { forwardRef } from 'react';

import { formatTime } from '@/utils/time';
import { EventSection } from '../Calendar/mockData';
import { Text } from '../Text/Text';
import { Box } from '../Box/Box';
import AgendaTimeSection from '../Calendar/Agenda/AgendaTimeSection';

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

    return (
        <Box>
            <Text variant="body-1" fontWeight={'bold'} marginTop="l">
                {formatSectionHeader(section.date)}
            </Text>
            {keys.map((key, index) => {
                return (
                    <AgendaTimeSection
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
