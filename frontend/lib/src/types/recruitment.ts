import { z } from "zod";

// Enums:
const recruitmentCycleEnum = z.enum(["fall", "spring", "fallSpring", "always"]);
const recruitmentTypeEnum = z.enum(["unrestricted", "tryout", "application"]);

export const applicationSchema = z.object({
  applicationTitle: z.string(),
  applicationLink: z.string().url(),
})

export const recruitmentSchema = z.object({
  recruitment_cycle: recruitmentCycleEnum,
  recruitment_type: recruitmentTypeEnum,
  is_recruiting: z.boolean(),
  applications: z.array(applicationSchema),
});

export type RecruitmentCycle = z.infer<typeof recruitmentCycleEnum>;
export type RecruitmentType = z.infer<typeof recruitmentTypeEnum>;
export type Application = z.infer<typeof applicationSchema>;
export type Recruitment = z.infer<typeof recruitmentSchema>;
