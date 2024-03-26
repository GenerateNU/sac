import { getAllCategories } from './../services/categories';
import {useQuery, type UseQueryResult} from '@tanstack/react-query';
import {Category} from '@/types/category';

export const useCategories = (): UseQueryResult<Category[], Error> => {
    return useQuery<Category[], Error>({
        queryKey: ['category'],
        queryFn: getAllCategories,
    });
};