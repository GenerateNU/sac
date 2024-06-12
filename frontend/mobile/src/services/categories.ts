import { Category } from '@generatesac/lib';
import axios from 'axios';

import { URL } from './tags';

export const getCategories = async (): Promise<Category[]> => {
    try {
        const categories = await axios.get(`${URL}/categories`);
        return categories.data;
    } catch (error) {
        console.error('Failed to get categories', error);
        throw error;
    }
};
