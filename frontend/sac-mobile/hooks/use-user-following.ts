import { type UseQueryResult, useQuery } from '@tanstack/react-query';

import { GetUserFollowing } from '@/services/user-follower';
import { Club } from '@/types/club';

export const useUserFollowing = (
    userID: string
): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['user_club_followers'],
        queryFn: async () => await GetUserFollowing(userID)
    });
};
