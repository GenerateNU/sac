import { z } from 'zod';

import { baseApi, handleQueryParams } from '../api/base';
import {
    Category,
    CreateCategoryRequestBody,
    categorySchema
} from '@/types/category';
import { PaginationQueryParams } from '@/types/root';
import { Tag } from '@/types/tag';

const CATEGORY_API_BASE_URL = '/categories';

export const categoryApi = baseApi.injectEndpoints({
    endpoints: (builder) => ({
        createCategory: builder.mutation<Category, CreateCategoryRequestBody>({
            query: (body) => ({
                url: CATEGORY_API_BASE_URL,
                method: 'POST',
                body
            }),
            transformResponse: (response: Category) => {
                return categorySchema.parse(response);
            },
            invalidatesTags: ['Category']
        }),
        categories: builder.query<Category[], PaginationQueryParams>({
            query: (queryParams) => ({
                url: handleQueryParams(
                    `${CATEGORY_API_BASE_URL}/`,
                    queryParams
                ),
                method: 'GET'
            }),
            transformResponse: (response: Category[]) => {
                return z.array(categorySchema).parse(response);
            },
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((category) => ({
                        type: 'Category',
                        id: category.id
                    }))
                    : ['Category']
        }),
        category: builder.query<Category, string>({
            query: (id) => ({
                url: `${CATEGORY_API_BASE_URL}/${id}/`,
                method: 'GET'
            }),
            transformResponse: (response: Category) => {
                return categorySchema.parse(response);
            },
            providesTags: (result, _, id) =>
                result ? [{ type: 'Category', id }] : []
        }),
        deleteCategory: builder.mutation<void, string>({
            query: (id) => ({
                url: `${CATEGORY_API_BASE_URL}/${id}/`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, id) => [{ type: 'Category', id }]
        }),
        updateCategory: builder.mutation<
            Category,
            { id: string; body: CreateCategoryRequestBody }
        >({
            query: ({ id, body }) => ({
                url: `${CATEGORY_API_BASE_URL}/${id}/`,
                method: 'PUT',
                body
            }),
            transformResponse: (response: Category) => {
                return categorySchema.parse(response);
            },
            invalidatesTags: (result, _, { id }) =>
                result ? [{ type: 'Category', id }] : []
        }),
        categoryTags: builder.query<
            Tag[],
            { id: string; queryParams?: PaginationQueryParams }
        >({
            query: ({ id, queryParams }) => ({
                url: handleQueryParams(
                    `${CATEGORY_API_BASE_URL}/${id}/tags/`,
                    queryParams
                ),
                method: 'GET'
            }),
            providesTags: (result, _, _arg) =>
                result
                    ? result.map((tag) => ({ type: 'Tag', id: tag.id }))
                    : ['Tag']
        }),
        categoryTag: builder.query<Tag, { categoryID: string; tagID: string }>({
            query: ({ categoryID, tagID }) => ({
                url: `${CATEGORY_API_BASE_URL}/${categoryID}/tags/${tagID}/`,
                method: 'GET'
            }),
            providesTags: (result, _, { tagID }) =>
                result ? [{ type: 'Tag', id: tagID }] : []
        })
    })
});

export const {
    useCreateCategoryMutation,
    useCategoriesQuery,
    useCategoryQuery,
    useDeleteCategoryMutation,
    useUpdateCategoryMutation,
    useCategoryTagsQuery,
    useCategoryTagQuery
} = categoryApi;
