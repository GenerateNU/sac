import axios from 'axios';

import { Tokens, User } from '@/types/user';
import { API_BASE_URL } from '@/utils/const';

/**
 * Logins the user with the given email and password.
 * @param email The email of the user.
 * @param password The password of the user.
 * @returns The user that was logged in.
 */
export const loginByEmail = async (
    email: string,
    password: string
): Promise<{ user: User; tokens: Tokens }> => {
    try {
        const response = await axios.post(`${API_BASE_URL}/auth/login`, {
            email,
            password
        });
        const cookies = response.headers['set-cookie'];

        let accessToken = '';
        let refreshToken = '';
        cookies?.forEach((cookie: string) => {
            if (cookie.includes('access_token')) {
                accessToken = cookie.split('=')[1].split(';')[0];
            }
            if (cookie.includes('refresh_token')) {
                refreshToken = cookie.split('=')[1].split(';')[0];
            }
        });

        const user = response.data;

        console.log('[auth.ts] accessToken', accessToken);
        console.log('[auth.ts] refreshToken', refreshToken);

        return { user, tokens: { accessToken, refreshToken } };
    } catch (error) {
        console.error(error);
        throw new Error('Error logging in');
    }
};
