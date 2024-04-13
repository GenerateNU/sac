import { API_BASE_URL } from "@/const";
import { User } from "@/types/user";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const usersApi = createApi({
  reducerPath: "usersApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  endpoints: (builder) => ({
    getCurrentUser: builder.query<User, void>({
      query: () => "users/me",
    }),
  }),
});

export const { useGetCurrentUserQuery } = usersApi;
