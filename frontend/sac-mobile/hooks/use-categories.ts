import { type UseQueryResult, useQuery } from '@tanstack/react-query';

import { Category } from '@/types/category';

import { getAllCategories } from './../services/categories';

export const useCategories = (): UseQueryResult<Category[], Error> => {
    return useQuery<Category[], Error>({
        queryKey: ['category'],
        queryFn: getAllCategories
    });
};
