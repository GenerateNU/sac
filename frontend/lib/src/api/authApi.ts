import { LoginResponse, RefreshTokenRequestBody } from "../types/auth";
import {
  EmailRequestBody,
  VerifyEmailRequestBody,
  VerifyPasswordResetTokenRequestBody,
} from "../types/verification";
import { baseApi } from "./base";

const AUTH_API_BASE_URL = "/auth";

export const authApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    login: builder.query<LoginResponse, void>({
      query: () => ({
        url: `${AUTH_API_BASE_URL}/login`,
        method: "GET",
        responseHandler: 'text',
      }),
      transformResponse: async (_, meta) => {
        const redirectUri = meta?.response?.headers.get("Redirect") as string;
        const sac_session = meta?.response?.headers.get("_sac_session") as string;

        return {
          redirect_uri: redirectUri,
          sac_session,
        }
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
