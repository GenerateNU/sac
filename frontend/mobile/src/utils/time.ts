import { FormattedTime } from '@/src/types/time';

import { Options } from './timeOptions';

// calculate duration of event
export const calculateDuration = (start: string, end: string) => {
    const startTime = new Date(start);
    const endTime = new Date(end);
    const difference = endTime.getTime() - startTime.getTime();
    const hours = Math.floor(difference / (1000 * 60 * 60));
    const minutes = Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
    return `${hours}h ${minutes}m`;
};

// format time from Date type
export const formatTime = (timeString: string): FormattedTime => {
    const time = new Date(timeString);

    const daysOfWeek = [
        'Sunday',
        'Monday',
        'Tuesday',
        'Wednesday',
        'Thursday',
        'Friday',
        'Saturday'
    ];
    const months = [
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
    ];
    const dayOfWeek = daysOfWeek[time.getDay()];
    const date = time.getDate();
    const month = months[time.getMonth()];
    const year = time.getFullYear();

    let hours = time.getHours();
    const minutes = time.getMinutes().toString().padStart(2, '0');
    const ampm = hours >= 12 ? 'PM' : 'AM';

    const formattedHours = hours % 12 || 12;
    const formattedTime =
        minutes === '00' ? `${formattedHours}` : `${formattedHours}:${minutes}`;

    return {
        dayOfWeek,
        date: date.toString(),
        month,
        year: year.toString(),
        formattedTime,
        ampm,
        hour: formattedHours.toString(),
        minute: minutes
    };
};

// return true if an event is happening now
export const happeningNow = (startTime: string, endTime: string) => {
    const start = new Date(startTime);
    const end = new Date(endTime);
    const currentTime = new Date();
    return currentTime >= start && currentTime <= end;
};

// select what to display on event time
export const createOptions = <T extends keyof Options>(
    ...keys: T[]
): Partial<Record<T, true>> => {
    return keys.reduce(
        (acc, key) => {
            acc[key] = true;
            return acc;
        },
        {} as Partial<Record<T, true>>
    );
};

// select options to display on event time: day of the week, month and date, year, startTime, endTime
// e.g. eventTime(startTime, endTime, createOptions('dayOfWeek', 'monthAndDate'))
export const eventTime = (
    startTime: string,
    endTime: string,
    options: Partial<Record<keyof Options, true>> = {}
) => {
    const { dayOfWeek, monthAndDate, year, start, end } = options;

    const startTimeFormat = formatTime(startTime);
    const endTimeFormat = formatTime(endTime);

    const dayOfWeekString = dayOfWeek ? startTimeFormat.dayOfWeek : '';
    const monthAndDateString = monthAndDate
        ? `${startTimeFormat.month} ${startTimeFormat.date}`
        : '';
    const yearString = year ? startTimeFormat.year : '';

    let timeString = '';
    if (start && end) {
        timeString = `${startTimeFormat.formattedTime}-${endTimeFormat.formattedTime} ${endTimeFormat.ampm}`;
    } else if (start) {
        timeString = `${startTimeFormat.formattedTime} ${startTimeFormat.ampm}`;
    }

    const parts = [
        dayOfWeekString + (dayOfWeek ? ',' : ''),
        monthAndDateString + (timeString && monthAndDate ? ',' : ''),
        yearString + (yearString && timeString ? ' •' : ''),
        timeString
    ].filter(Boolean);

    return parts.join(' ');
};

export const googleCalendarTime = (timeString: string) => {
    const date = new Date(timeString);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');

    const formattedTime = `${year}${month}${day}T${hours}${minutes}${seconds}Z`;
    return formattedTime;
};
