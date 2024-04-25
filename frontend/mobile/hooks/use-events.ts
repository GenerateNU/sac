import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { uuid } from '@/types/uuid';

export type Event = {
    title: string;
    time: string;
    date: string;
    location: string;
};

export const useEvents = (clubID: uuid): UseQueryResult<Event[], Error> => {
    return useQuery<Event[], Error>({
        queryKey: ['events', clubID],
        queryFn: async (): Promise<Event[]> => {
            return [
                {
                    title: 'Orientation',
                    time: '8:00 am - 9:30 am',
                    date: 'February 29',
                    location: 'Ryder Hall, Room 294'
                },
                {
                    title: 'Breakfast',
                    time: '9:00 am - 10:30 am',
                    date: 'March 1',
                    location: 'Ryder Hall, Room 294'
                },
                {
                    title: 'Orientation',
                    time: '7:00 am - 11:30 am',
                    date: 'March 2',
                    location: 'Ryder Hall, Room 294'
                }
            ];
        }
    });
};
