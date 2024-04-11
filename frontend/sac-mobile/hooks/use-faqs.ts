import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { uuid } from '@/types/uuid';

export type Faq = {
    question: string;
    answer: string;
};

export const useFAQs = (clubID: uuid): UseQueryResult<Faq[], Error> => {
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

export const useFAQ = (faqID: uuid): UseQueryResult<Faq, Error> => {
    return useQuery<Faq, Error>({
        queryKey: ['faq', faqID],
        queryFn: async (): Promise<Faq> => {
            return {
                question: 'What is the purpose of the club?',
                answer: 'The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field. The purpose of the club is to provide a platform for like-minded individuals to come together and pursue their interests in a specific field.'
            };
        }
    });
};
