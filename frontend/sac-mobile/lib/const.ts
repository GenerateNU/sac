import { Club, HomepageItem } from '@/types/item';

export const API_BASE_URL = 'http://localhost:8080/api/v1';

export const college = [
    { label: 'College of Arts, Media and Design', value: 'CAMD' },
    { label: "D'Amore-McKim School of Business", value: 'DMSB' },
    { label: 'Khoury College of Computer Sciences', value: 'KCCS' },
    { label: 'College of Engineering', value: 'CE' },
    { label: 'Bouvé College of Health Sciences', value: 'BCHS' },
    { label: 'School of Law', value: 'SL' },
    { label: 'College of Professional Studies', value: 'CPS' },
    { label: 'College of Science', value: 'CS' },
    { label: 'College of Social Sciences and Humanities', value: 'CSSH' }
];

export const majorArr = [
    'Africana Studies',
    'American Sign Language',
    'American Sign Language – English Interpreting',
    'Applied Physics',
    'Architectural Studies',
    'Architecture',
    'Art: Art, Visual Studies',
    'Behavioral Neuroscience*',
    'Biochemistry',
    'Bioengineering',
    'Biology',
    'Biomedical Physics',
    'Bouvé Undeclared',
    'Business Administration: Accounting',
    'Business Administration: Accounting and Advisory Services',
    'Business Administration: Brand Management',
    'Business Administration: Business Analytics',
    'Business Administration: Corporate Innovation',
    'Business Administration: Entrepreneurial Startups',
    'Business Administration: Family Business',
    'Business Administration: Finance',
    'Business Administration: Fintech',
    'Business Administration: Healthcare Management and Consulting',
    'Business Administration: Management',
    'Business Administration: Management Information Systems',
    'Business Administration: Marketing',
    'Business Administration: Marketing Analytics',
    'Business Administration: Social Innovation and Entrepreneurship',
    'Business Administration: Supply Chain Management',
    'Business Administration: Undeclared',
    'Business Administration',
    'CAMD Undeclared',
    'Cell and Molecular Biology',
    'Chemical Engineering',
    'Chemistry',
    'Civil Engineering',
    'COE Undeclared',
    'Communication Studies',
    'Computer Engineering',
    'Computer Science',
    'Computing and Law',
    'Criminology and Criminal Justice',
    'Cultural Anthropology',
    'Cybersecurity',
    'Data Science',
    'Design',
    'Economics',
    'Electrical Engineering',
    'English',
    'Environmental and Sustainability Sciences',
    'Environmental Engineering',
    'Environmental Science',
    'Environmental Studies',
    'Explore Program for undeclared students',
    'Game Art and Animation',
    'Game Design',
    'Global Asian Studies',
    'Health Science',
    'History',
    'History, Culture, and Law',
    'Human Services',
    'Industrial Engineering',
    'International Affairs',
    'International Business: Accounting',
    'International Business: Accounting and Advisory Services',
    'International Business: Brand Management',
    'International Business: Business Analytics',
    'International Business: Corporate Innovation',
    'International Business: Entrepreneurial Startups',
    'International Business: Family Business',
    'International Business: Finance',
    'International Business: Fintech',
    'International Business: Healthcare Management and Consulting',
    'International Business: Management',
    'International Business: Management Information Systems',
    'International Business: Marketing',
    'International Business: Marketing Analytics',
    'International Business: Social Innovation and Entrepreneurship',
    'International Business: Supply Chain Management',
    'International Business: Undeclared',
    'Journalism',
    'Landscape Architecture',
    'Linguistics',
    'Marine Biology',
    'Mathematics',
    'Mechanical Engineering',
    'Media and Screen Studies',
    'Media Arts',
    'Music',
    'Music Industry',
    'Music Technology',
    'Northeastern Explore Program (Undeclared)',
    'Nursing',
    'Pharmaceutical Sciences',
    'Pharmacy (PharmD)',
    'Philosophy',
    'Physics',
    'Political Science',
    'Politics, Philosophy, Economics',
    'Psychology',
    'Public Health',
    'Public Relations',
    'Religious Studies',
    'Sociology',
    'Spanish',
    'Speech Language Pathology and Audiology',
    'Theatre'
];

