import { Event } from '@generatesac/lib';

import { DAY_EPOCH_TIME } from '@/app/(design-system)/components/Calendar/Calendar';
import {
    EventPreview,
    EventSection
} from '@/app/(design-system)/components/Calendar/DayTimeSection';

export function parseData(start: number, end: number, data?: Event[]) {
    // Parsed data:
    const parsedData: EventSection[] = [];

    // Compute the number of dates between start and end:
    const numDates = Math.floor((end - start) / DAY_EPOCH_TIME);

    // For each day, create an instance in the data array:
    for (let i = 0; i < numDates; i++) {
        const additionalTime = i * DAY_EPOCH_TIME;

        // Parse the date string:
        const date = new Date(start + additionalTime)
            .toISOString()
            .split('T')[0];

        // Create a new EventSection:
        const eventSection: EventSection = {
            date: date,
            data: {}
        };

        // Add the EventSection to the parsed data:
        parsedData.push(eventSection);
    }

    // Loop over the data and add the events to the correct date:
    data?.forEach((event) => {
        // Parse the event date:
        const date = new Date(event.start_time);

        // Parse the event start date:
        const tz = Intl.DateTimeFormat().resolvedOptions().timeZone;
        const formattedDate = new Intl.DateTimeFormat('en-US', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            timeZone: tz
        }).format(date);
        const [month, day, year] = formattedDate.split('/');
        const isoFormattedDate = `${year}-${month}-${day}`;

        // Parse the event start time:
        const startTime = parseTime(date);

        // Find the index of the date in the parsed data:
        const index = parsedData.findIndex(
            (section) => section.date === isoFormattedDate
        );

        // Create an event preview:
        const eventPreview: EventPreview = {
            title: event.name,
            startTime: event.start_time,
            endTime: event.end_time,
            host: event.host,
            tags: []
        };

        // If the date is found, add the event to the data:
        if (index !== -1) {
            // Check if the start time already exists:
            if (parsedData[index].data[startTime]) {
                // Check to see if the events are overflowing:
                if (parsedData[index].data[startTime].events.length >= 3) {
                    parsedData[index].data[startTime].overflow = true;
                } else {
                    parsedData[index].data[startTime].events.push(eventPreview);
                }
            } else {
                // Create a new time:
                parsedData[index].data[startTime] = {
                    overflow: false,
                    events: [eventPreview]
                };
            }
        }
    });

    return parsedData;
}

function parseTime(date: Date) {
    const minutes =
        date.getMinutes() === 0 ? '00' : date.getMinutes().toString();
    const time = date.getHours().toString() + minutes;
    return parseInt(time);
}
