import axios from 'axios';
import { Category } from '@/types/category';
import { API_BASE_URL } from '@/lib/const';


export const getAllCategories = async (): Promise<Category[]> => {
    try {
        const response = await axios.get(`${API_BASE_URL}/categories`)
        return response.data as Category[]
    } catch (error) {
        console.log('Error fetching categories');
        throw new Error('Error fetching categories')
    }
}