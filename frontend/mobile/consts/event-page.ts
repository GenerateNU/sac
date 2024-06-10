import { Event, Tag } from '@generatesac/lib';

export const description =
    "Generate is Northeastern's only student-led product development studio, working with industry clients and Northeastern founders from all backgrounds to turn their vision into reality. We provide hardware and software development expertise, supporting over 16 clients annually. Generate offers a unique and supportive learning opportunity for all members through professional product development experience, leadership, and training in their respective disciplines. We provide our clients with the engineering, design, and prototyping they need to take their venture to the next stage.";
export const tags: Tag[] = [
    {
        id: '1',
        created_at: new Date('2023-01-01T10:00:00Z'),
        updated_at: new Date('2023-01-10T10:00:00Z'),
        name: 'Engineering',
        category_id: '100'
    },
    {
        id: '2',
        created_at: new Date('2023-02-01T11:00:00Z'),
        updated_at: new Date('2023-02-10T11:00:00Z'),
        name: 'Computer Science',
        category_id: '101'
    },
    {
        id: '3',
        created_at: new Date('2023-03-01T12:00:00Z'),
        updated_at: new Date('2023-03-10T12:00:00Z'),
        name: 'Design',
        category_id: '102'
    },
    {
        id: '4',
        created_at: new Date('2023-04-01T13:00:00Z'),
        updated_at: new Date('2023-04-10T13:00:00Z'),
        name: 'Tabling',
        category_id: '103'
    },
    {
        id: '5',
        created_at: new Date('2023-05-01T14:00:00Z'),
        updated_at: new Date('2023-05-10T14:00:00Z'),
        name: 'Showcase',
        category_id: '104'
    }
];

export const events: Event[] = [
    {
        id: 'event1',
        created_at: new Date('2024-06-01T08:00:00Z'),
        updated_at: new Date('2024-06-01T08:00:00Z'),
        name: 'Networking Mixer',
        preview: 'Join us for an evening of networking and socializing!',
        content:
            'Our Networking Mixer event provides a casual environment for professionals to connect and build relationships.',
        start_time: new Date('2024-06-15T18:00:00Z'),
        end_time: new Date('2024-06-15T20:00:00Z'),
        location: 'City Center Lounge',
        event_type: 'open',
        is_recurring: false,
        host: 'Business Networking Group',
        meeting_link: 'https://example.com/networking-mixer'
    },
    {
        id: 'event2',
        created_at: new Date('2024-06-01T09:00:00Z'),
        updated_at: new Date('2024-06-01T09:00:00Z'),
        name: 'Exclusive Seminar',
        preview: 'Limited seating available! Reserve your spot now.',
        content:
            'Our Exclusive Seminar offers in-depth insights into the latest industry trends and best practices.',
        start_time: new Date('2024-06-20T10:00:00Z'),
        end_time: new Date('2024-06-20T12:00:00Z'),
        location: 'Grand Ballroom, XYZ Hotel',
        event_type: 'membersOnly',
        is_recurring: false,
        host: 'Industry Experts Association'
    },
    {
        id: 'event3',
        created_at: new Date('2024-06-02T10:00:00Z'),
        updated_at: new Date('2024-06-02T10:00:00Z'),
        name: 'Weekly Yoga Class',
        preview: 'Relax, rejuvenate, and find your inner balance.',
        content:
            'Join us every Thursday for a refreshing yoga session suitable for all levels.',
        start_time: new Date('2024-06-06T17:30:00Z'),
        end_time: new Date('2024-06-06T18:30:00Z'),
        location: 'Sunset Yoga Studio',
        event_type: 'open',
        is_recurring: true,
        host: 'Sunset Yoga Studio'
    },
    {
        id: 'event4',
        created_at: new Date('2024-06-02T11:00:00Z'),
        updated_at: new Date('2024-06-02T11:00:00Z'),
        name: "Beginner's Coding Workshop",
        preview: 'Learn the basics of coding in a supportive environment.',
        content:
            'This workshop is designed for beginners with no prior coding experience.',
        start_time: new Date('2024-06-12T09:00:00Z'),
        end_time: new Date('2024-06-12T12:00:00Z'),
        location: 'Tech Hub Co-working Space',
        event_type: 'open',
        is_recurring: false,
        host: 'Tech Skills Academy'
    }
];
