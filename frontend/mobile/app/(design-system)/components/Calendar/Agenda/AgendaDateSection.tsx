import { formatTime } from "@/utils/time";
import { Box } from "../../Box/Box";
import { Text } from "../../Text/Text";
import { EventSection } from "../mockData"
import AgendaTimeSection from "./AgendaTimeSection";
import { forwardRef, useEffect } from "react";

type AgendaSectionProps = {
    section: EventSection;
}

function formatSectionHeader(date: string) {
    let addedZero = false;
    const { date: dayOfMonth, dayOfWeek } = formatTime(new Date(date + 'T00:00:00'));
    if (dayOfMonth.length < 2) {
        addedZero = true;
    }

    return `${dayOfWeek} ${addedZero ? '0' : ''}${dayOfMonth}`;
}

export const AgendaDateSection = forwardRef(function AgendaDateSection({ section }: AgendaSectionProps, ref) {
    const keys = Object.keys(section.data).sort();

    return (
        <Box flex={1}>
            <Text 
                variant="body-1"
                fontWeight={'bold'}
                marginTop='l'
                ref={ref}
            >
                {formatSectionHeader(section.date)}
            </Text>
            {
                keys.map((key, index) => {
                    return (
                        <AgendaTimeSection 
                            key={`${key}${index}`}
                            time={parseInt(key)}
                            data={section.data[key]} />
                    )
                })
            }
            {keys.length === 0 && 
                <Text textAlign="center" marginVertical="s">
                    No Events Today
                </Text>}
        </Box>
    )
});