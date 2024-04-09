import { z } from 'zod';

export const clubSchema = z.object({
    id: z.string().uuid(),
    name: z.string().min(1)
});

export type Club = z.infer<typeof clubSchema>;
