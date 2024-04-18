import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';
import { UseQueryResult, useQuery } from '@tanstack/react-query';

export type TTag = {
    name: string;
};

export const useClubTags = (
    clubID: uuid
): UseQueryResult<TTag[], Error> => {
    return useQuery<TTag[], Error>({
        queryKey: ['tags', clubID],
        queryFn: async (): Promise<TTag[]> => {
            return [
                { name: 'Science' },
                { name: 'Technology' },
                { name: 'Engineering' },
                { name: 'Mathematics' }
            ];
        }
    });
}

export const useTags = (): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['tags'],
        queryFn: getAllTags
    });
};
