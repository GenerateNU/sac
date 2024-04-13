import { API_BASE_URL } from "@/const";
import { Club } from "@/types/club";
import { Tag } from "@/types/tag";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const eventsApi = createApi({
  reducerPath: "eventsApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  endpoints: (builder) => ({
    getEventByID: builder.query<Event, uuid>({
      query: (eventID) => `events/${eventID}`,
    }),
    getEventHostsByEventID: builder.query<Club[], uuid>({
      query: (eventID) => `events/${eventID}/hosts`,
    }),
    getEventTagsByEventID: builder.query<Tag[], uuid>({
      query: (eventID) => `events/${eventID}/tags`,
    }),
    getEvents: builder.query<Event[], void>({
      query: () => "events",
    }),
  }),
});

export const {
  useGetEventByIDQuery,
  useGetEventHostsByEventIDQuery,
  useGetEventTagsByEventIDQuery,
  useGetEventsQuery,
} = eventsApi;
