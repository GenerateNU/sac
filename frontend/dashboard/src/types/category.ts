import { z } from "zod";

import { rootModelSchema } from "./root";
import { tagSchema } from "./tag";

// Schemas:
export const createCategoryRequestBodySchema = z.object({
    name: z.string().min(1),
});

export const categorySchemaIntermediate = z.object({
  name: z.string(),
  tags: z.array(tagSchema),
});

export const categorySchema = categorySchemaIntermediate.merge(rootModelSchema);

// Types:
export type CreateCategoryRequestBody = z.infer<typeof createCategoryRequestBodySchema>;
export type Category = z.infer<typeof categorySchema>;