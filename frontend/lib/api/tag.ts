import { z } from "zod";

import { rootModelSchema } from "@/types/root";

export const tagSchemaIntermediate = z.object({
  name: z.string().max(255),
});

export const tagSchema = tagSchemaIntermediate.merge(rootModelSchema);
export type Tag = z.infer<typeof tagSchema>;
