import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClubContacts } from '@/services/club-contact';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';

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
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        type: 'email',
                        content: 'generatenu@gmail.com'
                    } as Contact,
                    {
                        id: 'tester-slack',
                        createdAt: new Date(),
                        updatedAt: new Date(),
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
