import { queryClubs } from '@/services/club';
import { Club } from '@/types/club';
import { useQuery, UseQueryResult } from '@tanstack/react-query';

export const useClubSearch = (query: string): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['useClubSearchQuery', query],
        queryFn: () => {
            return queryClubs(query);
        }
    });
};
