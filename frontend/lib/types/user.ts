import { z } from "zod";
import { rootModelSchema } from "./root";

export const userSchemaIntermediate = z.object({
    role: z.enum(["super", "student"]),
    first_name: z.string().min(1),
    last_name: z.string().min(1),
    email: z.string().email(),
    college: z.string().optional(),
    graduation_cycle: z.string().optional(),
    graduation_year: z.number().optional(),
    is_verified: z.boolean(),
});

export const collegeSchema = z.enum([
    "CAMD",
    "DMSB",
    "KCCS",
    "CE",
    "BCHS",
    "SL",
    "CPS",
    "CS",
    "CSSH",
]);
export type College = z.infer<typeof collegeSchema>;

export const yearSchema = z.enum(["1", "2", "3", "4", "5"]);
export type Year = z.infer<typeof yearSchema>;

export const userSchema = userSchemaIntermediate.merge(rootModelSchema);
export type User = z.infer<typeof userSchema>;