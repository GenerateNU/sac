import {
    UseMutationResult,
    type UseQueryResult,
    useMutation,
    useQuery
} from '@tanstack/react-query';

import { queryClient } from '@/app/_layout';
import {
    createUserFollowing,
    createUserTags,
    deleteUserFollowing,
    fetchUserFollowing,
    updateUser
} from '@/services/user';
import { Club } from '@/types/club';

const useUserFollowing = (userID: string): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['user', userID, 'following'],
        queryFn: async () => await fetchUserFollowing(userID)
    });
};

const useUpdateUser = () => {
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
        }) => updateUser(userId, updateUserRequestBody),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['users'] });
        }
    });
};

const useCreateUserTags = (): UseMutationResult<
    void,
    Error,
    {
        userId: string;
        tagIDs: string[];
    },
    unknown
> => {
    return useMutation({
        mutationFn: ({ userId, tagIDs }) => createUserTags(userId, tagIDs),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user_tags'] });
        }
    });
};

const useCreateClubFollowing = (): UseMutationResult<
    void,
    Error,
    {
        userId: string;
        clubId: string;
    },
    unknown
> => {
    return useMutation({
        mutationFn: ({ userId, clubId }) => createUserFollowing(userId, clubId),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user_following'] });
        }
    });
};

const useDeleteClubFollowing = () => {
    return useMutation({
        mutationFn: ({ userId, clubId }: { userId: string; clubId: string }) =>
            deleteUserFollowing(userId, clubId),
        onSuccess: (_, variables: { userId: string; clubId: string }, __) => {
            return queryClient.invalidateQueries({
                queryKey: [`user${variables.userId}following`]
            });
        }
    });
};

export {
    useUserFollowing,
    useUpdateUser,
    useCreateUserTags,
    useCreateClubFollowing,
    useDeleteClubFollowing
};
