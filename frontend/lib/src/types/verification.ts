import { z } from "zod";

// Enums:
const verificationTypeEnum = z.enum(["email_verification", "password_reset"]);

// Schemas:
const emailRequestBodySchema = z.object({
  email: z.string().email(),
});

const verifyEmailRequestBodySchema = z.object({
  email: z.string().email(),
  token: z.string().length(6),
});

const verifyPasswordResetTokenRequestBodySchema = z.object({
  token: z.string(),
  new_password: z.string().min(8),
  verify_new_password: z.string().min(8),
});

const verificationSchema = z.object({
  user_id: z.string().uuid(),
  token: z.string(),
  expires_at: z.date(),
  type: verificationTypeEnum,
});

// Types:
export type VerificationType = z.infer<typeof verificationTypeEnum>;
export type EmailRequestBody = z.infer<typeof emailRequestBodySchema>;
export type VerifyEmailRequestBody = z.infer<
  typeof verifyEmailRequestBodySchema
>;
export type VerifyPasswordResetTokenRequestBody = z.infer<
  typeof verifyPasswordResetTokenRequestBodySchema
>;
export type Verification = z.infer<typeof verificationSchema>;
