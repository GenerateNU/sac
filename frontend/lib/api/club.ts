import { API_BASE_URL } from '@/const';
import { Club } from '@/types/club';
import { Contact } from '@/types/contact';
import { Tag } from '@/types/tag';
import { uuid } from '@/types/uuid';
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const clubsApi = createApi({
  reducerPath: "clubsApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  tagTypes: ["Clubs"],
  endpoints: (builder) => ({
    clubs: builder.query<Club[], void>({
      query: () => "clubs",
      // transformResponse: (rawResponse: unknown) => {
      //   return z.array(clubSchema).parse(rawResponse);
      // },
      providesTags: [{ type: "Clubs", id: "LIST" }],
    }),
    club: builder.query<Club, uuid>({
      query: (clubID: uuid) => `clubs/${clubID}`,
      // transformResponse: (rawResponse: unknown) => {
      //   return clubSchema.parse(rawResponse);
      // },
      providesTags: (result, _, clubID) =>
        result ? [{ type: "Clubs", id: clubID }] : [],
    }),
    clubContacts: builder.query<Contact[], uuid>({
      query: (clubID: uuid) => `clubs/${clubID}/contacts`,
      // transformResponse: (rawResponse: unknown) => {
      //   return z.array(contactSchema).parse(rawResponse);
      // },
      providesTags: (result, _, clubID) =>
        result ? [{ type: "Clubs", id: clubID }] : [],
    }),
    clubTags: builder.query<Tag[], uuid>({
      query: (clubID: uuid) => `clubs/${clubID}/tags`,
      // transformResponse: (rawResponse: unknown) => {
      //   return z.array(tagSchema).parse(rawResponse);
      // },
      providesTags: (result, _, clubID) =>
        result ? [{ type: "Clubs", id: clubID }] : [],
    }),

  }),
});

export const {
  useClubsQuery,
  useClubQuery,
  useClubContactsQuery,
  useClubTagsQuery
} = clubsApi;