import { uuid } from '@/types/uuid';
import { UseQueryResult, useQuery } from '@tanstack/react-query';

export type Tag = {
    name: string;
};

export const useTags = (
    clubID: uuid
): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['tags', clubID],
        queryFn: async (): Promise<Tag[]> => {
            return [
                { name: 'Science'},
                { name: 'Technology' },
                { name: 'Engineering' },
                { name: 'Mathematics' }
            ];
        }
    });
};
