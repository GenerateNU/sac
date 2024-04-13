import { z } from 'zod';

import { rootModelSchema } from './root';

export const tagIntermediate = z.object({
    name: z.string().min(1),
    category_id: z.string().uuid()
});

export const tagSchema = tagIntermediate.merge(rootModelSchema);
export type Tag = z.infer<typeof tagSchema>;
