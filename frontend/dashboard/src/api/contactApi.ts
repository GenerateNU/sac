import { z } from "zod";

import { Contact, contactSchema } from "../types/contact";
import { PaginationQueryParams } from "../types/root";
import { baseApi, handleQueryParams } from "./api";

const CONTACT_API_BASE_URL = "/contacts";

const contactApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    contact: builder.query<Contact, string>({
      query: (id) => ({
        url: `${CONTACT_API_BASE_URL}/${id}`,
        method: "GET",
      }),
      providesTags: (result, _, id) =>
        result ? [{ type: "Contact", id }] : [],
      transformResponse: (response) => {
        return contactSchema.parse(response);
      },
    }),
    contacts: builder.query<Contact[], PaginationQueryParams>({
      query: (queryParams) => ({
        url: handleQueryParams(`${CONTACT_API_BASE_URL}/`, queryParams),
        method: "GET",
      }),
      providesTags: (result) =>
        result
          ? result.map(({ id }) => ({ type: "Contact", id }))
          : ["Contact"],
      transformResponse: (response) => {
        return z.array(contactSchema).parse(response);
      },
    }),
    deleteContact: builder.mutation<void, string>({
      query: (id) => ({
        url: `${CONTACT_API_BASE_URL}/${id}`,
        method: "DELETE",
      }),
      invalidatesTags: (_result, _, id) => [{ type: "Contact", id }],
    }),
  }),
});

export const { useContactQuery, useContactsQuery, useDeleteContactMutation } =
  contactApi;
