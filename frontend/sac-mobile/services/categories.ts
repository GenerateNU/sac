import axios from 'axios';

import { API_BASE_URL } from '@/lib/const';
import { Category } from '@/types/category';

import { api } from './api';

export const getAllCategories = async (): Promise<Category[]> => {
    try {
        const response = await api.get(`${API_BASE_URL}/categories`);
        return response.data as Category[];
    } catch (error) {
        console.log('Error fetching categories');
        throw new Error('Error fetching categories');
    }
};
