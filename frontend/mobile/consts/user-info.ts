import { Item } from '@/types/item';

export const COLLEGE: Item[] = [
    { value: 'College of Arts, Media and Design', key: 'CAMD' },
    { value: "D'Amore-McKim School of Business", key: 'DMSB' },
    { value: 'Khoury College of Computer Sciences', key: 'KCCS' },
    { value: 'College of Engineering', key: 'CE' },
    { value: 'Bouvé College of Health Sciences', key: 'BCHS' },
    { value: 'School of Law', key: 'SL' },
    { value: 'College of Professional Studies', key: 'CPS' },
    { value: 'College of Science', key: 'CS' },
    { value: 'College of Social Sciences and Humanities', key: 'CSSH' }
];

export const GRADUATION_CYCLE: Item[] = [
    { value: 'December', key: 'december' },
    { value: 'May', key: 'may' }
];

export const MAJOR_ARR = [
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

export const graduationYear = () => {
    var year = new Date().getFullYear();
    const range = 5;
    const graduationYears: Item[] = [];
    for (let i = 0; i < range; i++) {
        graduationYears.push({
            key: String(year + i),
            value: String(year + i)
        });
    }
    return graduationYears;
};

export const major = () => {
    const majors: Item[] = [];
    for (let i = 0; i < MAJOR_ARR.length; i++) {
        majors.push({
            key: MAJOR_ARR[i],
            value: MAJOR_ARR[i]
        });
    }
    return majors;
};
