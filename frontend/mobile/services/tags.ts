import { AxiosError } from 'axios';

import { ErrorResponse } from '@/types/error';
import { Tag } from '@/types/tag';

import { api } from './api';

/**
 * Fetches all tags.
 *
 * @returns A Promise that resolves to the fetched tags or rejects with an error
 * @throws Error if the tags cannot be fetched
 * @see Tag
 */
const fetchTags = async (): Promise<Tag[]> => {
    return api
        .get<Tag[]>('/tags')
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error getting all tags');
        });
};

export { fetchTags };
