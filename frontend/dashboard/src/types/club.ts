import { z } from "zod";
import { rootModelSchema } from "./root";

// Enums:
const recruitmentCycleEnum = z.enum(["fall", "spring", "fallSpring", "always"]);
const recruitmentTypeEnum = z.enum(["unrestricted", "tryout", "application"]);

// Schemas:
export const createClubRequestBodySchema = z.object({
    user_id: z.string().uuid(),
    name: z.string(),
    preview: z.string(),
    description: z.string(),
    is_recruiting: z.boolean(),
    recruitment_cycle: recruitmentCycleEnum,
    recruitment_type: recruitmentTypeEnum,
    application_link: z.string(),
    logo: z.string().optional(),
    weekly_time_committment: z.number().optional(),
    one_word_to_describe_us: z.string().max(20).optional(),
});

export const updateClubRequestBodySchema = z.object({
    name: z.string().optional(),
    preview: z.string().optional(),
    description: z.string().optional(),
    is_recruiting: z.boolean().optional(),
    recruitment_cycle: recruitmentCycleEnum.optional(),
    recruitment_type: recruitmentTypeEnum.optional(),
    application_link: z.string().optional(),
    logo: z.string().optional(),
    weekly_time_committment: z.number().optional(),
    one_word_to_describe_us: z.string().max(20).optional(),
});

export const createClubTagsRequestBodySchema = z.object({
    tags: z.array(z.string().uuid()),
});

const clubSchemaIntermediate = z.object({
    name: z.string().max(255),
    preview: z.string().max(255),
    description: z.string().max(255),
    num_members: z.number(),
    is_recruiting: z.boolean(),
    recruitment_cycle: recruitmentCycleEnum,
    recruitment_type: recruitmentTypeEnum,
    application_link: z.string().max(255),
    logo: z.string().max(255).optional(),
    weekly_time_committment: z.number().optional(),
    one_word_to_describe_us: z.string().max(20).optional(),
});

export const clubSchema = clubSchemaIntermediate.merge(rootModelSchema);

// Types:
export type RecruitmentCycle = z.infer<typeof recruitmentCycleEnum>;
export type RecruitmentType = z.infer<typeof recruitmentTypeEnum>;
export type CreateClubRequestBody = z.infer<typeof createClubRequestBodySchema>;
export type UpdateClubRequestBody = z.infer<typeof updateClubRequestBodySchema>;
export type CreateClubTagsRequestBody = z.infer<typeof createClubTagsRequestBodySchema>;
export type Club = z.infer<typeof clubSchema>;