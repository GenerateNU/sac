import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const API_BASE_URL = 'http://localhost:8080/api/v1';

export const baseQuery = fetchBaseQuery({
    baseUrl: API_BASE_URL,
    credentials: 'include',
    prepareHeaders: async (headers, { getState }) => {
        console.log((getState() as { auth: any })?.auth);
        const token = (getState() as { auth: any })?.auth.token || null;
        if (token) {
            headers.set('Authorization', `Bearer ${token}`);
        }
        return headers;
    }
});

export const baseApi = createApi({
    baseQuery: baseQuery,
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
