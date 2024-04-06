import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchEvent } from '@/services/event';
import { Event } from '@/types/event';
import { uuid } from '@/types/uuid';

export const useEvent = (eventID: uuid): UseQueryResult<Event, Error> => {
    return useQuery<Event, Error>({
        queryKey: ['useEventeventID', eventID],
        queryFn: () => {
            if (eventID === 'tester') {
                return {
                    id: 'tester',
                    created_at: new Date(),
                    updated_at: new Date(),
                    name: 'Tester Event',
                    preview: 'This is a preview',
                    content:
                        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut',
                    start_time: new Date(2024, 2, 15, 18, 30),
                    end_time: new Date(2024, 2, 15, 20, 0),
                    location: 'Here',
                    meeting_link: 'https://foo.com',
                    event_type: 'open',
                    is_recurring: false
                } as Event;
            } else {
                return fetchEvent(eventID);
            }
        }
    });
};
