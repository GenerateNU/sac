import {useMutation} from '@tanstack/react-query';
import { queryClient } from '@/app/_layout';
import { UpdateUser } from '@/services/user';

export const useUpdateUser = () => {
    return useMutation({
        mutationFn: ({userId, tagIDs}: {userId: string; updateRequestBody: string[];}) => UpdateUser(userId, tagIDs),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['user'] });
        }
    });
};
