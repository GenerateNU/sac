import {
    UseMutationResult,
    type UseQueryResult,
    useMutation,
    useQuery
} from '@tanstack/react-query';

import { queryClient } from '@/app/_layout';
import { UpdateUser } from '@/services/user';
import {
    CreateUserFollowing,
    DeleteUserFollowing,
    GetUserFollowing
} from '@/services/user-follower';
import { CreateUserTags } from '@/services/user-tag';
import { Club } from '@/types/club';

export const useUserFollowing = (
    userID: string
): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['user', userID, 'following'],
        queryFn: async () => await GetUserFollowing(userID)
    });
};

export const useUpdateUser = () => {
    return useMutation({
        mutationFn: ({
            userId,
            updateUserRequestBody
        }: {
            userId: string;
            updateUserRequestBody: {
                graduation_year: number;
                college: string;
                graduation_cycle: string;
            };
        }) => UpdateUser(userId, updateUserRequestBody),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['users'] });
        }
    });
};

export const useCreateUserTags = (): UseMutationResult<
    void,
    Error,
    {
        userId: string;
        tagIDs: string[];
    },
    unknown
> => {
    return useMutation({
        mutationFn: ({
            userId,
            tagIDs
        }: {
            userId: string;
            tagIDs: string[];
        }) => CreateUserTags(userId, tagIDs),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user_tags'] });
        }
    });
};

export const useCreateClubFollowing = () => {
    return useMutation({
        mutationFn: ({ userId, clubId }: { userId: string; clubId: string }) =>
            CreateUserFollowing(userId, clubId),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user_following'] });
        }
    });
};

export const useDeleteClubFollowing = () => {
    return useMutation({
        mutationFn: ({ userId, clubId }: { userId: string; clubId: string }) =>
            DeleteUserFollowing(userId, clubId),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user_following'] });
        }
    });
};
