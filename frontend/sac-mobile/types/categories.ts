import { z } from 'zod';

import { rootModelSchema } from './root';

export const categorySchema = z.object({
    name: z.string(),
});

const Category = categorySchema.merge(rootModelSchema);
export type Category = z.infer<typeof Category>;
