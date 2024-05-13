import { SET_EMAIL, SET_CLUB_NAME, SET_CLUB_ID, SET_CALENDAR_TOKEN, SET_BE_TOKEN, RESET } from './constants';

export const setEmail = (email: string) => ({
    type: SET_EMAIL,
    payload: email,
})

export const setClubName = (clubName: string) => ({
    type: SET_CLUB_NAME,
    payload: clubName,
})

export const setClubId = (clubId: string) => ({
    type: SET_CLUB_ID,
    payload: clubId,
})

export const setCalendarAccessToken = (calendarAccessToken: string) => ({
    type: SET_CALENDAR_TOKEN,
    payload: calendarAccessToken,
})

export const setBEAccessToken = (beAccessToken: string) => ({
    type: SET_BE_TOKEN,
    payload: beAccessToken,
})

export const reset = () => ({
    type: RESET,
})