import { UseQueryResult, useQuery } from '@tanstack/react-query';

import {
    fetchEvent,
    fetchEventTags,
    fetchEvents,
    fetchHosts
} from '@/services/event';
import { Club } from '@/types/club';
import { Event } from '@/types/event';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';

export const useEvent = (eventID: uuid): UseQueryResult<Event, Error> => {
    return useQuery<Event, Error>({
        queryKey: ['useEventeventID', eventID],
        queryFn: () => {
            if (eventID === 'tester') {
                return {
                    id: 'tester',
                    created_at: new Date(),
                    updated_at: new Date(),
                    name: 'Tester Event',
                    preview: 'This is a preview',
                    content:
                        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut',
                    start_time: new Date(2024, 2, 15, 18, 30),
                    end_time: new Date(2024, 2, 15, 20, 0),
                    location: '1234 Tester St, Boston, MA 02115',
                    meeting_link: 'https://foo.com',
                    event_type: 'open',
                    is_recurring: false
                } as Event;
            } else {
                return fetchEvent(eventID);
            }
        }
    });
};

export const useEvents = (): UseQueryResult<Event[], Error> => {
    return useQuery<Event[], Error>({
        queryKey: ['events'],
        queryFn: () => {
            return fetchEvents();
        }
    });
};

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
                        name: 'Tester',
                        created_at: new Date(),
                        updated_at: new Date()
                    } as Tag,
                    {
                        id: 'tester1',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Tester1',
                        created_at: new Date(),
                        updated_at: new Date()
                    } as Tag,
                    {
                        id: 'tester2',
                        createdAt: new Date(),
                        updatedAt: new Date(),
                        name: 'Tester1',
                        created_at: new Date(),
                        updated_at: new Date()
                    } as Tag
                ];
            } else {
                return fetchEventTags(eventID);
            }
        }
    });
};

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
                        applicationLink: 'https://www.generatenu.com',
                        created_at: new Date(),
                        updated_at: new Date()
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
                        applicationLink: 'https://www.generatenu.com',
                        created_at: new Date(),
                        updated_at: new Date()
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
                        applicationLink: 'https://www.generatenu.com',
                        created_at: new Date(),
                        updated_at: new Date()
                    } as Club
                ];
            } else {
                return fetchHosts(eventID);
            }
        }
    });
};
