import { z } from 'zod';

import { baseApi, handleQueryParams } from './base';
import { UpdatePasswordRequestBody } from '../types/auth';
import { Club, clubSchema } from '../types/club';
import { PaginationQueryParams } from '../types/root';
import { Tag, tagSchema } from '../types/tag';
import {
    CreateUserRequestBody,
    CreateUserTagsRequestBody,
    UpdateUserRequestBody,
    User,
    userSchema
} from '../types/user';

const USER_API_BASE_URL = '/users';

export const userApi = baseApi.injectEndpoints({
    endpoints: (builder) => ({
        users: builder.query<User[], PaginationQueryParams>({
            query: (queryParams) => ({
                url: handleQueryParams(`${USER_API_BASE_URL}/`, queryParams),
                method: 'GET'
            }),
            providesTags: (result) =>
                result
                    ? result.map(({ id }) => ({ type: 'User', id }))
                    : ['User'],
            transformResponse: (response) => {
                return z.array(userSchema).parse(response);
            }
        }),
        user: builder.query<User, string>({
            query: (id) => ({
                url: `${USER_API_BASE_URL}/${id}/`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result ? [{ type: 'User', id }] : [],
            transformResponse: (response) => {
                return userSchema.parse(response);
            }
        }),
        currentUser: builder.query<User, void>({
            query: () => ({
                url: `${USER_API_BASE_URL}/me`,
                method: 'GET'
            }),
            providesTags: (result) =>
                result ? [{ type: 'User', id: result.id }] : [],
            transformResponse: (response) => {
                return userSchema.parse(response);
            }
        }),
        createUser: builder.mutation<User, CreateUserRequestBody>({
            query: (body) => ({
                url: `${USER_API_BASE_URL}/`,
                method: 'POST',
                body
            }),
            invalidatesTags: ['User'],
            transformResponse: (response) => {
                return userSchema.parse(response);
            }
        }),
        updateUser: builder.mutation<
            User,
            { id: string; body: UpdateUserRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${USER_API_BASE_URL}/${id}/`,
                method: 'PATCH',
                body
            }),
            invalidatesTags: (result, _, { id }) =>
                result ? [{ type: 'User', id }] : [],
            transformResponse: (response) => {
                return userSchema.parse(response);
            }
        }),
        deleteUser: builder.mutation<void, string>({
            query: (id) => ({
                url: `${USER_API_BASE_URL}/${id}/`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, id) => [{ type: 'User', id }]
        }),
        updatePassword: builder.mutation<
            void,
            { id: string; body: UpdatePasswordRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${USER_API_BASE_URL}/${id}/password`,
                method: 'PATCH',
                body
            })
        }),
        userFollowing: builder.query<Club[], string>({
            query: (id) => ({
                url: `${USER_API_BASE_URL}/${id}/follower/`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result
                    ? [{ type: 'Follower', id }, 'Club']
                    : [{ type: 'Follower', id }],
            transformResponse: (response) => {
                return z.array(clubSchema).parse(response);
            }
        }),
        createUserFollowing: builder.mutation<
            void,
            { userID: string; clubID: string }
        >({
            query: ({ userID, clubID }) => ({
                url: `${USER_API_BASE_URL}/${userID}/follower/${clubID}`,
                method: 'POST'
            }),
            invalidatesTags: (_result, _, { userID }) => [
                { type: 'Follower', id: userID }
            ]
        }),
        deleteUserFollowing: builder.mutation<
            void,
            { userID: string; clubID: string }
        >({
            query: ({ userID, clubID }) => ({
                url: `${USER_API_BASE_URL}/${userID}/follower/${clubID}`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, { userID }) => [
                { type: 'Follower', id: userID }
            ]
        }),
        userMembership: builder.query<Club[], string>({
            query: (id) => ({
                url: `${USER_API_BASE_URL}/${id}/member/`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result
                    ? [{ type: 'Member', id }, 'Club']
                    : [{ type: 'Member', id }],
            transformResponse: (response) => {
                return z.array(clubSchema).parse(response);
            }
        }),
        createUserMembership: builder.mutation<
            void,
            { userID: string; clubID: string }
        >({
            query: ({ userID, clubID }) => ({
                url: `${USER_API_BASE_URL}/${userID}/member/${clubID}`,
                method: 'POST'
            }),
            invalidatesTags: (_result, _, { userID }) => [
                { type: 'Follower', id: userID },
                { type: 'Member', id: userID }
            ]
        }),
        deleteUserMembership: builder.mutation<
            void,
            { userID: string; clubID: string }
        >({
            query: ({ userID, clubID }) => ({
                url: `${USER_API_BASE_URL}/${userID}/member/${clubID}`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, { userID }) => [
                { type: 'Member', id: userID }
            ]
        }),
        userTags: builder.query<Tag[], void>({
            query: () => ({
                url: `${USER_API_BASE_URL}/tags/`,
                method: 'GET'
            }),
            providesTags: (result) =>
                result
                    ? result.map((tag) => ({ type: 'Tag', id: tag.id }))
                    : ['Tag'],
            transformResponse: (response) => {
                return z.array(tagSchema).parse(response);
            }
        }),
        createUserTags: builder.mutation<Tag[], CreateUserTagsRequestBody>({
            query: (body) => ({
                url: `${USER_API_BASE_URL}/tags/`,
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
        deleteUserTag: builder.mutation<void, string>({
            query: (id) => ({
                url: `${USER_API_BASE_URL}/tags/${id}/`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, id) => [{ type: 'Tag', id }]
        })
    })
});
