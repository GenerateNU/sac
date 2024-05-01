import { z } from 'zod';

import { baseApi, handleQueryParams } from '@/api/base';
import { Club, clubSchema } from '@/types/club';
import {
    CreateEventRequestBody,
    Event,
    UpdateEventRequestBody,
    eventSchema
} from '@/types/event';
import { PaginationQueryParams } from '@/types/root';
import { Tag } from '@/types/tag';

const EVENT_API_BASE_URL = '/events';

export const eventApi = baseApi.injectEndpoints({
    endpoints: (builder) => ({
        events: builder.query<Event[], PaginationQueryParams>({
            query: (queryParams) => ({
                url: handleQueryParams(`${EVENT_API_BASE_URL}/`, queryParams),
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((event) => ({ type: 'Event', id: event.id }))
                    : ['Event'],
            transformResponse: (response) => {
                return z.array(eventSchema).parse(response);
            }
        }),
        event: builder.query<Event, string>({
            query: (id) => ({
                url: `${EVENT_API_BASE_URL}/${id}/`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result ? [{ type: 'Event', id }] : [],
            transformResponse: (response: Event) => {
                return eventSchema.parse(response);
            }
        }),
        createEvent: builder.mutation<Event, CreateEventRequestBody>({
            query: (body) => ({
                url: `${EVENT_API_BASE_URL}/`,
                method: 'POST',
                body
            }),
            invalidatesTags: ['Event'],
            transformResponse: (response) => {
                return eventSchema.parse(response);
            }
        }),
        updateEvent: builder.mutation<
            Event,
            { id: string; body: UpdateEventRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${EVENT_API_BASE_URL}/${id}`,
                method: 'PATCH',
                body
            }),
            invalidatesTags: (result, _, { id }) =>
                result ? [{ type: 'Event', id }] : [],
            transformResponse: (response) => {
                return eventSchema.parse(response);
            }
        }),
        deleteEvent: builder.mutation<void, string>({
            query: (id) => ({
                url: `${EVENT_API_BASE_URL}/${id}`,
                method: 'DELETE'
            }),
            invalidatesTags: (result, _, id) =>
                result ? [{ type: 'Event', id }] : []
        }),
        eventHosts: builder.query<Club[], string>({
            query: (id) => ({
                url: `${EVENT_API_BASE_URL}/${id}/hosts`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result ? [{ type: 'Event', id }, 'Club'] : ['Event'],
            transformResponse: (response) => {
                return z.array(clubSchema).parse(response);
            }
        }),
        eventTags: builder.query<Tag[], string>({
            query: (id) => ({
                url: `${EVENT_API_BASE_URL}/${id}/tags`,
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((tag) => ({ type: 'Tag', id: tag.id }))
                    : ['Tag']
        })
    })
});

export const {
    useEventsQuery,
    useEventQuery,
    useCreateEventMutation,
    useUpdateEventMutation,
    useDeleteEventMutation,
    useEventHostsQuery,
    useEventTagsQuery
} = eventApi;
