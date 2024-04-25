import { baseApi } from "./api";
import { LoginRequestBody, RefreshTokenRequestBody } from "../types/auth";
import { User, userSchema } from "../types/user";
import {
  EmailRequestBody,
  VerifyEmailRequestBody,
  VerifyPasswordResetTokenRequestBody,
} from "../types/verification";

const AUTH_API_BASE_URL = "/auth";

const authApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    login: builder.mutation<User, LoginRequestBody>({
      query: (body) => ({
        url: `${AUTH_API_BASE_URL}/login`,
        method: "POST",
        body,
      }),
      transformResponse: (response: User) => {
        return userSchema.parse(response);
      },
    }),
    logout: builder.mutation<void, void>({
      query: () => ({
        url: `${AUTH_API_BASE_URL}/logout`,
        method: "POST",
      }),
    }),
    refresh: builder.mutation<void, RefreshTokenRequestBody>({
      query: (body) => ({
        url: "refresh",
        method: "POST",
        body,
      }),
    }),
    forgotPassword: builder.mutation<void, EmailRequestBody>({
      query: (body) => ({
        url: `${AUTH_API_BASE_URL}/forgot-password`,
        method: "POST",
        body,
      }),
    }),
    verifyPasswordResetToken: builder.mutation<
      void,
      VerifyPasswordResetTokenRequestBody
    >({
      query: (body) => ({
        url: `${AUTH_API_BASE_URL}/verify-reset`,
        method: "POST",
        body,
      }),
    }),
    sendCode: builder.mutation<void, EmailRequestBody>({
      query: (body) => ({
        url: `${AUTH_API_BASE_URL}/send-code`,
        method: "POST",
        body,
      }),
    }),
    verifyEmail: builder.mutation<void, VerifyEmailRequestBody>({
      query: (body) => ({
        url: `${AUTH_API_BASE_URL}/verify-email`,
        method: "POST",
        body,
      }),
    }),
  }),
});

export const {
  useLoginMutation,
  useLogoutMutation,
  useRefreshMutation,
  useForgotPasswordMutation,
  useVerifyPasswordResetTokenMutation,
  useSendCodeMutation,
  useVerifyEmailMutation,
} = authApi;
