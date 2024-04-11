import { z } from "zod";
import { tagSchema } from "./tag";
import { rootModelSchema } from "./root";

export const categorySchema = z.object({
    name: z.string(),
    tags: z.array(tagSchema),
});

const Category = categorySchema.merge(rootModelSchema);
export type Category = z.infer<typeof Category>;