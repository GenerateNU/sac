import { useMutation } from '@tanstack/react-query';

import { queryClient } from '@/app/_layout';
import { UpdateUser } from '@/services/user';

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
            queryClient.invalidateQueries({ queryKey: ['update_user'] });
        }
    });
};
