import { FetchBaseQueryMeta } from '@reduxjs/toolkit/query';

export const extractToken = (
    metadata: FetchBaseQueryMeta | undefined
): string => {
    const access_token = metadata?.response?.headers
        .get('Authorization')
        ?.split(' ')[1];

    if (!access_token) {
        throw new Error('Token not found in response');
    }

    return access_token;
};