export const categories = [
    {
        name: 'Pre-Professional',
        tags: ['Pre-med', 'Pre-law', 'Other']
    },
    {
        name: 'Cultural and Identity',
        tags: [
            'Judaism',
            'Christianity',
            'Hinduism',
            'Islam',
            'Latin America',
            'African American',
            'Asian American',
            'LGBTQ',
            'Other'
        ]
    },
    {
        name: 'Arts and Creativity',
        tags: [
            'Performing Arts',
            'Visual Arts',
            'Creative Writing',
            'Music',
            'Other'
        ]
    },
    {
        name: 'Sports and Recreation',
        tags: ['Soccer', 'Hiking', 'Climbing', 'Lacrosse', 'Other']
    },
    {
        name: 'Science and Technology',
        tags: [
            'Mathematics',
            'Physics',
            'Biology',
            'Chemistry',
            'Environmental Science',
            'Geology',
            'Neuroscience',
            'Psychology',
            'Software Engineering',
            'Artificial Intelligence',
            'DataScience',
            'Mechanical Engineering',
            'Electrical Engineering',
            'Industrial Engineering',
            'Other'
        ]
    },
    {
        name: 'Community Service and Advocacy',
        tags: [
            'Volunteerism',
            'Environmental Advocacy',
            'Human Rights',
            'Community Outreach',
            'Other'
        ]
    },
    {
        name: 'Media and Communication',
        tags: [
            'Journalism',
            'Broadcasting',
            'Film',
            'Public Relations',
            'Other'
        ]
    }
];

export const HomepageList: HomepageItem[] = [
    {
        type: 'club',
        name: "Nor'easters A Capella",
        description:
            'The Nor’easters are Northeastern University’s premier a cappella group. As the oldest group on campus, the Nor’easters have continuously striven to maintain the highest standards of musicality while still keeping the fun and love alive.',
    },
    {
        type: 'event',
        clubName: "Nor'easters A Capella",
        eventName: 'Best of Northeastern Region',
        location: 'Blackman Auditorium',
        description: 'It’s that time of year again... BONR is one week away.',
        time: 'March 18, 7-8 PM',
    },
    {
        type: 'club',
        name: 'Cheese Club',
        description:
            'Yes, it’s exactly as it sounds: Cheese Club is a club where we learn about and eat delicious cheeses every week.',
    },
    {
        type: 'club',
        name: 'AerospaceNU',
        description:
            'This group is for anyone and everyone who wants to design and build projects with the goal of getting them in the air.',
    },
    {
        type: 'faq',
        clubName: "Nor'easters A Capella",
        question: 'Is the ticket free?',
        answer: 'Yes, the event is opened to all students, and the tickets are free. You can purchase them here.',
    },
    {
        type: 'club',
        name: 'Super Smash Bros Club',
        description:
            'We are the smash bros club, we hold Smash Ultimate tournaments every Friday, and Smash Melee tournaments every Monday.',
    },
    {
        type: 'event',
        clubName: 'Super Smash Bros Club',
        eventName: 'Big Northeastern Ultimate Tournament 4',
        location:
            'West Village H, Rooms 108 & 110, 440 Huntington Ave, Boston, MA',
        description:
            'The Northeastern Super Smash Bros. Club is pleased to announce that it will be hosting its fourth BIG NORTHEASTERN ULTIMATE TOURNAMENT (aka BIG NUT 4).',
        time: 'March 20, 6-10 PM',
    },
    {
        type: 'club',
        name: 'Oasis',
        description:
            'In the simplest terms, Oasis is a full-fledged project accelerator where every semester, a new cohort of students build a software project with the support of upperclassmen mentors.',
    },
    {
        type: 'faq',
        clubName: 'Oasis',
        question: 'How do you assess applications?',
        answer: "We don't assess them in the traditional sense. We pride ourselves on being open to students from all backgrounds and experience levels, so our application is first-come first-serve to keep it simple and fair for everybody.",
    },
    {
        type: 'faq',
        clubName: 'Oasis',
        question: 'How many students are there in normal cohort?',
        answer: 'A typical semester is roughly 80 students. We target 10 mentors each semester, and each mentor works with two groups of four students each.',
    }
];

