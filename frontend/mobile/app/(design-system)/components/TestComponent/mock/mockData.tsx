import { Tag } from "@generatesac/lib";
import { EventSection } from "../../Calendar/mockData";

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

export const DAY = {
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
};

export async function getEvents(startTime: number, endTime: number) {
    let time = startTime;
    const increment = 1000 * 60 * 60 * 24;
    const events: EventSection[] = [];

    while (time <= endTime) {
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
            }
        });
        time += increment;
    }

    return events;
}