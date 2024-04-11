import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { fetchClub, fetchClubContacts, fetchClubs } from '@/services/club';
import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';

// export const useClubs = (): UseQueryResult<Club[], Error> => {
//     return useQuery<Club[], Error>({
//         queryKey: ['clubs'],
//         queryFn: () => {
//             return fetchClubs();
//         }
//     })
// }

export const useClub = (clubID: uuid): UseQueryResult<Club, Error> => {
    return useQuery<Club, Error>({
        queryKey: ['club', clubID],
        queryFn: () => {
            return fetchClub(clubID);
        }
    });
};

export const useClubContacts = (
    clubID: uuid
): UseQueryResult<Contact[], Error> => {
    return useQuery<Contact[], Error>({
        queryKey: ['club', clubID, 'contacts'],
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
