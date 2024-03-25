import axios from 'axios';
import { Tag } from '@/types/tag';
import { API_BASE_URL } from '@/lib/const';

export const getAllTags = async (): Promise<Tag[]> => {
    try {
        const response = await axios.get(`${API_BASE_URL}/tags`)
        return response.data as Tag[]
    } catch (error) {
        console.log('Error fetching tags');
        throw new Error('Error fetching tags')
    }
}