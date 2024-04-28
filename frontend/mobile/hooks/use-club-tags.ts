import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClubTags } from '@/services/club';
import { Tag } from '@/types/tag';

export const useClubTags = (clubUUID: string): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['useClubTagsQueryID', clubUUID],
        queryFn: () => {
            return fetchClubTags(clubUUID);
        }
    });
};
