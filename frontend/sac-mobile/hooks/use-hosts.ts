import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchHosts } from '@/services/event';
import { Club } from '@/types/club';
import { uuid } from '@/types/uuid';

export const useHosts = (eventID: uuid): UseQueryResult<Club[], Error> => {
    return useQuery<Club[], Error>({
        queryKey: ['useHostsEventID', eventID],
        queryFn: () => {
            if (eventID === 'tester') {
                return [
                    {
                        id: 'tester0',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Generate0',
                        preview: 'Generate',
                        description: 'Generate',
                        numMembers: 69420,
                        isRecruiting: true,
                        recruitmentCycle: 'always',
                        recruitmentType: 'unrestricted',
                        applicationLink: 'https://www.generatenu.com'
                    } as Club,
                    {
                        id: 'tester1',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Generate1',
                        preview: 'Generate',
                        description: 'Generate',
                        numMembers: 69420,
                        isRecruiting: true,
                        recruitmentCycle: 'always',
                        recruitmentType: 'unrestricted',
                        applicationLink: 'https://www.generatenu.com'
                    } as Club,
                    {
                        id: 'tester2',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Generate2',
                        preview: 'Generate',
                        description: 'Generate',
                        numMembers: 69420,
                        isRecruiting: true,
                        recruitmentCycle: 'always',
                        recruitmentType: 'unrestricted',
                        applicationLink: 'https://www.generatenu.com'
                    } as Club
                ];
            } else {
                return fetchHosts(eventID);
            }
        }
    });
};
