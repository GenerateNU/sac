import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClub } from '@/services/club';
import { Club } from '@/types/club';
import { uuid } from '@/types/uuid';

export const useClub = (clubID: uuid): UseQueryResult<Club, Error> => {
    return useQuery<Club, Error>({
        queryKey: ['useClubclubID', clubID],
        queryFn: () => {
            if (clubID === 'tester') {
                return {
                    id: 'tester',
                    createdAt: new Date(),
                    updatedAt: new Date(),
                    name: 'Generate',
                    preview: 'Generate',
                    description: 'Generate',
                    numMembers: 69420,
                    isRecruiting: true,
                    recruitmentCycle: 'always',
                    recruitmentType: 'unrestricted',
                    applicationLink: 'https://www.generatenu.com',
                    logo: 'https://www.generatenu.com/logo.png'
                } as Club;
            } else {
                return fetchClub(clubID);
            }
        }
    });
};
