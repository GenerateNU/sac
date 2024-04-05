import { z } from 'zod';

export const userSchema = z.object({
    id: z.string().uuid(),
    role: z.enum(['super', 'student']),
    email: z.string().email(),
    first_name: z.string().min(3),
    last_name: z.string().min(3),
    created_at: z.date(),
    updated_at: z.date(),
    is_verified: z.boolean()
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

export type User = z.infer<typeof userSchema>;
export type UserWithPassword = User & { password: string };

export const InitialUser: User = {
    id: '',
    role: 'student',
    email: '',
    first_name: '',
    last_name: '',
    created_at: new Date(),
    updated_at: new Date(),
    is_verified: false
};

export type Tokens = {
    accessToken: string;
    refreshToken: string;
};
