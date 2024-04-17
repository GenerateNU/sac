import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClub, fetchClubContacts, fetchClubTags, queryClubs as searchForClubs } from '@/services/club';
import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';
import { Tag } from '@/types/tag';

export const useClubTags = (clubUUID: string): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['useClubTagsQueryID', clubUUID],
        queryFn: () => {
            return fetchClubTags(clubUUID);
        }
    });
};