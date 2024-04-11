import { z } from 'zod';

import { rootModelSchema } from './root';

export const tagSchema = z.object({
    name: z.string().min(1),
    category_id: z.string().uuid()
});

const Tag = tagSchema.merge(rootModelSchema);
export type Tag = z.infer<typeof Tag>;
