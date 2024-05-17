import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Club, CreateClubRequestBody, UpdateClubRequestBody } from '../../../types/club';
import { produce } from 'immer';

const initialState: Club = {
    userId: "",
    clubId: "",
    clubName: "",
    accessToken: "",
    csrfToken: "",
    preview: "",
    description: "",
    weeklyTimeCommitment: 0,
    oneWordToDescribeUs: "",
    logo: "",
};

const clubSlice = createSlice({
    name: 'club',
    initialState,
    reducers: {
        setClub(state, action: PayloadAction<CreateClubRequestBody>) {
            return produce(state, draft => {
                Object.assign(draft, action.payload);
            });
        },
        updateClub(state, action: PayloadAction<UpdateClubRequestBody>) {
            return produce(state, draft => {
                Object.assign(draft, action.payload);
            });
        },
        resetClub() {
            return initialState;
        }
    }
})

export const { setClub, updateClub, resetClub } = clubSlice.actions;
export const clubReducer = clubSlice.reducer;