import { CalendarLink } from '@/src/types/calendarLink';

import { googleCalendarTime } from './time';

export const firstLetterUppercase = (str: string) => {
    return str.charAt(0).toUpperCase() + str.slice(1);
};

export const httpFormat = (str: string) => {
    str = encodeURIComponent(str);
    const specialCharactersMap: { [key: string]: string } = {
        "'": '%27'
    };

    for (const char in specialCharactersMap) {
        if (str.includes(char)) {
            str = str.replace(
                new RegExp(char, 'g'),
                specialCharactersMap[char]
            );
        }
    }
    return str;
};

export const createGoogleCalendarLink = ({
    eventName,
    eventDetail,
    location,
    startTime,
    endTime
}: CalendarLink) => {
    const googleCalendar = `https://calendar.google.com/calendar/u/0/r/eventedit?text=${httpFormat(eventName)}&dates=${googleCalendarTime(startTime)}/${googleCalendarTime(endTime)}&location=${httpFormat(location)}&details=${httpFormat(eventDetail)}&sf=true&output=xml`;
    return googleCalendar;
};

export const formatCategoryOrTag = (str: string) => {
    if (str.includes('Pre')) {
        const index = str.indexOf('Pre');
        str = str.slice(0, index) + 'Pre-' + str.slice(index + 3);
        return str;
    }
    if (str.includes('And')) {
        str = str.split('And').join(' and');
    }
    if (str === str.toUpperCase()) {
        return str;
    }
    return str.replace(/([A-Z])/g, ' $1').trim();
};