import { createSlice } from '@reduxjs/toolkit';

type GlobalState = {
    platform: 'web' | 'mobile';
    loggedIn: boolean;
    accessToken?: string;
};

const initialState: GlobalState = {
    platform: 'mobile',
    loggedIn: false
};

export const globalSlice = createSlice({
    name: 'global',
    initialState,
    reducers: {
        setLoggedIn: (state, action) => {
            state.loggedIn = action.payload;
        },
        setAccessToken: (state, action) => {
            state.accessToken = action.payload;
        },
        resetAccessToken: (state) => {
            state.accessToken = undefined;
            state.loggedIn = false;
        }
    }
});

export const { setLoggedIn, setAccessToken, resetAccessToken } =
    globalSlice.actions;
export default globalSlice.reducer;
