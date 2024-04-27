import { z } from "zod";

import { API_BASE_URL } from "@/const";
import { Category, categorySchema } from "@/types/category";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const categoriesApi = createApi({
  reducerPath: "categoriesApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  tagTypes: ["Categories"],
  endpoints: (builder) => ({
    categories: builder.query<Category[], void>({
      query: () => "categories",
      transformResponse: (rawResponse: unknown) => {
        return z.array(categorySchema).parse(rawResponse);
      },
      providesTags: [{ type: "Categories", id: "LIST" }],
    }),
    category: builder.query<Category, uuid>({
      query: (categoryID: uuid) => `categories/${categoryID}`,
      transformResponse: (rawResponse: unknown) => {
        return categorySchema.parse(rawResponse);
      },
      providesTags: (result, _, categoryID) =>
        result ? [{ type: "Categories", id: categoryID }] : [],
    }),
  }),
});

export const { useCategoriesQuery, useCategoryQuery } = categoriesApi;
