import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchTags } from '@/services/tags';
import { Tag } from '@/types/tag';

const useTags = (): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['tags'],
        queryFn: fetchTags
    });
};

export { useTags };
