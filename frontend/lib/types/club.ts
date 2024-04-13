import { z } from "zod";

import { rootModelSchema } from "./root";

const clubSchema = z.object({
  name: z.string().max(255),
  preview: z.string().max(255),
  description: z.string().max(255),
  num_members: z.number(),
  is_recruiting: z.boolean(),
  recruitment_cycle: z.enum(["fall", "spring", "fallSpring", "always"]),
  recruitment_type: z.enum(["unrestricted", "tryout", "application"]),
  application_link: z.string().max(255),
  logo: z.string().max(255).optional(),
});

const Club = clubSchema.merge(rootModelSchema);
export type Club = z.infer<typeof Club>;
