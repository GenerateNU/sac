import { z } from 'zod';

export const eventSchema = z.object({
    id: z.string().uuid(),
    name: z.string().max(255),
    preview: z.string().max(255),
    content: z.string().max(255),
    startTime: z.date(),
    endTime: z.date(),
    location: z.string().max(255),
    eventType: z.enum(['open', 'membersOnly']),
    isRecurring: z.boolean(),
    createdAt: z.date(),
    updatedAt: z.date()
});

export type Event = z.infer<typeof eventSchema>;
