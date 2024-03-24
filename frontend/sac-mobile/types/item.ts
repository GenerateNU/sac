export type Item = {
    label: string;
    value: string;
};

export type Event = {
    type: 'event';
    clubName: string;
    eventName: string;
    location: string;
    description: string;
    time: string;
};

export type Club = {
    type: 'club';
    name: string;
    description: string;
};

export type FAQ = {
    type: 'faq';
    clubName: string;
    question: string;
    answer: string;
};

export type HomepageItem = Event | Club | FAQ;
