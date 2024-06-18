import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const API_BASE_URL = "http://127.0.0.1:8080/api/v1";

// BaseAPI for the entire application:
export const baseApi = createApi({
  baseQuery: fetchBaseQuery({
    baseUrl: API_BASE_URL,
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
