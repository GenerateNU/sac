import { z } from 'zod';

import { File, fileSchema } from '../types/file';
import { PaginationQueryParams } from '../types/root';

import { baseApi, handleQueryParams } from '../api/base';

const FILE_API_BASE_URL = '/files';

export const fileApi = baseApi.injectEndpoints({
    endpoints: (builder) => ({
        files: builder.query<File[], PaginationQueryParams>({
            query: (queryParams) => ({
                url: handleQueryParams(`${FILE_API_BASE_URL}/`, queryParams),
                method: 'GET'
            }),
            providesTags: (result) =>
                result
                    ? result.map(({ id }) => ({ type: 'File', id }))
                    : ['File'],
            transformResponse: (response) => {
                return z.array(fileSchema).parse(response);
            }
        }),
        file: builder.query<File, string>({
            query: (id) => ({
                url: `${FILE_API_BASE_URL}/${id}`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result ? [{ type: 'File', id }] : [],
            transformResponse: (response) => {
                return fileSchema.parse(response);
            }
        }),
        createFile: builder.mutation<File, FormData>({
            query: (body) => ({
                url: `${FILE_API_BASE_URL}/`,
                method: 'POST',
                body
            }),
            invalidatesTags: ['File'],
            transformResponse: (response) => {
                return fileSchema.parse(response);
            }
        }),
        deleteFile: builder.mutation<void, string>({
            query: (id) => ({
                url: `${FILE_API_BASE_URL}/${id}`,
                method: 'DELETE'
            }),
            invalidatesTags: (_result, _, id) => [{ type: 'File', id }]
        })
    })
});

export const {
    useFilesQuery,
    useFileQuery,
    useCreateFileMutation,
    useDeleteFileMutation
} = fileApi;
