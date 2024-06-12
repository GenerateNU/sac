import { Tag } from '@generatesac/lib';
import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { getTags } from '@/src/services/tags';

export const useTags = (): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['tags'],
        queryFn: getTags
    });
};