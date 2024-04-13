import { API_BASE_URL } from "@/const";
import { Club } from "@/types/club";
import { Contact } from "@/types/contact";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const clubsApi = createApi({
  reducerPath: "clubsApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  endpoints: (builder) => ({
    getClubByID: builder.query<Club, uuid>({
      query: (clubID) => `clubs/${clubID}`,
    }),
    getClubContactsByClubID: builder.query<Contact[], uuid>({
      query: (clubID) => `clubs/${clubID}/contacts`,
    }),
  }),
});

export const { useGetClubByIDQuery, useGetClubContactsByClubIDQuery } =
  clubsApi;
