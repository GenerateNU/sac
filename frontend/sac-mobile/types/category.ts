import { z } from 'zod';

import { tagSchema } from '@/types/tag';

import { rootModelSchema } from './root';

export const categorySchema = z.object({
    name: z.string().min(1),
    tags: z.array(tagSchema)
});

const Category = categorySchema.merge(rootModelSchema);
export type Category = z.infer<typeof Category>;
