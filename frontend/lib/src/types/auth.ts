import { z } from "zod";

export const loginResponseSchema = z.object({
  redirect_uri: z.string(),
  sac_session: z.string(),
});

export const updatePasswordRequestBodySchema = z.object({
  old_password: z.string().min(8),
  new_password: z.string().min(8),
});

export const refreshTokenRequestBodySchema = z.object({
  refresh_token: z.string(),
});

export const tokensSchema = z.object({
  access_token: z.string(),
  refresh_token: z.string(),
});

// Types:
export type LoginResponse = z.infer<typeof loginResponseSchema>;
export type UpdatePasswordRequestBody = z.infer<
  typeof updatePasswordRequestBodySchema
>;
export type RefreshTokenRequestBody = z.infer<
  typeof refreshTokenRequestBodySchema
>;
export type Tokens = z.infer<typeof tokensSchema>;
