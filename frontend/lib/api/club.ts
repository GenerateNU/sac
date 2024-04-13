import { z } from "zod";

import { Contact } from "@/api/contact";
import { API_BASE_URL } from "@/const";
import { rootModelSchema } from "@/types/root";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const clubSchemaIntermediate = z.object({
  name: z.string().max(255),
  preview: z.string().max(255),
  description: z.string().max(255),
  num_members: z.number(),
  is_recruiting: z.boolean(),
  recruitment_cycle: z.enum(["fall", "spring", "fallSpring", "always"]),
  recruitment_type: z.enum(["unrestricted", "tryout", "application"]),
  application_link: z.string().max(255),
  logo: z.string().max(255).optional(),
});

export const clubSchema = clubSchemaIntermediate.merge(rootModelSchema);
export type Club = z.infer<typeof clubSchema>;

export const clubsApi = createApi({
  reducerPath: "clubsApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  endpoints: (builder) => ({
    clubByID: builder.query<Club, uuid>({
      query: (clubID: uuid) => `clubs/${clubID}`,
      transformResponse: (rawResponse: unknown) => {
        return clubSchema.parse(rawResponse);
      },
    }),
    clubContactsByClubID: builder.query<Contact[], uuid>({
      query: (clubID: uuid) => `clubs/${clubID}/contacts`,
    }),
  }),
});

export const { useClubByIDQuery, useClubContactsByClubIDQuery } = clubsApi;
