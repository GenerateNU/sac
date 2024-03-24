import { z } from 'zod';

import { rootModelSchema } from './root';

const eventSchema = z.object({
    name: z.string().max(255),
    preview: z.string().max(255),
    content: z.string().max(255),
    startTime: z.date(),
    endTime: z.date(),
    location: z.string().max(255),
    meetingLink: z.string().max(255).optional(),
    eventType: z.enum(['open', 'membersOnly']),
    isRecurring: z.boolean(),
    hosts: z.array(z.string().max(255)).min(1)
});

const Event = eventSchema.merge(rootModelSchema);
export type Event = z.infer<typeof Event>;
