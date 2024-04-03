import clsx, { ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

import { Item } from '@/types/item';
import { addDays, getDay, getDaysInMonth, isSameWeek, setDate, subDays } from 'date-fns';
import calendarize from 'calendarize';

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
 * Accepts a date in exchange for a calendar view of that day, weekly or monthly.
 * @returns an array of dates that corresponds to a weekly or monthly view of a date.
 */
export const calendarizeHelper = (target: Date, type: 'week' | 'month'): Date[][] => {
    let rows = calendarize(target);

    // Convert numbers to dates
    let dates = rows.map((week, i, _) => {
        return week.map((day, j, _) => {
            if (day == 0) {
                return undefined;
            } else {
                return setDate(target, day);
            }
        })
    });

    // Add adjacent months to starting and ending row
    for (let i = 5; i >= 0; i--) {
        if (dates[0][i] === undefined && dates[0][i + 1] !== undefined) {
            dates[0][i] = subDays(dates[0][i + 1]!, 1);
        }
    }

    let endIdx = dates.length - 1;
    for (let i = 1; i < 7; i++) {
        if (dates[endIdx][i - 1] !== undefined && dates[endIdx][i] === undefined) {
            dates[endIdx][i] = addDays(dates[endIdx][i - 1]!, 1);
        }
    }

    // Return only the week or the full month
    if (type == 'week') {
        for (const week of dates) {
            if (isSameWeek(target, week[0]!)) {
                return [week as Date[]];
            }
        }
    }

    return dates as Date[][];
} 
