import { z } from 'zod';

export const tagSchema = z.object({
    id: z.string().uuid(),
    name: z.string().min(1),
    categoryId: z.string().uuid(),
    createdAt: z.date(),
    updatedAt: z.date()
});

export type Category = z.infer<typeof tagSchema>;

export type CategoryDisplay = {
    name: string;
    tags: Array<string>;
};