import { type UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchCategories } from '@/services/categories';
import { Category } from '@/types/category';

export const useCategories = (): UseQueryResult<Category[], Error> => {
    return useQuery<Category[], Error>({
        queryKey: ['category'],
        queryFn: fetchCategories
    });
};
