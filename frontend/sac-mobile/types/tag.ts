import { z } from 'zod';

import { rootModelSchema } from './root';

const tagIntermediate = z.object({
    name: z.string().min(1),
    category_id: z.string().uuid()
});

const tagSchema = tagIntermediate.merge(rootModelSchema);
type Tag = z.infer<typeof tagSchema>;

export { tagSchema, Tag };
