import { z } from 'zod';

import { rootModelSchema } from './root';

const eventSchema = z.object({
    name: z.string().max(255),
    preview: z.string().max(255),
    content: z.string().max(255),
    start_time: z.date(),
    end_time: z.date(),
    location: z.string().max(255),
    meeting_link: z.string().max(255).optional(),
    event_type: z.enum(['open', 'membersOnly']),
    is_recurring: z.boolean(),
    host: z.string().max(255),
});

const Event = eventSchema.merge(rootModelSchema);
export type Event = z.infer<typeof Event>;
