import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchEvent } from '@/services/event';
import { Event } from '@/types/event';
import { uuid } from '@/types/uuid';

export const useEvent = (eventID: uuid): UseQueryResult<Event, Error> => {
    return useQuery({
        queryKey: ['event', eventID],
        queryFn: () => {
            if (eventID == 'tester') {
                return {
                    id: 'tester',
                    name: 'Tester Event',
                    preview: 'This is a preview',
                    content: 'This is the content',
                    startTime: new Date(),
                    endTime: new Date(),
                    location: 'Here',
                    eventType: 'open',
                    isRecurring: false,
                    createdAt: new Date(),
                    updatedAt: new Date()
                };
            } else {
                return fetchEvent(eventID);
            }
        }
    });
};
