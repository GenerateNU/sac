import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClub, fetchClubContacts, queryClubs as searchForClubs } from '@/services/club';
import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';
import { Tag } from '@/types/tag';

export const useClubSearch = (query: string): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['useClubSearchQuery', query],
        queryFn: () => {
            return searchForClubs(query);
        }
    });
};