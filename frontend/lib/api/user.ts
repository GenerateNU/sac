import { z } from "zod";

import { API_BASE_URL } from "@/const";
import { Club, clubSchema } from "@/types/club";
import { User, userSchema } from "@/types/user";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const usersApi = createApi({
  reducerPath: "usersApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  tagTypes: ["Users"],
  endpoints: (builder) => ({
    users: builder.query<User[], void>({
      query: () => "users",
      transformResponse: (rawResponse: unknown) => {
        return z.array(userSchema).parse(rawResponse);
      },
      providesTags: [{ type: "Users", id: "LIST" }],
    }),
    user: builder.query<User, uuid>({
      query: (userID: uuid) => `users/${userID}`,
      transformResponse: (rawResponse: unknown) => {
        return userSchema.parse(rawResponse);
      },
      providesTags: (result, _, userID) =>
        result ? [{ type: "Users", id: userID }] : [],
    }),
    userFollowing: builder.query<Club[], uuid>({
      query: (userID: uuid) => `users/${userID}/follower`,
      transformResponse: (rawResponse: unknown) => {
        return z.array(clubSchema).parse(rawResponse);
      },
      providesTags: (result, _, userID) =>
        result ? [{ type: "Users", id: userID }] : [],
    }),
    createUserClubFollowing: builder.mutation<
      void,
      { userID: uuid; clubID: uuid }
    >({
      query: ({ userID, clubID }) => ({
        url: `users/${userID}/follower/${clubID}`,
        method: "POST",
      }),
      invalidatesTags: ["Users"],
    }),
    deleteUserClubFollowing: builder.mutation<
      void,
      { userID: uuid; clubID: uuid }
    >({
      query: ({ userID, clubID }) => ({
        url: `users/${userID}/follower/${clubID}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Users"],
    }),
  }),
});

export const {
  useUsersQuery,
  useUserQuery,
  useUserFollowingQuery,
  useCreateUserClubFollowingMutation,
  useDeleteUserClubFollowingMutation,
} = usersApi;
