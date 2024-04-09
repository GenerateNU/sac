import { CreateUserTags } from '@/services/user-tag';
import { type UseQueryResult, useQuery, useMutation, UseMutationResult, QueryClient } from '@tanstack/react-query';
import { User } from '@/types/user';
import { queryClient } from '@/app/_layout';

export const useCreateUserTags = () => {
    return useMutation({
        mutationFn: ({ userId, tagIDs }: { userId: string, tagIDs: string[] }) => CreateUserTags(userId, tagIDs),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user_tags']})
        }
    });
};