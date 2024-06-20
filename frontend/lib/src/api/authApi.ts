import { User, userSchema } from "../types";
import { LoginResponse, OAuthCallbackRequestQueryParams } from "../types/auth";
import { baseApi, handleQueryParams } from "./base";

const AUTH_API_BASE_URL = "/auth";
const PROVIDER = "microsoftonlineweb";

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
    logout: builder.query<void, void>({
      query: () => ({
        url: `${AUTH_API_BASE_URL}/logout`,
        method: "GET",
      }),
    }),
    callback: builder.query<User, OAuthCallbackRequestQueryParams>({
      query: (params) => ({
        url: handleQueryParams(`${AUTH_API_BASE_URL}/${PROVIDER}/callback`, params),
        method: "GET",
      }),
      transformResponse: (response) => {
        return userSchema.parse(response);
      }
    }),
  }),
});
