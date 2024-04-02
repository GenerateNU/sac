import axios from 'axios';

import { API_BASE_URL } from '@/lib/const';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';

/**
 * Fetches the tags associated with an event.
 *
 * @param eventID The ID of the event to fetch tags for
 * @returns A Promise that resolves to the fetched tags or rejects with an error
 * @throws Error if the tags cannot be fetched
 * @see Tag
 */
export const fetchEventTags = async (eventID: uuid): Promise<Tag[]> => {
    return axios
        .get(`${API_BASE_URL}/events/${eventID}/tags`)
        .then((response) => response.data as Tag[])
        .catch((error) => {
            console.error(error);
            throw new Error('Failed to fetch event tags');
        });
};
