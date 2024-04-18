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
import { uuid } from '@/types/uuid';

const useUserFollowing = (userID: uuid): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['user', userID, 'following'],
        queryFn: async () => await fetchUserFollowing(userID)
    });
};

const useUpdateUser = () => {
    return useMutation({
        mutationFn: ({
            userID,
            updateUserRequestBody
        }: {
                userID: uuid;
            updateUserRequestBody: {
                graduation_year: number;
                college: string;
                graduation_cycle: string;
            };
            }) => updateUser(userID, updateUserRequestBody),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user'] });
        }
    });
};

const useCreateUserTags = (): UseMutationResult<
    void,
    Error,
    {
        userID: uuid;
        tagIDs: uuid[];
    },
    unknown
> => {
    return useMutation({
        mutationFn: ({ userID, tagIDs }) => createUserTags(userID, tagIDs),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user_tags'] });
        }
    });
};

const useCreateClubFollowing = (): UseMutationResult<
    void,
    Error,
    {
        userID: uuid;
        clubID: uuid;
    },
    unknown
> => {
    return useMutation({
        mutationFn: ({ userID, clubID }) => createUserFollowing(userID, clubID),
        onSuccess: (_, variables: { userID: uuid; clubID: uuid }, __) => {
            queryClient.invalidateQueries({ queryKey: [`user${variables.userID}following`] });
        }
    });
};

const useDeleteClubFollowing = () => {
    return useMutation({
        mutationFn: ({ userID, clubID }: { userID: uuid; clubID: uuid }) =>
            deleteUserFollowing(userID, clubID),
        onSuccess: (_, variables: { userID: string; clubID: string }, __) => {
            return queryClient.invalidateQueries({
                queryKey: [`user${variables.userID}following`]
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
