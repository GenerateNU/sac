import { setRecruitment, updateRecruitment, resetRecruitment, recruitmentReducer } from '../state/recruitment/recruitmentSlice';
import { Recruitment, RecruitmentCycle, RecruitmentType, UpdateRecruitmentRequestBody } from '../../types/recruitment';

describe('recruitment slice', () => {
    const initialState: Recruitment = {
        recruitmentCycle: RecruitmentCycle.Always,
        recruitmentType: [],
        isRecruiting: false,
        applications: []
    }

    it('set recruitment', () => {
        const recruitment: Recruitment = {
            recruitmentCycle: RecruitmentCycle.Spring,
            recruitmentType: [RecruitmentType.Apply],
            isRecruiting: true,
            applications: []
        }

        const action = setRecruitment(recruitment);
        const nextState = recruitmentReducer(initialState, action);
        expect(nextState).toEqual({
            ...initialState,
            ...recruitment
        });
    })

    it('update recruitment', () => {
        const currentState: Recruitment = {
            recruitmentCycle: RecruitmentCycle.Spring,
            recruitmentType: [RecruitmentType.Apply],
            isRecruiting: true,
            applications: []
        }

        const updatedFields: UpdateRecruitmentRequestBody = {
            isRecruiting: false,
        }

        const action = updateRecruitment(updatedFields);
        const nextState = recruitmentReducer(currentState, action);
        expect(nextState).toEqual({
            ...currentState,
            ...updatedFields,
        })
    })

    it('reset recruitment', () => {
        const currentState: Recruitment = {
            recruitmentCycle: RecruitmentCycle.Spring,
            recruitmentType: [RecruitmentType.Apply],
            isRecruiting: true,
            applications: []
        }
    
        const action = resetRecruitment();
        const nextState = recruitmentReducer(currentState, action);
        expect(nextState).toEqual(initialState);
      });
})