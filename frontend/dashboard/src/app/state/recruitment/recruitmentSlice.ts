import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Recruitment, RecruitmentCycle } from '../../../types/recruitment';


const initialState: Recruitment = {
    recruitmentCycle: RecruitmentCycle.Always,
    recruitmentType: [],
    isRecruiting: false,
    applications: []
};

const recruitmentSlice = createSlice({
    name: 'recruitment',
    initialState,
    reducers: {
        setRecruitment(_, action: PayloadAction<Recruitment>) {
            return action.payload;
        },
        updateRecruitment(state, action: PayloadAction<Partial<Recruitment>>) {
            return { ...state, ...action.payload };
        },
        resetRecruitment() {
            return initialState;
        }
    }
});
  
export const { setRecruitment, updateRecruitment, resetRecruitment } = recruitmentSlice.actions;
export const recruitmentReducer = recruitmentSlice.reducer;