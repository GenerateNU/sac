import { z } from 'zod';

export const rootModelSchema = z.object({
    id: z.string().uuid(),
    createdAt: z.date(),
    updatedAt: z.date()
});
