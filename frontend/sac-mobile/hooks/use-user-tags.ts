import {
    QueryClient,
    UseMutationResult,
    type UseQueryResult,
    useMutation,
    useQuery
} from '@tanstack/react-query';

import { queryClient } from '@/app/_layout';
import { CreateUserTags } from '@/services/user-tag';

export const useCreateUserTags = () => {
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
