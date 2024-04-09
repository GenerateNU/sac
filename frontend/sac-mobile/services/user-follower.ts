import { Club } from '@/types/club';
import { api } from './api';

export const GetUserFollowing = async (userId: string): Promise<Club[]> => {
    try {
        const response = await api.get(
            `/users/${userId}/follower`
        );
        return response.data as Club[];
    } catch (error) {
        console.log('Error fetching followed clubs');
        throw new Error('Error fetching followed clubs');
    }
};
