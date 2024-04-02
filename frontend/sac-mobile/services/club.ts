import axios from 'axios';

import { API_BASE_URL } from '@/lib/const';
import { Club } from '@/types/club';
import { uuid } from '@/types/uuid';

/**
 * Fetches a club by its ID.
 *
 * @param clubID The ID of the club to fetch
 * @returns A Promise that resolves to the fetched club or rejects with an error
 * @throws Error if the club cannot be fetched
 * @see Club
 */
export const fetchClub = async (clubID: uuid): Promise<Club> => {
    return axios
        .get(`${API_BASE_URL}/clubs/${clubID}`)
        .then((response) => response.data as Club)
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching club');
        });
};
