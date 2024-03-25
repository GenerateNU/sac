import {useQuery, type UseQueryResult} from '@tanstack/react-query';
import {Tag} from '@/types/tag';
import { getAllTags } from '@/services/tags';

export const useTags = (): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['tags'],
        queryFn: getAllTags
    });
};