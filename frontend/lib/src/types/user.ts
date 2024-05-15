import { z } from "zod";

import { rootModelSchema } from "./root";

// Enums:
export const userRoleEnum = z.enum(["super", "student"]);

export const collegeEnum = z.enum([
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

export const majorEnum = z.enum([
  "africanaStudies",
  "americanSignLanguage",
  "americanSignLanguage-EnglishInterpreting",
  "appliedPhysics",
  "architecturalStudies",
  "architecture",
  "art:ArtVisualStudies",
  "behavioralNeuroscience",
  "biochemistry",
  "bioengineering",
  "biology",
  "biomedicalPhysics",
  "businessAdministration",
  "businessAdministration:Accounting",
  "businessAdministration:AccountingAndAdvisoryServices",
  "businessAdministration:BrandManagement",
  "businessAdministration:BusinessAnalytics",
  "businessAdministration:CorporateInnovation",
  "businessAdministration:EntrepreneurialStartups",
  "businessAdministration:FamilyBusiness",
  "businessAdministration:Finance",
  "businessAdministration:Fintech",
  "businessAdministration:HealthcareManagementAndConsulting",
  "businessAdministration:Management",
  "businessAdministration:ManagementInformationSystems",
  "businessAdministration:Marketing",
  "businessAdministration:MarketingAnalytics",
  "businessAdministration:SocialInnovationAndEntrepreneurship",
  "businessAdministration:SupplyChainManagement",
  "cellAndMolecularBiology",
  "chemicalEngineering",
  "chemistry",
  "civilEngineering",
  "communicationStudies",
  "computerEngineering",
  "computerScience",
  "computingAndLaw",
  "criminologyAndCriminalJustice",
  "culturalAnthropology",
  "cybersecurity",
  "dataScience",
  "design",
  "economics",
  "electricalEngineering",
  "english",
  "environmentalAndSustainabilityStudies",
  "environmentalEngineering",
  "environmentalScience",
  "environmentalStudies",
  "gameArtAndAnimation",
  "gameDesign",
  "globalAsianStudies",
  "healthScience",
  "history",
  "historyCultureAndLaw",
  "humanServices",
  "industrialEngineering",
  "internationalAffairs",
  "internationalBusiness",
  "internationalBusiness:Accounting",
  "internationalBusiness:AccountingAndAdvisoryServices",
  "internationalBusiness:BrandManagement",
  "internationalBusiness:BusinessAnalytics",
  "internationalBusiness:CorporateInnovation",
  "internationalBusiness:EntrepreneurialStartups",
  "internationalBusiness:FamilyBusiness",
  "internationalBusiness:Finance",
  "internationalBusiness:Fintech",
  "internationalBusiness:HealthcareManagementAndConsulting",
  "internationalBusiness:Management",
  "internationalBusiness:ManagementInformationSystems",
  "internationalBusiness:Marketing",
  "internationalBusiness:MarketingAnalytics",
  "internationalBusiness:SocialInnovationAndEntrepreneurship",
  "internationalBusiness:SupplyChainManagement",
  "journalism",
  "landscapeArchitecture",
  "linguistics",
  "marineBiology",
  "mathematics",
  "mechanicalEngineering",
  "mediaAndScreenStudies",
  "mediaArts",
  "music",
  "musicTechnology",
  "nursing",
  "pharmaceuticalSciences",
  "pharmacy(PharmD)",
  "philosophy",
  "physics",
  "politicalScience",
  "politicsPhilosophyEconomics",
  "psychology",
  "publicHealth",
  "publicRelations",
  "religiousStudies",
  "sociology",
  "spanish",
  "speechLanguagePathologyAndAudiology",
  "theatre",
]);

export const graduationCycleEnum = z.enum(["december", "may"]);

export const yearEnum = z.enum(["1", "2", "3", "4", "5"]);

// Schemas:
export const createUserRequestBodySchema = z.object({
  first_name: z.string().min(1),
  last_name: z.string().min(1),
  email: z.string().email(),
  password: z.string().min(8),
  major0: majorEnum.optional(),
  major1: majorEnum.optional(),
  major2: majorEnum.optional(),
  college: collegeEnum,
  graduation_cycle: graduationCycleEnum.optional(),
  graduation_year: z.number().optional(),
});

export const updateUserRequestBodySchema = z.object({
  first_name: z.string().min(1).optional(),
  last_name: z.string().min(1).optional(),
  major0: majorEnum.optional(),
  major1: majorEnum.optional(),
  major2: majorEnum.optional(),
  college: collegeEnum.optional(),
  graduation_cycle: graduationCycleEnum.optional(),
  graduation_year: z.number().optional(),
});

export const createUserTagsRequestBodySchema = z.object({
  tags: z.array(z.string().uuid()),
});

const userSchemaIntermediate = z.object({
  role: userRoleEnum,
  first_name: z.string().min(1),
  last_name: z.string().min(1),
  email: z.string().email(),
  major0: majorEnum.optional(),
  major1: majorEnum.optional(),
  major2: majorEnum.optional(),
  college: collegeEnum.optional(),
  graduation_cycle: graduationCycleEnum.optional(),
  graduation_year: z.number().optional(),
  is_verified: z.boolean(),
});

export const userSchema = userSchemaIntermediate.merge(rootModelSchema);

// Types:
export type UserRole = z.infer<typeof userRoleEnum>;
export type College = z.infer<typeof collegeEnum>;
export type Major = z.infer<typeof majorEnum>;
export type GraduationCycle = z.infer<typeof graduationCycleEnum>;
export type Year = z.infer<typeof yearEnum>;

export type CreateUserRequestBody = z.infer<typeof createUserRequestBodySchema>;
export type UpdateUserRequestBody = z.infer<typeof updateUserRequestBodySchema>;
export type CreateUserTagsRequestBody = z.infer<
  typeof createUserTagsRequestBodySchema
>;
export type User = z.infer<typeof userSchema>;
