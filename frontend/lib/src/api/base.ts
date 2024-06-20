import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const LOCAL_API_BASE_URL = "http://127.0.0.1:8080/api/v1";
export const PROD_API_BASE_URL = "https://studentactivitycalendar.xyz/api/v1";

type GlobalState = {
  platform: "web" | "mobile";
  loggedIn: boolean;
  accessToken?: string;
};

// BaseAPI for the entire application:
export const baseApi = createApi({
  baseQuery: fetchBaseQuery({
    baseUrl: LOCAL_API_BASE_URL,
    credentials: "include",
    prepareHeaders: async (headers, { getState }) => {
      // User slice existing must exist in all dependent apps:
      const globalState = (getState() as { global: GlobalState })?.global;
      if (globalState.accessToken) {
        // Pass the authentication header:
        headers.set("_sac_session", globalState.accessToken);
      }
      // Set the platform header:
      headers.set("Platform", globalState.platform);
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
