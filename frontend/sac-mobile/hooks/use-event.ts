import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchEvent } from '@/services/event';
import { Event } from '@/types/event';
import { uuid } from '@/types/uuid';

export const useEvent = (eventID: uuid): UseQueryResult<Event, Error> => {
    return useQuery({
        queryKey: ['event', eventID],
        queryFn: () => fetchEvent(eventID)
    });
};
