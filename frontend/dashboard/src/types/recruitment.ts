import { Application } from "./application";

export enum RecruitmentCycle {
    Fall = 'fall',
    FallSpring = 'fallSpring',
    Spring = 'spring',
    Always = 'always'
}
  
export enum RecruitmentType {
    Unrestricted = 'unrestricted',
    Tryout = 'tryout',
    Apply = 'application'
}
  
export type Recruitment = {
    recruitmentCycle: RecruitmentCycle;
    recruitmentType: RecruitmentType[];
    isRecruiting: boolean;
    applications: Application[];
}

export type UpdateRecruitmentRequestBody = {
    recruitmentCycle?: RecruitmentCycle;
    recruitmentType?: RecruitmentType[];
    isRecruiting?: boolean;
    applications?: Application[];
}

