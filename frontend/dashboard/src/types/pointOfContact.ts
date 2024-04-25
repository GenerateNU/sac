import { z } from "zod";
import { fileSchema } from "./file";
import { rootModelSchema } from "./root";

// Schemas:
const createPointOfContactRequestBodySchema = z.object({
<<<<<<< HEAD
  name: z.string(),
  email: z.string().email(),
  position: z.string(),
});

const updatePointOfContactRequestBodySchema = z.object({
  name: z.string().optional(),
  email: z.string().email().optional(),
  position: z.string().optional(),
});

const pointOfContactSchemaIntermediate = z.object({
  name: z.string().max(255),
  email: z.string().email(),
  position: z.string().max(255),
  photo_file: fileSchema,
});

export const pointOfContactSchema =
  pointOfContactSchemaIntermediate.merge(rootModelSchema);

// Types:
export type CreatePointOfContactRequestBody = z.infer<
  typeof createPointOfContactRequestBodySchema
>;
export type UpdatePointOfContactRequestBody = z.infer<
  typeof updatePointOfContactRequestBodySchema
>;
=======
    name: z.string(),
    email: z.string().email(),
    position: z.string(),
})

const updatePointOfContactRequestBodySchema = z.object({
    name: z.string().optional(),
    email: z.string().email().optional(),
    position: z.string().optional(),
})

const pointOfContactSchemaIntermediate = z.object({
    name: z.string().max(255),
    email: z.string().email(),
    position: z.string().max(255),
    photo_file: fileSchema,
});

export const pointOfContactSchema = pointOfContactSchemaIntermediate.merge(rootModelSchema);

// Types:
export type CreatePointOfContactRequestBody = z.infer<typeof createPointOfContactRequestBodySchema>;
export type UpdatePointOfContactRequestBody = z.infer<typeof updatePointOfContactRequestBodySchema>;
>>>>>>> a28089a2 (feat: frontend lib, rtk migration, & dashboard init (#632))
export type PointOfContact = z.infer<typeof pointOfContactSchema>;
