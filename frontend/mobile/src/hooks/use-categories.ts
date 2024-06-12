import { Category } from '@generatesac/lib';
import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { getCategories } from '@/src/services/categories';

export const useCategories = (): UseQueryResult<Category[], Error> => {
    return useQuery<Category[], Error>({
        queryKey: ['categories'],
        queryFn: getCategories
    });
};