import { z } from 'zod';

import { PointOfContact, pointOfContactSchema } from '../types/pointOfContact';
import { PaginationQueryParams } from '../types/root';
import { baseApi, handleQueryParams } from './base';

const POC_API_BASE_URL = '/pocs';

export const pointOfContactApi = baseApi.injectEndpoints({
    endpoints: (builder) => ({
        pointOfContacts: builder.query<PointOfContact[], PaginationQueryParams>(
            {
                query: (queryParams) => ({
                    url: handleQueryParams(`${POC_API_BASE_URL}/`, queryParams),
                    method: 'GET'
                }),
                providesTags: (result) =>
                    result
                        ? result.map(({ id }) => ({
                              type: 'PointOfContact',
                              id
                          }))
                        : ['PointOfContact'],
                transformResponse: (response) => {
                    return z.array(pointOfContactSchema).parse(response);
                }
            }
        ),
        pointOfContact: builder.query<PointOfContact, string>({
            query: (id) => ({
                url: `${POC_API_BASE_URL}/${id}`,
                method: 'GET'
            }),
            providesTags: (result, _, id) =>
                result ? [{ type: 'PointOfContact', id }] : [],
            transformResponse: (response) => {
                return pointOfContactSchema.parse(response);
            }
        })
    })
});
