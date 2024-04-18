import { Category } from '@/types/category';

import { api } from './api';
import { AxiosError } from 'axios';
import { ErrorResponse } from '@/types/error';

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