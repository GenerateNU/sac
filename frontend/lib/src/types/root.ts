import { z } from "zod";

export type uuid = string;

export const rootModelSchema = z.object({
  id: z.string().uuid(),
  created_at: z.string(),
  updated_at: z.string(),
});

const paginationQueryParams = z
  .object({
    page: z.number().int().positive().optional(),
    limit: z.number().int().positive().optional(),
    start: z.string().optional(),
    end: z.string().optional(),
  })
  .optional();

export type PaginationQueryParams = z.infer<typeof paginationQueryParams>;