export const FollowedClubs: Club[] = [
    {
        type: 'club',
        name: 'Oasis',
        description:
            'In the simplest terms, Oasis is a full-fledged project accelerator where every semester, a new cohort of students build a software project with the support of upperclassmen mentors.',
    },
    {
        type: 'club',
        name: 'Super Smash Bros',
        description:
            'We are the smash bros club, we hold Smash Ultimate tournaments every Friday, and Smash Melee tournaments every Monday.',
    },
    {
        type: 'club',
        name: 'AerospaceNU',
        description:
            'This group is for anyone and everyone who wants to design and build projects with the goal of getting them in the air.',
    },
    {
        type: 'club',
        name: 'Cheese Club',
        description:
            'Yes, it’s exactly as it sounds: Cheese Club is a club where we learn about and eat delicious cheeses every week.',
    },
    {
        type: 'club',
        name: "Nor'easters",
        description:
            'The Nor’easters are Northeastern University’s premier a cappella group. As the oldest group on campus, the Nor’easters have continuously striven to maintain the highest standards of musicality while still keeping the fun and love alive.',
    },
    {
        type: 'club',
        name: 'Skate Club',
        description: '',
    },
    {
        type: 'club',
        name: 'Kinematix',
        description:
            'The team recognizes dance as a powerful, evolving culture and provides a channel for students to grow in that richness.',
    },
    {
        type: 'club',
        name: 'Sandbox',
        description:
            "Sandbox is Northeastern's student-led software consultancy. Sandbox members work in teams on projects that typically last two semesters or longer.",
    },
    {
        type: 'club',
        name: 'Swim Club',
        description:
            'NUSC is a Northeastern University club sport that competes on both regional and national levels while providing members with fun and rewarding experiences.',
    }
];

export const ChronologicalList: HomepageItem[] = [
    {
        type: 'event',
        clubName: "Nor'easters A Capella",
        eventName: 'Best of Northeastern Region',
        location: 'Blackman Auditorium',
        description: 'It’s that time of year again... BONR is one week away.',
        time: 'March 18, 7-8 PM',
    },
    {
        type: 'faq',
        clubName: 'Oasis',
        question: 'How do you assess applications?',
        answer: "We don't assess them in the traditional sense. We pride ourselves on being open to students from all backgrounds and experience levels, so our application is first-come first-serve to keep it simple and fair for everybody.",
    },
    {
        type: 'faq',
        clubName: 'Oasis',
        question: 'How many students are there in normal cohort?',
        answer: 'A typical semester is roughly 80 students. We target 10 mentors each semester, and each mentor works with two groups of four students each.',
    },
    {
        type: 'faq',
        clubName: "Nor'easters A Capella",
        question: 'Is the ticket free?',
        answer: 'Yes, the event is opened to all students, and the tickets are free. You can purchase them here.',
    },
    {
        type: 'club',
        name: 'Super Smash Bros Club',
        description:
            'We are the smash bros club, we hold Smash Ultimate tournaments every Friday, and Smash Melee tournaments every Monday.' ,
    },
    {
        type: 'event',
        clubName: 'Super Smash Bros Club',
        eventName: 'Big Northeastern Ultimate Tournament 4',
        location:
            'West Village H, Rooms 108 & 110, 440 Huntington Ave, Boston, MA',
        description:
            'The Northeastern Super Smash Bros. Club is pleased to announce that it will be hosting its fourth BIG NORTHEASTERN ULTIMATE TOURNAMENT (aka BIG NUT 4).',
        time: 'March 20, 6-10 PM',
    },
    {
        type: 'club',
        name: 'Oasis',
        description:
            'In the simplest terms, Oasis is a full-fledged project accelerator where every semester, a new cohort of students build a software project with the support of upperclassmen mentors.',
    },
    {
        type: 'club',
        name: 'AerospaceNU',
        description:
            'This group is for anyone and everyone who wants to design and build projects with the goal of getting them in the air.',
    }
];
