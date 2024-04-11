import * as SecureStore from 'expo-secure-store';

import { AxiosResponseHeaders, RawAxiosResponseHeaders } from 'axios';
import { Buffer } from 'buffer';
import clsx, { ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

import { majorArr } from '@/lib/const';
import { categories } from '@/lib/const';
import { Tokens } from '@/types/auth';
import { Item } from '@/types/item';

/**
 * Nativewind CSS classnames generator
 * @param inputs - a list of classnames
 * @returns a string of Nativewind CSS classnames
 */
const cn = (...inputs: ClassValue[]) => {
    return twMerge(clsx(inputs));
};

/**
 * Generates an array of graduation years
 * @returns an Item array of graduation years from the current year to the next 5 years
 */
const graduationYear = () => {
    var year = new Date().getFullYear();
    const range = 5;
    const graduationYears: Item[] = [];
    for (let i = 0; i < range; i++) {
        graduationYears.push({
            label: String(year + i),
            value: String(year + i)
        });
    }
    return graduationYears;
};

/**
 * Generates an array of item of major
 * @returns an Item array of majors
 */
export const major = () => {
    const majors: Item[] = [];
    for (let i = 0; i < majorArr.length; i++) {
        majors.push({
            label: majorArr[i],
            value: majorArr[i]
        });
    }
    return majors;
};

/**
 * Combine all tags of different categories into one for All tab
 * @return a list of tags
 */
export const allTags = () => {
    let allTags: string[] = [];
    for (let i = 0; i < categories.length; i++) {
        allTags = allTags.concat(categories[i].tags);
    }
    return allTags;
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
const tokenCache = {
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
const experationTime = (token: string, bufferTime: number = 0) => {
    try {
        if (!token) {
            return 0;
        }
        const jwtBuffer: Buffer = Buffer.from(token.split('.')[1], 'base64');
        const jwtPayload: { exp: number } = JSON.parse(
            jwtBuffer.toString('utf-8')
        );
        return jwtPayload.exp * 1000 - (Date.now() + bufferTime);
    } catch (error) {
        return 0;
    }
};

/**
 * Extracts the access token and refresh token from the cookies and headers of an Axios response.
 * @param cookies The cookies from the Axios response.
 * @param headers The headers from the Axios response.
 * @returns The access token and refresh token.
 */
const extractTokens = (
    cookies: string[] | undefined,
    headers: AxiosResponseHeaders | RawAxiosResponseHeaders
): Tokens => {
    if (!cookies || !headers.authorization) {
        throw new Error('No cookies or authorization header found');
    }

    let refreshToken = '';
    cookies?.forEach((cookie: string) => {
        if (cookie.includes('refresh_token')) {
            refreshToken = cookie.split('=')[1].split(';')[0];
        }
    });

    const accessToken = headers.authorization.split(' ')[1];

    return { accessToken, refreshToken };
};

export { cn, graduationYear, tokenCache, experationTime, extractTokens };
