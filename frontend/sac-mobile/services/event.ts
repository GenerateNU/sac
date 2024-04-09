import { Club } from '@/types/club';
import { Event } from '@/types/event';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';

import { api } from './api';

/**
 * Fetches an event by its ID.
 *
 * @param eventID The ID of the event to fetch
 * @returns A Promise that resolves to the fetched event or rejects with an error
 * @throws Error if the event cannot be fetched
 * @see Event
 */
export const fetchEvent = async (eventID: uuid): Promise<Event> => {
    return api
        .get<Event>(`/events/${eventID}`)
        .then((response) => response.data as Event)
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching event');
        });
};

/**
 * Fetches all events.
 *
 * @returns A Promise that resolves to the fetched events or rejects with an error
 * @throws Error if the events cannot be fetched
 * @see Event
 */
export const fetchEvents = async (): Promise<Event[]> => {
    return api
        .get('/events')
        .then((response) => response.data as Event[])
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching events');
        });
};

/**
 * Fetches the hosts of an event by its ID.
 *
 * @param eventID The ID of the event whose hosts to fetch
 * @returns A Promise that resolves to the fetched hosts or rejects with an error
 * @throws Error if the hosts cannot be fetched
 * @see Club
 */
export const fetchHosts = async (eventID: uuid): Promise<Club[]> => {
    return api
        .get(`/events/${eventID}/hosts`)
        .then((response) => response.data as Club[])
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching hosts');
        });
};

/**
 * Fetches the tags associated with an event.
 *
 * @param eventID The ID of the event to fetch tags for
 * @returns A Promise that resolves to the fetched tags or rejects with an error
 * @throws Error if the tags cannot be fetched
 * @see Tag
 */
export const fetchEventTags = async (eventID: uuid): Promise<Tag[]> => {
    return api
        .get(`/events/${eventID}/tags`)
        .then((response) => response.data as Tag[])
        .catch((error) => {
            console.error(error);
            throw new Error('Failed to fetch event tags');
        });
};
