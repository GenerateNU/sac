/**
 * Given a date, returns the day of the week as a string
 * @param date the date to get the day of the week from
 * @returns the day of the week as a string
 */
export const getDayOfWeek = (date: Date) => {
    return [
        'Monday',
        'Tuesday',
        'Wednesday',
        'Thursday',
        'Friday',
        'Saturday',
        'Sunday'
    ][date.getDay()];
};

/**
 * Given a date, returns the month as a string
 * @param date the date to get the month from
 * @returns the month as a string
 */
export const getMonth = (date: Date) => {
    return [
        'January',
        'February',
        'March',
        'April',
        'May',
        'June',
        'July',
        'August',
        'September',
        'October',
        'November',
        'December'
    ][date.getMonth()];
};

/**
 * Given a number, returns the ordinal suffix of the number
 * @param n the number to get the ordinal suffix from
 * @returns the ordinal suffix of the number
 */
export const getOrdinalSuffix = (n: number) => {
    const s = ['th', 'st', 'nd', 'rd'];
    const v = n % 100;
    return s[(v - 20) % 10] || s[v] || s[0];
};

/**
 * Given a date, returns the formatted date as a string
 * @param date the date to format
 * @returns the formatted date as a string
 * @example getFormattedDate(new Date(2024, 2, 15)) // '1:25 PM'
 */
export const getFormattedTime = (date: Date) => {
    return date.toLocaleTimeString('en-US', {
        hour: 'numeric',
        minute: '2-digit',
        hour12: true
    });
};
