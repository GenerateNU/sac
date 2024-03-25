import axios from 'axios';

import { API_BASE_URL } from '@/lib/const';
import { Event } from '@/types/event';
import { uuid } from '@/types/uuid';

/**
 * Fetches an event by its ID using functional style with `andThen` syntax.
 *
 * @param eventID The ID of the event to fetch
 * @returns A Promise that resolves to the fetched event or rejects with an error
 */
export const fetchEvent = async (eventID: uuid): Promise<Event> => {
    return axios
        .get(`${API_BASE_URL}/events/${eventID}`)
        .then((response) => response.data as Event)
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching event');
        });
};
