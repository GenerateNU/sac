import { z } from 'zod';

import { rootModelSchema } from './root';

const clubSchema = z.object({
    name: z.string().max(255),
    preview: z.string().max(255),
    description: z.string().max(255),
    numMembers: z.number(),
    isRecruiting: z.boolean(),
    recruitmentCycle: z.enum(['fall', 'spring', 'fallSpring', 'always']),
    recruitmentType: z.enum(['unrestricted', 'tryout', 'application']),
    applicationLink: z.string().max(255),
    logo: z.string().max(255).optional()
});

const Club = clubSchema.merge(rootModelSchema);
export type Club = z.infer<typeof Club>;
