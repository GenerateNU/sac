import { z } from 'zod';

import { rootModelSchema } from './root';

export const userSchema = z.object({
    role: z.enum(['super', 'student']),
    email: z.string().email(),
    username: z.string().min(3),
    password: z.string().min(8),
    firstName: z.string().min(3),
    lastName: z.string().min(3)
});

export const collegeSchema = z.enum([
    'CAMD',
    'DMSB',
    'KCCS',
    'CE',
    'BCHS',
    'SL',
    'CPS',
    'CS',
    'CSSH'
]);
export type College = z.infer<typeof collegeSchema>;

export const yearSchema = z.enum(['1', '2', '3', '4', '5']);
export type Year = z.infer<typeof yearSchema>;

const User = userSchema.merge(rootModelSchema);
export type User = z.infer<typeof User>;
