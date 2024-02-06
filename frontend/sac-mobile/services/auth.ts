import { Tokens, User } from "@/types/user";
import { API_BASE_URL } from "@/utils/const";
import axios from "axios";

/**
 * Logins the user with the given email and password.
 * @param email The email of the user.
 * @param password The password of the user.
 * @returns The user that was logged in.
 */
export const loginByEmail = async (email: string, password: string): Promise<{ user: User, tokens: Tokens }> => {
    try {
        const response = await axios.post(`${API_BASE_URL}/auth/login`, { email, password });
        // response headers
        const cookies = response.headers['set-cookie']

        // extract tokens from cookies
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

        console.log('accessToken', accessToken);
        console.log('refreshToken', refreshToken);

        const user = response.data;
        return { user, tokens: { accessToken, refreshToken } };

    } catch (error) {
        console.error(error);
        throw new Error('Error logging in');
    }
}