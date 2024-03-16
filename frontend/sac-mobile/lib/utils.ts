import clsx, { ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';
import {majorArr} from '@/lib/const';
/**
 * Nativewind CSS classnames generator
 * @param inputs - a list of classnames
 * @returns a string of Nativewind CSS classnames
 */
export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

// list of items for dropdown menu
export type Item = {
    label: string;
    value: string;
};

// list of graduation year
export const graduationYear = () => {
    var year = new Date().getFullYear();
    const graduationYear: Item[] = [];
    for (let i = 0; i < 5; i++) {
        graduationYear.push({
            label: String(year + i),
            value: String(year + i)
        });
    }
    return graduationYear;
};

export const major = () => {
    const major: Item[] = [];
    for (let i = 0; i < majorArr.length; i++) {
        major.push({
            label: majorArr[i],
            value: majorArr[i]
        });
    }
    return major;
}