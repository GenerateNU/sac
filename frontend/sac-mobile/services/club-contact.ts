import axios from 'axios';

import { API_BASE_URL } from '@/lib/const';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';

/**
 * Fetches all contacts for a club by its ID.
 *
 * @param clubID The ID of the club to fetch contacts for
 * @returns A Promise that resolves to the fetched contacts or rejects with an error
 * @throws Error if the contacts cannot be fetched
 * @see Contact
 * @see Contact[]
 */
export const fetchClubContacts = async (clubID: uuid): Promise<Contact[]> => {
    return axios
        .get(`${API_BASE_URL}/clubs/${clubID}/contacts`)
        .then((response) => response.data as Contact[])
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching contacts');
        });
};
