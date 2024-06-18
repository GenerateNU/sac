import { User } from '@generatesac/lib';
import { createSlice } from '@reduxjs/toolkit';

export type UserState = {
    loggedIn: boolean;
    accessToken?: string;
};

const initialState: UserState & User = {
    id: '',
    email: '',
    loggedIn: false,
    created_at: '',
    updated_at: '',
    name: '',
    role: 'student'
};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUser: (state, action) => {
            state.id = action.payload.id;
            state.email = action.payload.email;
            state.college = action.payload.college;
            state.graduation_cycle = action.payload.graduation_cycle;
            state.graduation_year = action.payload.graduation_year;
            state.major0 = action.payload.major0;
            state.major1 = action.payload.major1;
            state.major2 = action.payload.major2;
            state.name = action.payload.name;
            state.role = action.payload.role;
        },
        resetUser: () => initialState,
        setAccessToken: (state, action) => {
            state.accessToken = action.payload;
        },
        resetAccessToken: (state) => {
            state.accessToken = undefined;
            state.loggedIn = false;
        },
        setLoggedInState: (state, action) => {
            state.loggedIn = action.payload;
        }
    }
});

export const {
    setUser,
    resetUser,
    setAccessToken,
    resetAccessToken,
    setLoggedInState
} = userSlice.actions;
export default userSlice.reducer;
