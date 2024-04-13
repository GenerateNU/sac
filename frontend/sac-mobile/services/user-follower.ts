import { Club } from '@/types/club';

import { api } from './api';

export const GetUserFollowing = async (userId: string): Promise<Club[]> => {
    try {
        const response = await api.get(`/users/${userId}/follower`);
        return response.data as Club[];
    } catch (error) {
        console.log('Error fetching followed clubs');
        throw new Error('Error fetching followed clubs');
    }
};

export const CreateUserFollowing = async (
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

export const DeleteUserFollowing = async (
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
