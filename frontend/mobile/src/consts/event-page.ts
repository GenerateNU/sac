import { Event, Tag } from '@generatesac/lib';

export const description =
    "Generate is Northeastern's only student-led product development studio, working with industry clients and Northeastern founders from all backgrounds to turn their vision into reality. We provide hardware and software development expertise, supporting over 16 clients annually. Generate offers a unique and supportive learning opportunity for all members through professional product development experience, leadership, and training in their respective disciplines. We provide our clients with the engineering, design, and prototyping they need to take their venture to the next stage.";
    export const tags: Tag[] = [
        {
            id: '1',
            created_at: '2023-01-01T10:00:00Z',
            updated_at: '2023-01-10T10:00:00Z',
            name: 'Engineering',
            category_id: '100'
        },
        {
            id: '2',
            created_at: '2023-02-01T11:00:00Z',
            updated_at: '2023-02-10T11:00:00Z',
            name: 'Computer Science',
            category_id: '101'
        },
        {
            id: '3',
            created_at: '2023-03-01T12:00:00Z',
            updated_at: '2023-03-10T12:00:00Z',
            name: 'Design',
            category_id: '102'
        },
        {
            id: '4',
            created_at: '2023-04-01T13:00:00Z',
            updated_at: '2023-04-10T13:00:00Z',
            name: 'Tabling',
            category_id: '103'
        },
        {
            id: '5',
            created_at: '2023-05-01T14:00:00Z',
            updated_at: '2023-05-10T14:00:00Z',
            name: 'Showcase',
            category_id: '104'
        }
    ];
    
    
    export const events: Event[] = [
        {
            id: 'event1',
            created_at: '2024-06-01T08:00:00Z',
            updated_at: '2024-06-01T08:00:00Z',
            name: 'Web Development Workshop',
            preview: 'Learn the latest web development techniques and frameworks!',
            description: 'This workshop covers topics such as HTML, CSS, JavaScript, and popular frontend frameworks like React and Angular.',
            start_time: '2024-07-15T10:00:00Z',
            end_time: '2024-07-15T14:00:00Z',
            location: 'Tech Hub Co-working Space',
            event_type: 'in_person',
            host: 'Tech Skills Academy',
            is_public: true,
            is_draft: false,
            is_archived: false,
            link: 'https://example.com/web-development-workshop',
            is_recurring: false
        },
        {
            id: 'event2',
            created_at: '2024-06-15T09:00:00Z',
            updated_at: '2024-06-15T09:00:00Z',
            name: 'AI Conference',
            preview: 'Explore the latest advancements in artificial intelligence!',
            description: 'Join experts from academia and industry as they discuss topics such as machine learning, deep learning, natural language processing, and more.',
            start_time: '2024-08-20T09:00:00Z',
            end_time: '2024-08-21T17:00:00Z',
            location: 'Convention Center',
            event_type: 'hybrid',
            host: 'AI Institute',
            is_public: true,
            is_draft: false,
            is_archived: false,
            link: 'https://example.com/ai-conference',
            is_recurring: false
        },
        {
            id: 'event3',
            created_at: '2024-07-01T10:00:00Z',
            updated_at: '2024-07-01T10:00:00Z',
            name: 'Hackathon',
            preview: 'Put your coding skills to the test in our annual hackathon!',
            description: 'Gather with fellow developers to collaborate on projects, solve challenges, and showcase your creativity.',
            start_time: '2024-09-10T12:00:00Z',
            end_time: '2024-09-12T12:00:00Z',
            location: 'University Campus',
            event_type: 'in_person',
            host: 'Developer Community',
            is_public: true,
            is_draft: false,
            is_archived: false,
            link: 'https://example.com/hackathon',
            is_recurring: false
        }
    ];
    