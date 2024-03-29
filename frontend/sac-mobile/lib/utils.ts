import clsx, { ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';
import { Buffer } from 'buffer';
import { Item } from '@/types/item';
import * as SecureStore from 'expo-secure-store';
import { AxiosResponseHeaders, RawAxiosResponseHeaders } from 'axios';
import { Tokens } from '@/types/user';

/**
 * Nativewind CSS classnames generator
 * @param inputs - a list of classnames
 * @returns a string of Nativewind CSS classnames
 */
export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

/**
 * Generates an array of graduation years
 * @returns an Item array of graduation years from the current year to the next 5 years
 */
export const graduationYear = () => {
    var year = new Date().getFullYear();
    const graduationYears: Item[] = [];
    for (let i = 0; i < 5; i++) {
        graduationYears.push({
            label: String(year + i),
            value: String(year + i)
        });
    }
    return graduationYears;
};

/**
 * Defines the token cache object.
 * This object is used to interact with the SecureStore API.
 * @func getToken: Retrieves a token from the SecureStore.
 * @func saveToken: Saves a token to the SecureStore.
 * @func deleteToken: Deletes a token from the SecureStore.
 * @returns The token cache object.
 * @example
 * const token = await tokenCache.getToken('accessToken');
 * await tokenCache.saveToken('accessToken', 'token');
 * await tokenCache.deleteToken('accessToken');
 * @see https://docs.expo.dev/versions/latest/sdk/securestore/
 */
export const tokenCache = {
    async getToken(key: string) {
        try {
            return SecureStore.getItemAsync(key);
        } catch (error) {
            return null;
        }
    },
    async saveToken(key: string, value: string) {
        try {
            return SecureStore.setItemAsync(key, value);
        } catch (error) {
            return null;
        }
    },
    async deleteToken(key: string) {
        try {
            return SecureStore.deleteItemAsync(key);
        } catch (error) {
            return null;
        }
    }
};

/**
 * Calculates the expiration time of a token.
 * @param token The token to calculate the expiration time of.
 * @param bufferTime The buffer time to add to the expiration time.
 * @returns The expiration time of the token.
 */
export const experationTime = (token: string, bufferTime: number) => {
    try {
        if (!token) {
            return 0;
        }
        const jwtBuffer = Buffer.from(token.split('.')[1], 'base64')
        const jwtPayload = JSON.parse(jwtBuffer.toString('utf-8'));
        const experationTime = jwtPayload.exp * 1000;
        return experationTime - (Date.now() + bufferTime);
    } catch (error) {
        return 0;
    }
}

/**
 * Extracts the access token and refresh token from the cookies and headers of an Axios response.
 * @param cookies The cookies from the Axios response.
 * @param headers The headers from the Axios response.
 * @returns The access token and refresh token.
 */
export const extractTokens = (cookies: string[] | undefined, headers: AxiosResponseHeaders | RawAxiosResponseHeaders): Tokens => {
    if (!cookies || !headers['authorization']) {
        throw new Error('No cookies or authorization header found');
    }

    let refreshToken = '';
    cookies?.forEach((cookie: string) => {
        if (cookie.includes('refresh_token')) {
            refreshToken = cookie.split('=')[1].split(';')[0];
        }
    });

    const accessToken = headers['authorization'].split(' ')[1]

    return { accessToken, refreshToken };
}
