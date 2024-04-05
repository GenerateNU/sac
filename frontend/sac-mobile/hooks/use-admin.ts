import { UseQueryResult, useQuery } from '@tanstack/react-query';

import { uuid } from '@/types/uuid';

export type Admin = {
    id: uuid;
    name: string;
    title: string;
};

export const useAdmin = (
    clubID: uuid
): UseQueryResult<Admin[], Error> => {
    return useQuery<Admin[], Error>({
        queryKey: ['eboard-members', clubID],
        queryFn: async (): Promise<Admin[]> => {
            return [
                { id: 'alder', name: 'Alder Whiteford', title: 'Project Lead' },
                { id: 'amanda', name: 'Amanda Kerr', title: 'Design Lead' },
                { id: 'david', name: 'David Oduneye', title: 'Tech Lead' },
                { id: 'garrett', name: 'Garrett Ladley', title: 'Tech Lead' }
            ];
        }
    });
};
