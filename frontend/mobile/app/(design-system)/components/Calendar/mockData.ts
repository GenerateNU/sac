/**
 * TODO:
 * - Change the background color the app to white
 */
import { Tag } from '@generatesac/lib';

const tags: Tag[] = [
    {
        id: '1',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Software',
        category_id: '1'
    },
    {
        id: '2',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Free Food',
        category_id: '2'
    },
    {
        id: '3',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Panel Discussion',
        category_id: '3'
    },
    {
        id: '4',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Seminar',
        category_id: '4'
    },
    {
        id: '5',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Hackathon',
        category_id: '5'
    }
];

export type EventSection = {
    date: string;
    data: EventSectionTimes;
};

export type EventSectionTimes = {
    [key: string]: EventSectionData;
};

export type EventSectionData = {
    overflow: boolean;
    events: EventPreview[];
};

export type EventPreview = {
    title: string;
    startTime: string;
    endTime: string;
    host: string;
    tags: Tag[];
};

async function generateEvents(startTime: number, dateSpan: number) {
    let time = startTime;
    const increment = 1000 * 60 * 60 * 24;
    const events: EventSection[] = [];

    while (time <= startTime + dateSpan) {
        const date = new Date(time).toISOString().split('T')[0];
        events.push({
            date,
            data: {
                2000: {
                    overflow: true,
                    events: [
                        {
                            title: 'Gym Workout',
                            startTime: date,
                            endTime: new Date(
                                new Date(date).getTime() + 1000 * 60 * 60
                            ).toISOString(),
                            host: 'Generate',
                            tags: tags
                        }
                    ]
                },
                2015: {
                    overflow: false,
                    events: [
                        {
                            title: 'Gym Workout',
                            startTime: date,
                            endTime: new Date(
                                new Date(date).getTime() + 1000 * 60 * 60
                            ).toISOString(),
                            host: 'Generate',
                            tags: tags
                        }
                    ]
                }
            }
        });
        time += increment;
    }

    return events;
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
//                     tags: tags,
//                 }]
//             }
//         })
//     }

//     return { events, startTime };
// }

export async function fetchEvents(startTime: number, dateSpan: number) {
    const events = await generateEvents(startTime, dateSpan);
    return events;
}

export const mockEvents: EventSection[] = [
    {
        date: '2024-06-03',
        data: {}
    },
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
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            },
            2015: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            },
            2100: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            }
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
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            },
            1800: {
                overflow: false,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            },
            2100: {
                overflow: false,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            }
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
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            },
            2015: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            },
            2100: {
                overflow: true,
                events: [
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    },
                    {
                        title: 'Gym Workout',
                        startTime: '12am',
                        endTime: '1am',
                        host: 'Generate',
                        tags: tags
                    }
                ]
            }
        }
    }
];
