import { AxiosError } from 'axios';

import { api } from '@/services/api';
import { Club } from '@/types/club';
import { ErrorResponse } from '@/types/error';
import { User } from '@/types/user';

/**
 * Gets the user from the access token.
 * @param accessToken The access token of the user.
 * @returns The user that was retrieved.
 */
const getCurrentUser = async (): Promise<User> => {
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

const createUserTags = async (
    userId: string,
    tagsId: string[]
): Promise<void> => {
    try {
        await api.post<User>(`/users/${userId}/tags`, { tags: tagsId });
    } catch (error) {
        console.log('Error associating users to a tag');
        throw new Error('Error associating users to a tag');
    }
};

/**
 * Update user details after signup.
 * @param userId ID of this specific user
 * @param updateUserBody JSON body of updated information
 * @returns The user whose info is updated
 */
const updateUser = async (
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
            updateUserBody
        );
        return response.data as User;
    } catch (error) {
        console.log('Error updating user');
        throw new Error('Error updating user');
    }
};

/**
 * Fetches the clubs that the user is following.
 * @param userId The ID of the user.
 * @returns The clubs that the user is following.
 */
const fetchUserFollowing = async (userId: string): Promise<Club[]> => {
    try {
        const response = await api.get(`/users/${userId}/follower`);
        return response.data as Club[];
    } catch (error) {
        console.log('Error fetching followed clubs');
        throw new Error('Error fetching followed clubs');
    }
};

/**
 * Follows a club.
 * @param userId The ID of the user.
 * @param clubId The ID of the club.
 */
const createUserFollowing = async (
    userId: string,
    clubId: string
): Promise<void> => {
    try {
        await api.post(`/users/${userId}/follower/${clubId}`);
    } catch (error) {
        console.log('Error following clubs');
        throw new Error('Error following clubs');
    }
};

/**
 * Unfollows a club.
 * @param userId The ID of the user.
 * @param clubId The ID of the club.
 */
const deleteUserFollowing = async (
    userId: string,
    clubId: string
): Promise<void> => {
    try {
        await api.delete(`/users/${userId}/follower/${clubId}`);
    } catch (error) {
        console.log('Error unfollowing clubs');
        throw new Error('Error unfollowing clubs');
    }
};

export {
    getCurrentUser,
    createUserTags,
    updateUser,
    fetchUserFollowing,
    createUserFollowing,
    deleteUserFollowing
};
