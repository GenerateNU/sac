import { AxiosError } from 'axios';

import { extractTokens } from '@/lib/utils';
import { api } from '@/services/api';
import { Tokens } from '@/types/auth';
import { ErrorResponse } from '@/types/error';
import { User } from '@/types/user';

/**
 * Creates a new user with the given first name, last name, email, and password.
 * @param first_name The first name of the user.
 * @param last_name The last name of the user.
 * @param email The email of the user.
 * @param password The password of the user.
 * @returns The user that was created.
 */
export const signUp = async (
    first_name: string,
    last_name: string,
    email: string,
    password: string
): Promise<{ user: User; tokens: Tokens }> => {
    try {
        const response = await api.post<User>('/users/', {
            first_name: first_name,
            last_name: last_name,
            email: email,
            password: password
        });

        const tokens = extractTokens(
            response.headers['set-cookie'],
            response.headers
        );
        const user = response.data;

        return { user, tokens };
    } catch (error: any) {
        const axiosErr: AxiosError = error;
        const response = axiosErr.response;
        if (response) {
            const data = response.data as ErrorResponse;
            throw new Error(data.error);
        } else {
            throw new Error('Error signing up');
        }
    }
};

/**
 * Requests a verification code for the user with the given email.
 * @param email The email of the user.
 */
export const requestVerification = async (email: string): Promise<void> => {
    try {
        await api.post<void>('/auth/send-code', {
            email
        });
    } catch (error: any) {
        const axiosErr: AxiosError = error;
        const response = axiosErr.response;
        if (response) {
            const data = response.data as ErrorResponse;
            throw new Error(data.error);
        } else {
            throw new Error('Error requesting verification code');
        }
    }
};

/**
 * Completes the verification process for the user with the given email and code.
 * @param email The email of the user.
 * @param code The verification code of the user.
 * @returns The user that was verified.
 */
export const completeVerification = async (
    email: string,
    code: string
): Promise<void> => {
    try {
        await api.post<void>('/auth/verify-email/', {
            email,
            token: code
        });
    } catch (error: any) {
        const axiosErr: AxiosError = error;
        const response = axiosErr.response;
        if (response) {
            const data = response.data as ErrorResponse;
            throw new Error(data.error);
        } else {
            throw new Error('Error completing verification');
        }
    }
};

/**
 * Logins the user with the given email and password.
 * @param email The email of the user.
 * @param password The password of the user.
 * @returns The user that was logged in.
 */
export const login = async (
    email: string,
    password: string
): Promise<{ user: User; tokens: Tokens }> => {
    try {
        const response = await api.post<User>('/auth/login', {
            email,
            password
        });

        const tokens = extractTokens(
            response.headers['set-cookie'],
            response.headers
        );
        const user = response.data;

        return { user, tokens };
    } catch (error: any) {
        console.log('Error logging in:', error);
        const axiosErr: AxiosError = error;
        const response = axiosErr.response;
        if (response) {
            const data = response.data as ErrorResponse;
            throw new Error(data.error);
        } else {
            throw new Error('Error logging in');
        }
    }
};

/**
 * Logout the user.
 */
export const logout = async (): Promise<void> => {
    try {
        await api.post<void>('/auth/logout');
    } catch (error: any) {
        console.log('Error logging out:', error);
        throw new Error('Error logging out');
    }
};

/**
 * Refreshes the access token with the given refresh token.
 * @param refreshToken The refresh token of the user.
 * @returns The new tokens.
 */
export const refreshAccessToken = async (
    refreshToken: string
): Promise<Tokens> => {
    try {
        const response = await api.post<void>('/auth/refresh', {
            refresh_token: refreshToken
        });

        const tokens = extractTokens(
            response.headers['set-cookie'],
            response.headers
        );

        return tokens;
    } catch (error: any) {
        console.log('Error refreshing token:', error);
        const axiosErr: AxiosError = error;
        const response = axiosErr.response;
        if (response) {
            const data = response.data as ErrorResponse;
            throw new Error(data.error);
        } else {
            throw new Error('Error refreshing token');
        }
    }
};

// import { AxiosError } from 'axios';

// import { extractTokens } from '@/lib/utils';
// import { api } from '@/services/api';
// import { Tokens } from '@/types/auth';
// import { ErrorResponse } from '@/types/error';
// import { User } from '@/types/user';

