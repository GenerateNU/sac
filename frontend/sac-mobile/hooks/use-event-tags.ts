import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchEventTags } from '@/services/event-tag';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';

export const useEventTags = (eventID: uuid): UseQueryResult<Tag[], Error> => {
    return useQuery<Tag[], Error>({
        queryKey: ['useEventTagseventID', eventID],
        queryFn: () => {
            if (eventID === 'tester') {
                return [
                    {
                        id: 'tester0',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Tester'
                    } as Tag,
                    {
                        id: 'tester1',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Tester1'
                    } as Tag,
                    {
                        id: 'tester2',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Tester1'
                    } as Tag
                ];
            } else {
                return fetchEventTags(eventID);
            }
        }
    });
};
