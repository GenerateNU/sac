import { z } from 'zod';

import { baseApi, handleQueryParams } from '@/api/base';
import {
    Club,
    CreateClubRequestBody,
    CreateClubTagsRequestBody,
    UpdateClubRequestBody,
    clubSchema
} from '@/types/club';
import {
    Contact,
    CreateContactRequestBody,
    contactSchema
} from '@/types/contact';
import { Event, eventSchema } from '@/types/event';
import {
    PointOfContact,
    UpdatePointOfContactRequestBody,
    pointOfContactSchema
} from '@/types/pointOfContact';
import { PaginationQueryParams } from '@/types/root';
import { Tag, tagSchema } from '@/types/tag';
import { User, userSchema } from '@/types/user';

export const CLUB_API_BASE_URL = '/clubs';

export const clubApi = baseApi.injectEndpoints({
    endpoints: (builder) => ({
        createClub: builder.mutation<Club, CreateClubRequestBody>({
            query: (body) => ({
                url: `${CLUB_API_BASE_URL}/`,
                method: 'POST',
                body
            }),
            invalidatesTags: ['Club'],
            transformResponse: (response: Club) => {
                return clubSchema.parse(response);
            }
        }),
        clubs: builder.query<Club[], PaginationQueryParams>({
            query: (queryParams) => ({
                url: handleQueryParams(`${CLUB_API_BASE_URL}/`, queryParams),
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((club) => ({ type: 'Club', id: club.id }))
                    : ['Club'],
            transformResponse: (response: Club) => {
                return z.array(clubSchema).parse(response);
            }
        }),
        club: builder.query<Club, string>({
            query: (id) => ({
                url: `${CLUB_API_BASE_URL}/${id}/`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result ? [{ type: 'Club', id }] : [],
            transformResponse: (response: Club) => {
                return clubSchema.parse(response);
            }
        }),
        updateClub: builder.mutation<
            Club,
            { id: string; body: UpdateClubRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${CLUB_API_BASE_URL}/${id}/`,
                method: 'PATCH',
                body
            }),
            invalidatesTags: (result, _, { id }) =>
                result ? [{ type: 'Club', id }] : [],
            transformResponse: (response: Club) => {
                return clubSchema.parse(response);
            }
        }),
        deleteClub: builder.mutation<void, string>({
            query: (id) => ({
                url: `${CLUB_API_BASE_URL}/${id}/`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, id) => [{ type: 'Club', id }]
        }),
        clubContacts: builder.query<Contact[], string>({
            query: (id) => ({
                url: `${CLUB_API_BASE_URL}/${id}/contacts/`,
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((contact) => ({
                        type: 'Contact',
                        id: contact.id
                    }))
                    : ['Contact'],
            transformResponse: (response: Contact) => {
                return z.array(contactSchema).parse(response);
            }
        }),
        createContact: builder.mutation<
            Contact,
            { id: string; body: CreateContactRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${CLUB_API_BASE_URL}/${id}/contacts/`,
                method: 'PUT',
                body
            }),
            invalidatesTags: ['Contact'],
            transformResponse: (response: Contact) => {
                return contactSchema.parse(response);
            }
        }),
        clubEvents: builder.query<
            Event[],
            { id: string; queryParams?: PaginationQueryParams }
        >({
            query: ({ id, queryParams }) => ({
                url: handleQueryParams(
                    `${CLUB_API_BASE_URL}/${id}/events/`,
                    queryParams
                ),
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
        clubFollowers: builder.query<
            User[],
            { id: string; queryParams?: PaginationQueryParams }
        >({
            query: ({ id, queryParams }) => ({
                url: handleQueryParams(
                    `${CLUB_API_BASE_URL}/${id}/followers/`,
                    queryParams
                ),
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((follower) => ({
                        type: 'Follower',
                        id: follower.id
                    }))
                    : ['Follower'],
            transformResponse: (response) => {
                return z.array(userSchema).parse(response);
            }
        }),
        clubMembers: builder.query<
            User[],
            { id: string; queryParams?: PaginationQueryParams }
        >({
            query: ({ id, queryParams }) => ({
                url: handleQueryParams(
                    `${CLUB_API_BASE_URL}/${id}/members/`,
                    queryParams
                ),
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((member) => ({ type: 'User', id: member.id }))
                    : ['User'],
            transformResponse: (response) => {
                return z.array(userSchema).parse(response);
            }
        }),
        clubPointOfContacts: builder.query<PointOfContact[], string>({
            query: (id) => ({
                url: `${CLUB_API_BASE_URL}/${id}/pocs/`,
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((poc) => ({
                        type: 'PointOfContact',
                        id: poc.id
                    }))
                    : ['PointOfContact'],
            transformResponse: (response) => {
                return z.array(pointOfContactSchema).parse(response);
            }
        }),
        clubPointOfContact: builder.query<
            PointOfContact,
            { clubId: string; pocId: string }
        >({
            query: ({ clubId, pocId }) => ({
                url: `${CLUB_API_BASE_URL}/${clubId}/pocs/${pocId}`,
                method: 'GET'
            }),
            providesTags: (result, _, { pocId }) =>
                result ? [{ type: 'PointOfContact', id: pocId }] : [],
            transformResponse: (response) => {
                return pointOfContactSchema.parse(response);
            }
        }),
        createClubPointOfContact: builder.mutation<
            PointOfContact,
            { id: string; body: FormData }
        >({
            query: ({ id, body }) => ({
                url: `${CLUB_API_BASE_URL}/${id}/pocs/`,
                method: 'POST',
                body
            }),
            invalidatesTags: ['PointOfContact'],
            transformResponse: (response) => {
                return pointOfContactSchema.parse(response);
            }
        }),
        updateClubPointOfContactPhoto: builder.mutation<
            PointOfContact,
            { clubId: string; pocId: string; body: FormData }
        >({
            query: ({ clubId, pocId, body }) => ({
                url: `${CLUB_API_BASE_URL}/${clubId}/pocs/${pocId}/photo`,
                method: 'PATCH',
                body
            }),
            invalidatesTags: (result, _, { pocId }) =>
                result ? [{ type: 'PointOfContact', id: pocId }] : [],
            transformResponse: (response) => {
                return pointOfContactSchema.parse(response);
            }
        }),
        updateClubPointOfContact: builder.mutation<
            PointOfContact,
            {
                clubId: string;
                pocId: string;
                body: UpdatePointOfContactRequestBody;
            }
        >({
            query: ({ clubId, pocId, body }) => ({
                url: `${CLUB_API_BASE_URL}/${clubId}/pocs/${pocId}`,
                method: 'PATCH',
                body
            }),
            invalidatesTags: (result, _, { pocId }) =>
                result ? [{ type: 'PointOfContact', id: pocId }] : [],
            transformResponse: (response) => {
                return pointOfContactSchema.parse(response);
            }
        }),
        deleteClubPointOfContact: builder.mutation<
            void,
            { clubId: string; pocId: string }
        >({
            query: ({ clubId, pocId }) => ({
                url: `${CLUB_API_BASE_URL}/${clubId}/pocs/${pocId}`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, { pocId }) => [
                { type: 'PointOfContact', id: pocId }
            ]
        }),
        createClubTags: builder.mutation<
            Tag[],
            { id: string; body: CreateClubTagsRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${CLUB_API_BASE_URL}/${id}/tags/`,
                method: 'POST',
                body
            }),
            invalidatesTags: (result, _, _arg) =>
                result
                    ? result.map((tag) => ({ type: 'Tag', id: tag.id }))
                    : ['Tag'],
            transformResponse: (response) => {
                return z.array(tagSchema).parse(response);
            }
        }),
        clubTags: builder.query<Tag[], string>({
            query: (id) => ({
                url: `${CLUB_API_BASE_URL}/${id}/tags/`,
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((tag) => ({ type: 'Tag', id: tag.id }))
                    : ['Tag'],
            transformResponse: (response) => {
                return z.array(tagSchema).parse(response);
            }
        }),
        deleteClubTag: builder.mutation<
            void,
            { clubId: string; tagId: string }
        >({
            query: ({ clubId, tagId }) => ({
                url: `${CLUB_API_BASE_URL}/${clubId}/tags/${tagId}`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, { tagId }) => [
                { type: 'Tag', id: tagId }
            ]
        })
    })
});

export const {
    useCreateClubMutation,
    useClubsQuery,
    useClubQuery,
    useUpdateClubMutation,
    useDeleteClubMutation,
    useClubContactsQuery,
    useCreateContactMutation,
    useClubEventsQuery,
    useClubFollowersQuery,
    useClubMembersQuery,
    useClubPointOfContactsQuery,
    useClubPointOfContactQuery,
    useCreateClubPointOfContactMutation,
    useUpdateClubPointOfContactPhotoMutation,
    useUpdateClubPointOfContactMutation,
    useDeleteClubPointOfContactMutation,
    useCreateClubTagsMutation,
    useClubTagsQuery,
    useDeleteClubTagMutation
} = clubApi;
