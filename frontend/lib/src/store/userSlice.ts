import { createSlice } from '@reduxjs/toolkit';

import { authApi } from '../api/authApi';
import { User } from '../types/user';

interface UserState {
    user: User | null;
    loading: boolean;
    error: any; // TODO: define error type
}

const initialState: UserState = {
    user: null,
    loading: false,
    error: null
};

const userSlice = createSlice({
    name: 'userSlice',
    initialState,
    reducers: {
        setUser: (state, action) => {
            state.user = action.payload;
        }
    },
    extraReducers: (builder) => {
        builder
            .addMatcher(
                authApi.endpoints.login.matchFulfilled,
                (state, action) => {
                    state.user = action.payload.user;
                }
            )
            .addMatcher(
                authApi.endpoints.login.matchRejected,
                (state, action) => {
                    state.error = action.error;
                }
            )
            .addMatcher(
                authApi.endpoints.register.matchFulfilled,
                (state, action) => {
                    state.user = action.payload.user;
                }
            )
            .addMatcher(authApi.endpoints.logout.matchFulfilled, (state) => {
                state.user = null;
            });
    }
});

export const { setUser } = userSlice.actions;
export default userSlice.reducer;
