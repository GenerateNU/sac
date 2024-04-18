import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClub, fetchClubContacts, fetchClubs } from '@/services/club';
import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';

const useClubs = (): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['clubs'],
        queryFn: () => {
            return fetchClubs();
        }
    });
};

const useClub = (clubID: uuid): UseQueryResult<Club, Error> => {
    return useQuery<Club, Error>({
        queryKey: ['club', clubID],
        queryFn: () => {
            return fetchClub(clubID);
        }
    });
};

const useClubContacts = (clubID: uuid): UseQueryResult<Contact[], Error> => {
    return useQuery<Contact[], Error>({
        queryKey: ['club', clubID, 'contacts'],
        queryFn: () => {
            return fetchClubContacts(clubID);
        }
    });
};

export type Faq = {
    question: string;
    answer: string;
};

export const useClubFAQs = (clubID: uuid): UseQueryResult<Faq[], Error> => {
    return useQuery<Faq[], Error>({
        queryKey: ['faqs', clubID],
        queryFn: async (): Promise<Faq[]> => {
            return [
                {
                    question: 'What is the purpose of the club?',
                    answer: 'The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field.'
                },
                {
                    question: 'How can I join the club?',
                    answer: 'To join the club, you can attend our regular meetings and express your interest in becoming a member. We welcome anyone who shares our passion and is willing to contribute.'
                },
                {
                    question: 'What activities does the club organize?',
                    answer: 'The club organizes various activities such as workshops, guest lectures, competitions, and networking events to enhance the learning and networking opportunities for its members.'
                }
            ];
        }
    });
};

export const useClubFAQ = (faqID: uuid): UseQueryResult<Faq, Error> => {
    return useQuery<Faq, Error>({
        queryKey: ['faq', faqID],
        queryFn: async (): Promise<Faq> => {
            let question = '';
            let answer = '';

            if (faqID === '1') {
                question = 'What is the purpose of the club?';
                answer = 'The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field.';
            } else if (faqID === '2') {
                question = 'How can I join the club?';
                answer = 'To join the club, you can attend our regular meetings and express your interest in becoming a member. We welcome anyone who shares our passion and is willing to contribute.';
            } else if (faqID === '3') {
                question = 'What activities does the club organize?';
                answer = 'The club organizes various activities such as workshops, guest lectures, competitions, and networking events to enhance the learning and networking opportunities for its members.';
            }

            return {
                question,
                answer
            };
        }
    });
};

export type Admin = {
    id: uuid;
    name: string;
    title: string;
};

export const useClubAdmin = (clubID: uuid): UseQueryResult<Admin[], Error> => {
    return useQuery<Admin[], Error>({
        queryKey: ['eboard-members', clubID],
        queryFn: async (): Promise<Admin[]> => {
            return [
                { id: 'alder', name: 'Alder Whiteford', title: 'Project Lead' },
                { id: 'amanda', name: 'Amanda Kerr', title: 'Design Lead' },
                { id: 'david', name: 'David Oduneye', title: 'Tech Lead' },
                { id: 'garrett', name: 'Garrett Ladley', title: 'Tech Lead' }
            ];
        }
    });
};
export { useClub, useClubs, useClubContacts };
