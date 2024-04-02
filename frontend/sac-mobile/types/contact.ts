import { z } from 'zod';

import { rootModelSchema } from './root';

const ContactTypeValues = [
    'facebook',
    'instagram',
    'twitter',
    'linkedin',
    'youtube',
    'github',
    'slack',
    'discord',
    'email',
    'customSite'
] as const;

export type ContactType = (typeof ContactTypeValues)[number];

const contactSchema = z.object({
    type: z.enum(ContactTypeValues),
    content: z.string().max(255)
});

const Contact = contactSchema.merge(rootModelSchema);
export type Contact = z.infer<typeof Contact>;
