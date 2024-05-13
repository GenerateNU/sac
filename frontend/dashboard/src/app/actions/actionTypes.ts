import { SET_EMAIL, SET_CLUB_NAME, SET_CLUB_ID, SET_CALENDAR_TOKEN, SET_BE_TOKEN, RESET } from './constants';

// club state interface
interface ClubState {
    email: string;
    clubName: string;
    clubId: string;
    calendarAccessToken: string;
    beAccessToken: string;
}

// club actions
type ClubAction = SetEmailAction | SetClubNameAction | SetClubId | SetCalendarAccessToken | SetBEAccessToken | ResetClub;

interface SetEmailAction {
    type: typeof SET_EMAIL,
    payload: string,
}

interface SetClubNameAction {
    type: typeof SET_CLUB_NAME,
    payload: string,
}

interface SetClubId {
    type: typeof SET_CLUB_ID,
    payload: string,
}

interface SetCalendarAccessToken {
    type: typeof SET_CALENDAR_TOKEN,
    payload: string,
}

interface SetBEAccessToken {
    type: typeof SET_BE_TOKEN,
    payload: string,
}

interface ResetClub {
    type: typeof RESET
}

export type {
    ClubState,
    ClubAction,
    SetEmailAction,
    SetClubNameAction,
    SetClubId,
    SetCalendarAccessToken,
    SetBEAccessToken,
    ResetClub
};
