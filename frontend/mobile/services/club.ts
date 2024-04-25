import { AxiosError } from 'axios';

import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
import { ErrorResponse } from '@/types/error';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';

import { api } from './api';

/**
 * Fetches all clubs.
 *
 * @returns A Promise that resolves to the fetched clubs or rejects with an error
 * @throws Error if the clubs cannot be fetched
 * @see Club
 */
// TODO: deal with pagination + deal with other params
const fetchClubs = async (): Promise<Club[]> => {
    return api
        .get<Club[]>('/clubs')
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error fetching clubs');
        });
};

/**
 * Fetches a club by its ID.
 *
 * @param clubID The ID of the club to fetch
 * @returns A Promise that resolves to the fetched club or rejects with an error
 * @throws Error if the club cannot be fetched
 * @see Club
 */
const fetchClub = async (clubID: uuid): Promise<Club> => {
    return api
        .get<Club>(`/clubs/${clubID}`)
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
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
 * @see Club
 */
const fetchClubContacts = async (clubID: uuid): Promise<Contact[]> => {
    return api
        .get<Contact[]>(`/clubs/${clubID}/contacts`)
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error fetching contacts');
        });
};

/**
 * Fetches all tags for a club by its ID.
 *
 * @param _clubID The ID of the club to fetch tags for
 * @returns A Promise that resolves to the fetched tags or rejects with an error
 * @throws Error if the tags cannot be fetched
 * @see Tag
 * @see Club
 */
const fetchClubTags = async (_clubID: uuid): Promise<Tag[]> => {
    // return api
    //     .get<Tag[]>(`/clubs/${clubID}/tags`)
    //     .then((response) => response.data)
    //     .catch((error: AxiosError<ErrorResponse>) => {
    //         console.error(error);
    //         throw new Error('Error fetching tags');
    //     });

    return [
        {
            id: '1',
            name: 'Software',
            created_at: new Date(),
            updated_at: new Date(),
            category_id: '1'
        },
        {
            id: '2',
            name: 'Design',
            created_at: new Date(),
            updated_at: new Date(),
            category_id: '2'
        },
        {
            id: '3',
            name: 'Professional',
            created_at: new Date(),
            updated_at: new Date(),
            category_id: '3'
        },
        {
            id: '4',
            name: 'Product Development',
            created_at: new Date(),
            updated_at: new Date(),
            category_id: '4'
        }
    ];
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
const queryClubs = async (queryString: string): Promise<Club[]> => {
    return api
        .get(`/clubs/`, {
            params: { search: queryString }
        })
        .then((response) => response.data as Club[])
        .catch((error) => {
            console.error(error);
            throw new Error('Error querying clubs');
        });
};

export { fetchClubContacts, fetchClub, fetchClubs, fetchClubTags, queryClubs };
