import { Club, Contact, Event, PointOfContact, Tag } from '@generatesac/lib';

const mockClub: Club = {
    id: '1',
    name: 'Mock Club',
    preview: 'This is a mock club for demonstration purposes.',
    description:
        'A mock club that serves as an example for the ClubPage component.',
    is_recruiting: true,
    recruitment_cycle: 'fall',
    recruitment_type: 'application',
    application_link: 'https://example.com/apply',
    num_members: 50,
    created_at: new Date(),
    updated_at: new Date(),
    logo: 'https://example.com/logo.png',
    weekly_time_committment: 5,
    one_word_to_describe_us: 'Innovative'
};

const mockTags: Tag[] = [
    {
        id: '1',
        name: 'Technology',
        category_id: '1',
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        id: '2',
        name: 'Community',
        category_id: '2',
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        id: '3',
        name: 'Social',
        category_id: '3',
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        id: '4',
        name: 'Professional',
        category_id: '4',
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        id: '5',
        name: 'Academic',
        category_id: '5',
        created_at: new Date(),
        updated_at: new Date()
    }
];

const mockEvents: Event[] = [
    {
        id: '1',
        name: 'Mock Event 1',
        preview: 'This is the first mock event.',
        content: 'Details about the first mock event.',
        start_time: new Date(),
        end_time: new Date(),
        location: 'Online',
        event_type: 'open',
        is_recurring: false,
        created_at: new Date(),
        updated_at: new Date(),
        host: 'Mock Host 1',
        meeting_link: 'https://example.com/meeting1'
    },
    {
        id: '2',
        name: 'Mock Event 2',
        preview: 'This is the second mock event.',
        content: 'Details about the second mock event.',
        start_time: new Date(),
        end_time: new Date(),
        location: 'Offline',
        event_type: 'membersOnly',
        is_recurring: true,
        created_at: new Date(),
        updated_at: new Date(),
        host: 'Mock Host 2',
        meeting_link: 'https://example.com/meeting2'
    },
    {
        id: '3',
        name: 'Mock Event 3',
        preview: 'This is the third mock event.',
        content: 'Details about the third mock event.',
        start_time: new Date(),
        end_time: new Date(),
        location: 'Online',
        event_type: 'open',
        is_recurring: false,
        created_at: new Date(),
        updated_at: new Date(),
        host: 'Mock Host 3',
        meeting_link: 'https://example.com/meeting3'
    },
    {
        id: '4',
        name: 'Mock Event 4',
        preview: 'This is the fourth mock event.',
        content: 'Details about the fourth mock event.',
        start_time: new Date(),
        end_time: new Date(),
        location: 'Offline',
        event_type: 'membersOnly',
        is_recurring: true,
        created_at: new Date(),
        updated_at: new Date(),
        host: 'Mock Host 4',
        meeting_link: 'https://example.com/meeting4'
    },
    {
        id: '5',
        name: 'Mock Event 5',
        preview: 'This is the fifth mock event.',
        content: 'Details about the fifth mock event.',
        start_time: new Date(),
        end_time: new Date(),
        location: 'Online',
        event_type: 'open',
        is_recurring: false,
        created_at: new Date(),
        updated_at: new Date(),
        host: 'Mock Host 5',
        meeting_link: 'https://example.com/meeting5'
    }
];

const mockContacts: Contact[] = [
    {
        id: '1',
        type: 'email',
        content: 'contact@example.com',
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        id: '2',
        type: 'discord',
        content: 'https://discord.gg/example',
        created_at: new Date(),
        updated_at: new Date()
    }
];

const pointOfContactMocks: PointOfContact[] = [
    {
        name: 'Jane Smith',
        email: 'jane.smith@example.com',
        position: 'Director',
        photo_file: {
            owner_id: 'd4f5a6e7-8b9c-10d1-1121-314151617181',
            owner_type: 'user',
            file_name: 'photo1.jpg',
            file_type: 'image/jpeg',
            file_size: 2048,
            file_url: '/photos/photo1.jpg',
            object_key: 'photos/photo1.jpg',
            id: '12345',
            created_at: new Date('2023-05-01T00:00:00Z'),
            updated_at: new Date('2023-05-02T00:00:00Z')
        },
        id: '54321',
        created_at: new Date('2023-05-01T00:00:00Z'),
        updated_at: new Date('2023-05-02T00:00:00Z')
    },
    {
        name: 'John Doe',
        email: 'john.doe@example.com',
        position: 'Manager',
        photo_file: {
            owner_id: 'd4f5a6e7-8b9c-10d1-1121-314151617182',
            owner_type: 'user',
            file_name: 'photo2.jpg',
            file_type: 'image/jpeg',
            file_size: 1024,
            file_url: '/photos/photo2.jpg',
            object_key: 'photos/photo2.jpg',
            id: '12346',
            created_at: new Date('2023-05-01T00:00:00Z'),
            updated_at: new Date('2023-05-02T00:00:00Z')
        },
        id: '54322',
        created_at: new Date('2023-05-01T00:00:00Z'),
        updated_at: new Date('2023-05-02T00:00:00Z')
    },
    {
        name: 'Alice Johnson',
        email: 'alice.johnson@example.com',
        position: 'Team Lead',
        photo_file: {
            owner_id: 'd4f5a6e7-8b9c-10d1-1121-314151617183',
            owner_type: 'user',
            file_name: 'photo3.jpg',
            file_type: 'image/jpeg',
            file_size: 3072,
            file_url: '/photos/photo3.jpg',
            object_key: 'photos/photo3.jpg',
            id: '12347',
            created_at: new Date('2023-05-01T00:00:00Z'),
            updated_at: new Date('2023-05-02T00:00:00Z')
        },
        id: '54323',
        created_at: new Date('2023-05-01T00:00:00Z'),
        updated_at: new Date('2023-05-02T00:00:00Z')
    },
    {
        name: 'Bob Brown',
        email: 'bob.brown@example.com',
        position: 'Senior Developer',
        photo_file: {
            owner_id: 'd4f5a6e7-8b9c-10d1-1121-314151617184',
            owner_type: 'user',
            file_name: 'photo4.jpg',
            file_type: 'image/jpeg',
            file_size: 4096,
            file_url: '/photos/photo4.jpg',
            object_key: 'photos/photo4.jpg',
            id: '12348',
            created_at: new Date('2023-05-01T00:00:00Z'),
            updated_at: new Date('2023-05-02T00:00:00Z')
        },
        id: '54324',
        created_at: new Date('2023-05-01T00:00:00Z'),
        updated_at: new Date('2023-05-02T00:00:00Z')
    },
    {
        name: 'Clara White',
        email: 'clara.white@example.com',
        position: 'Product Manager',
        photo_file: {
            owner_id: 'd4f5a6e7-8b9c-10d1-1121-314151617185',
            owner_type: 'user',
            file_name: 'photo5.jpg',
            file_type: 'image/jpeg',
            file_size: 5120,
            file_url: '/photos/photo5.jpg',
            object_key: 'photos/photo5.jpg',
            id: '12349',
            created_at: new Date('2023-05-01T00:00:00Z'),
            updated_at: new Date('2023-05-02T00:00:00Z')
        },
        id: '54325',
        created_at: new Date('2023-05-01T00:00:00Z'),
        updated_at: new Date('2023-05-02T00:00:00Z')
    }
];

export { mockClub, mockTags, mockEvents, mockContacts, pointOfContactMocks };
