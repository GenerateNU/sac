import { z } from 'zod';

import { rootModelSchema } from './root';
import { tagSchema } from './tag';

export const categorySchema = z.object({
    name: z.string(),
    tags: z.array(tagSchema)
});

const Category = categorySchema.merge(rootModelSchema);
export type Category = z.infer<typeof Category>;
