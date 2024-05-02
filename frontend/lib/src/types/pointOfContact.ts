import { z } from 'zod';

import { fileSchema } from '@/src/types/file';
import { rootModelSchema } from '@/src/types/root';

// Schemas:
const createPointOfContactRequestBodySchema = z.object({
    name: z.string(),
    email: z.string().email(),
    position: z.string()
});

const updatePointOfContactRequestBodySchema = z.object({
    name: z.string().optional(),
    email: z.string().email().optional(),
    position: z.string().optional()
});

const pointOfContactSchemaIntermediate = z.object({
    name: z.string().max(255),
    email: z.string().email(),
    position: z.string().max(255),
    photo_file: fileSchema
});

export const pointOfContactSchema =
    pointOfContactSchemaIntermediate.merge(rootModelSchema);

// Types:
export type CreatePointOfContactRequestBody = z.infer<
    typeof createPointOfContactRequestBodySchema
>;
export type UpdatePointOfContactRequestBody = z.infer<
    typeof updatePointOfContactRequestBodySchema
>;
export type PointOfContact = z.infer<typeof pointOfContactSchema>;
