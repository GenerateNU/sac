import { z } from 'zod';

import { rootModelSchema } from './root';

export const tagSchema = z.object({
    name: z.string().max(255)
});

const Tag = tagSchema.merge(rootModelSchema);
export type Tag = z.infer<typeof Tag>;
