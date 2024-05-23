export type Club = {
    userId: string;
    clubId: string;
    parentId?: string;
    clubName: string;
    accessToken: string;
    csrfToken: string;
    preview: string;
    description: string;
    weeklyTimeCommitment: number;
    oneWordToDescribeUs: string;
    logo: string;
}

export enum collegeEnum {
    CAMD = "CAMD",
	DMSB = "DMSB",
	KCCS = "KCCS",
	CE   = "CE",
	BCHS = "BCHS",
	SL   = "SL",
	CPS  = "CPS",
	CS   = "CS",
	CSSH = "CSSH",
}

export interface CreateClubRequestBody {
    userId: string;
    name: string;
    preview: string;
    description: string;
    weeklyTimeCommitment: number;
    oneWordToDescribeUs: string;
    logo: string;
}
  
export interface UpdateClubRequestBody {
    preview?: string;
    description?: string;
    weeklyTimeCommitment?: number;
    oneWordToDescribeUs?: string;
    logo?: string;
}