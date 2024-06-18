import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const LOCAL_API_BASE_URL = "http://localhost:3000/api/v1";
export const PROD_API_BASE_URL = "https://studentactivitycalendar.xyz/api/v1";

// BaseAPI for the entire application:
export const baseApi = createApi({
  baseQuery: fetchBaseQuery({
    baseUrl: PROD_API_BASE_URL,
    credentials: "include",
    prepareHeaders: async (headers, { getState }) => {
      // User slice existing must exist in all dependent apps:
      const token = (getState() as { user: { accessToken: string } })?.user?.accessToken;
      if (token) {
        headers.set("_sac_session", token);
      }
      return headers;
    },
  }),
  tagTypes: [
    "User",
    "Tag",
    "Category",
    "Club",
    "Contact",
    "Event",
    "PointOfContact",
    "File",
    "Follower",
    "Member",
  ],
  endpoints: () => ({}),
});

// Add query params to a request:
export function handleQueryParams(
  baseUrl: string,
  queryParams?: Record<string, string | number | boolean>,
): string {
  let url = `${baseUrl}?`;
  if (queryParams) {
    Object.entries(queryParams).forEach(([key, value]) => {
      url += `${key}=${value}&`
    });
  }
  return url;
}
