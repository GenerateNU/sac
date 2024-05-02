import { z } from 'zod';

import { rootModelSchema } from '@/src/types/root';

// Schemas:
export const createTagRequestBodySchema = z.object({
    name: z.string(),
    categoryID: z.string().uuid()
});

export const updateTagRequestBodySchema = z.object({
    name: z.string().optional(),
    categoryID: z.string().uuid().optional()
});

export const tagSchemaIntermediate = z.object({
    name: z.string().max(255),
    category_id: z.string().uuid()
});

export const tagSchema = tagSchemaIntermediate.merge(rootModelSchema);

// Types:
export type UpdateTagRequestBody = z.infer<typeof updateTagRequestBodySchema>;
export type CreateTagRequestBody = z.infer<typeof createTagRequestBodySchema>;
export type Tag = z.infer<typeof tagSchema>;
