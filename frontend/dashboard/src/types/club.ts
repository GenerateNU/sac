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