import { fetchEvent } from '@/services/event';
import { Event } from '@/types/event';
import { uuid } from '@/types/uuid';
import { useQuery, UseQueryResult } from '@tanstack/react-query';

export const useEvent = (eventID: uuid): UseQueryResult<Event, Error> => {
    return useQuery({
        queryKey: ['event', eventID],
        queryFn: () => {
            if (eventID === 'tester') {
                return {
                    id: 'tester',
                    name: 'Tester Event',
                    preview: 'This is a preview',
                    content: 'This is the content',
                    startTime: new Date(2024, 2, 15, 18, 30),
                    endTime: new Date(2024, 2, 15, 20, 0),
                    location: 'Here',
                    meetingLink: 'https://foo.com',
                    eventType: 'open',
                    isRecurring: false,
                    createdAt: new Date(),
                    updatedAt: new Date(),
                    hosts: ['Generate', 'Generate1', 'Generate2', 'Generate3']
                } as Event;
            } else {
                return fetchEvent(eventID);
            }
        }
    });
};
