import { EventPreview, User } from '@generatesac/lib';
import { createSlice } from '@reduxjs/toolkit';

type UserState = {
    following: string[];
    rsvps: EventPreview[];
};

const initialState: User & UserState = {
    id: '',
    email: '',
    created_at: '',
    updated_at: '',
    name: '',
    role: 'student',
    following: [],
    rsvps: []
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
        setUserFollowing: (state, action) => {
            state.following = action.payload;
        },
        setUserRSVPs: (state, action) => {
            state.rsvps = action.payload;
        }
    }
});

export const { setUser, resetUser, setUserFollowing, setUserRSVPs } =
    userSlice.actions;
export default userSlice.reducer;
