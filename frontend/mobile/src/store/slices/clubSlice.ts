import { Club, Event, Tag } from '@generatesac/lib';
import { createSlice } from '@reduxjs/toolkit';

type ClubState = {
    shouldPreview?: boolean;
    tags: Tag[];
    events: Event[];
};

const initialState: Club & ClubState = {
    id: '',
    created_at: '',
    updated_at: '',
    name: '',
    description: '',
    preview: '',
    num_members: 0,
    logo: '',
    weekly_time_committment: 0,
    one_word_to_describe_us: '',
    tags: [],
    events: [],
};

export const clubSlice = createSlice({
    name: 'club',
    initialState,
    reducers: {
        setClubId: (state, action) => {
            state.id = action.payload;
        },
        setClub: (state, action) => {
            state.id = action.payload.id;
            state.created_at = action.payload.created_at;
            state.updated_at = action.payload.updated_at;
            state.name = action.payload.name;
            state.description = action.payload.description;
            state.preview = action.payload.preview;
            state.num_members = action.payload.num_members;
            state.logo = action.payload.logo;
            state.weekly_time_committment = action.payload.weekly_time_committment;
            state.one_word_to_describe_us = action.payload.one_word_to_describe_us;
            state.recruitment = action.payload.recruitment;
        },
        setClubTags: (state, action) => {
            state.tags = action.payload;
        },
        setClubEvents: (state, action) => {
            state.events = action.payload;
        },
        resetClub: () => initialState,
        setClubShouldPreview: (state, action) => {
            state.shouldPreview = action.payload;
        }
    }
});

export const { 
    setClubId, 
    setClub, 
    setClubTags, 
    setClubEvents,
    resetClub, 
    setClubShouldPreview 
} = clubSlice.actions;
export default clubSlice.reducer;
