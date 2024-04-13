import { z } from "zod";

import { API_BASE_URL } from "@/const";
import { rootModelSchema } from "@/types/root";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const userSchemaIntermediate = z.object({
  role: z.enum(["super", "student"]),
  email: z.string().email(),
  first_name: z.string().min(3),
  last_name: z.string().min(3),
  created_at: z.date(),
  updated_at: z.date(),
  is_verified: z.boolean(),
});

export const collegeSchema = z.enum([
  "CAMD",
  "DMSB",
  "KCCS",
  "CE",
  "BCHS",
  "SL",
  "CPS",
  "CS",
  "CSSH",
]);
export type College = z.infer<typeof collegeSchema>;

export const yearSchema = z.enum(["1", "2", "3", "4", "5"]);
export type Year = z.infer<typeof yearSchema>;

export const userSchema = userSchemaIntermediate.merge(rootModelSchema);
export type User = z.infer<typeof userSchema>;

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
