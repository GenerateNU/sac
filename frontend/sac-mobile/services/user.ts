import { AxiosError } from 'axios';

import { api } from '@/services/api';
import { ErrorResponse } from '@/types/error';
import { User } from '@/types/user';

/**
 * Gets the user from the access token.
 * @param accessToken The access token of the user.
 * @returns The user that was retrieved.
 */
export const getCurrentUser = async (): Promise<User> => {
    try {
        const response = await api.get<User>('/users/me');

        return response.data as User;
    } catch (error: any) {
        const axiosErr: AxiosError = error;
        const response = axiosErr.response;
        if (response) {
            const data = response.data as ErrorResponse;
            throw new Error(data.error);
        } else {
            throw new Error('Error getting current user');
        }
    }
};
