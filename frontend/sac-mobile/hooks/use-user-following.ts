import {useQuery, type UseQueryResult} from '@tanstack/react-query';
import {Club} from '@/types/club';
import { GetUserFollowing } from '@/services/user-follower';

export const useUserFollowing = (userID: string): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['user_club_followers'],
        queryFn: async () => await GetUserFollowing(userID)
    });
};