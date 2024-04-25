import { z } from "zod";
import { PaginationQueryParams } from "../types/root";
import {
  CreateUserRequestBody,
  CreateUserTagsRequestBody,
  UpdateUserRequestBody,
  User,
  userSchema,
} from "../types/user";
import { baseApi, handleQueryParams } from "./api";
import { UpdatePasswordRequestBody } from "../types/auth";
import { Club, clubSchema } from "../types/club";
import { Tag, tagSchema } from "../types/tag";

const USER_API_BASE_URL = "/users";

const userApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    users: builder.query<User[], PaginationQueryParams>({
      query: (queryParams) => ({
        url: handleQueryParams(`${USER_API_BASE_URL}/`, queryParams),
        method: "GET",
      }),
      providesTags: (result) =>
        result ? result.map(({ id }) => ({ type: "User", id })) : ["User"],
      transformResponse: (response) => {
        return z.array(userSchema).parse(response);
      },
    }),
    user: builder.query<User, string>({
      query: (id) => ({
        url: `${USER_API_BASE_URL}/${id}/`,
        method: "GET",
      }),
      providesTags: (result, _, id) => (result ? [{ type: "User", id }] : []),
      transformResponse: (response) => {
        return userSchema.parse(response);
      },
    }),
    currentUser: builder.query<User, void>({
      query: () => ({
        url: `${USER_API_BASE_URL}/me`,
        method: "GET",
      }),
      providesTags: (result) =>
        result ? [{ type: "User", id: result.id }] : [],
      transformResponse: (response) => {
        return userSchema.parse(response);
      },
    }),
    createUser: builder.mutation<User, CreateUserRequestBody>({
      query: (body) => ({
        url: `${USER_API_BASE_URL}/`,
        method: "POST",
        body,
      }),
      invalidatesTags: ["User"],
      transformResponse: (response) => {
        return userSchema.parse(response);
      },
    }),
    updateUser: builder.mutation<
      User,
      { id: string; body: UpdateUserRequestBody }
    >({
      query: ({ id, body }) => ({
        url: `${USER_API_BASE_URL}/${id}/`,
        method: "PATCH",
        body,
      }),
      invalidatesTags: (result, _, { id }) =>
        result ? [{ type: "User", id }] : [],
      transformResponse: (response) => {
        return userSchema.parse(response);
      },
    }),
    deleteUser: builder.mutation<void, string>({
      query: (id) => ({
        url: `${USER_API_BASE_URL}/${id}/`,
        method: "DELETE",
      }),
      invalidatesTags: (_result, _, id) => [{ type: "User", id }],
    }),
    updatePassword: builder.mutation<
      void,
      { id: string; body: UpdatePasswordRequestBody }
    >({
      query: ({ id, body }) => ({
        url: `${USER_API_BASE_URL}/${id}/password`,
        method: "PATCH",
        body,
      }),
    }),
    userFollowing: builder.query<Club[], string>({
      query: (id) => ({
        url: `${USER_API_BASE_URL}/${id}/follower/`,
        method: "GET",
      }),
      providesTags: (result, _, id) =>
        result
          ? [{ type: "Follower", id }, "Club"]
          : [{ type: "Follower", id }],
      transformResponse: (response) => {
        return z.array(clubSchema).parse(response);
      },
    }),
    createUserFollowing: builder.mutation<
      void,
      { userId: string; clubId: string }
    >({
      query: ({ userId, clubId }) => ({
        url: `${USER_API_BASE_URL}/${userId}/follower/${clubId}`,
        method: "POST",
      }),
      invalidatesTags: (_result, _, { userId }) => [
        { type: "Follower", id: userId },
      ],
    }),
    deleteUserFollowing: builder.mutation<
      void,
      { userId: string; clubId: string }
    >({
      query: ({ userId, clubId }) => ({
        url: `${USER_API_BASE_URL}/${userId}/follower/${clubId}`,
        method: "DELETE",
      }),
      invalidatesTags: (_result, _, { userId }) => [
        { type: "Follower", id: userId },
      ],
    }),
    userMembership: builder.query<Club[], string>({
      query: (id) => ({
        url: `${USER_API_BASE_URL}/${id}/member/`,
        method: "GET",
      }),
      providesTags: (result, _, id) =>
        result ? [{ type: "Member", id }, "Club"] : [{ type: "Member", id }],
      transformResponse: (response) => {
        return z.array(clubSchema).parse(response);
      },
    }),
    createUserMembership: builder.mutation<
      void,
      { userId: string; clubId: string }
    >({
      query: ({ userId, clubId }) => ({
        url: `${USER_API_BASE_URL}/${userId}/member/${clubId}`,
        method: "POST",
      }),
      invalidatesTags: (_result, _, { userId }) => [
        { type: "Follower", id: userId },
        { type: "Member", id: userId },
      ],
    }),
    deleteUserMembership: builder.mutation<
      void,
      { userId: string; clubId: string }
    >({
      query: ({ userId, clubId }) => ({
        url: `${USER_API_BASE_URL}/${userId}/member/${clubId}`,
        method: "DELETE",
      }),
      invalidatesTags: (_result, _, { userId }) => [
        { type: "Member", id: userId },
      ],
    }),
    userTags: builder.query<Tag[], void>({
      query: () => ({
        url: `${USER_API_BASE_URL}/tags/`,
        method: "GET",
      }),
      providesTags: (result) =>
        result ? result.map((tag) => ({ type: "Tag", id: tag.id })) : ["Tag"],
      transformResponse: (response) => {
        return z.array(tagSchema).parse(response);
      },
    }),
    createUserTags: builder.mutation<Tag[], CreateUserTagsRequestBody>({
      query: (body) => ({
        url: `${USER_API_BASE_URL}/tags/`,
        method: "POST",
        body,
      }),
      invalidatesTags: (result, _, _arg) =>
        result ? result.map((tag) => ({ type: "Tag", id: tag.id })) : ["Tag"],
      transformResponse: (response) => {
        return z.array(tagSchema).parse(response);
      },
    }),
    deleteUserTag: builder.mutation<void, string>({
      query: (id) => ({
        url: `${USER_API_BASE_URL}/tags/${id}/`,
        method: "DELETE",
      }),
      invalidatesTags: (_result, _, id) => [{ type: "Tag", id }],
    }),
  }),
});

export const {
  useUsersQuery,
  useUserQuery,
  useCurrentUserQuery,
  useCreateUserMutation,
  useUpdateUserMutation,
  useDeleteUserMutation,
  useUpdatePasswordMutation,
  useUserFollowingQuery,
  useCreateUserFollowingMutation,
  useDeleteUserFollowingMutation,
  useUserMembershipQuery,
  useCreateUserMembershipMutation,
  useDeleteUserMembershipMutation,
  useUserTagsQuery,
  useCreateUserTagsMutation,
  useDeleteUserTagMutation,
} = userApi;
