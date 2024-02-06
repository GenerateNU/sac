import { z } from 'zod';

export const userSchema = z.object({
    id: z.string().uuid(),
    role: z.enum(['super', 'student']),
    email: z.string().email(),
    username: z.string().min(3),
    password: z.string().min(8),
    firstName: z.string().min(3),
    lastName: z.string().min(3),
    createdAt: z.date(),
    updatedAt: z.date(),
});

export type User = z.infer<typeof userSchema>;

export type Tokens = {
    accessToken: string;
    refreshToken: string;
};