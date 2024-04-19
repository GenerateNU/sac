import { AxiosError } from 'axios';

import { Category } from '@/types/category';
import { ErrorResponse } from '@/types/error';

import { api } from './api';

const fetchCategories = async (): Promise<Category[]> => {
    return api
        .get<Category[]>(`/categories`)
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error fetching categories');
        });
};

export { fetchCategories };