// /**
//  * Creates a new user with the given user information.
//  *
//  * @param first_name The first name of the user.
//  * @param last_name The last name of the user.
//  * @param email The email of the user.
//  * @param password The password of the user.
//  * @returns The user that was created and the tokens.
//  * @throws Error if the user cannot be created.
//  * @see User
//  * @see Tokens
//  */
// const register = async (
//     first_name: string,
//     last_name: string,
//     email: string,
//     password: string
// ): Promise<{ user: User; tokens: Tokens }> => {
//     return api
//         .post<User>('/users/', {
//             first_name: first_name,
//             last_name: last_name,
//             email: email,
//             password: password
//         })
//         .then((response) => {
//             const tokens = extractTokens(
//                 response.headers['set-cookie'],
//                 response.headers
//             );
//             const user = response.data;

//             return { user, tokens };
//         })
//         .catch((error: AxiosError<ErrorResponse>) => {
//             console.error(error);
//             throw new Error('Error fetching clubs');
//         });
// };

// /**
//  * Requests a verification code for the user with the given email.
//  *
//  * @param email The email of the user.
//  * @throws Error if the verification code cannot be requested.
//  * @see User
//  * @see Tokens
//  */
// const requestVerification = async (email: string): Promise<void> => {
//     return api
//         .post<void>('/auth/send-code', { email })
//         .then(() => { })
//         .catch((error: AxiosError<ErrorResponse>) => {
//             console.error(error);
//             throw new Error('Error requesting verification code');
//         });
// };


// /**
//  * Completes the verification process for the user with the given email and code.
//  *
//  * @param email The email of the user.
//  * @param code The verification code of the user.
//  * @throws Error if the verification process cannot be completed.
//  * @see User
//  * @see Tokens
//  */
// const verifyEmail = async (
//     email: string,
//     code: string
// ): Promise<void> => {
//     return api
//         .post<void>('/auth/verify-email/', {
//             email,
//             token: code
//         })
//         .then(() => { })
//         .catch((error: AxiosError<ErrorResponse>) => {
//             console.error(error);
//             throw new Error('Error completing verification');
//         });
// };

// /**
//  * Logins the user with the given email and password.
//  *
//  * @param email The email of the user.
//  * @param password The password of the user.
//  * @returns The user that was logged in.
//  * @throws Error if the user cannot be logged in.
//  * @see User
//  * @see Tokens
//  */
// const login = async (
//     email: string,
//     password: string
// ): Promise<{ user: User; tokens: Tokens }> => {
//     return api
//         .post<User>('/auth/login', {
//             email,
//             password
//         })
//         .then((response) => {
//             const tokens = extractTokens(
//                 response.headers['set-cookie'],
//                 response.headers
//             );
//             const user = response.data;

//             return { user, tokens };
//         })
//         .catch((error: AxiosError<ErrorResponse>) => {
//             console.error(error);
//             throw new Error('Error logging in');
//         });
// };


// /**
//  * Logout the user.
//  *
//  * @throws Error if the user cannot be logged out.
//  * @see User
//  * @see Tokens
//  */
// const logout = async (): Promise<void> => {
//     return api
//         .post<void>('/auth/logout')
//         .then(() => { })
//         .catch((error: AxiosError<ErrorResponse>) => {
//             console.error(error);
//             throw new Error('Error logging out');
//         });
// }

// /**
//  * Refreshes the access token with the given refresh token.
//  *
//  * @param refreshToken The refresh token of the user.
//  * @returns The new tokens.
//  * @throws Error if the tokens cannot be refreshed.
//  * @see Tokens
//  */
// const refreshAccessToken = async (
//     refreshToken: string
// ): Promise<Tokens> => {
//     return api
//         .post<void>('/auth/refresh', {
//             refresh_token: refreshToken
//         })
//         .then((response) => {
//             const tokens = extractTokens(
//                 response.headers['set-cookie'],
//                 response.headers
//             );

//             return tokens;
//         })
//         .catch((error: AxiosError<ErrorResponse>) => {
//             console.error(error);
//             throw new Error('Error refreshing token');
//         });
// };

// export { register, requestVerification, verifyEmail, login, logout, refreshAccessToken };