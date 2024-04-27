import { z } from "zod";

import { rootModelSchema } from "./root";

/**
 * TODO: Re-Write After BE Fixes
 */

// Enums:
const eventTypeEnum = z.enum(['open', 'membersOnly']);

// Schemas:
const createEventRequestBodySchema = z.object({
    name: z.string(),
    preview: z.string(),
    content: z.string(),
    start_time: z.date(),
    end_time: z.date(),
    location: z.string(),
    event_type: eventTypeEnum,
    is_recurring: z.boolean()
});

const updateEventRequestBodySchema = z.object({
    name: z.string().optional(),
    preview: z.string().optional(),
    content: z.string().optional(),
    start_time: z.date().optional(),
    end_time: z.date().optional(),
    location: z.string().optional(),
    event_type: eventTypeEnum.optional(),
    host: z.string().uuid().optional()
});

const eventSchemaIntermediate = z.object({
    name: z.string().max(255),
    preview: z.string().max(255),
    content: z.string().max(255),
    start_time: z.date(),
    end_time: z.date(),
    location: z.string().max(255),
    meeting_link: z.string().max(255).optional(),
    event_type: eventTypeEnum,
    is_recurring: z.boolean(),
    host: z.string().uuid()
});

export const eventSchema = eventSchemaIntermediate.merge(rootModelSchema);

// Types:
export type CreateEventRequestBody = z.infer<
    typeof createEventRequestBodySchema
>;
export type UpdateEventRequestBody = z.infer<
    typeof updateEventRequestBodySchema
>;
export type Event = z.infer<typeof eventSchema>;
