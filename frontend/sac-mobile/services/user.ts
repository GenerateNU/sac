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

/**
 * Update user details after signup.
 * @param userId ID of this specific user
 * @param updateUserBody JSON body of updated information
 * @returns The user whose info is updated
 */
export const UpdateUser = async (
    userId: string,
    updateUserBody: {
        graduation_year: number;
        college: string;
        graduation_cycle: string;
    }
): Promise<User> => {
    try {
        const response = await api.patch<User>(
            `/users/${userId}/`,
            JSON.stringify(updateUserBody)
        );
        return response.data as User;
    } catch (error) {
        console.log('Error updating user');
        throw new Error('Error updating user');
    }
};
