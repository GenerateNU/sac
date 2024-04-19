import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { queryClubs } from '@/services/club';
import { Club } from '@/types/club';

export const useClubSearch = (query: string): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['useClubSearchQuery', query],
        queryFn: () => {
            return queryClubs(query);
        }
    });
};
