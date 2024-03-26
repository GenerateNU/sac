import { z } from 'zod';
import { tagSchema } from '@/types/tag';

export const categorySchema = z.object({
    id: z.string().uuid(),
    name: z.string().min(1),
    createdAt: z.date(),
    updatedAt: z.date(),
    tags: z.array(tagSchema),
});

export type Category = z.infer<typeof categorySchema>;

export type CategoryDisplay = {
    name: string;
    tags: Array<string>;
};