import { z } from 'zod';

import { API_BASE_URL } from '@/const';
import { Contact, contactSchema } from '@/types/contact';
import { uuid } from '@/types/uuid';
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const contactsApi = createApi({
  reducerPath: "contactApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_BASE_URL }),
  endpoints: (builder) => ({
    contacts: builder.query<Contact[], void>({
      query: () => "contacts",
      transformResponse: (rawResponse: unknown) => {
        return z.array(contactSchema).parse(rawResponse);
      },
    }),
    contact: builder.query<Contact, uuid>({
      query: (contactID: uuid) => `contacts/${contactID}`,
      transformResponse: (rawResponse: unknown) => {
        return contactSchema.parse(rawResponse);
      },
    }),
  }),
});

export const { useContactsQuery } = contactsApi;
