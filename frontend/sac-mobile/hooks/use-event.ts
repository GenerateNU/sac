import { UseQueryResult, useQuery } from '@tanstack/react-query';

import {
    fetchEvent,
    fetchEventTags,
    fetchEvents,
    fetchHosts
} from '@/services/event';
import { Club } from '@/types/club';
import { Event } from '@/types/event';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';

const useEvent = (eventID: string): UseQueryResult<Event, Error> => {
    return useQuery<Event, Error>({
        queryKey: ['event', eventID],
        queryFn: () => {
            return fetchEvent(eventID);
        }
    });
};

const useEvents = (): UseQueryResult<Event[], Error> => {
    return useQuery<Event[], Error>({
        queryKey: ['events'],
        queryFn: () => {
            return fetchEvents();
        }
    });
};

const useEventTags = (eventID: uuid): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['event', eventID, 'tags'],
        queryFn: () => {
            return fetchEventTags(eventID);
        }
    });
};

const useEventHosts = (eventID: uuid): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['event', eventID, 'hosts'],
        queryFn: () => {
            return fetchHosts(eventID);
        }
    });
};

export {
    useEvent,
    useEvents,
    useEventTags,
    useEventHosts
};