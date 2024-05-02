import { z } from 'zod';

import { rootModelSchema } from './root';

// Schemas:
export const createFileRequestBodySchema = z.object({
    owner_id: z.string().uuid(),
    owner_type: z.string().max(255)
});

const fileSchemaIntermediate = z.object({
    owner_id: z.string().uuid(),
    owner_type: z.string().max(255),
    file_name: z.string().max(255),
    file_type: z.string().max(255),
    file_size: z.number().min(1),
    file_url: z.string().max(255),
    object_key: z.string().max(255)
});

export const fileSchema = fileSchemaIntermediate.merge(rootModelSchema);

// Types:
export type CreateFileRequestBody = z.infer<typeof createFileRequestBodySchema>;
export type File = z.infer<typeof fileSchema>;
