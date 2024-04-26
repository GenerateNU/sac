import { z } from "zod";

export const rootModelSchema = z.object({
  id: z.string().uuid(),
  created_at: z.date(),
  updated_at: z.date(),
});
