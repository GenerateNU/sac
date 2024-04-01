// import { useQuery, UseQueryResult } from '@tanstack/react-query';
// import { uuid } from '@/types/uuid';
// import { Club } from '@/types/club';

// export const useClub = (clubID: uuid): UseQueryResult<Club, Error> => {
//     return useQuery({
//         queryKey: ['event', eventID],
//         queryFn: () => {
//             if (clubID === 'tester') {
//                 return {
//                     id: 'tester',
//                     name: 'Tester Event',
//                     preview: 'This is a preview',
//                     content: 'This is the content',
//                     startTime: new Date(2024, 2, 15, 18, 30),
//                     endTime: new Date(2024, 2, 15, 20, 0),
//                     location: 'Here',
//                     meetingLink: 'https://foo.com',
//                     eventType: 'open',
//                     isRecurring: false,
//                     createdAt: new Date(),
//                     updatedAt: new Date(),
//                     hosts: ['Generate', 'Generate1', 'Generate2', 'Generate3']
//                 } as Event;
//             } else {
//                 return fetchEvent(eventID);
//             }
//         }
//     });
// };