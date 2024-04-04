import axios from 'axios';
import { Category } from '@/types/category';
import { API_BASE_URL } from '@/lib/const';
import { Club } from '@/types/club';

export const GetUserFollowing = async (userId : string): Promise<Club[]> => {
    try {
        const response = await axios.get(`${API_BASE_URL}/users/${userId}/follower`)
        return response.data as Club[]
    } catch (error) {
        console.log('Error fetching followed clubs');
        throw new Error('Error fetching followed clubs')
    }
}