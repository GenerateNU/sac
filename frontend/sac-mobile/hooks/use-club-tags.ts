import { fetchClubTags } from '@/services/club';
import { Tag } from '@/types/tag';
import { useQuery, UseQueryResult } from '@tanstack/react-query';

export const useClubTags = (clubUUID: string): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['useClubTagsQueryID', clubUUID],
        queryFn: () => {
            return fetchClubTags(clubUUID);
        }
    });
};
