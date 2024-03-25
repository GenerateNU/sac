import { z } from 'zod';

export const tagSchema = z.object({
    id: z.string().uuid(),
    name: z.string().min(1),
    categoryId: z.string().uuid(),
    createdAt: z.date(),
    updatedAt: z.date()
});

export type Tag = z.infer<typeof tagSchema>;
