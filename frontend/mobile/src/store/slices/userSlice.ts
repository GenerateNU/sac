import { createSlice } from '@reduxjs/toolkit';
import { set } from 'react-hook-form';

export type UserState = {
    userId: string;
    email: string;
    accessToken?: string;
    sacLoginSession?: string;
};

const initialState: UserState = {
    userId: '',
    email: ''
};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUser: (state, action) => {
            state.userId = action.payload.userId;
            state.email = action.payload.email;
            state.accessToken = action.payload.accessToken;
        },
        resetUser: () => initialState,
        setAccessToken: (state, action) => {
            state.accessToken = action.payload;
        },
        resetAccessToken: (state) => {
            state.accessToken = undefined;
        },
        setSacLoginSession: (state, action) => {
            state.sacLoginSession = action.payload;
        }
    }
});

export const { 
    setUser, 
    resetUser, 
    setAccessToken, 
    resetAccessToken, 
    setSacLoginSession 
} = userSlice.actions;
export default userSlice.reducer;
