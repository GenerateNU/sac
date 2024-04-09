import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClub, fetchClubContacts } from '@/services/club';
import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
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
                    logo: 'https://www.generatenu.com/logo.png',
                    created_at: new Date(),
                    updated_at: new Date()
                } as Club;
            } else {
                return fetchClub(clubID);
            }
        }
    });
};

export const useClubContacts = (
    clubID: uuid
): UseQueryResult<Contact[], Error> => {
    return useQuery<Contact[], Error>({
        queryKey: ['useClubContactsclubID', clubID],
        queryFn: () => {
            if (clubID.startsWith('tester')) {
                return [
                    {
                        id: 'tester-email',
                        created_at: new Date(),
                        updated_at: new Date(),
                        type: 'email',
                        content: 'generatenu@gmail.com'
                    } as Contact,
                    {
                        id: 'tester-slack',
                        created_at: new Date(),
                        updated_at: new Date(),
                        type: 'slack',
                        content: 'generatenu.slack.com'
                    }
                ];
            } else {
                return fetchClubContacts(clubID);
            }
        }
    });
};
