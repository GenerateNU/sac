import { type UseQueryResult, useQuery } from '@tanstack/react-query';

import { getAllTags } from '@/services/tags';
import { Tag } from '@/types/tag';

export const useTags = (): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['tag'],
        queryFn: getAllTags
    });
};
