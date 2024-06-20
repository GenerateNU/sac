import { z } from "zod";

export const loginResponseSchema = z.object({
  redirect_uri: z.string(),
  sac_session: z.string(),
});

export const oauthCallbackRequestQueryParams = z.object({
  code: z.string(),
  session_state: z.string(),
  state: z.string(),
});

// Types:
export type LoginResponse = z.infer<typeof loginResponseSchema>;
export type OAuthCallbackRequestQueryParams = z.infer<typeof oauthCallbackRequestQueryParams>;
