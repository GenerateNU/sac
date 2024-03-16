import clsx, { ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

import { Item } from '@/types/item';

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
