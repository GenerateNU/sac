/**
 * TODO:
 * - Change the background color the app to white
 */

export type EventSection = {
    date: string;
    data: EventSectionTimes;
}

export type EventSectionTimes = {
    [key: string]: EventSectionData
}

export type EventSectionData = {
    overflow: boolean
    events: EventPreview[];
}

export type EventPreview = {
    title: string;
    startTime: string;
    endTime: string;
    host: string;
    tags: string[];
}

export async function fetchDateRange(start: string, end: string) {

}

// function generateEvents(endTime: number) {
//     let startTime = endTime;
//     const events: EventSection[] = [];
//     const randomNum = Math.floor(Math.random() * 10) + 1

//     for (let i = 0; i < randomNum; i++) {
//         startTime += 1000 * 60 * 60 * 24;
//         const dateStart = new Date(startTime);

//         events.push({
//             date: dateStart.toISOString().split('T')[0],
//             data: {
//                 2000: [{
//                     title: 'Gym Workout',
//                     startTime: dateStart.toISOString(),
//                     endTime: dateStart.toISOString(),
//                     host: 'Generate',
//                     tags: ['Fitness'],
//                 }]
//             }
//         })
//     }
    
//     return { events, startTime };
// }

export async function fetchEvents(time: number) {
    const events = await mockEvents;
    return events
}

export const mockEvents: EventSection[] = [
    {
        date: '2024-06-05',
        data: {
            2000: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                ],
            },
            2015: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                ],
            },
            2100: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    }
                ],
            },
        }
    },
    {
        date: '2024-06-07',
        data: {
            1900: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                ],
            },
            1800: {
                overflow: false,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    }
                ],
            },
            2100: {
                overflow: false,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    }
                ],
            },
        }
    },
    {
        date: '2024-06-08',
        data: {
            2000: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    }
                ],
            },
            2015: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    }
                ],
            },
            2100: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: ['Fitness'],
                    }
                ],
            },
        }
    }
]