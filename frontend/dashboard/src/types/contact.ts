import { z } from "zod";

import { rootModelSchema } from "./root";

// Enums:
const contactTypeEnum = z.enum([
    "facebook",
    "instagram",
    "x",
    "linkedin",
    "youtube",
    "github",
    "slack",
    "discord",
    "email",
    "customSite",
]);

// Schemas:
export const createContactRequestBodySchema = z.object({
    type: contactTypeEnum,
    content: z.string().max(255),
})

export const contactSchema = createContactRequestBodySchema.merge(rootModelSchema);

// Types:
export type ContactType = z.infer<typeof contactTypeEnum>;
export type CreateContactRequestBody = z.infer<typeof createContactRequestBodySchema>;
export type Contact = z.infer<typeof contactSchema>;