import axios from 'axios';

import { API_BASE_URL } from '@/lib/const';
import { Tag } from '@/types/tag';
import { api } from './api';

export const getAllTags = async (): Promise<Tag[]> => {
    try {
        const response = await api.get('/tags');
        return response.data as Tag[];
    } catch (error) {
        console.log('Error fetching tags');
        throw new Error('Error fetching tags');
    }
};
