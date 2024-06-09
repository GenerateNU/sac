import { Event } from "@generatesac/lib";
import { DAY_EPOCH_TIME } from "@/app/(design-system)/components/Calendar/Calendar";
import { EventPreview, EventSection } from "@/app/(design-system)/components/Calendar/DayTimeSection";

export const mockData: Event[] = 
    [
        {
            "id": "ae6e7813-7f9a-48d9-a83d-8d313268617b",
            "created_at": "2024-06-07T23:36:29.326493-04:00",
            "updated_at": "2024-06-07T23:36:29.326493-04:00",
            "name": "All Hands",
            "preview": "Join us for mid-semester presentations on all the amazing projects!",
            "description": "Wowza",
            "host": "464f73bc-198b-48a1-8c38-77bd9ed173be",
            "event_type": "hybrid",
            "location": "ISEC Auditorium",
            "link": "http://www.applytogenerate.com",
            "is_public": true,
            "is_draft": false,
            "is_archived": false,
            "start_time": "2024-02-27T14:00:00-05:00",
            "end_time": "2024-02-27T16:00:00-05:00"
        },
        {
            "id": "4b8123f8-82be-4096-b5c9-1c325d97b733",
            "created_at": "2024-06-07T23:42:15.588637-04:00",
            "updated_at": "2024-06-07T23:42:15.588637-04:00",
            "name": "All Hands",
            "preview": "Join us for mid-semester presentations on all the amazing projects!",
            "description": "Wowza",
            "host": "464f73bc-198b-48a1-8c38-77bd9ed173be",
            "event_type": "hybrid",
            "location": "ISEC Auditorium",
            "link": "http://www.applytogenerate.com",
            "is_public": true,
            "is_draft": false,
            "is_archived": false,
            "start_time": "2024-02-27T14:00:00-05:00",
            "end_time": "2024-02-27T16:00:00-05:00"
        }
    ]

export function parseData(start: number, end: number, data?: Event[]) {
    // Parsed data:
    const parsedData: EventSection[] = [];

    // Compute the number of dates between start and end:
    const numDates = Math.floor((end - start) / DAY_EPOCH_TIME);

    // For each day, create an instance in the data array:
    for (let i = 0 ; i < numDates ; i++) {
        const additionalTime = i * DAY_EPOCH_TIME;

        // Parse the date string:
        const date = new Date(start + additionalTime).toISOString().split('T')[0]; 

        // Create a new EventSection:
        const eventSection: EventSection = {
            date: date,
            data: {}
        };

        // Add the EventSection to the parsed data:
        parsedData.push(eventSection);
    }

    // Loop over the data and add the events to the correct date:
    data?.forEach(( event ) => {
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
        const index = parsedData.findIndex((section) => section.date === isoFormattedDate);

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
    const minutes = date.getMinutes() === 0 ? '00' : date.getMinutes().toString();
    const time = date.getHours().toString() + minutes;
    return parseInt(time);
}

// export const DAY = {
//     date: '2024-06-05',
//     data: {
//         2000: {
//             overflow: true,
//             events: [
//                 {
//                     title: 'Gym Workout',
//                     startTime: '12am',
//                     endTime: '1am',
//                     host: 'Generate',
//                     tags: tags
//                 },
//                 {
//                     title: 'Gym Workout',
//                     startTime: '12am',
//                     endTime: '1am',
//                     host: 'Generate',
//                     tags: tags
//                 },
//                 {
//                     title: 'Gym Workout',
//                     startTime: '12am',
//                     endTime: '1am',
//                     host: 'Generate',
//                     tags: tags
//                 }
//             ]
//         },
//         2015: {
//             overflow: true,
//             events: [
//                 {
//                     title: 'Gym Workout',
//                     startTime: '12am',
//                     endTime: '1am',
//                     host: 'Generate',
//                     tags: tags
//                 },
//                 {
//                     title: 'Gym Workout',
//                     startTime: '12am',
//                     endTime: '1am',
//                     host: 'Generate',
//                     tags: tags
//                 }
//             ]
//         },
//         2100: {
//             overflow: true,
//             events: [
//                 {
//                     title: 'Gym Workout',
//                     startTime: '12am',
//                     endTime: '1am',
//                     host: 'Generate',
//                     tags: tags
//                 }
//             ]
//         }
//     }
// };

// export async function getEvents(startTime: number, endTime: number) {
//     let time = startTime;
//     const increment = 1000 * 60 * 60 * 24;
//     const events: EventSection[] = [];

//     while (time <= endTime) {
//         const date = new Date(time).toISOString().split('T')[0];
//         events.push({
//             date,
//             data: {
//                 2000: {
//                     overflow: true,
//                     events: [
//                         {
//                             title: 'Gym Workout',
//                             startTime: date,
//                             endTime: new Date(
//                                 new Date(date).getTime() + 1000 * 60 * 60
//                             ).toISOString(),
//                             host: 'Generate',
//                             tags: tags
//                         }
//                     ]
//                 },
//             }
//         });
//         time += increment;
//     }

//     return events;
// }