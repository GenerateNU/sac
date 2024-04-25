import { z } from "zod";

export const rootModelSchema = z.object({
  id: z.string().uuid(),
  created_at: z.date(),
  updated_at: z.date(),
});

const paginationQueryParams = z.object({
  page: z.number().int().positive().optional(),
  limit: z.number().int().positive().optional(),
});

export type PaginationQueryParams = z.infer<typeof paginationQueryParams>;
