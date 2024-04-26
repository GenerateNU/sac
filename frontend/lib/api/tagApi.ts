import { z } from 'zod';

import { baseApi, handleQueryParams } from '../api/base';
import { PaginationQueryParams } from '@/types/root';
import {
    CreateTagRequestBody,
    Tag,
    UpdateTagRequestBody,
    tagSchema
} from '@/types/tag';

const TAG_API_BASE_URL = '/tags';

export const tagApi = baseApi.injectEndpoints({
    endpoints: (builder) => ({
        tags: builder.query<Tag[], PaginationQueryParams>({
            query: (queryParams) => ({
                url: handleQueryParams(`${TAG_API_BASE_URL}/`, queryParams),
                method: 'GET'
            }),
            providesTags: (result) =>
                result
                    ? result.map(({ id }) => ({ type: 'Tag', id }))
                    : ['Tag'],
            transformResponse: (response) => {
                return z.array(tagSchema).parse(response);
            }
        }),
        tag: builder.query<Tag, string>({
            query: (id) => ({
                url: `${TAG_API_BASE_URL}/${id}/`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result ? [{ type: 'Tag', id }] : [],
            transformResponse: (response) => {
                return tagSchema.parse(response);
            }
        }),
        createTag: builder.mutation<Tag, CreateTagRequestBody>({
            query: (body) => ({
                url: `${TAG_API_BASE_URL}/`,
                method: 'POST',
                body
            }),
            invalidatesTags: ['Tag'],
            transformResponse: (response) => {
                return tagSchema.parse(response);
            }
        }),
        updateTag: builder.mutation<
            Tag,
            { id: string; body: UpdateTagRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${TAG_API_BASE_URL}/${id}/`,
                method: 'PATCH',
                body
            }),
            invalidatesTags: (result, _, { id }) =>
                result ? [{ type: 'Tag', id }] : [],
            transformResponse: (response) => {
                return tagSchema.parse(response);
            }
        }),
        deleteTag: builder.mutation<void, string>({
            query: (id) => ({
                url: `${TAG_API_BASE_URL}/${id}/`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, id) => [{ type: 'Tag', id }]
        })
    })
});

export const {
    useTagsQuery,
    useTagQuery,
    useCreateTagMutation,
    useUpdateTagMutation,
    useDeleteTagMutation
} = tagApi;
