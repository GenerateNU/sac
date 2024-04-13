import { z } from "zod";

import { rootModelSchema } from "@/types/root";

import { tagSchema } from "./tag";

export const categorySchemaIntermediate = z.object({
  name: z.string(),
  tags: z.array(tagSchema),
});

export const categorySchema = categorySchemaIntermediate.merge(rootModelSchema);
export type Category = z.infer<typeof categorySchema>;
