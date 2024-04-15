import { z } from 'zod';

import { tagSchema } from '@/types/tag';

import { rootModelSchema } from './root';

const categorySchemaIntermediate = z.object({
    name: z.string().min(1),
    tags: z.array(tagSchema)
});

const categorySchema = categorySchemaIntermediate.merge(rootModelSchema);
type Category = z.infer<typeof categorySchema>;

export { categorySchema, Category };