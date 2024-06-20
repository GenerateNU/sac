import { Event, Tag } from '@generatesac/lib';
import { createSlice } from '@reduxjs/toolkit';

type EventClubState = {
    shouldPreview?: boolean;
    tags: Tag[];
    rsvps: string[];
};

const initialState: Event & EventClubState = {
    id: '',
    created_at: '',
    updated_at: '',
    name: '',
    description: '',
    start_time: '',
    end_time: '',
    location: '',
    event_type: 'hybrid',
    is_archived: false,
    is_draft: false,
    is_recurring: false,
    is_public: false,
    preview: '',
    host: '',
    tags: [],
    rsvps: []
};

export const eventSlice = createSlice({
    name: 'event',
    initialState,
    reducers: {
        setEventId: (state, action) => {
            state.id = action.payload;
        },
        setEvent: (state, action) => {
            state.id = action.payload.id;
            state.created_at = action.payload.created_at;
            state.updated_at = action.payload.updated_at;
            state.name = action.payload.name;
            state.description = action.payload.description;
            state.start_time = action.payload.start_time;
            state.end_time = action.payload.end_time;
            state.location = action.payload.location;
            state.event_type = action.payload.event_type;
            state.image = action.payload.image;
            state.is_archived = action.payload.is_archived;
            state.is_draft = action.payload.is_draft;
            state.is_recurring = action.payload.is_recurring;
            state.is_public = action.payload.is_public;
            state.preview = action.payload.preview;
            state.host = action.payload.host;
        },
        setEventTags: (state, action) => {
            state.tags = action.payload;
        },
        resetEvent: () => initialState,
        setEventShouldPreview: (state, action) => {
            state.shouldPreview = action.payload;
        },
        setEventAttendees: (state, action) => {
            state.rsvps = action.payload;
        }
    }
});

export const {
    setEvent,
    resetEvent,
    setEventShouldPreview,
    setEventTags,
    setEventId,
    setEventAttendees
} = eventSlice.actions;
export default eventSlice.reducer;
