import { fetchTags } from '@/services/tags';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';
import { UseQueryResult, useQuery } from '@tanstack/react-query';

const useTags = (): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['tags'],
        queryFn: fetchTags
    });
};

export { useTags }