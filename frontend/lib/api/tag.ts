import { z } from "zod";

import { API_BASE_URL } from "@/const";
import { Tag, tagSchema } from "@/types/tag";
import { uuid } from "@/types/uuid";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const tagsApi = createApi({
  reducerPath: "tagsApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  tagTypes: ["Tags"],
  endpoints: (builder) => ({
    tags: builder.query<Tag[], void>({
      query: () => "tags",
      transformResponse: (rawResponse: unknown) => {
        return z.array(tagSchema).parse(rawResponse);
      },
      providesTags: [{ type: "Tags", id: "LIST" }],
    }),
    tag: builder.query<Tag, uuid>({
      query: (tagID: uuid) => `tags/${tagID}`,
      transformResponse: (rawResponse: unknown) => {
        return tagSchema.parse(rawResponse);
      },
      providesTags: (result, _, tagID) =>
        result ? [{ type: "Tags", id: tagID }] : [],
    }),
  }),
});

export const { useTagsQuery, useTagQuery } = tagsApi;
