import { createSlice } from '@reduxjs/toolkit';

export type UserState = {
    userId: string;
    email: string;
    accessToken?: string;
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
        }
    }
});

export const { setUser, resetUser, setAccessToken, resetAccessToken } =
    userSlice.actions;
export default userSlice.reducer;
