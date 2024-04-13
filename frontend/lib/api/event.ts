import { z } from "zod";

import { Tag } from "@/api/tag";
import { API_BASE_URL } from "@/const";
import { rootModelSchema } from "@/types/root";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

import { Club } from "./club";

const eventSchemaIntermediate = z.object({
  name: z.string().max(255),
  preview: z.string().max(255),
  content: z.string().max(255),
  start_time: z.date(),
  end_time: z.date(),
  location: z.string().max(255),
  meeting_link: z.string().max(255).optional(),
  event_type: z.enum(["open", "membersOnly"]),
  is_recurring: z.boolean(),
  // hosts: z.array(z.string().max(255)).min(1)
});

export const eventSchema = eventSchemaIntermediate.merge(rootModelSchema);
export type Event = z.infer<typeof eventSchema>;

export const eventsApi = createApi({
  reducerPath: "eventsApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  endpoints: (builder) => ({
    eventByID: builder.query<Event, uuid>({
      query: (eventID) => `events/${eventID}`,
    }),
    eventHostsByEventID: builder.query<Club[], uuid>({
      query: (eventID) => `events/${eventID}/hosts`,
    }),
    eventTagsByEventID: builder.query<Tag[], uuid>({
      query: (eventID) => `events/${eventID}/tags`,
    }),
    events: builder.query<Event[], void>({
      query: () => "events",
    }),
  }),
});

export const {
  useEventByIDQuery,
  useEventHostsByEventIDQuery,
  useEventTagsByEventIDQuery,
  useEventsQuery,
} = eventsApi;
