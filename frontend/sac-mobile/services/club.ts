import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';

import { api } from './api';
import { Tag } from '@/types/tag';

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
    return api
        .get(`/clubs/${clubID}/contacts`)
        .then((response) => response.data as Contact[])
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching contacts');
        });
};

/**
 * Fetch clubs according to a search query.
 * 
 * @param clubParams The parameters you'd like to search with.
 * @returns A Promise that resolves to the fetched clubs or rejects with an error
 * @throws Error if the clubs cannot be fetched
 * @see Club
 * @see Club[]
 */
// TODO: deal with pagination + deal with other params
export const queryClubs = async (queryString: string): Promise<Club[]> => {
    return api
        .get(`/clubs/`, {
            params: { search: queryString }
        })
        .then((response) => response.data as Club[])
        .catch((error) => {
            console.error(error);
            throw new Error('Error querying clubs');
        });
}

/** Get the tags of a club */
export const fetchClubTags = async (clubID: uuid): Promise<Tag[]> => {
    return api
        .get(`/clubs/${clubID}/tags`)
        .then((response) => response.data as Tag[])
        .catch((error) => {
            console.error(error);
            throw new Error('Error fetching contacts');
        });
};
