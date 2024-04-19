import * as SecureStore from 'expo-secure-store';

import { AxiosResponseHeaders, RawAxiosResponseHeaders } from 'axios';
import { Buffer } from 'buffer';
import calendarize from 'calendarize';
import clsx, { ClassValue } from 'clsx';
import { addDays, isSameWeek, setDate, subDays } from 'date-fns';
import { twMerge } from 'tailwind-merge';

import { majorArr } from '@/lib/const';
import { Tokens } from '@/types/auth';
import { Club } from '@/types/club';
import { FAQ } from '@/types/faq';
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
 * Checks if a date is today
 * @param date - the date to check
 * @returns a boolean if the date is today
 */
const isToday = (date: Date) => {
    const today = new Date();
    const eventDate = new Date(date);

    return today.toDateString() === eventDate.toDateString();
};

/**
 * Updates a string if it's plural or not
 * @param count - the number of items
 * @param singular - the singular form of the item
 * @param plural - the plural form of the item
 * @returns a string of the item in singular or plural form
 */
export const pluralize = (count: number, singular: string, plural: string) => {
    return count === 1 ? singular : plural;
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
 * Accepts a date in exchange for a calendar view of that day, weekly or monthly.
 * @returns an array of dates that corresponds to a weekly or monthly view of a date.
 */
export const calendarizeHelper = (
    target: Date,
    type: 'week' | 'month'
): Date[][] => {
    let rows = calendarize(target);

    // Convert numbers to dates
    let dates = rows.map((week) => {
        return week.map((day) => {
            if (day === 0) {
                return undefined;
            } else {
                return setDate(target, day);
            }
        });
    });

    // Add adjacent months to starting and ending row
    for (let i = 5; i >= 0; i--) {
        if (dates[0][i] === undefined && dates[0][i + 1] !== undefined) {
            dates[0][i] = subDays(dates[0][i + 1]!, 1);
        }
    }

    let endIdx = dates.length - 1;
    for (let i = 1; i < 7; i++) {
        if (
            dates[endIdx][i - 1] !== undefined &&
            dates[endIdx][i] === undefined
        ) {
            dates[endIdx][i] = addDays(dates[endIdx][i - 1]!, 1);
        }
    }

    // Return only the week or the full month
    if (type === 'week') {
        for (const week of dates) {
            if (isSameWeek(target, week[0]!)) {
                return [week as Date[]];
            }
        }
    }

    return dates as Date[][];
};
/* Generates an array of item of major
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

const isClub = (item: any): Club => {
    return typeof item === 'object' && item.hasOwnProperty('is_recruiting');
};

const isEvent = (item: any): Event => {
    return (
        typeof item === 'object' &&
        item.hasOwnProperty('start_time') &&
        item.hasOwnProperty('end_time')
    );
};

const isFAQ = (item: any): FAQ => {
    return (
        typeof item === 'object' &&
        item.hasOwnProperty('question') &&
        item.hasOwnProperty('answer')
    );
};

const randomEventColorData = [
    {
        image: 'https://i.imgur.com/pLtAoSt.png',
        color: '#EDA990'
    },
    {
        image: 'https://i.imgur.com/1Auiabb.png',
        color: '#C5D9A4'
    },
    {
        image: 'https://i.imgur.com/pLtAoSt.png',
        color: '#EA90BB'
    },
    {
        image: 'https://i.imgur.com/pLtAoSt.png',
        color: '#EBC767'
    }
];

export {
    cn,
    isToday,
    isClub,
    isEvent,
    isFAQ,
    graduationYear,
    tokenCache,
    experationTime,
    extractTokens,
    randomEventColorData
};
