import { createSlice } from '@reduxjs/toolkit';

import { authApi } from '../api/authApi';

interface AuthState {
    isAuthenticated: boolean;
    loading: boolean;
    token: string | null;
    error: any; // TODO: define error type
}

const initialState: AuthState = {
    isAuthenticated: false,
    loading: false,
    token: null,
    error: null
};

const authSlice = createSlice({
    name: 'authSlice',
    initialState,
    reducers: {
        refreshToken: (state, action) => {
            state.token = action.payload.access_token;
        },
        setError: (state, action) => {
            state.error = action.payload;
        },
        setIsAuthenticated: (state, action) => {
            state.isAuthenticated = action.payload;
        },
        setToken: (state, action) => {
            state.token = action.payload;
        },
        logout: (state) => {
            authApi.endpoints.logout.initiate();
            state.token = null;
            state.isAuthenticated = false;
        }
    },
    extraReducers: (builder) => {
        builder
            .addMatcher(
                authApi.endpoints.login.matchFulfilled,
                (state, action) => {
                    console.log(
                        '[login fulfilled] action.payload',
                        action.payload
                    );
                    state.token = action.payload.token;
                    state.isAuthenticated = true;

                    // add to local storage
                    localStorage.setItem('access_token', action.payload.token);
                }
            )
            .addMatcher(
                authApi.endpoints.login.matchRejected,
                (state, action) => {
                    console.log(
                        '[login rejected] action.payload',
                        action.payload
                    );
                    state.error = action.error;
                }
            )
            .addMatcher(
                authApi.endpoints.register.matchFulfilled,
                (state, action) => {
                    console.log(
                        '[register fulfilled] action.payload',
                        action.payload
                    );
                    state.token = action.payload.token;
                    state.isAuthenticated = true;
                }
            )
            .addMatcher(
                authApi.endpoints.register.matchRejected,
                (state, action) => {
                    console.log(
                        '[register rejected] action.payload',
                        action.payload
                    );
                    state.error = action.error;
                }
            );
    }
});

export default authSlice.reducer;
export const { refreshToken, setError, setIsAuthenticated, setToken, logout } =
    authSlice.actions;
