import { Club } from '@/types/club';
import { uuid } from '@/types/uuid';
import { api } from './api';
import { Contact } from '@/types/contact';
import { clubContacts } from '@/data/clubs';

/**
 * Fetches a club by its ID.
 *
 * @param clubID The ID of the club to fetch
 * @returns A Promise that resolves to the fetched club or rejects with an error
 * @throws Error if the club cannot be fetched
 * @see Club
 */
export const fetchClub = async (clubID: uuid): Promise<Club> => {
    return api
        .get(`/clubs/${clubID}`)
        .then((response) => response.data as Club)
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching club');
        });
};


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
    // return api
    //     .get(`/clubs/${clubID}/contacts`)
    //     .then((response) => response.data as Contact[])
    //     .catch((error) => {
    //         console.error(error);
    //         throw new Error('Error fetching contacts');
    //     });

    const allContacts = clubContacts[clubID as keyof typeof clubContacts] as Contact[];
    if (!allContacts) {
        throw new Error('Error fetching contacts');
    }

    return allContacts;
}
