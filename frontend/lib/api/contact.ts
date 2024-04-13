import { z } from "zod";

import { rootModelSchema } from "@/types/root";

// TODO: each one needs a svg icons
const ContactTypeValues = [
  // 'facebook',
  // 'instagram',
  // 'twitter',
  // 'linkedin',
  // 'youtube',
  // 'github',
  "slack",
  // 'discord',
  "email",
  // 'customSite'
] as const;

export type ContactType = (typeof ContactTypeValues)[number];

const contactSchemaIntermediate = z.object({
  type: z.enum(ContactTypeValues),
  content: z.string().max(255),
});

export const contactSchema = contactSchemaIntermediate.merge(rootModelSchema);
export type Contact = z.infer<typeof contactSchema>;
