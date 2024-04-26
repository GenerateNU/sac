import { API_BASE_URL } from "@/const";
import { Club } from "@/types/club";
import { Event } from "@/types/event";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const eventsApi = createApi({
  reducerPath: "eventsApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  tagTypes: ["Events"],
  endpoints: (builder) => ({
    events: builder.query<Event[], void>({
      query: () => "events",
      // transformResponse: (rawResponse: unknown) => {
      //   return z.array(eventSchema).parse(rawResponse);
      // },
      providesTags: [{ type: "Events", id: "LIST" }],
    }),
    event: builder.query<Event, uuid>({
      query: (eventID: uuid) => `events/${eventID}`,
      // transformResponse: (rawResponse: unknown) => {
      //   return eventSchema.parse(rawResponse);
      // },
      providesTags: (result, _, eventID) =>
        result ? [{ type: "Events", id: eventID }] : [],
    }),
    eventTags: builder.query<Club[], uuid>({
      query: (eventID: uuid) => `events/${eventID}/tags`,
      // transformResponse: (rawResponse: unknown) => {
      //   return z.array(clubSchema).parse(rawResponse);
      // },
      providesTags: (result, _, eventID) =>
        result ? [{ type: "Events", id: eventID }] : [],
    }),
    eventHosts: builder.query<Club[], uuid>({
      query: (eventID: uuid) => `events/${eventID}/hosts`,
      // transformResponse: (rawResponse: unknown) => {
      //   return z.array(clubSchema).parse(rawResponse);
      // },
      providesTags: (result, _, eventID) =>
        result ? [{ type: "Events", id: eventID }] : [],
    }),
  }),
});

export const {
  useEventsQuery,
  useEventQuery,
  useEventTagsQuery,
  useEventHostsQuery,
} = eventsApi;
