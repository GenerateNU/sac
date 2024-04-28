import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const API_BASE_URL = 'http://127.0.0.1:8080/api/v1';

// BaseAPI for the entire application:
export const baseApi = createApi({
    baseQuery: fetchBaseQuery({
        baseUrl: API_BASE_URL,
        credentials: 'include',
        prepareHeaders: async (headers, { getState }) => {
            const token = (getState() as { auth: { token: string } })?.auth.token;
            if (token) {
                headers.set('Authorization', `Bearer ${token}`);
            }
            return headers;
        }
    }),
    tagTypes: [
        'User',
        'Tag',
        'Category',
        'Club',
        'Contact',
        'Event',
        'PointOfContact',
        'File',
        'Follower',
        'Member'
    ],
    endpoints: () => ({})
});

// Add query params to a request:
export function handleQueryParams(
    baseUrl: string,
    queryParams?: Record<string, string | number | boolean>
): string {
    const url = new URL(baseUrl);
    if (queryParams) {
        Object.entries(queryParams).forEach(([key, value]) => {
            url.searchParams.append(key, value.toString());
        });
    }
    return url.toString();
}