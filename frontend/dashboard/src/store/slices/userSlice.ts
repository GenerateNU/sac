import { createSlice } from '@reduxjs/toolkit';

export type UserState = {
    userId: string;
    email: string;
    clubId: string;
    clubName: string;
    clubLogo: string;
    accessToken: string;
}

const initialState: UserState = {
    userId: "",
    email: "",
    clubId: "",
    clubName: "",
    clubLogo: "",
    accessToken: "",
};

const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUser: (state, action) => {
            state.userId = action.payload.userId;
            state.email = action.payload.email;
            state.clubId = action.payload.clubId;
            state.clubName = action.payload.clubName;
            state.clubLogo = action.payload.clubLogo;
        },
        resetUser: () => initialState,
        setAccessToken: (state, action) => {
            state.accessToken = action.payload;
        },
        resetAccessToken: state => {
            state.accessToken = "";
        }
    }
})

export const { 
    setUser, 
    resetUser,
    setAccessToken,
    resetAccessToken
} = userSlice.actions;
export default userSlice.reducer;