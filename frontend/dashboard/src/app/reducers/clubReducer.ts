import { Reducer } from "@reduxjs/toolkit/react";
import { ClubState, ClubAction } from '@/src/app/actions/actionTypes';

// default initial state
const initialState: ClubState = {
    email: '',
    clubName: '',
    clubId: '',
    calendarAccessToken: '',
    beAccessToken: '',
}

// create a new club state given an initial state and action
export const clubReducer: Reducer<ClubState, ClubAction> = (state = initialState, action) => {
    switch (action.type) {
        case 'SET_EMAIL': 
            return {...state, email: action.payload};
        case 'SET_CLUB_NAME': 
            return {...state, clubName: action.payload};
        case 'SET_CLUB_ID': 
            return {...state, clubId: action.payload};
        case 'SET_CALENDAR_TOKEN': 
            return {...state, calendarAccessToken: action.payload};
        case 'SET_BE_TOKEN': 
            return {...state, beAccessToken: action.payload};
        case 'RESET': 
            return initialState;
        default:
            return state;
    }
};