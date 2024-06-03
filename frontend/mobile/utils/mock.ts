export type MarkedDates = {
    [key: string]: { marked?: boolean };
};

export type EventInformation = {
    hour: string;
    duration: string;
    title: string;
    location: string;
};

export const dates = [
    '2024-04-01',
    '2024-04-02',
    '2024-04-03',
    '2024-04-04',
    '2024-04-05',
    '2024-04-06',
    '2024-04-07',
    '2024-04-08',
    '2024-04-09',
    '2024-04-10',
    '2024-04-11',
    '2024-04-12',
    '2024-04-13',
    '2024-04-14',
    '2024-04-15',
    '2024-04-16',
    '2024-04-17',
    '2024-04-18',
    '2024-04-19',
    '2024-04-20',
    '2024-04-21',
    '2024-04-22',
    '2024-04-23',
    '2024-04-24',
    '2024-04-25',
    '2024-04-26',
    '2024-04-27',
    '2024-04-28',
    '2024-04-29',
    '2024-04-30'
];

export const agendaItems = [
    {
        title: dates[0],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Club GM',
                location: 'West Village G 102'
            }
        ]
    },
    {
        title: dates[1],
        data: [
            {
                hour: '4pm',
                duration: '1h',
                title: 'Intramural game',
                location: 'Marino Center'
            },
            {
                hour: '5pm',
                duration: '1h',
                title: 'Generate Team Meeting',
                location: 'Sherman Center'
            },
            {
                hour: '5pm',
                duration: '1h',
                title: 'Other commitment',
                location: 'Online'
            }
        ]
    },
    {
        title: dates[2],
        data: [
            {
                hour: '1pm',
                duration: '1h',
                title: 'Dance Crew Rehearsal',
                location: 'Curry Ballroom'
            },
            {
                hour: '2pm',
                duration: '1h',
                title: 'Intramural game',
                location: 'Carter Field 1'
            },
            {
                hour: '3pm',
                duration: '1h',
                title: 'Club GM',
                location: 'West Village G 102'
            }
        ]
    },
    {
        title: dates[3],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Ashtanga Yoga',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[5],
        data: [
            {
                hour: '11pm',
                duration: '1h',
                title: 'Generate Team Meeting',
                location: 'Big Sherm'
            },
            {
                hour: '12pm',
                duration: '1h',
                title: 'Running Group',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[6],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Intramural game',
                location: 'Carter Field 1'
            }
        ]
    },
    {
        title: dates[7],
        data: []
    },
    {
        title: dates[8],
        data: [
            {
                hour: '9pm',
                duration: '1h',
                title: 'Cheese Club GM',
                location: 'Richards 325'
            },
            {
                hour: '10pm',
                duration: '1h',
                title: 'Club GM',
                location: 'Ryder Hall 289'
            },
            {
                hour: '11pm',
                duration: '1h',
                title: 'Intramural game',
                location: 'Carter Field 1'
            },
            {
                hour: '12pm',
                duration: '1h',
                title: 'Running Group',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[9],
        data: [
            {
                hour: '1pm',
                duration: '1h',
                title: 'Intramural game',
                location: 'Carter Field 1'
            },
            {
                hour: '2pm',
                duration: '1h',
                title: 'Club GM',
                location: 'West Village G 102'
            },
            {
                hour: '3pm',
                duration: '1h',
                title: 'Marino Yoga',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[10],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[11],
        data: [
            {
                hour: '1pm',
                duration: '1h',
                title: 'Intramural game',
                location: 'Carter Field 1'
            },
            {
                hour: '2pm',
                duration: '1h',
                title: 'Club GM',
                location: 'Ryder Hall 289'
            },
            {
                hour: '3pm',
                duration: '1h',
                title: 'Marino Yoga',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[12],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[15],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[16],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[18],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            },
            {
                hour: '1pm',
                duration: '1h',
                title: 'Intramural game',
                location: 'Carter Field 1'
            },
            {
                hour: '2pm',
                duration: '1h',
                title: 'Club GM',
                location: 'West Village G 102'
            },
            {
                hour: '3pm',
                duration: '1h',
                title: 'Marino Yoga',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[19],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[20],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[22],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[24],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
    {
        title: dates[25],
        data: [
            {
                hour: '12am',
                duration: '1h',
                title: 'Gym Workout',
                location: 'Marino Center'
            }
        ]
    },
];

export function getMarkedDates() {
    const marked: MarkedDates = {};

    agendaItems.forEach((item) => {
        // NOTE: only mark dates with data
        if (item.data.length) {
            marked[item.title] = { marked: true };
        }
    });
    return marked;
}
